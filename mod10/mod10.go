package mod10

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

// Removes all whitespaces
func Strip(payload string) string {
	return strings.ReplaceAll(payload, " ", "")
}

func IntToSlice(n byte, sequence []int) []int {
	if n != 0 {
		i := n % 10
		// sequence = append(sequence, i) // reverse order output
		sequence = append([]int{int(i)}, sequence...)
		return IntToSlice(n/10, sequence)
	}
	return sequence
}

// StringToAsciiBytes Converts all chars to bytes except numbers
func StringToAsciiBytes(str string) []int {
	// converting and printing Byte array
	data := []byte(str)
	fmt.Printf("Converting %s to ascii", str)
	fmt.Println()
	fmt.Println(data)
	var series []int
	var a []int
	//Checking if ascii code is a number
	for _, b := range data {
		if number, err := strconv.Atoi(string(b)); err == nil {
			fmt.Printf("%q looks like a number.\n", string(b))
			//if so just add the integer to the array instead of spreading byte
			series = append(series, number)
		} else {
			slice := IntToSlice(b, a)
			series = append(series, slice...)
		}
	}

	fmt.Printf("The finished series has a capacity of %d", len(series))

	return series

}

func SplitNumbersFromArray(array []byte) int {

	//todo: Alle getallen worden gesplitst en los in een array geplaatst
	//fmt.Println(res)
	//for _, number := range array {
	//	fmt.Println(number)
	//	bytes.SplitN(number, []byte(""), 3)
	//	res := bytes.Split(byte(number), []byte(""))
	//
	//}

	return 0
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

func ChunkSlice(slice []int, chunkSize int) [][]int {
	var chunks [][]int
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func HashBlock(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	hash := h.Sum(nil)
	fmt.Println(input)
	return string(hash)
}

func IntegerSliceToString(slice []int) string {
	var IDs []string
	for _, i := range slice {
		IDs = append(IDs, strconv.Itoa(i))
	}
	return strings.Join(IDs, "")
}

func PopulateArray(remainder int, series []int) []int {
	start := len(series) - remainder

	fmt.Println()
	fmt.Printf("STARt: %d end: %d", start, len(series)-1)
	fmt.Println()

	roundedSerie := series[:len(series)-remainder]
	fmt.Printf("Rounded values")

	fmt.Println(series)
	fmt.Println(len(roundedSerie))
	fmt.Println()
	sliceToPopulate := series[start:]
	fmt.Println()

	fmt.Printf("sliceToPopulate %d", sliceToPopulate)
	count := 10 - len(sliceToPopulate)
	fmt.Println(count)

	var a = make([]int, count)
	for i := 0; i <= count-1; i++ {
		a[i] = i
	}
	fmt.Println(a)

	populatedArray := append(sliceToPopulate, a...)
	return append(roundedSerie, populatedArray...)
}

func SumSequence(chunks [][]int) [][]int {
	var sequence []int
	remainder := chunks[2:]
	fmt.Println("inital")
	fmt.Println(chunks[0])
	fmt.Println("+")
	fmt.Println(chunks[1])
	fmt.Println("remainder")
	fmt.Println(remainder)

	for i, number := range chunks[0] {
		sequence = append(sequence, (number+chunks[1][i])%10)
	}
	fmt.Println("sequence")
	fmt.Println(sequence)
	fmt.Println("Finisshed")
	finished := append([][]int{sequence}, remainder...)
	fmt.Printf("NEW LeNFTH: %d", len(finished))

	fmt.Println("--------------")
	return finished
}

func Sum(arr [][]int, start int, end int) [][]int {
	if end == 0 {
		end = len(arr)
		fmt.Printf("NEW END: %d", end)

	}
	//fmt.Println(len(arr) % 2)
	fmt.Printf("Current count: %d", start)
	fmt.Printf("MAX CAP count: %d", cap(arr))
	if start > len(arr) {
		return arr
	}
	if start < len(arr) {
		arr = SumSequence(arr)
		fmt.Println("-------------------------")
		fmt.Printf("Adding block %d", start)
		fmt.Println("-------------------------")
		fmt.Println()
		fmt.Println(arr)
		return Sum(arr, start+1, end)
	}
	return Sum(arr, start+1, end)
}

func HashPayload(payload string) string {
	strippedPayload := Strip(payload)
	convertedBytes := StringToAsciiBytes(strippedPayload)
	fmt.Println(convertedBytes)
	////todo: Get first block of 10 digits
	remainder := len(convertedBytes) % 10
	////todo: Get second block of 10 items
	if remainder != 0 {
		fmt.Printf("Remainder %d", remainder)
		convertedBytes = PopulateArray(remainder, convertedBytes)
		fmt.Println("Rounded series")
		fmt.Println(convertedBytes)
		fmt.Println(len(convertedBytes))
	}
	chunks := ChunkSlice(convertedBytes, 10)
	fmt.Println("Chunks lenght:")
	fmt.Println(len(chunks))
	finishedSequence := Sum(chunks, 0, 0)
	//We select first two blocks and sum the sequences
	fmt.Println(finishedSequence)
	fmt.Println(HashBlock(IntegerSliceToString(finishedSequence[0])))
	str := IntegerSliceToString(finishedSequence[0])
	return HashBlock(str)
}
