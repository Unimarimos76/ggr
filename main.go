package main

import (
	"flag"
	"os/exec"
	"strings"
)

func main() {
	flag.Parse()
	var str string
	var pri = flag.Bool("p", true, "private")

	args := flag.Args()
	str = strings.Join(args, " ")
	err := exec.Command("open", "https://www.google.com/search?q="+str).Start()
	if err != nil {
		panic(err)
	}
}
