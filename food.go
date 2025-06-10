package main

import "strings"
import "math/rand"

type Food struct {
	char     string
	position Point
}

type FoodFactory struct {
	charSets []string
	position Point
}

func newFoodFactory(canvasSize Point, charSets string) *FoodFactory {
	return &FoodFactory{
		charSets: strings.Split(charSets, ""),
		position: Point{
			x: rand.Intn(canvasSize.x),
			y: rand.Intn(canvasSize.y),
		},
	}
}

func (f *FoodFactory) getFood() Food {
	return Food{
		char:     f.charSets[rand.Intn(len(f.charSets))],
		position: f.position,
	}
}
