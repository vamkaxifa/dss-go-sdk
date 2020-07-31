package dss

// DsnProperties database connection parameters
type DsnProperties struct {
	// IP
	IP string `yaml:"ip"`
	// Port
	Port int `yaml:"port"`
	// Sid
	Sid string `yaml:"sid"`
	// UserName
	UserName string `yaml:"username"`
	// Password
	Password string `yaml:"password"`
}
