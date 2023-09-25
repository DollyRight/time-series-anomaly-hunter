package config

import "fmt"

type Server struct {
	Host       string `json:"host,omitempty" yaml:"Host"`
	Port       int    `json:"port,omitempty" yaml:"port"`
	ApiPrefix  string `json:"apiPrefix,omitempty" yaml:"api-prefix"`
	ApiVersion string `json:"apiVersion,omitempty" yaml:"api-version"`
	Mode       string `json:"mode" yaml:"mode"`
}

func (s *Server) GetAddr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
