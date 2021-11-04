package main

import (
	"fmt"

	"github.com/digital-polyglot/intervalsets/domain"
)

func main() {
	interval := domain.NewInterval(0, 1)

	fmt.Printf("interval: %v", interval)
}
