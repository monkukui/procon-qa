# Simple Web Application

内部での処理内容については Qiita の [ここ](https://qiita.com/x-color/items/24ff2491751f55e866cf) にまとめてあります。

## ディレクトリ構成

```
simple-webapp/
├── db
│   └── sample.db
├── handler
│   ├── auth.go
│   └── handler.go
├── model
│   ├── db.go
│   ├── todo.go
│   └── user.go
├── public
│   ├── assets
│   │   └── js
│   │       ├── login.js
│   │       ├── signup.js
│   │       └── todoList.js
│   ├── index.html
│   ├── login.html
│   ├── signup.html
│   └── todos.html
├── README.md
├── go.mod
├── go.sum
├── main.go
└── router.go
```

## サーバーの実行

```
$ go run main.go router.go
```

## ページ表示用URL

### トップページを表示

`GET /`

### ユーザーの登録画面を表示

`GET /signup`

### ユーザーのログイン画面を表示

`GET /login`

### ユーザーのTodo一覧を表示

`GET /todos`

## ユーザー認証API

### ユーザーの登録

`POST /signup`

以下の値をJSON形式で送信する

name     | value
---------|------
name     | Bob
password | bob-pass

登録に成功すると以下のようなJSON形式のデータが返る

```json
{
    "id":1,
    "name":"Bob",
    "password":""
}
```

### ユーザーの認証

`POST /login`

以下の値をJSON形式で送信する

name     | value
---------|------
name     | Bob
password | bob-pass

認証が成功すると以下のような形式でJWTが返る

```json
{
    "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsIm5hbWUiOiJCb2IiLCJleHAiOjE1NTQ4ODA0NDh9._Z5Aq7_bOWAz99HXrVPbiWwF3euijMvG__f8p9sWXu8"
}
```

## Todo操作API

以下はJWTをリクエストに含めないとアクセスできない (認証済である必要がある)

### ユーザーのTodo一覧の取得

`GET /api/todos`

以下のようなJSON形式のTodo一覧が返る

```json
[
    {
        "uid":1,
        "id":1,
        "name":"first todo",
        "completed":true
    },
    {
        "uid":1,
        "id":2,
        "name":"second todo",
        "completed":false
    }
]
```

### 新たなTodoの登録

`POST /api/todos`

以下の値をJSON形式で送信する

name     | value
---------|------
name     | third todo

登録に成功した場合、以下のようなJSON形式登録されたTodoが返る

```json
{
    "uid":1,
    "id":3,
    "name":"third todo",
    "completed":false
}
```

### 指定されたIDのTodoの削除

`DELETE /api/todos/:id`

### 指定されたIDのTodoの完了状態の変更

`PUT /api/todos/:id/completed`

## 使用技術

- Go言語 (https://golang.org/)
- Echo (https://echo.labstack.com/)
- GORM (http://gorm.io/)
- SQLite (https://www.sqlite.org/index.html)
- JWT (https://jwt.io/)
