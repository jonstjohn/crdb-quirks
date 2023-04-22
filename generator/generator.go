package generator

import (
	"github.com/google/uuid"
	"math"
	"math/rand"
)

type Generator struct {
	Rand *rand.Rand
}

func New(seed int64) *Generator {
	return &Generator{rand.New(rand.NewSource(seed))}
}

func (generator *Generator) UUID() (string, error) {
	uuid, err := uuid.NewRandomFromReader(generator.Rand)
	if err != nil {
		return "", err
	}
	return uuid.String(), err
}

const lowerLetters = "abcdefghijklmnopqrstuvwxyz"
const lowerLetterCnt = len(lowerLetters)
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const letterCnt = len(letters)
const UUIDChars = "0123456789abcdefg"

func String(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func StringFromRand(length int, rand *rand.Rand) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func LowerString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = lowerLetters[rand.Intn(len(lowerLetters))]
	}
	return string(b)
}

func UuidBytes() []byte {
	b := make([]byte, 32)
	for i := range b {
		b[i] = UUIDChars[rand.Intn(len(UUIDChars))]
	}
	return b
}

func Uuid() string {
	b := make([]byte, 32)
	for i := range b {
		b[i] = UUIDChars[rand.Intn(len(UUIDChars))]
	}
	return string(b)
}

func GenerateOrderedString(cnt int, max int) string {

	chars := 1

	// Add an extra character for each power
	for int(math.Pow(float64(letterCnt), float64(chars))) < max {
		chars++
	}

	result := ""

	for remainingChars := chars; remainingChars > 0; remainingChars-- {

		lindex := (int(math.Ceil(float64(cnt)/math.Pow(float64(letterCnt), float64(remainingChars-1)))) - 1) % letterCnt
		l := string(letters[lindex])
		result = result + l
	}

	return result
}

func GenerateOrderedLowerString(cnt int, max int) string {

	chars := 1

	// Add an extra character for each power
	for int(math.Pow(float64(lowerLetterCnt), float64(chars))) < max {
		chars++
	}

	result := ""

	for remainingChars := chars; remainingChars > 0; remainingChars-- {

		lindex := (int(math.Ceil(float64(cnt)/math.Pow(float64(lowerLetterCnt), float64(remainingChars-1)))) - 1) % lowerLetterCnt
		l := string(lowerLetters[lindex])
		result = result + l
	}

	return result
}
