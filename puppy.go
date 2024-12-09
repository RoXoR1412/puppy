package puppy

import (
	"github.com/RoXoR1412/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}
func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}
func Form11() string {
	return "I am from v1.1.0"
}
func From12() string {
	return ("I'm from version 1.2.0")
}
