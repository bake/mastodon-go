package mastodon

import "fmt"

type Notifications struct {
	api *API
}

// Get returns a list of notifications for the authenticated user.
func (notifications Notifications) Get() ([]Notification, error) {
	n := []Notification{}
	if err := notifications.api.Get("notifications", nil, &n); err != nil {
		return nil, err
	}
	return n, nil
}

// GetSingle returns the notification.
func (notifications Notifications) GetSingle(id string) (Notification, error) {
	n := Notification{}
	end := fmt.Sprintf("notifications/%s", id)
	if err := notifications.api.Get(end, nil, &n); err != nil {
		return n, err
	}
	return n, nil
}

// Clear deletes all notifications from the Mastodon server for the
// authenticated user.
func (notifications Notifications) Clear() error {
	return notifications.api.Get("notifications/clear", nil, nil)
}
