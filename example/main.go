package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bakerolls/mastodon"
)

func main() {
	token := flag.String("token", "", "Token")
	flag.Parse()

	app, err := mastodon.NewApp("mastodon-go", "urn:ietf:wg:oauth:2.0:oob", []string{"read", "write", "follow"}, "")
	if err != nil {
		log.Fatal(err)
	}

	if *token == "" {
		url := app.AuthCodeURL()
		fmt.Printf("goto: %s\ntoken: ", url)
		fmt.Scanf("%s", token)

		if err := app.Exchange(*token); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("access token: %s\n", app.Token.AccessToken)
	} else {
		app.SetToken(*token)
	}

	user, err := app.VerifyCredentials()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("signed in as %s\n", user.Username)
}
