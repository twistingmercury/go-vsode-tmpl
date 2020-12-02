package main

import (
	"fmt"
	"log"
	"os"
	"rest_api/config"
	"rest_api/httpd"

	flags "github.com/spf13/pflag"
)

// general flags
var (
	ver = flags.BoolP("version", "v", false, "Displays the build information for rest_api.")
	cfg = flags.StringP("cfg-file", "c", "config", "Sets the name of the configuration file to load.")
)

// versioning
var (
	buildDate    = "{not set}"
	buildVersion = "{not set}"
	buildCommit  = "{not set}"
)

func main() {
	flags.Parse()
	if err := config.Initialize(*cfg, buildVersion, buildCommit); err != nil {
		log.Fatalln(err)
	}

	if *ver {
		fmt.Printf("rest_api - version %s build date: %s commit: %s\n", buildVersion, buildDate, buildCommit)
		fmt.Println(config.String())
		os.Exit(0)
	}

	if err := httpd.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
