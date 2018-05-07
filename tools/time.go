package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.AddDate(0, 0, 2).UTC())
}
