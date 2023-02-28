package main

import (
	"context"
	"fmt"
	"github.com/mel3kings/event-driven-architecture/temporal_workflow"
	"go.temporal.io/sdk/client"
	"log"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: temporal_workflow.GreetingTaskQueue,
	}

	name := "World"
	we, err := c.ExecuteWorkflow(context.Background(), options, temporal_workflow.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}

	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}
	fmt.Printf("Workflow ID: %s, Run ID: %s : %s\n", we.GetID(), we.GetRunID(), greeting)
}
