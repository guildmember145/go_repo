package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", "GET", server.AddMiddleware(HandleRoot, Logging(), CheckAuth()))
	server.Handle("/home", "GET", HandleHome)
	server.Handle("/create", "POST", PostRequest)
	server.Handle("/user", "POST", server.AddMiddleware(CreateRequest, CheckCreateRequest()))
	server.Listen()
}