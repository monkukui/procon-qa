# PROCON-QA
Source code of https://procon-qa.herokuapp.com.


競技プログラミング専用の Q&A サイトです．

## Requirements
- npm 5.14.2
- go1.12.9 darwin/amd64
- postgres (PostgreSQL) 12.1
- Docker 19.03.8

## 検証環境 (docker-compose)
docker-composeを使います。
```
git clone https://github.com/monkukui/procon-qa.git
cd procon-qa
docker-compose build
docker-compose up
```
としてから `localhost:8080` にアクセスしてください。
検証が終わったら `docker-compose down` してください。

## Contributing
いつもサイトを使っていただきありがとうございます．
こんな機能があったらいいなとか，ここはこうしたほうがいいとかがあったら，気軽の issue に書き込んでください．

もし，開発の手助けをしてくれる方がいたら歓迎します．
