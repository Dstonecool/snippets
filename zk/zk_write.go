package main

import (
	"flag"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func main() {
	path := ""
	data := ""

	flag.StringVar(&path, "path", "", "zookeeper node path")
	flag.StringVar(&data, "data", "", "zookeeper node data")
	flag.Parse()

	fmt.Printf("connecting to %v\n", flag.Args())

	conn, _, err := zk.Connect(flag.Args(), time.Second)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	exists, _, err := conn.Exists(path)
	if err != nil {
		panic(err)
	}

	if !exists {
		fmt.Printf("creating %s\n", path)

		_, err := conn.Create(path, []byte(data), int32(0), zk.WorldACL(zk.PermAll))
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("reading %s\n", path)

		_, s, err := conn.Get(path)
		if err != nil {
			panic(err)
		}

		fmt.Printf("version %d\n", s.Version)

		fmt.Printf("writing %s\n", path)

		_, err = conn.Set(path, []byte(data), s.Version)
		if err != nil {
			panic(err)
		}
	}
}
