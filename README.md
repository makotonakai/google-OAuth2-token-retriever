# google-OAuth2-token-retriever

## 前提条件

Google CloudでOAuth 2.0 Clientを作成している


## 手順

1. 以下の通りに.envを作成する
2. `go run main.go`を実行する
3. `Visit the URL for the auth dialog: ` に続くURLにアクセスする
4. Googleログインした後に\[コード\] に当たる部分を`Enter your code (the one which starts from code= in the URL): `にペーストする
```http://localhost:8080/?state=state&code=[コード]&scope=```
5. アクセストークンを取得する
