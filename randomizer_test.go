package basic

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandomizer(t *testing.T) {
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Println(randomizer.Intn(42))
	fmt.Println(randomizer.Uint32())
	fmt.Println(randomizer.Float32())
	
 
}

func TestStringRandomizer(t *testing.T) {
	stringRandomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	letter := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	var amba = func (length int) string {
		b := make([]rune, length)
		for i := range b {
			b[i] = letter[stringRandomizer.Intn(len(letter))]
		}
		return string(b)
	}

	fmt.Println("Random String 5 karakter", amba(4))
	
}
