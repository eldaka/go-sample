package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eldaka/go-sample/api"
	config "github.com/eldaka/go-sample/config/xfers"
)

func main() {
	config.NewMainConfig()

	// load config and check config is ok
	debugConfig, err := json.MarshalIndent(config.Main, "", "   ")
	if err != nil {
		log.Fatal(fmt.Errorf("error marshalling config: %v", err))
	}
	log.Println(string(debugConfig))

	api.InitHTTPClientRequest(&config.Main)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	http.HandleFunc("/mockA", serveMockA)

	environ := os.Getenv("ENV")
	if environ == "" {
		environ = "development"
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
