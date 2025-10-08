package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	t, _ := time.ParseDuration("24h")
	fmt.Println(time.Now().Add(t).Unix())

	t, _ = time.ParseDuration("24h30m")
	fmt.Println(time.Now().Add(t).Unix())

	t, _ = time.ParseDuration("48h")
	fmt.Println(time.Now().Add(t).Unix())

	t, _ = time.ParseDuration("48h30m")
	fmt.Println(time.Now().Add(t).Unix())

	t, _ = time.ParseDuration("72h")
	fmt.Println(time.Now().Add(t).Unix())

	t, _ = time.ParseDuration("72h30m")
	fmt.Println(time.Now().Add(t).Unix())
}
