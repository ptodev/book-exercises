package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}

type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

func main() {
	A := Record{"String value", -12.123, Secret{"Mihalis", "Tsoukalos"}}
	r := reflect.ValueOf(A)
	fmt.Println("String value:", r.String())
}

// String value: <main.Record Value>
// i Type: main.Record
// The 3 fields of main.Record are
// 	Field1  with type: string
// 	Field2  with type: float64
// 	Field3  with type: main.Secret  and value _{Mihalis Tsoukalos}_
// main.Secret
// main.Secret
