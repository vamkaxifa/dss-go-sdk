package dss

// MiscProperties database connection parameters
type MiscProperties struct {
	// IP
	IP string `yaml:"ip"`
	// Port
	Port int `yaml:"port"`
	// Password
	Password string `yaml:"password"`
}