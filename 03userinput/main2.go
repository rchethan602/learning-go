package alter

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Main1() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for our service:")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Rating you entered is %s", input)
}
