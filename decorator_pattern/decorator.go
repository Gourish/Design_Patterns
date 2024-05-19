package main

type Ipizza interface {
	getPrice() int
}

type margareta struct{}

func (mg *margareta) getPrice() int {
	return 120
}

type cheeseToppings struct {
	pizza Ipizza
}

func (ct *cheeseToppings) getPrice() int {
	return ct.pizza.getPrice() + 50
}

type tomatoToppings struct {
	pizza Ipizza
}

func (ct *tomatoToppings) getPrice() int {
	return ct.pizza.getPrice() + 25
}
