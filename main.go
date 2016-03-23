package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/jerrac/servicecheckbeat/beater"
)

func main() {
	err := beat.Run("servicecheckbeat", "", beater.New())
	if err != nil {
		os.Exit(1)
	}
}
