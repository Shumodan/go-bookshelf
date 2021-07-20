package config

import (
	"flag"
	"os"
)

// Config represents the composition of yml settings.
type Config struct {
	Database struct {
		Dialect   string `default:"sqlite3"`
		Host      string `default:"book.db"`
		Port      string
		Dbname    string
		Username  string
		Password  string
		Migration bool `default:"true"`
	}
	Redis struct {
		Enabled            bool `default:"false"`
		ConnectionPoolSize int  `yaml:"connection_pool_size" default:"10"`
		Host               string
		Port               string
	}
	Extension struct {
		MasterGenerator bool `yaml:"master_generator" default:"true"`
		CorsEnabled     bool `yaml:"cors_enabled" default:"true"`
		SecurityEnabled bool `yaml:"security_enabled" default:"true"`
	}
	Log struct {
		RequestLogFormat string `yaml:"request_log_format" default:"${remote_ip} ${account_name} ${uri} ${method} ${status}"`
	}
	StaticContents struct {
		Path string `yaml:"path" default:"public"`
	}
	Security struct {
		AuthPath    []string `yaml:"auth_path"`
		ExcludePath []string `yaml:"exclude_path"`
		UserPath    []string `yaml:"user_path"`
		AdminPath   []string `yaml:"admin_path"`
	}
}

const (
	// DEV represents development environment
	DEV = "develop"
	// PRD represents production environment
	PRD = "production"
	// DOC represents docker container
	DOC = "docker"
)

// Load reads the settings written to the yml file
func Load() (*Config, string) {

	var env *string
	if value := os.Getenv("WEB_APP_ENV"); value != "" {
		env = &value
	} else {
		env = flag.String("env", "develop", "To switch configurations.")
		flag.Parse()
	}

	config := &Config{
		Database: struct {
			Dialect   string `default:"sqlite3"`
			Host      string `default:"book.db"`
			Port      string
			Dbname    string
			Username  string
			Password  string
			Migration bool `default:"true"`
		}{Dialect: "sqlite3", Host: "book.db", Port: "", Dbname: "", Username: "", Password: "", Migration: true},
		Redis: struct {
			Enabled            bool `default:"false"`
			ConnectionPoolSize int  `yaml:"connection_pool_size" default:"10"`
			Host               string
			Port               string
		}{Enabled: false, ConnectionPoolSize: 0, Host: "", Port: ""},
		Extension: struct {
			MasterGenerator bool `yaml:"master_generator" default:"true"`
			CorsEnabled     bool `yaml:"cors_enabled" default:"true"`
			SecurityEnabled bool `yaml:"security_enabled" default:"true"`
		}{MasterGenerator: true, CorsEnabled: true, SecurityEnabled: true},
		Log: struct {
			RequestLogFormat string `yaml:"request_log_format" default:"${remote_ip} ${account_name} ${uri} ${method} ${status}"`
		}{RequestLogFormat: "${remote_ip} ${account_name} ${uri} ${method} ${status}"},
		StaticContents: struct {
			Path string `yaml:"path" default:"public"`
		}{Path: "public"},
		Security: struct {
			AuthPath    []string `yaml:"auth_path"`
			ExcludePath []string `yaml:"exclude_path"`
			UserPath    []string `yaml:"user_path"`
			AdminPath   []string `yaml:"admin_path"`
		}{
			AuthPath:    []string{"/api/.*"},
			ExcludePath: []string{"/api/auth/login$", "/api/auth/logout$", "/api/health$"},
			UserPath:    []string{"/api/.*"},
			AdminPath:   []string{"/api/.*"},
		},
	}

	return config, *env
}
