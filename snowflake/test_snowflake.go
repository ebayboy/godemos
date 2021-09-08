package main

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

func main() {

	//+--------------------------------------------------------------------------+
	//| 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
	//+--------------------------------------------------------------------------+

	// Create a new Node with a Node number of 1
	sfnode, err := snowflake.NewNode(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	var steps []int64

	for i := 0; i < 15; i++ {
		// Generate a snowflake ID.
		id := sfnode.Generate()

		// Print out the ID in a few different ways.
		fmt.Printf("Int64  ID: %d\n", id)
		fmt.Printf("String ID: %s\n", id)
		fmt.Printf("Base2  ID: %s len:%d\n", id.Base2(), len(id.Base2()))
		fmt.Printf("Base64 ID: %s\n", id.Base64())

		// Print out the ID's timestamp
		fmt.Printf("ID Time  : %d\n", id.Time())

		// Print out the ID's sfnode number
		fmt.Printf("ID Node  : %d\n", id.Node())

		// Print out the ID's sequence number
		fmt.Printf("ID Step  : %d\n", id.Step())
		steps = append(steps, id.Step())

		// Generate and print, all in one.
		fmt.Printf("ID       : %d\n\n", sfnode.Generate().Int64())
	}

	fmt.Println("steps:", steps)
}
