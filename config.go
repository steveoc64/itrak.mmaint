package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

// Runtime variables, held in external file config.json
type iTrakMMaintConfig struct {
	Debug          bool
	DataSourceName string
	WebPort        int
}

var itrak iTrakMMaintConfig

// Load the config.json file, and override with runtime flags
func LoadConfig() {
	cf, err := os.Open("config.json")
	if err != nil {
		log.Println("Could not open config.json :", err.Error())
	}

	data := json.NewDecoder(cf)
	if err = data.Decode(&itrak); err != nil {
		log.Fatalln("Failed to load config.json :", err.Error())
	}

	flag.BoolVar(&itrak.Debug, "debug", itrak.Debug, "Enable Debugging")
	flag.StringVar(&itrak.DataSourceName, "", itrak.DataSourceName, "DataSourceName for SQLServer")
	flag.IntVar(&itrak.WebPort, "webport", itrak.WebPort, "Port Number for Web Server")
	flag.Parse()

	log.Printf("Starting\n\tDebug: \t\t%t\n\tSQLServer: \t%s\n\tWeb Port: \t%d\n",
		itrak.Debug,
		itrak.DataSourceName,
		itrak.WebPort)
}
