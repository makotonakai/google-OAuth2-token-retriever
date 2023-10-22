package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtubeanalytics/v2"
)

func main() {

	ctx := context.Background()
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Failed to load .env")
	}

	config := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH2_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH2_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_OAUTH2_REDIRECT_URL"),
		Scopes:       []string{youtubeanalytics.YoutubeScope, youtubeanalytics.YoutubeReadonlyScope, youtubeanalytics.YoutubepartnerScope, youtubeanalytics.YtAnalyticsMonetaryReadonlyScope, youtubeanalytics.YtAnalyticsReadonlyScope},
		Endpoint:     google.Endpoint,
	}

	authCodeURL := config.AuthCodeURL("state")
	fmt.Printf("Visit the URL for the auth dialog: %v\n", authCodeURL)

	fmt.Printf("Enter your code (the one which starts from code= in the URL): ")

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	dict, err := url.ParseQuery(code)
	if err != nil {
		fmt.Println(err)
	}

	keys := make([]string, len(dict), len(dict))
	i := 0
	for key := range dict {
		keys[i] = key
		i++
	}

	decodedCode := keys[0]
	tok, err := config.Exchange(ctx, string(decodedCode))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tok.AccessToken)
}
