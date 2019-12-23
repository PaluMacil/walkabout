package main

func main() {
	server := StartServer("12345")
	server.HandleSessionConnections()
	server.WaitForInterrupt()
}
