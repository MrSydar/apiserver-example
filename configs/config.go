package configs

import (
	stdlog "log"

	"mrsydar/apiserver/configs/auth0"
	"mrsydar/apiserver/configs/database"
	"mrsydar/apiserver/configs/environment"
	"mrsydar/apiserver/configs/log"
)

func init() {
	stdlog.Print("Initializing configs")

	if err := log.Init(); err != nil {
		stdlog.Fatalf("failed to initialize logger config: %v", err)
	}

	if err := environment.Init(); err != nil {
		stdlog.Fatalf("failed to initialize environment config: %v", err)
	}

	if err := database.Init(); err != nil {
		stdlog.Fatalf("failed to initialize database config: %v", err)
	}

	if err := auth0.Init(); err != nil {
		stdlog.Fatalf("failed to initialize auth0 config: %v", err)
	}
}
