package config

import (
	// "fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Constants - encapsulates the server configs environmental variables
type Constants struct {
	Server struct {
		URL    string
		Port   string
		AppKey string
	}
	Client struct {
		URL string
	}
	Auth struct {
		UserTokenSecret        string
		AccountUserTokenSecret string
		AdminTokenSecret       string
	}
	Database struct {
		Host    string
		Port    string
		Name    string
		User    string
		Pass    string
		Charset string
	}
	Pexportal struct {
		BaseURL          string
		ProductSearchURL string
		FlightSearchURL  string
	}
	SendGrid struct {
		ApiKey string
	}
	Paypal struct {
		URL       string
		AccessKey string
	}
}

// Config - Encapsulate the application configuration
type Config struct {
	Constants
}

// New NewConfig is used to generate a configuration instance which will be passed around the codebase
func New() (*Config, error) {
	config := Config{}
	constants, err := viperConfigInit() // for local dev
	// constants, err := viperStagingEnvInit() // for staging
	// constants, err := viperEnvInit() // for live production server
	config.Constants = constants
	if err != nil {
		return &config, err
	}
	return &config, err
}

// For local dev configuration
func viperConfigInit() (Constants, error) {
	viper.SetConfigName("app.config") // Configuration fileName without the .TOML or .YAML extension
	viper.AddConfigPath(".")          // Search the root directory for the configuration file
	err := viper.ReadInConfig()       // Find and read the config file
	if err != nil {                   // Handle errors reading the config file
		return Constants{}, err
	}
	viper.WatchConfig() // Watch for changes to the configuration file and recompile
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})
	viper.SetDefault("Port", "8080")
	if err = viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading config file, %s", err)
	}

	var constants Constants
	err = viper.Unmarshal(&constants)
	return constants, err
}

// For local dev configuration
func viperStagingEnvInit() (Constants, error) {
	viper.SetEnvPrefix("pr") // will be uppercased automatically so use PX in eenv settup
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)

	var constants Constants

	// main server configs
	constants.Server.URL = viper.GetString("staging_server_url")
	constants.Server.Port = viper.GetString("staging_server_port")
	constants.Server.AppKey = viper.GetString("staging_server_app_key")

	// Client/Front-end configs
	constants.Client.URL = viper.GetString("staging_client_url")

	// JWT Token Authentication configs
	constants.Auth.UserTokenSecret = viper.GetString("staging_auth_customer_token_secret")
	constants.Auth.AccountUserTokenSecret = viper.GetString("staging_auth_account_user_token_secret")
	constants.Auth.AdminTokenSecret = viper.GetString("staging_auth_admin_token_secret")
	// Database Configs
	constants.Database.Host = viper.GetString("staging_database_host")
	constants.Database.Port = viper.GetString("staging_database_port")
	constants.Database.Name = viper.GetString("staging_database_name")
	constants.Database.User = viper.GetString("staging_database_user")
	constants.Database.Pass = viper.GetString("staging_database_pass")
	constants.Database.Charset = viper.GetString("staging_database_charset")

	// Send Grid Configs
	constants.SendGrid.ApiKey = viper.GetString("staging_send_grid_api_key")

	return constants, nil
}

// For production configuration
func viperEnvInit() (Constants, error) {
	viper.SetEnvPrefix("px") // will be uppercased automatically so use PX in eenv settup
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)

	var constants Constants

	// main server configs
	constants.Server.URL = viper.GetString("server_url")
	constants.Server.Port = viper.GetString("server_port")
	constants.Server.AppKey = viper.GetString("server_app_key")

	// Client/Front-end configs
	constants.Client.URL = viper.GetString("client_url")

	// JWT Token Authentication configs
	constants.Auth.UserTokenSecret = viper.GetString("auth_customer_token_secret")
	constants.Auth.AccountUserTokenSecret = viper.GetString("auth_account_user_token_secret")
	constants.Auth.AdminTokenSecret = viper.GetString("auth_admin_token_secret")
	// Database Configs
	constants.Database.Host = viper.GetString("database_host")
	constants.Database.Port = viper.GetString("database_port")
	constants.Database.Name = viper.GetString("database_name")
	constants.Database.User = viper.GetString("database_user")
	constants.Database.Pass = viper.GetString("database_pass")
	constants.Database.Charset = viper.GetString("database_charset")

	// Send Grid Configs
	constants.SendGrid.ApiKey = viper.GetString("send_grid_api_key")

	return constants, nil
}
