package mod10

import (
	"fmt"
	"testing"
)

func TestHashing(t *testing.T) {
	t.Logf("This example test was running")

	slice := []int{1, 1, 6, 1, 0, 1, 1, 1, 5, 1, 1, 6, 9, 1, 1, 1, 1, 0, 7, 1, 0, 1, 0, 1, 2, 3, 4, 5, 6, 7}
	chunks := chunkSlice(slice, 10)
	fmt.Println(chunks)
	//finishedSequence := Sum(chunks, 0)
	if len(chunks) != 3 {
		t.Error("Slice is not split in to the correct amount of chunks expected")
	}
}
