package mastodon

import "fmt"

// Notifications implements methods under /notifications.
type Notifications struct {
	api *API
}

// Get returns a list of notifications for the authenticated user.
func (notifications Notifications) Get() ([]Notification, error) {
	n := []Notification{}
	return n, notifications.api.Get("notifications", nil, &n)
}

// GetSingle returns the notification.
func (notifications Notifications) GetSingle(id string) (Notification, error) {
	n := Notification{}
	end := fmt.Sprintf("notifications/%s", id)
	return n, notifications.api.Get(end, nil, &n)
}

// Clear deletes all notifications from the Mastodon server for the
// authenticated user.
func (notifications Notifications) Clear() error {
	return notifications.api.Get("notifications/clear", nil, nil)
}
