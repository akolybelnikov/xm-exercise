package main

import (
	"os"

	"github.com/akolybelnikov/xm-exercise/internal/app"

	"github.com/akolybelnikov/xm-exercise/internal/config"
)

func main() {
	// Get the environment name from the environment variable
	env := os.Getenv("APP_ENV")
	// Load the configuration
	cfg, cfgErr := config.NewConfig(env)
	if cfgErr != nil {
		panic(cfgErr)
	}
	// Run the main function of the application
	if err := app.Run(cfg); err != nil {
		panic(err)
	}
}
