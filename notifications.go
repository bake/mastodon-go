package mastodon

import (
	"fmt"
	"net/http"
)

// GetNotifications returns a list of notifications for the authenticated user.
func (app App) GetNotifications() ([]Notification, error) {
	n := []Notification{}
	if err := app.generic(http.MethodGet, "notifications", nil, &n); err != nil {
		return nil, err
	}
	return n, nil
}

// GetNotification returns the notification.
func (app App) GetNotification(id int) (Notification, error) {
	n := Notification{}
	end := fmt.Sprintf("notifications/%d", id)
	if err := app.generic(http.MethodGet, end, nil, &n); err != nil {
		return n, err
	}
	return n, nil
}

// ClearNotifications deletes all notifications from the Mastodon server for the authenticated user.
func (app App) ClearNotifications() error {
	if err := app.generic(http.MethodGet, "notifications/clear", nil, nil); err != nil {
		return err
	}
	return nil
}
