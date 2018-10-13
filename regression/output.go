package regression

import "fmt"

func output(mode string, cmd string, file string) {
	fmt.Println("------------------------------------------------------------------")
	fmt.Println(fmt.Sprintf("Liner regression -> %s", mode))
	fmt.Println("------------------------------------------------------------------")
	fmt.Println(fmt.Sprintf("SUCCESS: %s in %s", cmd, file))
	fmt.Println("------------------------------------------------------------------")
}
