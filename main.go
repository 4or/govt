package main

import (
	"fmt"

	"github.com/ahmedbejaouiJS/govt/dynamo"
)

// api string, resource string
func main() {
	resp := dynamo.Dbn()
	fmt.Println(resp)
}

// The Main function will host all Functions that filter, save to dynamodb, alert us
