package mastodon

import (
	"fmt"
	"net/url"
	"strconv"
)

type Accounts struct {
	api *API
}

// VerifyCredentials returns the authenticated user's account.
func (accounts Accounts) VerifyCredentials() (Account, error) {
	acc := Account{}
	return acc, accounts.api.Get("accounts/verify_credentials", nil, &acc)
}

// Get returns an account.
func (accounts Accounts) Get(id string) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d", id)
	return acc, accounts.api.Get(end, nil, &acc)
}

// Followers returns an slice of following accounts.
func (accounts Accounts) Followers(id string) ([]Account, error) {
	accs := []Account{}
	end := fmt.Sprintf("accounts/%d/followers", id)
	return accs, accounts.api.Get(end, nil, &accs)
}

// Following returns an slice of followed accounts.
func (accounts Accounts) Following(id string) ([]Account, error) {
	accs := []Account{}
	end := fmt.Sprintf("accounts/%d/following", id)
	return accs, accounts.api.Get(end, nil, &accs)
}

// Statuses returns an slice of statuses. Accepted params are:
// only_media: Only return statuses that have media attachments
// exclude_replies: Skip statuses that reply to other statuses
func (accounts Accounts) Statuses(id string, params url.Values) ([]Status, error) {
	end := fmt.Sprintf("accounts/%d/statuses", id)
	statuses := []Status{}
	return statuses, accounts.api.Get(end, params, &statuses)
}

// Follow an user.
func (accounts Accounts) Follow(id string) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d/follow", id)
	return acc, accounts.api.Get(end, nil, &acc)
}

// Unfollow an account
func (accounts Accounts) Unfollow(id string) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d/unfollow", id)
	return acc, accounts.api.Post(end, nil, &acc)
}

// Block an account
func (accounts Accounts) Block(id string) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d/block", id)
	return acc, accounts.api.Post(end, nil, &acc)
}

// Unblock an account.
func (accounts Accounts) Unblock(id string) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d/unblock", id)
	return acc, accounts.api.Post(end, nil, &acc)
}

// Mute an account.
func (accounts Accounts) Mute(id string) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d/mute", id)
	return acc, accounts.api.Post(end, nil, &acc)
}

// Unmute an user.
func (accounts Accounts) Unmute(id string) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d/unmute", id)
	return acc, accounts.api.Post(end, nil, &acc)
}

// Relationships returns an slice of Relationships of the current user to a
// list of given accounts.
func (accounts Accounts) Relationships(ids ...int) ([]Relationship, error) {
	idss := []string{}
	for _, id := range ids {
		idss = append(idss, strconv.Itoa(id))
	}
	rels := []Relationship{}
	v := url.Values{"id": idss}
	return rels, accounts.api.Get("accounts/relationships", v, &rels)
}

// Search returns an slice of matching Accounts. Will lookup an account
// remotely if the search term is in the username@domain format and not yet in
// the database.
func (accounts Accounts) Search(q string, limit int) ([]Account, error) {
	accs := []Account{}
	v := url.Values{
		"q":     {q},
		"limit": {strconv.Itoa(limit)},
	}
	return accs, accounts.api.Get("accounts/search", v, &accs)
}
