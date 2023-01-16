package main

import "nats-publisher/streaming"

func main() {
	publisher := streaming.New()
	publisher.Init()

}
