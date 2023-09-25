package config

type Postgres struct {
	Host     string `json:"host,omitempty" yaml:"host"`
	Port     int    `json:"port,omitempty" yaml:"port"`
	Username string `json:"username,omitempty" yaml:"username"`
	Password string `json:"password,omitempty" yaml:"password"`
	DbName   string `json:"db-name,omitempty" yaml:"db-name"`
}
