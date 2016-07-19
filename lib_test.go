package random_project_generator

import (
	"fmt"
	"regexp"
	"testing"
)

func TestDifferentResults(t *testing.T) {
	t.Parallel()
	word1 := Generate()
	word2 := Generate()
	if word1 == word2 {
		t.Errorf("Generate() got same results for different calls: %s, %s", word1, word2)
	}
}

func TestReturnedValue(t *testing.T) {
	t.Parallel()
	if len(Generate()) == 0 {
		t.Errorf("Generate() returned an empty string")
	}
	if len(GenerateNumber(2)) == 0 {
		t.Errorf("GenerateNumber() returned an empty string")
	}
}

var validRandomNumber = regexp.MustCompile(`^[a-z]+-[a-z]+-[0-9]{4}$`)

func TestGenerateNumber(t *testing.T) {
	t.Parallel()
	for i := 0; i < 10000; i++ {
		word := GenerateNumber(4)
		if !validRandomNumber.MatchString(word) {
			t.Errorf("generated word did not match regex: %s", word)
		}
	}
}

func TestToThe(t *testing.T) {
	t.Parallel()
	if tothe(1) != 10 {
		t.Errorf("tothe(1) should equal 10, got %d", tothe(1))
	}
	if tothe(2) != 100 {
		t.Errorf("tothe(2) should equal 100, got %d", tothe(2))
	}
	if tothe(3) != 1000 {
		t.Errorf("tothe(3) should equal 1000, got %d", tothe(3))
	}
}

var validTests = []struct {
	in       string
	expected bool
}{
	{"foo-bar-3", false},
	{"", false},
	{"foo-bar", false},
	{"aback-yam", true},
	{"absent-zoo-23", true},
	{"absent-zoo-yam", false},
	{"absent-zoo--23", false},
	{"absent-zoo-0", false},
}

func TestValid(t *testing.T) {
	t.Parallel()
	for _, tt := range validTests {
		if Valid(tt.in) != tt.expected {
			t.Errorf("Valid(\"%s\"): got %t, expected %t", tt.in, !tt.expected, tt.expected)
		}
	}
}

var benchWord string
var benchWordNumber string
var benchValid bool

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchWord = Generate()
	}
}

func BenchmarkGenerateNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchWordNumber = GenerateNumber(5)
	}
}

func BenchmarkValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchValid = Valid("zonked-zoo-12334343")
	}
}

func ExampleGenerate() {
	// Will print words like "smiling-fold"
	fmt.Println(Generate())
}

func ExampleGenerateNumber() {
	// Will print a string like "viable-action-218"
	fmt.Println(GenerateNumber(3))
}
