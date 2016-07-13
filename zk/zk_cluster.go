package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"os"
)

func main() {
	cluster, err := zk.StartTestCluster(3, os.Stdout, os.Stderr)
	if err != nil {
		panic(err)
	}
	defer cluster.Stop()

	fmt.Printf("Listening at 127.0.0.1:%d\n", cluster.Servers[0].Port)

	select {}
}
