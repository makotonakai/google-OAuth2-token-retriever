# google-OAuth2-token-retriever

## 前提条件

Google Cloud上で、収益化できているYouTubeチャンネルに紐づいているGoogleアカウントのOAuth 2.0 Clientを作成している

作成方法は[こちら](https://developers.google.com/identity/protocols/oauth2/web-server?hl=ja)を参照

## 手順

1. 以下の通りに.envを作成する
   ```
   GOOGLE_OAUTH2_CLIENT_ID=(OAuth 2.0 ClientのClient ID)
   GOOGLE_OAUTH2_CLIENT_SECRET=(OAuth 2.0 ClientのClient Secret)
   GOOGLE_OAUTH2_REDIRECT_URL=(OAuth 2.0 ClientのRedirect URL e.g. http://localhost:8080)
   ```
2. `go run main.go`を実行する
3. `Visit the URL for the auth dialog: ` に続くURLにアクセスする
4. YouTubeのアカウントに紐づいているGoogleアカウントにログインした後にリダイレクト先の\[コード\] に当たる部分を`Enter your code (the one which starts from code= in the URL): `にペーストする
   

```(GOOGLE_OAUTH2_REDIRECT_URL)/?state=state&code=[コード]&scope=```

5. アクセストークンを取得する
