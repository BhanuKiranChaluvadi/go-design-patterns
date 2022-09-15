package main

import "fmt"

func makeAmericano(size float32) {
	fmt.Println("\nMaking an Americano\n--------------------")

	// TODO: make an americano coffee using the coffeemachine API
	americano := &CoffeeMachine{}

	// determine beans amount to use - 5oz for every 8oz size
	beanAmount := (size / 8.0) * 5.0
	americano.startCoffee(beanAmount, 2)
	americano.grindBeans()
	americano.useHotWater(size)
	americano.endCoffee()

	fmt.Println("Americano is ready!")
}

func makeLatte(size float32, foam bool) {
	fmt.Println("\nMaking a Latte\n--------------------")

	// TODO: make a latte coffee using the coffeemachine API
	latte := &CoffeeMachine{}

	// determine beans amount to use - 5oz for every 8oz size
	beanAmount := (size / 8.0) * 5.0
	latte.startCoffee(beanAmount, 4)
	latte.grindBeans()
	latte.useHotWater(size)

	// determine milk amount to use - 2oz for every 8oz size
	milk := (size / 8.0) * 2.0
	latte.useMilk(milk)
	latte.doFoam(true)
	latte.endCoffee()
	fmt.Println("Latte is ready!")
}
