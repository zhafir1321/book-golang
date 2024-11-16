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
(If you use postman, use 'x-www-urlencoded' as a body content-type)

### /api/auth/register (POST)

If you dont have an account, you can register

```json
{
  "name": "test",
  "email": "test@mail.com",
  "password": "12345"
}


{
    "message": "Success create user"
}
```

```json
{
    "name": "",
    "email": "test1@mail.com",
    "password": "12345"
}

{
    "message": "Name is required"
}
```

If theres a same email

```json
{
    "name": "test",
    "email": "test@mail.com",
    "password": "12345"
}

{
    "message": "Email already exists"
}
```

### /api/auth/login (POST)

```json
{
  "email": "test@mail.com",
  "password": "12345"
}

{
    "data": "<token_string>"
}
```

### /api/book (GET)

```json
{
  "data": ["<data-books>"]
}
```

### /api/book/{id} (GET)

```json
{
  "data": "<data-book>"
}
```

### /api/book/ (POST) Admin ONLY

You can change manually role of user in your database to admin

```json
{
  "title": "test",
  "category": "comic"
}

{
    "message": "Success create book"
}
```

### /api/book/{id} (PUT) Admin ONLY

You can change manually role of user in your database to admin

```json
{
    "title": "Football Skill"
}

{
    "message": "Success update book"
}
```

### /api/book/{id} (DELETE) Admin ONLY

You can change manually role of user in your database to admin'

```json
{
  "message": "Success delete book"
}
```

### /api/book/search?title={title} (GET)

```json
{
  "data": ["<data-books>"]
}
```

### /api/book/filter?category={category} (GET)

```json
{
  "data": ["<data-books>"]
}
```

### /api/borrow/{bookId} (POST)

```json
{
    "message": "Success borrow book"
}

{
    "message": "Book not found"
}
```

### /api/borrow/{id} (PUT)

```json
{
  "message": "Success return book"
}

{
    "message": "ID not exist"
}
```

This is all, thanks for reading and try it. Let me know if theres a idea or advice on this project.
