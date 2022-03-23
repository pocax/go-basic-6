package main

import "ginproject/routers"

func main() {
	var PORT = ":8181"

	routers.StartServer().Run(PORT)
}
