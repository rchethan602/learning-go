package main

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your password:")

	//readpass, _ := reader.ReadString('\n')
	bytePass, _ := term.ReadPassword(int(syscall.Stdin))
	fmt.Println("THis is your Password:", string(bytePass))
}
