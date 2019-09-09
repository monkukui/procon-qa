サーバ起動
```
go run main.go router.go
```

### DB 設計
```
Questions {
   id: int (primary_key)
   uid: int
   title: string
   body: string
   tagid: int
   state (未完了, 回答中, 解決済み)
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

### TODO

#### 優先度高

- [ ] (GET) 質問の全取得
- [ ] (GET) 未回答の質問を全取得
- [ ] (GET) 特定のユーザが投稿した質問の全取得
- [ ] (POST) 質問の投稿
- [ ] (PUT) 質問の編集
- [ ] (DELETE) 質問の削除

####  優先度中
