package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput (prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions (b bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose an option (a - add item, s - save bill, t - add tip) ", reader)

	switch option {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Price: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		b.addItem(name, p)
		fmt.Println("Item added: ", name, p)
		promptOptions(b)
		// break

		// fmt.Println(name, price)
	case "t":
		tip, _ := getInput("Enter the tip amount ($): ", reader)
		
		t, err := strconv.ParseFloat(tip, 64)
		
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		
		b.updateTip(t)
		fmt.Println("Tip added: ", t)
		promptOptions(b)
		// break
		
	case "s":
		b.save()
		fmt.Println("You saved the file ", b.name)
		// break
	default:
		fmt.Println("That was not a valid option")
		promptOptions(b)
	}

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b

}

func main() {
	mybill := createBill()
	promptOptions(mybill)
	// fmt.Println(mybill)
}