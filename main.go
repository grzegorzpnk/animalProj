package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var animals map[string]Animal
	animals = make(map[string]Animal)

	bird := Animal{"worms", "fly", "peep"}
	cow := Animal{"grass", "walk", "moo"}
	snake := Animal{"mice", "slither", "hsss"}

	animals["bird"] = bird
	animals["cow"] = cow
	animals["snake"] = snake

	fmt.Println(animals)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("please provide:  animal_name space request_type ")
		text, _ := reader.ReadString('\n')
		splitedText := strings.Fields(text)
		fmt.Println(splitedText)
		switch splitedText[1] {
		case "food":
			fmt.Println(animals[splitedText[0]].Eat())
		case "locomotion":
			fmt.Println(animals[splitedText[0]].Move())
		case "noise":
			fmt.Println(animals[splitedText[0]].Speak())
		}
	}
}

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
