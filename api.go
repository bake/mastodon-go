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

// Do executes an API request. The method is a HTTP method, e.g. GET or POST.
func (app App) Do(method string, endpoint string, values url.Values) (io.ReadCloser, error) {
	client := &http.Client{}
	r := bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest(method, base+endpoint, r)
	if err != nil {
		return nil, fmt.Errorf("could not create request to %s: %v", endpoint, err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", app.Token.AccessToken))

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
		err := app.getError(res.Body)
		return nil, fmt.Errorf("%s: %v", res.Status, err)
	}
}

// Get request
func (app App) Get(endpoint string) (io.ReadCloser, error) {
	return app.Do(http.MethodGet, endpoint, nil)
}

// Post request
func (app App) Post(endpoint string) (io.ReadCloser, error) {
	return app.Do(http.MethodPost, endpoint, nil)
}

func (app App) generic(method, endpoint string, values url.Values, dest interface{}) error {
	r, err := app.Do(method, endpoint, values)
	if err != nil {
		return fmt.Errorf("could not %s %s: %v", method, endpoint, err)
	}
	defer r.Close()

	if err := json.NewDecoder(r).Decode(dest); err != nil {
		return fmt.Errorf("could not decode %s: %v", endpoint, err)
	}

	return nil
}

func (app App) getError(r io.ReadCloser) error {
	defer r.Close()
	res := Error{}
	if err := json.NewDecoder(r).Decode(&res); err != nil {
		return fmt.Errorf("could not decode error: %v", err)
	}
	return errors.New(res.Error)
}
