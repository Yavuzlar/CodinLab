package config

import (
	"time"

	"github.com/spf13/viper"
)

/*
There are functions that convert the settings and environment variables
found in the configuration files according to the appropriate working environment.
*/
const (
	defaultConfigDir              = "./config"
	defaultHTTPPort               = "8081"
	defaultHTTPRWTimeout          = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1
	defaultSessionExpiration      = 24 * time.Hour
	defaultManagmentPath          = "/management"
	defaultAppMode                = "self"
)

var APP_VERSION string = "0.0.1"

type Managment struct {
	ManagmentUsername string `mapstructure:"username"`
	ManagmentPassword string `mapstructure:"password"`
}
type Application struct {
	DevMode        bool      `mapstructure:"devMode"`
	Mode           string    `mapstructure:"mode"`
	Managment      Managment `mapstructure:"managment"`
	Version        string
	MigrationsPath string `mapstructure:"migrationsPath"`
}
type Config struct {
	HTTP        HTTPConfig  `mapstructure:"http"`
	Application Application `mapstructure:"app"`
}

type HTTPConfig struct {
	Host               string        `mapstructure:"host"`
	Port               string        `mapstructure:"port"`
	ReadTimeout        time.Duration `mapstructure:"readTimeout"`
	WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
	MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	SessionExpiration  time.Duration `mapstructure:"sessionExpiration"`
	AllowedOrigins     []string      `mapstructure:"allowedOrigins"`
	AllowedHeaders     []string      `mapstructure:"allowedHeaders"`
	AllowedMethods     []string      `mapstructure:"allowedMethods"`
	ExposedHeaders     []string      `mapstructure:"exposedHeaders"`
	AllowCredentials   bool          `mapstructure:"allowCredentials"`
	ProxyHeader        string        `mapstructure:"proxyHeader"`
}

func Init(configsDir ...string) (cfg *Config, err error) {
	cfg = new(Config)
	viper.SetDefault("app.version", APP_VERSION)
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.max_header_megabytes", defaultHTTPMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", defaultHTTPRWTimeout)
	viper.SetDefault("http.timeouts.write", defaultHTTPRWTimeout)
	viper.SetDefault("http.session_expiration", defaultSessionExpiration)
	viper.SetDefault("managment.managmentPath", defaultManagmentPath)
	viper.SetDefault("mode", defaultAppMode)

	dir := ""
	if len(configsDir) > 0 {
		dir = configsDir[0]
	} else {
		dir = defaultConfigDir
	}
	viper.AddConfigPath(dir)
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	if err = viper.MergeInConfig(); err != nil {
		return
	}
	if err = viper.Unmarshal(&cfg); err != nil {
		return
	}
	return
}
