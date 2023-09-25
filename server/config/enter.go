package config

type Configs struct {
	Server   Server   `json:"server" yaml:"server"`
	Postgres Postgres `json:"postgres" yaml:"postgres"`
}
