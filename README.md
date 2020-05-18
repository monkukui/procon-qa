# PROCON-QA
Source code of [https://procon-qa.herokuapp.com/].

# About
競技プログラミング専用の Q&A サイトです．

# Requi

### 検証環境 (docker-compose)
docker-composeを使います。
```
git clone https://github.com/monkukui/procon-qa.git
cd procon-qa
docker-compose build
docker-compose up
```
としてから `localhost:8080` にアクセスしてください。
検証が終わったら `docker-compose down` してください。
