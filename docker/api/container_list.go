package api

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ContainerList() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	var all_containers bool

	fmt.Scanf("%t", &all_containers)

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: all_containers})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		// fmt.Println("Name:", container.Names[0],
		// 	"Image:", container.Image,
		// 	"Status:", container.Status)
		fmt.Println(container)
	}
}
