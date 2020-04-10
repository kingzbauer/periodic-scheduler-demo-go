# Simple Periodic task scheduler demo with golang

## A lightweight attempt to mimic celery beat

### Requirements
> - github.com/RichardKnop/machinery/v1
> - github.com/robfig/cron/v3
> - docker

### How to run

`docker-compose up -d --build`

The docker-compose.yml file contains 4 services
- rabbit => rabbitmq service
- worker1 => Consumes tasks from the job queue
- worker2 => Consumes tasks from the job queue
- beat => Triggers periodic tasks, pushing them to the job queue

You can open 3 separate tabs and log the output of the different services

> - `docker-compose logs -f worker1`
> - `docker-compose logs -f worker2`
> - `docker-compose logs -f beat`

__The code presented here is by no means good code, just a simple demo of how you'd setup a job scheduler with support for periodic tasks__
