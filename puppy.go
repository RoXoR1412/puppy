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
	return dog.WhenGrowsUp(Bark())
}
func BigBark() string {
	return dog.WhenGrowsUp(Barks())
}
