package entity

// Config is the servers config
type Config struct {
	// AppID is the applicationID
	AppID string `yaml:"appid"`

	// Env Specifies what env we are in, this is obtained from an env variable
	Env string

	// ServerPort specifies the server port
	ServerPort int `yaml:"serverPort"`
}

// Secrets contains the secrets for the server
type Secrets struct {
	// PostgresInfo is all of the postgresql information
	PostgresInfo PostgresInfo `yaml:"postgres"`
}

// PostgresInfo ... temp
type PostgresInfo struct {
	// Name is the name for the postgresql server
	Name string `yaml:"name"`

	// Host is the host for the postgresql server
	Host string `yaml:"host"`

	// Port is the port number for your postgresql server
	Port int `yaml:"port"`

	// User is the user for your postgresql server
	User string `yaml:"user"`

	// Password is the password for your postgresql server
	Password string `yaml:"password"`
}
