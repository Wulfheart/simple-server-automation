package main

import (
	"fmt"
	"github.com/helloyi/go-sshclient"
	"strings"
	"sync"
)

func connect(ip string) *sshclient.Client {
	fmt.Println("Started", ip)
	client, err := sshclient.DialWithKey(ip+":22", "root", "/home/alex/.ssh/id_rsa")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected", ip)

	return client
}

func execute(client *sshclient.Client) {
	defer client.Close()
	res, _ := client.Cmd("hostname").SmartOutput()
	fmt.Println(strings.TrimSpace(string(res)))
}

func main() {
	ips := []string{"116.203.148.41", "116.203.191.33"}
	var cons []*sshclient.Client
	var wg sync.WaitGroup

	for _, ip := range ips {
		wg.Add(1)
		ip := ip
		go func() {
			defer wg.Done()
			cons = append(cons, connect(ip))
		}()

	}
	wg.Wait()

	for _, con := range cons {
		wg.Add(1)
		con := con
		go func() {
			defer wg.Done()
			execute(con)
		}()
	}

	wg.Wait()
}
