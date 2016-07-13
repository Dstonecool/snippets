package main

import (
	"flag"
	"github.com/samuel/go-zookeeper/zk"
	"strings"
	"time"
)

func main() {
	path := ""
	children := false

	flag.StringVar(&path, "path", "/", "zookeeper node path")
	flag.BoolVar(&children, "children", false, "monitor child nodes")
	flag.Parse()

	conn, _, err := zk.Connect(flag.Args(), time.Second)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	if children {
		for {
			children, _, ch, err := conn.ChildrenW(path)
			if err != nil {
				panic(err)
			}

			println("children: [" + strings.Join(children, ", ") + "]")

			e := <-ch
			println("event: " + e.Type.String())
		}
	} else {
		for {
			data, _, ch, err := conn.GetW(path)
			if err != nil {
				panic(err)
			}

			println("value: " + string(data))

			e := <-ch
			println("event: " + e.Type.String())
		}
	}
}
