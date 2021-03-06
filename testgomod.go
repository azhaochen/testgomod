package testgomod

import "github.com/xdg/stringprep"
import "fmt"

func print() {
	fmt.Println("hello world")
	userprep, err := stringprep.SASLprep.Prepare(username)
	if err != nil {
		fmt.Println("Error SASLprepping username '%s': %v", username, err)
	}
	fmt.Println(userprep)
}
