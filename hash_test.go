package basic

import (
	"crypto/sha1"
	"fmt"
	"testing"
)

func TestHashing(t *testing.T) {
	var text = "Hai Si imut"
	var sha = sha1.New()
	sha.Write([]byte(text))

	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)
	fmt.Println(encryptedString)
}
