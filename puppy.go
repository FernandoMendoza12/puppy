package puppy

import (
	dog "github.com/FernandoMendoza12/GoesToEleven"
)

func Bark() string {
	return "Woof"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrowUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrowUp(Barks())
}

func From11() string {
	return "Version number v1.1.0"
}

func From12() string {
	return "Version number v2.2.0"
}

func From13() string {
	return "Version number 3.3.0"
}
