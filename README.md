# mastodon-go

## Usage

```go
func main() {
	app, err := mastodon.NewApp("https://mastodon.social", "mastodon-go", "urn:ietf:wg:oauth:2.0:oob", []string{"read", "write", "follow"}, "")
	if err != nil {
		log.Fatal(err)
	}

	// redirect user to auth url (skip if you stored an auth token)
	url := app.AuthCodeURL()
	fmt.Printf("goto: %s\ntoken: ", url)
	code := ""
	fmt.Scanf("%s", &code)

	// exchange auth code for auth token
	token, err := app.Exchange(code)
	if err != nil {
		log.Fatal(err)
	}
	app.SetToken(token)
	fmt.Printf("access token: %s\n", token)

	// request authenticated user
	user, err := app.Accounts.VerifyCredentials()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("signed in as %s (%d)\n", user.Username, user.ID)

	// toot!
	status, err := app.Statuses.Update("toot!", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tooted: %s\n", status.URL)
}
```

See [example/main.go](example/main.go).
