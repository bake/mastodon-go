package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BakeRolls/mastodon-go"
)

func main() {
	token := flag.String("token", "", "Token")
	flag.Parse()

	app, err := mastodon.NewApp("https://mastodon.social", "mastodon-go", "urn:ietf:wg:oauth:2.0:oob", []string{"read", "write", "follow"}, "")
	if err != nil {
		log.Fatal(err)
	}

	if *token == "" {
		url := app.AuthCodeURL()
		fmt.Printf("goto: %s\ntoken: ", url)
		fmt.Scanf("%s", token)

		token, err := app.Exchange(*token)
		if err != nil {
			log.Fatal(err)
		}
		app.SetToken(token)
		fmt.Printf("access token: %s\n", app.API)
	} else {
		app.SetToken(*token)
	}

	user, err := app.Accounts.VerifyCredentials()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("signed in as %s (%s)\n", user.Username, user.ID)

	followers, err := app.Accounts.Followers(user.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d followers:\n", len(followers))
	for _, follower := range followers {
		fmt.Println(follower.Username)
	}
}
