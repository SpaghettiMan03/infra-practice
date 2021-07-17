package main

func main() {
	server := NewServer()

	err := server.Run("3000")
	if err != nil {
		panic(err)
	}
}
