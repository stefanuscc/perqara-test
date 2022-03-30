package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFraudNotif(t *testing.T) {
	assert.Equal(t, 1, fraudNotif(3, []int{10, 20, 30, 40, 50}))
	assert.Equal(t, 2, fraudNotif(5, []int{2, 3, 4, 2, 3, 6, 8, 4, 5}))
	assert.Equal(t, 0, fraudNotif(4, []int{1, 2, 3, 4, 4}))
}

func TestQueensAttack(t *testing.T) {
	assert.Equal(t, 24, queensAttack(8, 1, 4, 4, [][]int{{3, 5}}))
	assert.Equal(t, 10, queensAttack(5, 3, 4, 3, [][]int{{5, 5}, {4, 2}, {2, 3}}))
}
