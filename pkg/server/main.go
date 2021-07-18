package main

func main() {
	server, err := initialize()
	if err != nil {
		panic(err)
	}

	err = server.Run("8080")
	if err != nil {
		panic(err)
	}
}
