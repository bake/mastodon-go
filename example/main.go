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

	app, err := mastodon.NewApp("mastodon-go", "urn:ietf:wg:oauth:2.0:oob", []string{"read", "write", "follow"}, "")
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
		app.SetToken(token.AccessToken)
		fmt.Printf("access token: %s\n", app.Token.AccessToken)
	} else {
		app.SetToken(*token)
	}

	user, err := app.VerifyCredentials()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("signed in as %s\n", user.Username)

	// users, err := app.SearchAccount("bak", 100)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, user := range users {
	// 	fmt.Println(user.Username)
	// }

	// followers, err := app.GetFollowers(user.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, follower := range followers {
	// 	fmt.Printf("%d %s\n", follower.ID, follower.Username)
	// }

	// rels, err := app.Relationships(user.ID, 34733)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, rel := range rels {
	// 	fmt.Printf("%+v\n", rel)
	// }

	// statuses, err := app.GetStatuses(user.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, status := range statuses {
	// 	fmt.Printf("%s: %s\n", status.CreatedAt, status.Content)
	// }
}
