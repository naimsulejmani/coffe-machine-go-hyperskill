package main

import (
	"fmt"
	"strings"
)

var (
	water       int = 400
	milk        int = 540
	coffeeBeans int = 120
	cups        int = 9
	money       int = 550
)

func printState() {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", water)
	fmt.Printf("%d ml of milk\n", milk)
	fmt.Printf("%d g of coffee beans\n", coffeeBeans)
	fmt.Printf("%d disposable cups\n", cups)
	fmt.Printf("$%d of money\n\n", money)
}

func takeMoney() {
	fmt.Printf("I gave you $%d\n", money)
	money = 0
}

func requiredIngredientForOneCoffee(coffeeType int) (int, int, int, int) {
	water, milk, coffeeBeans, price := 0, 0, 0, 0

	switch coffeeType {
	case 1:
		water = 250
		coffeeBeans = 16
		price = 4
	case 2:
		water = 350
		milk = 75
		coffeeBeans = 20
		price = 7
	case 3:
		water = 200
		milk = 100
		coffeeBeans = 12
		price = 6
	case 4:
		water = 350
		milk = 0
		coffeeBeans = 0
		price = 5
	}
	return water, milk, coffeeBeans, price
}

func getUserResponse(question string) int {
	fmt.Println(question)
	var total int
	fmt.Scanln(&total)
	return total
}

func fillMachine() {
	addWater := getUserResponse("Write how many ml of water do you want to add:")
	addMilk := getUserResponse("Write how many ml of milk do you want to add:")
	addCoffeeBeans := getUserResponse("Write how many grams of coffee beans do you want to add:")
	addCups := getUserResponse("Write how many disposable cups do you want to add:")

	water += addWater
	milk += addMilk
	coffeeBeans += addCoffeeBeans
	cups += addCups
}

func checkCoffeeMachineForChoice(choice int) {
	neededWater, neededMilk, neededCoffeeBeans, neededPrice := requiredIngredientForOneCoffee(choice)

	if cups <= 0 {
		fmt.Println("Sorry, not enough cup!")
	} else if neededWater > water {
		fmt.Println("Sorry, not enough water!")
	} else if neededMilk > milk {
		fmt.Println("Sorry, not enough milk!")
	} else if neededCoffeeBeans > coffeeBeans {
		fmt.Println("Sorry, not enough coffee beans!")
	} else {
		cups--
		water -= neededWater
		milk -= neededMilk
		coffeeBeans -= neededCoffeeBeans
		money += neededPrice
	}
}

func buyCoffee() {
	choice := getUserResponse("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, 4 - tea:")
	if choice >= 1 && choice <= 4 {
		checkCoffeeMachineForChoice(choice)
	} else {
		fmt.Println("Invalid choice.")
	}

}

func main() {

	for {
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		var action string
		fmt.Scanln(&action)
		if strings.EqualFold(action, "exit") {
			break
		}
		switch strings.ToLower(action) {
		case "buy":
			buyCoffee()
		case "fill":
			fillMachine()
		case "take":
			takeMoney()
		case "remaining":
			printState()
		default:
			fmt.Println("Invalid action.")
		}
	}
}
