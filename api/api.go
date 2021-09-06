package api


func Start(port string) {
	
	r := InitRoute()
	
	server := newServer(port, r)

	server.Start()
}