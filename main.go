// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anduckhmt146/google-sso/cmd"
	"github.com/spf13/viper"
)

func main() {
	http.HandleFunc("/health", cmd.HandleHealthCheck)
	http.HandleFunc("/login", cmd.HandleGoogleLogin)
	http.HandleFunc("/callback", cmd.HandleGoogleCallback)

	port := viper.GetString("service.port")

	fmt.Println("Started running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
