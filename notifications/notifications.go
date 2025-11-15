package notifications

import alerts "github.com/tomasharkema/systemd-alert"

type creator func() alerts.Notifier

var Plugins = map[string]creator{}

func Add(name string, creator creator) {
	Plugins[name] = creator
}
