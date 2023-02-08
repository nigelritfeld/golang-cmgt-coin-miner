package mod10

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Removes all whitespaces
func Strip(payload string) string {
	return strings.ReplaceAll(payload, " ", "")
}

// Converts all chars to bytes except numbers
func StringToAsciiBytes(str string) []byte {
	fmt.Println(str)
	t := make([]uint8, utf8.RuneCountInString(str))
	fmt.Println(t)
	i := 0

	for _, item := range str {
		//Check whether the value in string is type of int
		if number, err := strconv.Atoi(string(item)); err == nil {
			fmt.Printf("%q looks like a number.\n", string(item))
			fmt.Println(number)
			//if so just add the integer to the array instead of byte
			t[i] = uint8(byte(number))
		} else {
			t[i] = uint8(byte(item))
		}
		i++
	}
	return t
}

func toArray(array []uint8) []string {
	//todo:  Alle getallen worden gesplitst en los in een array geplaatst
	return nil
}

func Calculate() int {
	//todo: Deze twee blokken worden bij elkaar opgeteld op de volgende manie
	//todo: Van beide blokken wordt het eerste getal genomen
	//todo: deze worden bij elkaar opgeteld
	//todo: van de uitkomt wordt modulus 10 genomen
	//todo: Bijv. (7 + 8) % 10 = 15 % 10 = 5
	//todo: Dit wordt 10 x gedaan voor alle getallen uit beide blokken van 10 getallen
	//todo: Dit levert een nieuwe reeks van 10 getallen op
	return 0
}

func HashBlock(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	hash := h.Sum(nil)
	fmt.Println(input)
	//fmt.Printf("%x\n", hash)
	return string(hash)
}
