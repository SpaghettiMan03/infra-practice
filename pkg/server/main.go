package main

func main() {
	server, err := initialize()
	if err != nil {
		panic(err)
	}

	err = server.Run("3000")
	if err != nil {
		panic(err)
	}
}
