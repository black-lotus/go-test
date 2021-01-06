package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getElements(n int) []int {
	result := make([]int, n)
	j := 0
	for i := n - 1; i > 0; i-- {
		result[j] = i
		j++
	}
	return result
}

func TestSort(t *testing.T) {
	elements := getElements(10)

	Sort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])
}

func TestSortMoreThan10000(t *testing.T) {
	elements := getElements(20001)

	Sort(elements)

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 20000 {
		t.Error("last element should be 20000")
	}
}

func BenchmarkBubbleSort10K(b *testing.B) {
	elements := getElements(20000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkBubbleSort100K(b *testing.B) {
	elements := getElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
