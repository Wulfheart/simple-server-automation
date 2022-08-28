package shared

import ssa "github.com/Wulfheart/simple-server-automation"

var Web1 = &ssa.Server{
	Name: "web-1",
	Ip:   "127.0.0.1",
	Port: "123",
	SSHConfig: &ssa.SSHConfig{
		AuthType: ssa.Key,
		Location: "",
	},
}

var _ = ssa.Group{Servers: []*ssa.Server{Web1}}
