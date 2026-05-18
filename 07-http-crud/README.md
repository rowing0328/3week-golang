# 07. HTTP CRUD

## 목표

- `net/http`로 HTTP 서버를 실행합니다.
- JSON 요청 본문을 구조체로 디코딩하고 JSON 응답을 반환합니다.
- 메모리 슬라이스를 저장소처럼 사용해 CRUD 흐름을 구현합니다.

## 실행

```bash
go run ./07-http-crud
```

서버가 실행되면 다른 터미널에서 아래 요청을 실행합니다.

## API 테스트

### 목록 조회

```bash
curl http://localhost:8080/users
```

### 나이 필터 목록 조회

```bash
curl http://localhost:8080/users?minAge=20
```

### 생성

```bash
curl -X POST http://localhost:8080/users \
  -H 'Content-Type: application/json' \
  -d '{"name":"Alice","age":25,"email":"alice@example.com"}'
```

### 단건 조회

```bash
curl http://localhost:8080/users/1
```

### 수정

```bash
curl -X PUT http://localhost:8080/users/1 \
  -H 'Content-Type: application/json' \
  -d '{"name":"Alice Kim","age":26,"email":"alice.kim@example.com"}'
```

### 삭제

```bash
curl -X DELETE http://localhost:8080/users/1
```

## 핵심 설명

- `http.HandleFunc`는 URL 경로와 처리 함수를 연결합니다.
- `json.NewDecoder(r.Body).Decode(&req)`는 요청 JSON을 구조체에 채웁니다.
- `json.NewEncoder(w).Encode(value)`는 응답 값을 JSON으로 씁니다.
- 현재 저장소는 메모리 슬라이스입니다. 서버를 재시작하면 데이터가 사라집니다.

## 실습

- `GET /users?minAge=20`처럼 특정 나이 이상만 조회하는 기능을 추가해 보세요.
- `Email` 필드는 생성, 수정, 조회 응답에 포함됩니다.
- `type Store struct`가 메모리 저장소 로직을 메서드로 관리합니다.
