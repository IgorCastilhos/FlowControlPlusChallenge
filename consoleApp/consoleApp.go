package consoleApp

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"strconv"
)

func ConsoleApplication() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappucino\r")
	fmt.Println("2 - Latte\r")
	fmt.Println("3 - Americano\r")
	fmt.Println("4 - Mocha\r")
	fmt.Println("5 - Macchiato\r")
	fmt.Println("6 - Espresso\r")
	fmt.Println("Q - Quit the program\r")

	char := ' '

	for char != 'q' && char != 'Q' {
		char, _, err = keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		i, _ := strconv.Atoi(string(char))
		fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))
		if char == 'q' || char == 'Q' {
			fmt.Println(fmt.Sprintf("You chose %c", char))
		}
	}
	fmt.Println("Program exiting...")
}
