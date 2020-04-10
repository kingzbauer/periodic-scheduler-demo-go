package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/robfig/cron/v3"
)

func chk(msg string, err error) {
	if err != nil {
		log.Panicf("Error: %s : %s", msg, err)
	}
}

func machineryServer() (*machinery.Server, error) {
	machineryConf := &config.Config{
		Broker:        os.Getenv("BROKER"),
		DefaultQueue:  "machinery",
		ResultBackend: os.Getenv("BROKER"),
		AMQP: &config.AMQPConfig{
			Exchange:     "machinery",
			ExchangeType: "direct",
			BindingKey:   "machinery",
		},
	}

	return machinery.NewServer(machineryConf)
}

func main() {
	fmt.Println("Waiting for rabbit to warm up...")
	// wait for a minute or two to give rabbitmq time to start up
	time.Sleep(time.Minute * 1)
	c := cron.New(cron.WithSeconds())
	srv, err := machineryServer()
	chk("Initializing server", err)

	log.Println("All wamrmed up now")

	// Runs every second
	count := 1
	c.AddFunc("*/1 * * * * *", func() {
		// create task signature
		sig, err := tasks.NewSignature(
			"echo",
			[]tasks.Arg{
				{Type: "string", Value: fmt.Sprintf("Message: %d", count)},
			})
		chk("Create signature", err)
		count++

		// send the task
		_, err = srv.SendTask(sig)
		chk("Send task", err)
	})

	c.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	// Waiting for interrupt signal
	<-sig
	// stop scheduler
	c.Stop()
	fmt.Println("Cron has closed successfully")
}
