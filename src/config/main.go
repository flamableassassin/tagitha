package config

import "log/slog"

var log = slog.With(slog.Group("config"))
