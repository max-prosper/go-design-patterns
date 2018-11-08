package decorator

import (
	"errors"
	"fmt"
)

type IngredientAdd interface {
	AddIngredient() (string, error)
}

type PizzaDecorator struct {
	Ingredient IngredientAdd
}

func (p *PizzaDecorator) AddIngredient() (string, error) {
	return "Pizza with the following ingredients:", nil
}

type Cheese struct {
	Ingredient IngredientAdd
}

func (c *Cheese) AddIngredient() (string, error) {
	if c.Ingredient == nil {
		return "", errors.New("An IngredientAdder is needed on the Ingredient field of the Cheese")
	}
	s, err := c.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "cheese"), nil
}

type Onion struct {
	Ingredient IngredientAdd
}

func (o *Onion) AddIngredient() (string, error) {
	if o.Ingredient == nil {
		return "", errors.New("An IngredientAdder is needed on the Ingredient field of the Onion")
	}
	s, err := o.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "onion"), nil
}
