package mastodon

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// API contains necessary informations to work with Mastodons API.
type API struct {
	Base        string
	Prefix      string
	AccessToken string
}

// Do executes an API request. The method is a HTTP method, e.g. GET or POST.
func (api API) Do(method string, endpoint string, values url.Values) (io.ReadCloser, error) {
	client := &http.Client{}
	r := bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest(method, api.Base+api.Prefix+endpoint, r)
	if err != nil {
		return nil, fmt.Errorf("could not create request to %s: %v", endpoint, err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api.AccessToken))

	switch method {
	case http.MethodGet:
		req.URL.RawQuery = values.Encode()
	case http.MethodPost:
		req.Form = values
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not execute %s: %v", endpoint, err)
	}

	switch res.StatusCode {
	case http.StatusOK:
		return res.Body, nil
	default:
		err := api.getError(res.Body)
		return nil, fmt.Errorf("%s: %v", res.Status, err)
	}
}

// Get request
func (api API) Get(endpoint string, values url.Values, dest interface{}) error {
	return api.generic(http.MethodGet, endpoint, values, dest)
}

// Post request
func (api API) Post(endpoint string, values url.Values, dest interface{}) error {
	return api.generic(http.MethodPost, endpoint, values, dest)
}

// Delete request
func (api API) Delete(endpoint string, values url.Values, dest interface{}) error {
	return api.generic(http.MethodDelete, endpoint, values, dest)
}

func (api API) generic(method, endpoint string, values url.Values, dest interface{}) error {
	r, err := api.Do(method, endpoint, values)
	if err != nil {
		return fmt.Errorf("could not %s %s: %v", method, endpoint, err)
	}
	defer r.Close()

	if err := json.NewDecoder(r).Decode(dest); err != nil {
		return fmt.Errorf("could not decode %s: %v", endpoint, err)
	}
	return nil
}

func (api API) getError(r io.ReadCloser) error {
	defer r.Close()
	res := Error{}
	if err := json.NewDecoder(r).Decode(&res); err != nil {
		return fmt.Errorf("could not decode error: %v", err)
	}
	return errors.New(res.Error)
}
