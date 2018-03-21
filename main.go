package main

import "flag"

var (
	// http port
	port = flag.String("p", "8081", "-p=8081")
	// upload full path
	storage = flag.String("s", "/opt/", "-s=/opt/")
	// relative path to web root
	url = flag.String("u", "tmp", "-u=tmp")
)

func main() {
	flag.Parse()
	a := App{}
	a.Initialize()
	a.Run(":" + *port)
}
