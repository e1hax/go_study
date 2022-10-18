package main

func main() {
	server := NewServer("192.168.169.1", 8899)
	server.Start()
}
