package config

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (
	once        sync.Once
	hostName    string
	serviceName string
	serviceVer  string
	gitHash     string
	settings    AppSettings
)

// Initialize initializes the configuration.
func Initialize(cfg, version, hash string) (err error) {
	hn, _ := os.Hostname()
	once.Do(func() {
		initEnvVars()

		hostName = hn
		serviceVer = version
		gitHash = hash
		viper.SetConfigType("yaml")
		viper.SetConfigName(cfg)
		viper.AddConfigPath("/etc/rest_api")
		viper.AddConfigPath("./conf")
		viper.AddConfigPath("./")
		err = viper.ReadInConfig()
		err = viper.Unmarshal(&settings)
	})
	return
}

func initEnvVars() {
	// viper.SetEnvPrefix("RAPI")
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// viper.AutomaticEnv()
	// viper.BindEnv("monitoring.logging.buffersize")
	// viper.BindEnv("api.port")
	// viper.BindEnv("monitoring.providerkey")
	// viper.BindEnv("monitoring.provider")
	// viper.BindEnv("monitoring.tracing.intakeurl")
	// viper.BindEnv("monitoring.tracing.intakeport")
	// viper.BindEnv("monitoring.tracing.logsinjection")
	// viper.BindEnv("monitoring.tracing.analyticsenabled")
	// viper.BindEnv("monitoring.logging.intakeport")
	// viper.BindEnv("monitoring.logging.timeout")
	// viper.BindEnv("environment")
	// viper.BindEnv("healthcheck.port")
	// viper.BindEnv("healthcheck.path")
	// viper.BindEnv("monitoring.logging.intakeurl")
}

// ServiceHost ..
func ServiceHost() string {
	return hostName
}

// ServiceName ..
func ServiceName() string {
	return settings.Name
}

// ServiceVer ..
func ServiceVer() string {
	return serviceVer
}

// GitHash ..
func GitHash() string {
	return gitHash
}

// TracingIntakeURL ..
func TracingIntakeURL() string {
	return settings.Monitoring.Tracing.AgentHost
}

// LoggingIntakeURL ..
func LoggingIntakeURL() string {
	return settings.Monitoring.Logging.IntakeURL
}

// HealthcheckPath ..
func HealthcheckPath() string {
	return settings.Healthcheck.Path
}

// LoggingIntakePort ..
func LoggingIntakePort() int {
	return settings.Monitoring.Logging.IntakePort
}

// TracingIntakePort ..
func TracingIntakePort() int {
	return settings.Monitoring.Tracing.IntakePort
}

// MonitorAPIKey ..
func MonitorAPIKey() string {
	return settings.Monitoring.APIKey
}

// APIPort ..
func APIPort() int {
	return settings.API.Port
}

// HealthcheckPort ..
func HealthcheckPort() int {
	return settings.Healthcheck.Port
}

// LogBufferSize ..
func LogBufferSize() int {
	return settings.Monitoring.Logging.BufferSize
}

// LogTimeout ..
func LogTimeout() time.Duration {
	to := settings.Monitoring.Logging.Timeout
	return time.Duration(to) * time.Second
}

// Environment ..
func Environment() string {
	return settings.Environment
}

// ClientCredentials ..
func ClientCredentials(id string) *ClientCreds {
	for _, c := range settings.ClientAuth {
		if strings.ToLower(id) == strings.ToLower(c.ID) {
			return &c
		}
	}
	return nil
}

// Dependencies ..
func Dependencies() []Dependency {
	return settings.API.Dependencies
}

// OAuthProvider ..
func OAuthProvider() AuthProvider {
	return settings.API.AuthProvider
}

// String causes the current configuration to be printed to the terminal.
func String() string {
	b, _ := yaml.Marshal(settings)
	return string(b)
}
