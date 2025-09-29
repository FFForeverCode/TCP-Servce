package main

import "github.com/FFForeverCode/ginx/gnet"

func main() {
	server := gnet.NewServer("myServer")
	server.Serve()
}
