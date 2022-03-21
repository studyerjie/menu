package main

import "fmt"

func main() {
	var cmd string
	for {
		fmt.Print("please input your command\n")
		fmt.Scanln(&cmd)
		if cmd == "quit" {
			break
		} else if cmd == "help" {
			fmt.Println("this is help command")
		} else {
			fmt.Println("wrong command")
		}
	}
}
