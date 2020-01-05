package init_test

import "fmt"

func init() {
	fmt.Println("Init in sub ")
}

func main() {
	fmt.Println("MAIN IN SUB")
}
