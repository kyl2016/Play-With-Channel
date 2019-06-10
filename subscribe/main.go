package main

func main() {
	server := Server{}

	_ = New("sub1", &server)
	_ = New("sub2", &server)

	server.Start()
}
