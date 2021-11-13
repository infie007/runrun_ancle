package main

func main() {
	router := BuildRouter()

	router.Run(":80")
}
