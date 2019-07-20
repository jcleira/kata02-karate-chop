package basic

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChop(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Target int
		Array  []int
		Result int
	}{
		{3, []int{}, -1},
		{1, []int{1}, 0},
		{3, []int{1}, -1},

		{1, []int{1, 3, 5}, 0},
		{3, []int{1, 3, 5}, 1},
		{5, []int{1, 3, 5}, 2},
		{0, []int{1, 3, 5}, -1},
		{2, []int{1, 3, 5}, -1},
		{4, []int{1, 3, 5}, -1},
		{6, []int{1, 3, 5}, -1},

		{1, []int{1, 3, 5, 7}, 0},
		{3, []int{1, 3, 5, 7}, 1},
		{5, []int{1, 3, 5, 7}, 2},
		{7, []int{1, 3, 5, 7}, 3},
		{0, []int{1, 3, 5, 7}, -1},
		{2, []int{1, 3, 5, 7}, -1},
		{4, []int{1, 3, 5, 7}, -1},
		{6, []int{1, 3, 5, 7}, -1},
		{8, []int{1, 3, 5, 7}, -1},
	}

	for _, test := range tests {
		assert.Equal(test.Result, chop(test.Target, test.Array))
	}
}

func BenchmarkChop(b *testing.B) {
	rand.Seed(int64(b.N))
	for n := 0; n < b.N; n++ {
		chop(rand.Intn(21), []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21})
	}
}
