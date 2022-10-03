package hello

import (
	"fmt"
)

func Hello() string {
	fmt.Printf("Hello %v", "friend")
	return "Hello friend"
}
