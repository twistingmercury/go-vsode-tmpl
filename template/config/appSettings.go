package config

import (
	"fmt"
)

// AppSettings ..
type AppSettings struct {
	Environment string              `yaml:"environment"`
	ClientAuth  []ClientCreds       `yaml:"clientauth"`
	API         APISettings         `yaml:"api"`
	Healthcheck HealthcheckSettings `yaml:"healthcheck"`
	Monitoring  MonitorSettings     `yaml:"monitoring"`
	Name        string              `yaml:"name"`
}

// ClientCreds ..
type ClientCreds struct {
	ID        string `yaml:"id"`
	Secret    string `yaml:"secret"`
	GrantType string `yaml:"granttype"`
	Scope     string `yaml:"scope"`
}

// APISettings ..
type APISettings struct {
	Port         int          `yaml:"port"`
	Dependencies []Dependency `yaml:"dependencies"`
	AuthProvider AuthProvider `yaml:"authprovider"`
}

// AuthProvider ..
type AuthProvider struct {
	Provider       string `yaml:"provider"`
	ID             string `yaml:"id"`
	Secret         string `yaml:"secret"`
	Scope          string `yaml:"scope"`
	BaseURL        string `yaml:"baseurl"`
	TokenPath      string `yaml:"tokenpath"`
	IntrospectPath string `yaml:"introspectpath"`
	Path           string `yaml:"path"`
}

// IntrospectAddr ..
func (ap *AuthProvider) IntrospectAddr() (a string) {
	return fmt.Sprintf("%s/%s", ap.BaseURL, ap.Path)
}

// Dependency is another service or resource this API depends upon.
type Dependency struct {
	Name       string `yaml:"name"`
	Connection string `yaml:"connection"`
	Type       string `yaml:"type"`
}

// HealthcheckSettings ..
type HealthcheckSettings struct {
	Path string `yaml:"path"`
	Port int    `yaml:"port"`
}

// MonitorSettings ..
type MonitorSettings struct {
	Provider string        `yaml:"provider"`
	APIKey   string        `yaml:"apikey"`
	Tracing  TraceSettings `yaml:"tracing"`
	Logging  LogSettings   `yaml:"logging"`
}

// TraceSettings ..
type TraceSettings struct {
	AgentHost        string `yaml:"agentHost"`
	IntakePort       int    `yaml:"intakePort"`
	LogsInjection    bool   `yaml:"injectLogs"`
	AnalyticsEnabled bool   `yaml:"enableAnalytics"`
}

// LogSettings ..
type LogSettings struct {
	IntakeURL  string `yaml:"intakeURL"`
	IntakePort int    `yaml:"intakePort"`
	BufferSize int    `yaml:"bufferSize"`
	Timeout    int    `yaml:"timeout"`
}
