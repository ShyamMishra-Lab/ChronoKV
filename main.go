package main

import (
	"fmt"
	"github.com/chronokv/ChronoKV/store"
)

func main() {
	// Right now just to test other user made packages
	fmt.Println(store.Show(1))
}