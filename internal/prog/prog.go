package main

import (
	"fmt"
	"os"
)

//noinspection ALL
const (
	A_ENV = "A_ENV"
	B_VAR = "B_VAR"
)


func main()  {
	fmt.Println("Test run for go-envdir.")
	fmt.Println("Prog should get A_ENV and B_VAR from \"./env\"")
	fmt.Println()
	fmt.Println("Result:")

	aEnv, ok := os.LookupEnv(A_ENV)
	printEnv(A_ENV, aEnv, ok)

	bVar, ok := os.LookupEnv(B_VAR)
	printEnv(B_VAR, bVar, ok)
}

func printEnv(envName, envValue string, ok bool)  {
	if ok {
		fmt.Println("Received: ", envName, "=", envValue)
	} else {
		fmt.Println(envName, "not received")
	}

}