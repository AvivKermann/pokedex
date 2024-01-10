package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Printf("pokedex> ")
		usrInput := bufio.NewScanner(os.Stdin)
		usrInput.Scan()
		//fmt.Println(usrInput.Text())
		if usrInput.Text() == "exit" {
			break
		}
		if usrInput.Text() == "help" {
			continue
		}

	}
}
