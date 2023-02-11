package mod10

import (
	"crypto/sha256"
)

// HashSHA256 Hashes input with SHA256 Algorithm
func HashSHA256(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	hash := h.Sum(nil)
	return hashToHexString(string(hash))
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
