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
	if err := accounts.api.Get("accounts/verify_credentials", nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Get returns an account.
func (accounts Accounts) Get(id int) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d", id)
	if err := accounts.api.Get(end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Followers returns an slice of following accounts.
func (accounts Accounts) Followers(id int) ([]Account, error) {
	accs := []Account{}
	end := fmt.Sprintf("accounts/%d/followers", id)
	if err := accounts.api.Get(end, nil, &accs); err != nil {
		return accs, err
	}
	return accs, nil
}

// Following returns an slice of followed accounts.
func (accounts Accounts) Following(id int) ([]Account, error) {
	accs := []Account{}
	end := fmt.Sprintf("accounts/%d/following", id)
	if err := accounts.api.Get(end, nil, &accs); err != nil {
		return accs, err
	}
	return accs, nil
}

// Statuses returns an slice of statuses. Accepted params are:
// only_media: Only return statuses that have media attachments
// exclude_replies: Skip statuses that reply to other statuses
func (accounts Accounts) Statuses(id int, params url.Values) ([]Status, error) {
	end := fmt.Sprintf("accounts/%d/statuses", id)
	statuses := []Status{}
	if err := accounts.api.Get(end, params, &statuses); err != nil {
		return statuses, err
	}
	return statuses, nil
}

// Follow an user.
func (accounts Accounts) Follow(id int) (Account, error) {
	end := fmt.Sprintf("accounts/%d/follow", id)
	acc := Account{}
	if err := accounts.api.Get(end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Unfollow an account
func (accounts Accounts) Unfollow(id int) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d/unfollow", id)
	if err := accounts.api.Post(end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Block an account
func (accounts Accounts) Block(id int) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d/block", id)
	if err := accounts.api.Post(end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Unblock an account.
func (accounts Accounts) Unblock(id int) (Account, error) {
	end := fmt.Sprintf("accounts/%d/unblock", id)
	acc := Account{}
	if err := accounts.api.Post(end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Mute an account.
func (accounts Accounts) Mute(id int) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d/mute", id)
	if err := accounts.api.Post(end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Unmute an user.
func (accounts Accounts) Unmute(id int) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d/unmute", id)
	if err := accounts.api.Post(end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Relationships returns an slice of Relationships of the current user to a
// list of given accounts.
func (accounts Accounts) Relationships(ids ...int) ([]Relationship, error) {
	idss := []string{}
	for _, id := range ids {
		idss = append(idss, strconv.Itoa(id))
	}
	v := url.Values{"id": idss}
	rels := []Relationship{}
	if err := accounts.api.Get("accounts/relationships", v, &rels); err != nil {
		return rels, err
	}
	return rels, nil
}

// Search returns an slice of matching Accounts. Will lookup an account
// remotely if the search term is in the username@domain format and not yet in
// the database.
func (accounts Accounts) Search(q string, limit int) ([]Account, error) {
	v := url.Values{
		"q":     {q},
		"limit": {strconv.Itoa(limit)},
	}
	accs := []Account{}
	if err := accounts.api.Get("accounts/search", v, &accs); err != nil {
		return nil, err
	}
	return accs, nil
}
