package api

func Start(port, port2 string) {

	r := InitRoute()

	server := newServer(port, r)
	server2 := newServer(port, r)

	server.Start()
	server2.Start()
}
