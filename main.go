package main

func main() {
	store := NewStore()
	api := NewAPI(store)
	StartServer(api)
}