package main

import "flag"

var (
	ConfPath  string
	configure Configure
)

func main() {
	flag.StringVar(&ConfPath, "conf", "configure.yaml", "-conf to specify configure file.")
	var err error
	configure, err = LoadConfigure(ConfPath)
	if err != nil {
		InitialConfigure(ConfPath)
	}
}
