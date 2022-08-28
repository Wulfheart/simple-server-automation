package ssa

const (
	Password SSHAuthType = "password"
	Key      SSHAuthType = "key"
)

type SSHAuthType string

type Executable interface {
}

type SSHConfig struct {
	AuthType SSHAuthType
	Location string
	Password string
}

type Server struct {
	Name      string
	Ip        string
	Port      string
	SSHConfig *SSHConfig
}

type Group struct {
	Name    string
	Servers []*Server
}
