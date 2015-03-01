package main

type empty struct{}

func main() {
	var ping empty
	channel := make(chan empty)
	//...
	channel <- ping
	//...
}
