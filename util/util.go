package util

import (
	"math/rand"
	"time"
)

func Shuffle(vals []int) []int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	dest := make([]int, len(vals))
	perm := r.Perm(len(vals))
	for i, v := range perm {
		dest[v] = vals[i]
	}
	return dest
}

func GetNewRandom() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().Unix()))
}

func GetRandomArrayString(arr []string) string {
	return arr[GetNewRandom().Intn(len(arr))]
}
