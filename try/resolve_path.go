package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := "$HOME/.ssh/id_rsa"
	res, _ := filepath.Abs(path)
	fmt.Println(res)
}
