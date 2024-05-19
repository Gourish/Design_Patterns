package main

import "fmt"

func main() {
	pizza := margareta{}

	pizzaWithCheeseToppings := &cheeseToppings{pizza: &pizza}

	fmt.Printf("Price of pizza with cheese topping %d\n", pizzaWithCheeseToppings.getPrice())

	pizzaWithTomatoAndCheeseToppings := &tomatoToppings{pizza: pizzaWithCheeseToppings}

	fmt.Printf("Price of pizza with cheese and tomato toppings %d\n", pizzaWithTomatoAndCheeseToppings.getPrice())
}
