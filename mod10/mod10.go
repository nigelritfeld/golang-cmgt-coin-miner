package mod10

import (
	"crypto/sha256"
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

// HashSHA256 Hashes input with SHA256 Algorithm
func HashSHA256(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	hash := h.Sum(nil)
	return hashToHexString(string(hash))
}

// IntegerSliceToString Transforms byte to slice with seperate integers
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

// SumSequence Sums the integers in the slice and returns the remainder
func SumSequence(chunks [][]int) []int {
	var sequence []int
	for i, number := range chunks[0] {
		sequence = append(sequence, (number+chunks[1][i])%10)
	}
	return sequence
}

// Sum performs calculations on nested slices
func Sum(arr [][]int, n int) [][]int {
	//Length of array is equal to one we know if there is nothing left to multiply
	if n == 1 {
		return arr
	}
	multiplySequences := arr[:2]
	remainingBlocks := arr[2:]
	newSequence := SumSequence(multiplySequences)
	sequences := append([][]int{newSequence}, remainingBlocks...)
	return Sum(sequences, len(sequences))
}

// HashPayload Hashes input using the Luhn algorithm and sha256
func HashPayload(payload string) string {
	strippedPayload := Strip(payload)
	bytes := StringToAsciiBytes(strippedPayload)
	remainder := len(bytes) % 10
	if remainder != 0 {
		bytes = PopulateSequence(remainder, bytes)
	}
	chunks := ChunkSlice(bytes, 10)
	finishedSequence := Sum(chunks, len(chunks))
	sequenceToHash := IntegerSliceToString(finishedSequence[0])
	return HashSHA256(sequenceToHash)
}
