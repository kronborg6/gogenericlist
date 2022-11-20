package main

import (
	"fmt"

	"github.com/kronborg6/gogenericlist/genericlist"
)

func main() {
	glist := genericlist.New[string]()

	glist.Insert("bob") // 0
	glist.Insert("foo") // 1
	glist.Insert("bar") // 2
	glist.Insert("lol") // 3
	glist.Insert("why") // 4

	/* 	glist.Remove(3)
	   	glist.Remove(1) */
	glist.RemoveByValue("bar")

	fmt.Printf("%+v\n", glist)
	// fmt.Println(glist)
}
