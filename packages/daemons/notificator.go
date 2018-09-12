package daemons

import (
	"context"

	"github.com/ug93tad/go-apla/packages/notificator"
)

// Notificate is sending notifications
func Notificate(ctx context.Context, d *daemon) error {
	notificator.SendNotifications()
	return nil
}
