package mastodon

import (
	"net/url"
)

type Reports struct {
	api *API
}

// Get returns a list of reports made by the authenticated user.
func (reports Reports) Get() ([]Report, error) {
	r := []Report{}
	return r, reports.api.Get("reports", nil, &r)
}

// Report reports a user and returns the finished report.
func (reports Reports) Report(account, status string, comment string) (Report, error) {
	r := Report{}
	v := url.Values{
		"account_id": {account},
		"status_ids": {status},
		"comment":    {comment},
	}
	return r, reports.api.Post("reports", v, &r)
}
