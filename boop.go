package main

import "github.com/spenserblack/boop/pkg/boop"

func main() {
	boop := boop.New()
	boop.Executable(true)
	if err := boop.Boop("deeply/nested/file"); err != nil {
		panic(err)
	}
}
