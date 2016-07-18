// Package random_project_generator generates random project names.
//
// Random names are generated by selecting a random adjective from a list of
// 1300 adjectives, and a random noun from a list of 870 nouns, giving you over
// a million possible combinations.
package random_project_generator

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Generate finds a random adjective and a random noun and joins them with
// a hyphen.
func Generate() string {
	adjective := Adjectives[r.Intn(len(Adjectives))]
	noun := Nouns[r.Intn(len(Nouns))]
	return adjective + "-" + noun
}

func tothe(power int) int {
	sum := 1
	for i := 0; i < power; i++ {
		sum = 10 * sum
	}
	return sum
}

// GenerateNumber returns a string with an adjective, a noun, and a number with
// d digits, joined together by hyphens. Each additional digit adds a factor of
// 9 to the number of possible results that can be generated, so 4 digits will
// yield roughly 7420491000 different possibilities.
func GenerateNumber(d int) string {
	if d < 0 {
		panic(fmt.Sprintf("random_project_generator: invalid negative length %d", d))
	}
	maxval := tothe(d) - 1
	minval := tothe(d - 1)
	return Generate() + "-" + strconv.Itoa(r.Intn(maxval-minval)+minval)
}
