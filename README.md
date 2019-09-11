### サーバ起動
```
go run main.go router.go
```

### DB 設計
```
Questions {
  id: int ( primary_key )
  user_id: int
  tag_id: int
  
  title: string
  body: string
  date: string
  state: string
  url: string
  state: int (未完了, 回答中, 解決済み)
}

Answers {
  id: int ( primary_key )
  user_id: int
  question_id

  body: string
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

質問を全取得する (GET)
```
/api/questions
```

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


### TODO

#### 優先度高

- [ ] (GET) 質問の全取得
- [ ] (GET) 質問の固定長取得 (10 個ほど)
- [ ] (GET) 質問を 1 つ取得
- [ ] (GET) 未回答の質問を全取得
- [ ] (GET) 特定のユーザが投稿した質問の全取得
- [ ] (GET) 特定の質問に対する, 回答の全取得
- [ ] (POST) 質問の投稿
- [ ] (POST) 回答の投稿
- [ ] (PUT) 質問の編集
- [ ] (PUT) 回答の編集
- [ ] (DELETE) 質問の削除
- [ ] (DELETE) 回答の削除

####  優先度中

- [ ] (GET) タグ検索

#### 優先度低
- [ ] (POST) ファボ機能
- [ ] (PUT) レート変動
