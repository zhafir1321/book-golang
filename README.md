# Book Golang

This is my first project using Go as a backend. This is only sending API as JSON, but you need to register and login to make it all works otherwise its just gonna send

```
{ "message": "Unauthorized"}
```

## 1. Clone this Repo

Clone this Repo to your computer

```
git clone https://github.com/zhafir1321/book-golang.git
```

## 2. Install air-verse

This is like nodemon, but its for Go. So, when inside project theres a change, you dont need to restart the server.

```
go install github.com/air-verse/air@latest

air init

air
```

## 2. Or you can use Docker

You can change my settings into yours

```
docker-compose up --build
```

## 3. Endpoints

Heres a list of my endpoints

### 1. /api/auth/register

If you dont have an account, you can register

```json
{
  "name": "test",
  "email": "test@mail.com",
  "password": "12345"
}
```
