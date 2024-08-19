package main

import (
	"github.com/alexferl/pk8s"
	"github.com/alexferl/pk8s/cmd"
)

func main() {
	pk8s.InitLog()
	cmd.Execute()
}
