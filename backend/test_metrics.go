package main

import (
	"fmt"
	"kubeintel/backend/internal/monitoring"
)

func main() {

	nodes, err := monitoring.GetNodeMetrics()

	if err != nil {
		panic(err)
	}

	fmt.Println("Nodes")

	for _, n := range nodes {
		fmt.Println(n)
	}

	pods, err := monitoring.GetPodMetrics()

	if err != nil {
		panic(err)
	}

	fmt.Println("Pods")

	for _, p := range pods {
		fmt.Println(p)
	}
}