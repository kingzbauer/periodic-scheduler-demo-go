package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
)

func chk(msg string, err error) {
	if err != nil {
		log.Panicf("Error: %s : %s", msg, err)
	}
}

func echo(msg string) error {
	log.Printf("Echo from: %s\n", getWorkerName())
	log.Printf("Message: %s\n", msg)
	return nil
}

func main() {
	fmt.Println("Waiting for rabbit to warm up...")
	time.Sleep(time.Minute * 1)
	conf := &config.Config{
		Broker:        os.Getenv("BROKER"),
		DefaultQueue:  "machinery",
		ResultBackend: os.Getenv("BROKER"),
		AMQP: &config.AMQPConfig{
			Exchange:     "machinery",
			ExchangeType: "direct",
			BindingKey:   "machinery",
		},
	}

	srv, err := machinery.NewServer(conf)
	chk("Initializing server", err)

	log.Println("All warmed up now")

	// register task
	srv.RegisterTask("echo", echo)

	worker := srv.NewWorker(getWorkerName(), 5)
	chk("Launching worker", worker.Launch())
	log.Println("Worker left")
}

func getWorkerName() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return "worker"
}
