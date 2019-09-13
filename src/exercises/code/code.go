package code

import (
	"bufio"
	"fmt"
	"os"
)

func init() {
	fmt.Println(`The "code" package was started!`)
}

// Do not forget: functions have to be capitalized
func Exercise01() {
	// short declaration
	x := 42
	y := `James Bond`
	z := true

	fmt.Println("x =", x)
	fmt.Printf("Type of x= %T\n", x)
	fmt.Println("-----------------------")

	fmt.Println(`y =`, y)
	fmt.Printf("Type of y = %T \n", y)
	fmt.Println("-----------------------")

	fmt.Println("z =", z)
	fmt.Printf("Type of z = %T\n", z)
	fmt.Println("-----------------------")
	fmt.Println("END OF EXERCISE 1...")
}


func Exercise02() (int, string, bool) {
	var x int
	var y string
	var z bool

	// Zero values
	fmt.Println("x =", x)
	fmt.Printf("Type of x = %T \n", x)
	fmt.Println("---------------------")

	fmt.Println("y =", y)
	fmt.Printf("Type of y = %T \n", y)
	fmt.Println("---------------------")

	fmt.Println("z =", z)
	fmt.Printf("Type of z = %T \n", z)
	fmt.Println("---------------------")
	fmt.Println("END OF EXERCISE 2...")

	return x, y, z
}

func Exercise03() {
	x, y, z := Exercise02()

	x = 42
	y = `James Bond`
	z = true

	result := fmt.Sprintf(y, x, z)
	fmt.Println(result)
}

func ReadInput() (string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite seu nome: ")
	text, err := reader.ReadString('\n')
	if (err != nil) {
		fmt.Println(err)
	}
	return text

}






























