package main

import (
	"context"
	"fmt"
)

func Context() {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}
