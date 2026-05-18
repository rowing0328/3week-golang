package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStoreCreatesListsGetsUpdatesAndDeletesUsers(t *testing.T) {
	store := NewStore()

	alice := store.Create(userRequest{Name: "Alice", Age: 25, Email: "alice@example.com"})
	bob := store.Create(userRequest{Name: "Bob", Age: 17, Email: "bob@example.com"})

	if alice.ID != 1 || bob.ID != 2 {
		t.Fatalf("사용자 ID는 1부터 순서대로 증가해야 합니다: alice=%d bob=%d", alice.ID, bob.ID)
	}

	minAge := 20
	adults := store.List(&minAge)
	if len(adults) != 1 || adults[0].Name != "Alice" {
		t.Fatalf("minAge 필터 결과가 예상과 다릅니다: %#v", adults)
	}

	got, ok := store.Get(alice.ID)
	if !ok || got.Email != "alice@example.com" {
		t.Fatalf("생성한 사용자를 조회해야 합니다: got=%#v ok=%v", got, ok)
	}

	updated, ok := store.Update(alice.ID, userRequest{Name: "Alice Kim", Age: 26, Email: "alice.kim@example.com"})
	if !ok || updated.Name != "Alice Kim" || updated.Age != 26 || updated.Email != "alice.kim@example.com" {
		t.Fatalf("사용자 수정 결과가 예상과 다릅니다: got=%#v ok=%v", updated, ok)
	}

	deleted, ok := store.Delete(bob.ID)
	if !ok || deleted.Name != "Bob" {
		t.Fatalf("사용자 삭제 결과가 예상과 다릅니다: got=%#v ok=%v", deleted, ok)
	}

	if _, ok := store.Get(bob.ID); ok {
		t.Fatalf("삭제한 사용자는 다시 조회되지 않아야 합니다")
	}
}

func TestHTTPCRUDFlow(t *testing.T) {
	handler := newRouter(NewStore())

	created := requestJSON[User](t, handler, http.MethodPost, "/users", `{"name":"Alice","age":25,"email":"alice@example.com"}`, http.StatusCreated)
	if created.ID != 1 || created.Name != "Alice" || created.Age != 25 || created.Email != "alice@example.com" {
		t.Fatalf("생성 응답이 예상과 다릅니다: %#v", created)
	}

	users := requestJSON[[]User](t, handler, http.MethodGet, "/users", "", http.StatusOK)
	if len(users) != 1 || users[0] != created {
		t.Fatalf("목록 응답이 예상과 다릅니다: %#v", users)
	}

	got := requestJSON[User](t, handler, http.MethodGet, "/users/1", "", http.StatusOK)
	if got != created {
		t.Fatalf("단건 조회 응답이 예상과 다릅니다: %#v", got)
	}

	updated := requestJSON[User](t, handler, http.MethodPut, "/users/1", `{"name":"Alice Kim","age":26,"email":"alice.kim@example.com"}`, http.StatusOK)
	if updated.ID != 1 || updated.Name != "Alice Kim" || updated.Age != 26 || updated.Email != "alice.kim@example.com" {
		t.Fatalf("수정 응답이 예상과 다릅니다: %#v", updated)
	}

	deleted := requestJSON[User](t, handler, http.MethodDelete, "/users/1", "", http.StatusOK)
	if deleted != updated {
		t.Fatalf("삭제 응답이 예상과 다릅니다: %#v", deleted)
	}

	requestJSON[map[string]string](t, handler, http.MethodGet, "/users/1", "", http.StatusNotFound)
}

func TestHTTPMinAgeFilter(t *testing.T) {
	handler := newRouter(NewStore())

	requestJSON[User](t, handler, http.MethodPost, "/users", `{"name":"Alice","age":25,"email":"alice@example.com"}`, http.StatusCreated)
	requestJSON[User](t, handler, http.MethodPost, "/users", `{"name":"Bob","age":17,"email":"bob@example.com"}`, http.StatusCreated)

	users := requestJSON[[]User](t, handler, http.MethodGet, "/users?minAge=20", "", http.StatusOK)
	if len(users) != 1 || users[0].Name != "Alice" {
		t.Fatalf("minAge 필터 응답이 예상과 다릅니다: %#v", users)
	}
}

func TestHTTPValidationErrors(t *testing.T) {
	handler := newRouter(NewStore())

	tests := []struct {
		name   string
		method string
		path   string
		body   string
		want   string
	}{
		{
			name:   "invalid json",
			method: http.MethodPost,
			path:   "/users",
			body:   `{`,
			want:   "invalid json body",
		},
		{
			name:   "missing name",
			method: http.MethodPost,
			path:   "/users",
			body:   `{"name":"","age":25,"email":"alice@example.com"}`,
			want:   "name is required",
		},
		{
			name:   "negative age",
			method: http.MethodPost,
			path:   "/users",
			body:   `{"name":"Alice","age":-1,"email":"alice@example.com"}`,
			want:   "age must be greater than or equal to 0",
		},
		{
			name:   "missing email",
			method: http.MethodPost,
			path:   "/users",
			body:   `{"name":"Alice","age":25,"email":""}`,
			want:   "email is required",
		},
		{
			name:   "invalid minAge",
			method: http.MethodGet,
			path:   "/users?minAge=-1",
			body:   "",
			want:   "minAge must be greater than or equal to 0",
		},
		{
			name:   "invalid id",
			method: http.MethodGet,
			path:   "/users/abc",
			body:   "",
			want:   "invalid user id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := requestJSON[map[string]string](t, handler, tt.method, tt.path, tt.body, http.StatusBadRequest)
			if got["error"] != tt.want {
				t.Fatalf("예상과 다른 에러 메시지입니다: got %q, want %q", got["error"], tt.want)
			}
		})
	}
}

func TestMethodNotAllowed(t *testing.T) {
	handler := newRouter(NewStore())

	got := requestJSON[map[string]string](t, handler, http.MethodPatch, "/users", "", http.StatusMethodNotAllowed)
	if got["error"] != "method not allowed" {
		t.Fatalf("예상과 다른 에러 메시지입니다: got %q", got["error"])
	}
}

func requestJSON[T any](t *testing.T, handler http.Handler, method, path, body string, wantStatus int) T {
	t.Helper()

	req := httptest.NewRequest(method, path, bytes.NewBufferString(body))
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	if recorder.Code != wantStatus {
		t.Fatalf("예상과 다른 HTTP 상태입니다: got %d body %s, want %d", recorder.Code, recorder.Body.String(), wantStatus)
	}

	var got T
	if err := json.NewDecoder(recorder.Body).Decode(&got); err != nil {
		t.Fatalf("JSON 응답을 디코딩할 수 없습니다: %v body %s", err, recorder.Body.String())
	}
	return got
}
