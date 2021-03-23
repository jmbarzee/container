package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	bbgrpc "github.com/jmbarzee/bitbox/grpc"
)

var cmdStart = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long:  "Start a process on the bitbox server",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			panic(errors.New("Require atleast a single command as an argument"))
		}

		job := jobStart{
			command:    args[0],
			parameters: args[0:],
		}
		ctx := context.Background()
		bbClient := getClient(ctx)
		if err := job.execute(ctx, bbClient); err != nil {
			panic(err)
		}
	},
}

type jobStart struct {
	command    string
	parameters []string
}

// Execute starts a job on the remote BibBox
func (j jobStart) execute(ctx context.Context, c bbgrpc.BitBoxClient) error {
	request := &bbgrpc.StartRequest{
		Command:    j.command,
		Parameters: j.parameters,
	}

	reply, err := c.Start(ctx, request)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to run %s: %w", j.command, err))
	}
	log.Println("Successfully started process: ", reply.GetID())
	return nil
}