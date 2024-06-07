package cmd

import (
	"fmt"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
	// Random string for oauth2 API to protect against CSRF
	oauthStateString = "pseudo-random"
)

func init() {
	// Initialize Viper and read the configuration
	viper.SetConfigName("config/local.yaml") // name of config file (without extension)
	viper.SetConfigType("yaml")              // or viper.SetConfigType("YAML")
	viper.AddConfigPath(".")                 // optionally look for config in the working directory
	err := viper.ReadInConfig()              // Find and read the config file
	if err != nil {                          // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// Initialize OAuth2 configuration from Viper
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  viper.GetString("oauth.redirectURL"),
		ClientID:     viper.GetString("oauth.clientID"),
		ClientSecret: viper.GetString("oauth.clientSecret"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}
