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

- [] hoge
