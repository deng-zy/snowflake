package main

import (
	"fmt"
	"snowflake"
	"time"
)

func main() {
	t, err := time.Parse(time.RFC3339, "2021-11-03T14:14:14+08:00")
	if err != nil {
		fmt.Println("time parse error:", err.Error())
		return
	}

	n, err := snowflake.NewNode(1, 1, t)
	if err != nil {
		fmt.Println("NewCode execute fail. error:", err.Error())
		return
	}

	var last, curr snowflake.ID

	for i := 0; i < 100000; i++ {
		curr = n.Generate()
		if curr == last {
			fmt.Println("x(%d) & y(%d) are the same", curr, last)
			break
		} else {
			fmt.Println(curr)
		}
	}
}
