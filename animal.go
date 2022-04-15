package main

//Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings.
type Animal struct {
	Food       string
	Locomotion string
	Noise      string
}

func (animal Animal) Eat() string {
	return animal.Food
}

func (animal Animal) Move() string {
	return animal.Locomotion
}

func (animal Animal) Speak() string {
	return animal.Noise
}
