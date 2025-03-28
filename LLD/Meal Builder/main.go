package main

import "fmt"

type Meal struct {
	main   string
	side   string
	drink  string
	desert string
}

func newMeal(options ...func(*Meal)) *Meal {
	meal := &Meal{}
	for _, option := range options {
		option(meal)
	}
	return meal
}

func main() {
	fmt.Println(newMeal(WithMain("North Indian Thali"), WithDrink(" Lassi "), WithDesert("Jamoon")))

	fmt.Println(newMeal(WithMain("South Indian thali"), WithSide(" Palya "), WithDesert("Jamoon")))
}

func WithMain(main string) func(*Meal) {
	return func(m *Meal) { m.main = main }
}

func WithSide(side string) func(*Meal) {
	return func(m *Meal) { m.side = side }
}

func WithDrink(drink string) func(*Meal) {
	return func(m *Meal) { m.drink = drink }
}

func WithDesert(desert string) func(*Meal) {
	return func(m *Meal) { m.desert = desert }
}
