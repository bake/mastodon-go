package mastodon

import (
	"net/url"
	"strconv"
)

type Reports struct {
	api *API
}

// Get returns a list of reports made by the authenticated user.
func (reports Reports) Get() ([]Report, error) {
	r := []Report{}
	if err := reports.api.Get("reports", nil, &r); err != nil {
		return r, err
	}
	return r, nil
}

// Report reports a user and returns the finished report.
func (reports Reports) Report(account, status int, comment string) (Report, error) {
	r := Report{}
	v := url.Values{
		"account_id": {strconv.Itoa(account)},
		"status_ids": {strconv.Itoa(status)},
		"comment":    {comment},
	}
	if err := reports.api.Post("reports", v, &r); err != nil {
		return r, err
	}
	return r, nil
}
