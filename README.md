# 페이히어 백엔드 엔지니어 과제

### 시작하기
```shell
docker compose up -d
```

서버 호스트: localhost

포트:80

## 회원가입
- 고객은 이메일과 비밀번호 입력을 통해서 회원 가입을 할 수 있습니다.
### Request
```http request
POST /register
```
### Json Sample
```json
{
  "email": "cloq@kakao.com",
  "name": "kwanok",
  "password": "password"
}
```

## 로그인
- 고객은 회원 가입이후, 로그인을 할 수 있습니다.
### Request
```http request
POST /login
```
### Json Sample
```json
{
  "email": "cloq@kakao.com",
  "password": "password"
}
```
### Response
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjMxYmU1NTg2LWE1YTMtNDlkMS1hMDYxLTI3ZDdiMmFiY2M1ZiIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTY1MDQ1ODM0MSwidXNlcl9pZCI6IjAwYjgxNmEzLWQxNjgtNDRlZC1hOTgzLTAzYmE4ODJhNWI0ZCJ9.tSABTkKP__sqMAkUFt4WgF_VC019D7wRwebq6n_GMZk",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTEwNjIyNDEsInJlZnJlc2hfdXVpZCI6IjA0ZjhiM2JlLThjNjMtNDllNy05ODNiLTE1NzU2OTgxMmZiNyIsInVzZXJfaWQiOiIwMGI4MTZhMy1kMTY4LTQ0ZWQtYTk4My0wM2JhODgyYTViNGQifQ.P1QxRW4gE_fzHY4oJiGEVbMBNcaDqY5weuasC7yA5XU"
}
```

## 토큰 갱신

### Request
```http request
POST /refresh
```
### Json Sample
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTEwNjIyNDEsInJlZnJlc2hfdXVpZCI6IjA0ZjhiM2JlLThjNjMtNDllNy05ODNiLTE1NzU2OTgxMmZiNyIsInVzZXJfaWQiOiIwMGI4MTZhMy1kMTY4LTQ0ZWQtYTk4My0wM2JhODgyYTViNGQifQ.P1QxRW4gE_fzHY4oJiGEVbMBNcaDqY5weuasC7yA5XU",
}
```
### Response
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6ImFiNzIyNjQyLWEzOGYtNGY0NS05ZWY1LTZiYWQ0ODEwMTZkNSIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTY1MDQ1ODQ2MywidXNlcl9pZCI6IjAwYjgxNmEzLWQxNjgtNDRlZC1hOTgzLTAzYmE4ODJhNWI0ZCJ9.Xa7nTLUOr00ca5cKetSEK5LaasEljV-PUr6yDaY59mc",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTEwNjIzNjMsInJlZnJlc2hfdXVpZCI6IjllNWFkMzg3LTA1MGMtNGQ3MS1iMGMxLTFmYzYyZWVhNDZkOSIsInVzZXJfaWQiOiIwMGI4MTZhMy1kMTY4LTQ0ZWQtYTk4My0wM2JhODgyYTViNGQifQ.Jel2LKdP5tDuyTFqv0T33pupkuY4LoJCOScqDHjIeYA"
}
```

## 로그아웃
- 고객은 회원 가입이후, 로그아웃을 할 수 있습니다.
### Request
```http request
POST /refresh
Authorization: Bearer ${access_token}
```

### Response
```json
"Successfully logged out"
```

## 가계부에 금액과 가격 남기기
- 가계부에 오늘 사용한 돈의 금액과 관련된 메모를 남길 수 있습니다.
### Request
```http request
POST /pays
Authorization: Bearer ${access_token}
```
### Json Sample
```json
{
  "price": 5000,
  "memo": "김밥"
}
```
### Response
- 소비내역의 ID(UUID)를 반환합니다.
```json
"00b816a3-d168-44ed-a983-03ba882a5b4d"
```

## 가계부 수정하기
- 가계부에서 수정을 원하는 내역은 금액과 메모를 수정 할 수 있습니다.
### Request
```http request
PATCH /pays/:id
Authorization: Bearer ${access_token}
```
### Json Sample
```json
{
  "price": 3000,
  "memo": "김밥이 아니었음"
}
```
### Response
- 성공시
```http request
204 No Contents
```

## 소비내역 삭제하기
- 가계부에서 삭제를 원하는 내역은 삭제 할 수 있습니다.
### Request
```http request
DELETE /pays/:id
Authorization: Bearer ${access_token}
```
### Response
- 성공시
```http request
204 No Contents
```

## 삭제내역 복구하기
- 삭제한 내역은 언제든지 다시 복구 할 수 있어야 한다.
### Request
```http request
PATCH /pays/restore/:id
Authorization: Bearer ${access_token}
```
### Response
- 성공시
```http request
200 OK
```
```json
"restored"
```

## 가계부 리스트 가져오기
- 가계부에서 이제까지 기록한 가계부 리스트를 볼 수 있습니다.
### Request
```http request
GET /pays
Authorization: Bearer ${access_token}
```
### Response
- 성공시
```http request
200 OK
```
```json
[
  {
    "id": "33960e0f-f5d9-4054-9775-6100e56110d8",
    "memo": "김밥",
    "price": 5000,
    "createdAt": "2022-04-20 21:37:12",
    "updatedAt": "2022-04-20 21:37:12"
  },
  {
    "id": "69edf581-902e-451f-90b8-4ac55f153530",
    "memo": "돈까스였나?",
    "price": 5000,
    "createdAt": "2022-04-20 21:35:43",
    "updatedAt": "2022-04-20 21:37:12"
  }
  ...
]
```

## 가계부 내역 하나 가져오기
- 가계부에서 상세한 세부 내역을 볼 수 있습니다.
### Request
```http request
GET /pays
Authorization: Bearer ${access_token}
```
### Response
- 성공시
```http request
200 OK
```
```json
{
  "id": "33960e0f-f5d9-4054-9775-6100e56110d8",
  "memo": "김밥",
  "price": 5000,
  "createdAt": "2022-04-20 21:37:12",
  "updatedAt": "2022-04-20 21:37:12"
}
```

## 미인증시
- 로그인하지 않은 고객은 가계부 내역에 대한 접근 제한 처리가 되어야 합니다.
### Request
```http request
GET /pays
```
```json
{
  "result": "there is invalid",
  "status": 401
}
```

## 유효하지 않은 토큰인 경우
### Request
```http request
GET /pays
Authorization: Bearer 11223344
```
```json
{
  "result": "unauthorized",
  "status": 401
}
```
