package mastodon

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// VerifyCredentials returns the authenticated user's account.
func (app App) VerifyCredentials() (Account, error) {
	acc := Account{}
	if err := app.generic(http.MethodGet, "accounts/verify_credentials", nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// GetAccount returns an account.
func (app App) GetAccount(id int) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d", id)
	if err := app.generic(http.MethodGet, end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// GetFollowers returns an slice of following accounts.
func (app App) GetFollowers(id int) ([]Account, error) {
	accs := []Account{}
	end := fmt.Sprintf("accounts/%d/followers", id)
	if err := app.generic(http.MethodGet, end, nil, &accs); err != nil {
		return accs, err
	}
	return accs, nil
}

// GetFollowing returns an slice of followed accounts.
func (app *App) GetFollowing(id int) ([]Account, error) {
	accs := []Account{}
	end := fmt.Sprintf("accounts/%d/following", id)
	if err := app.generic(http.MethodGet, end, nil, &accs); err != nil {
		return accs, err
	}
	return accs, nil
}

// GetStatuses returns an slice of statuses. Accepted params are:
// only_media: Only return statuses that have media attachments
// exclude_replies: Skip statuses that reply to other statuses
func (app *App) GetStatuses(id int, params url.Values) ([]Status, error) {
	end := fmt.Sprintf("accounts/%d/statuses", id)
	statuses := []Status{}
	if err := app.generic(http.MethodGet, end, params, &statuses); err != nil {
		return statuses, err
	}
	return statuses, nil
}

// Follow an user.
func (app App) Follow(id int) (Account, error) {
	end := fmt.Sprintf("accounts/%d/follow", id)
	acc := Account{}
	if err := app.generic(http.MethodPost, end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Unfollow an account
func (app App) Unfollow(id int) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d/unfollow", id)
	if err := app.generic(http.MethodPost, end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Block an account
func (app App) Block(id int) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d/block", id)
	if err := app.generic(http.MethodPost, end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Unblock an account.
func (app App) Unblock(id int) (Account, error) {
	end := fmt.Sprintf("accounts/%d/unblock", id)
	acc := Account{}
	if err := app.generic(http.MethodPost, end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Mute an account.
func (app App) Mute(id int) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d/mute", id)
	if err := app.generic(http.MethodPost, end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Unmute an user.
func (app App) Unmute(id int) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d/unmute", id)
	if err := app.generic(http.MethodPost, end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Relationships returns an slice of Relationships of the current user to a
// list of given accounts.
func (app App) Relationships(ids ...int) ([]Relationship, error) {
	idss := []string{}
	for _, id := range ids {
		idss = append(idss, strconv.Itoa(id))
	}
	v := url.Values{"id": idss}
	rels := []Relationship{}
	if err := app.generic(http.MethodGet, "accounts/relationships", v, &rels); err != nil {
		return rels, err
	}
	return rels, nil
}

// SearchAccount returns an slice of matching Accounts. Will lookup an account
// remotely if the search term is in the username@domain format and not yet in
// the database.
func (app App) SearchAccount(q string, limit int) ([]Account, error) {
	v := url.Values{
		"q":     {q},
		"limit": {strconv.Itoa(limit)},
	}
	accs := []Account{}
	if err := app.generic(http.MethodGet, "accounts/search", v, &accs); err != nil {
		return nil, err
	}
	return accs, nil
}
