### サーバ起動
```
go run main.go router.go
```
`dist/` 以下の静的なファイルを参照する.
client 側を変更したら, 
```
npm run build
```
を実行して, `dist/` に反映させる.

### DB 設計
```
Questions {
  id: int ( primary_key )
  uid: int
  tid: int
  
  title: string
  body: string
  date: string
  state: string
  url: string
  completed: boolean
}

Answers {
  id: int ( primary_key )
  uid: int
  qid: int

  body: string
  date: string
}

tags {
  id: int (primary_key)
  string: body ( "tag1,tag2,...,tagn" )
}

Users {
  id: int (primary_key)
  name: string
  pass: string (仮)
  rate: int (仮)
}
```

### API

## GET
質問を全取得する (GET)
```
/api/questions

[
  {
    "uid":1,
    "tid":2,
    "id":3,
    "title":"タイトル1",
    "body":"本文1",
    "url":"https://atcoder/abc123/proba",
    "state":"",
    "date":"1996/05/26",
    "completed":true
  },
  {
    "uid":1,
    "tid":0,
    "id":4,
    "title":"aa",
    "body":"bb",
    "url":"",
    "state":"",
    "date":"",
    "completed":false
  },
]
```

質問をページ取得する (GET)
```
/api/questions/<page(1-indexed)>

[
  {
    "uid":1,
    "tid":2,
    "id":3,
    "title":"タイトル1",
    "body":"本文1",
    "url":"https://atcoder/abc123/proba",
    "state":"",
    "date":"1996/05/26",
    "completed":true
  },
  {
    "uid":1,
    "tid":0,
    "id":4,
    "title":"aa",
    "body":"bb",
    "url":"",
    "state":"",
    "date":"",
    "completed":false
  },
]
```

質問を 1 つ取得する (GET)
```
/api/question/<question_id>
{
  "uid":1,
  "tid":2,
  "id":3,
  "title":"タイトル1",
  "body":"本文1",
  "url":"https://atcoder/abc123/proba",
  "state":"",
  "date":"1996/05/26",
  "completed":true
},
```
ユーザー 1 つを取得する (GET)
```
/api/user/<user_id>
{
  "id":3,
  "name":"monkukui",
  "password":"mon"
}
```

## POST
質問を投稿する (POST)
```
/api/questions
{
  title: "タイトル"
  body: "本文"
  date: "yyyy/mm/dd-hh:mm:ss"
  state: "WA"
  url: "https://atcoder.jp/contests/jsc2019-qual/submissions/7120555"
}
```

## DELETE
質問を削除する (DELETE)
自分の投稿じゃない場合, 404 がかえる

```
/api/question/<question_id>
```



### TODO

#### 優先度高

- [x] (GET) 質問の全取得
- [x] (GET) 質問の固定長取得 (10 個ほど)
- [x] (GET) 質問を 1 つ取得
- [ ] (GET) 未回答の質問を全取得
- [ ] (GET) 特定のユーザが投稿した質問の全取得
- [ ] (GET) 特定の質問に対する, 回答の全取得
- [x] (POST) 質問の投稿
- [ ] (POST) 回答の投稿
- [ ] (PUT) 質問の編集
- [ ] (PUT) 回答の編集
- [X] (DELETE) 質問の削除
- [ ] (DELETE) 回答の削除

####  優先度中

- [ ] (GET) タグ検索

#### 優先度低
- [ ] (POST) ファボ機能
- [ ] (PUT) レート変動
