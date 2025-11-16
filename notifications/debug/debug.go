package debug

import (
	"log"

	"github.com/coreos/go-systemd/v22/dbus"
	alerts "github.com/tomasharkema/systemd-alert"
	"github.com/tomasharkema/systemd-alert/notifications"
)

func init() {
	notifications.Add("debug", func() alerts.Notifier {
		return NewAlerter()
	})
}

// NewAlerter configures the Alerter
func NewAlerter() *Alerter {
	return &Alerter{}
}

// Alerter - sends an alert to a webhook.
type Alerter struct{}

// Alert about the provided units.
func (t Alerter) Alert(units ...*dbus.UnitStatus) {
	for _, unit := range units {
		log.Println("alert", unit)
	}
}
