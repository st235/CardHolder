package server

func Register() {
	NewServer().
		AddRoute("/cards/get", GetCards).
		AddRoute("/card/get", GetCard).
		Listen(determineListenAddress())
}
