package decorator

import (
	"strings"
	"testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza := &PizzaDecorator{}
	pizzaResult, err := pizza.AddIngredient()
	if err != nil {
		t.Error(err)
	}
	expectedText := "Pizza with the following ingredients:"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("When calling the add ingredient of the pizza decorator it "+
			"must return the text %s, not '%s'", expectedText, pizzaResult)
	}
}

func TestOnion_AddIngredient(t *testing.T) {
	onion := &Onion{}
	onionResult, err := onion.AddIngredient()
	if err == nil {
		t.Errorf("When calling AddIngredient on the onion decorator without "+
			"an IngredientAdder on its Ingredient field it must return an error, "+
			"not a string with '%s'", onionResult)
	}

	onion = &Onion{&PizzaDecorator{}}
	onionResult, err = onion.AddIngredient()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(onionResult, "onion") {
		t.Errorf("When calling the add ingredient of the onion decorator it "+
			"must return a text with the word 'onion', not '%s'", onionResult)
	}
}

func TestCheese_AddIngredient(t *testing.T) {
	cheese := &Cheese{}
	cheeseResult, err := cheese.AddIngredient()
	if err == nil {
		t.Error("When calling AddIngredient on the cheese decorator without " +
			"an IngredientAdder on its Ingredient field it must return an error")
	}

	cheese = &Cheese{&PizzaDecorator{}}
	cheeseResult, err = cheese.AddIngredient()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(cheeseResult, "cheese") {
		t.Errorf("When calling the add ingredient of the cheese decorator it "+
			"must return a text with the word 'meat', not '%s'", cheeseResult)
	}
}

func TestPizzaDecorator_FullStack(t *testing.T) {
	pizza := &Onion{&Cheese{&PizzaDecorator{}}}
	pizzaResult, err := pizza.AddIngredient()
	if err != nil {
		t.Error(err)
	}

	expectedText := "Pizza with the following ingredients: cheese, onion"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("When asking for a pizza with onion and cheese the returned "+
			"string must contain the text '%s' but '%s' didn't have it", expectedText,
			pizzaResult)
	}

	t.Log(pizzaResult)
}
