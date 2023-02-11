package mod10

import (
	"fmt"
	"strconv"
	"strings"
)

// Strip Removes all whitespaces
func Strip(payload string) string {
	return strings.ReplaceAll(payload, " ", "")
}

// byteToSlice Splits byte into slice with integers
func byteToSlice(n byte, sequence []int) []int {
	if n != 0 {
		i := n % 10
		// sequence = append(sequence, i) // reverse order output
		sequence = append([]int{int(i)}, sequence...)
		return byteToSlice(n/10, sequence)
	}
	return sequence
}

// StringToAsciiBytes Converts all chars to bytes except numbers
func StringToAsciiBytes(str string) []int {
	// converting and printing Byte array
	data := []byte(str)
	var series []int
	var a []int
	//Checking if ascii code is a number
	for _, b := range data {
		if number, err := strconv.Atoi(string(b)); err == nil {
			//fmt.Printf("%q looks like a number.\n", string(b))
			//if so just add the integer to the array instead of spreading byte
			series = append(series, number)
		} else {
			slice := byteToSlice(b, a)
			series = append(series, slice...)
		}
	}
	return series

}

// ChunkSlice Creates chunks of slice
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

// hashToHexString Formats hash to hex string. Makes it more readable for me and useful for testing
func hashToHexString(hash string) string {
	return fmt.Sprintf("%x", hash)
}

// IntegerSliceToString Transforms byte to slice with separate integers
func IntegerSliceToString(slice []int) string {
	var IDs []string
	for _, i := range slice {
		IDs = append(IDs, strconv.Itoa(i))
	}
	return strings.Join(IDs, "")
}

// PopulateSequence Takes sequence and remainder and populated remainder with mod10
func PopulateSequence(remainder int, sequence []int) []int {
	start := len(sequence) - remainder
	roundedSequence := sequence[:len(sequence)-remainder]
	sliceToPopulate := sequence[start:]
	count := 10 - len(sliceToPopulate)
	var a = make([]int, count)
	for i := 0; i <= count-1; i++ {
		a[i] = i
	}
	populatedSequence := append(sliceToPopulate, a...)
	return append(roundedSequence, populatedSequence...)
}
