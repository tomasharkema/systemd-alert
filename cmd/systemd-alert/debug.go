package main

import (
	"time"

	"github.com/coreos/go-systemd/v22/dbus"
	alerts "github.com/tomasharkema/systemd-alert"
	"github.com/tomasharkema/systemd-alert/notifications/debug"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type debugAlert struct {
	uconn, conn *dbus.Conn
	Frequency   time.Duration
}

func (t *debugAlert) configure(cmd *kingpin.CmdClause) {
	cmd.Action(t.execute)
	cmd.Flag("frequency", "frequency to emit events").Default("1s").DurationVar(&t.Frequency)
}

func (t *debugAlert) execute(c *kingpin.ParseContext) error {
	go alerts.Run(t.conn, alerts.AlertNotifiers(debug.NewAlerter()), alerts.AlertFrequency(t.Frequency))
	go alerts.SafeRun(t.uconn, alerts.AlertNotifiers(debug.NewAlerter()), alerts.AlertFrequency(t.Frequency))
	return nil
}
