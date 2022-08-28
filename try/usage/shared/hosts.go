package shared

import ssa "github.com/Wulfheart/simple-server-automation"

var defaultSSHConfig = &ssa.SSHConfig{
	AuthType: ssa.Key,
	Location: KeyLocation,
}

var Web1 = &ssa.Server{
	Name:      "web-1",
	Ip:        "127.0.0.1",
	Port:      "123",
	SSHConfig: defaultSSHConfig,
}

var Web2 = &ssa.Server{
	Name:      "web-2",
	Ip:        "127.0.0.1",
	Port:      "1234",
	SSHConfig: defaultSSHConfig,
}

var _ = ssa.Group{Servers: []*ssa.Server{Web1, Web2}}
