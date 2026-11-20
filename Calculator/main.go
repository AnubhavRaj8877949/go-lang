package main

import (
	"errors"
	"fmt"
)

func add(num1 int, num2 int) int {

	return num1 + num2
}

func subtract(num1, num2 int) int {

	return num1 - num2
}

func multiply(num1, num2 int) int {

	return num1 * num2
}

func divide(num1, num2 int) (float64, error) {

	if num2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return float64(num1) / float64(num2), nil
}

func main() {

	for {

		var num1 int
		var num2 int
		var choice int

		fmt.Print("Enter  num1 : ")
		if _, err := fmt.Scanln(&num1); err != nil {
			fmt.Println("Invalid input! Kripya sirf number enter karein.")
			return
		}

		fmt.Print("Enter num2: ")
		if _, err := fmt.Scanln(&num2); err != nil {
			fmt.Println("Invalid input! Kripya sirf number enter karein.")
			return
		}

		// Menu options display karein
		fmt.Println("\nChoose an operation:")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Print("Enter your choice (1-4): ")
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println("Invalid input! Kripya sirf number enter karein.")
			return
		}

		switch choice {
		case 1:
			ans := add(num1, num2)
			fmt.Println("Add: ", ans)
		case 2:
			ans := subtract(num1, num2)
			fmt.Println("sub: ", ans)
		case 3:
			ans := multiply(num1, num2)
			fmt.Println("mul: ", ans)
		case 4:
			ans, err := divide(num1, num2)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Divide:", ans)
			}

		default:
			fmt.Println("Invalid option")

		}
		var cont int

		fmt.Println("Do you want  to  continue ? ")
		fmt.Println("1.  Yes")
		fmt.Println("2.  No")
		fmt.Print("Enter your choice (1-2): ")

		if _, err := fmt.Scanln(&cont); err != nil {
			fmt.Println("Invalid input! Kripya sirf number enter karein.")
			return
		}

		switch cont {
		case 1:
			continue
		case 2:
			fmt.Println("Calculator closed.")
			return

		default:
			fmt.Println("Invalid choice. Calculator will continue.")
		}

	}

}
