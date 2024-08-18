package main

import (
	"fmt"
	"iter"
	"math/rand/v2"

	crand "crypto/rand"
)

func randomItr[T any](seed [32]byte, slice []T) iter.Seq[T] {
	ordered := make([]T, len(slice))
	copy(ordered, slice)
	rnd := rand.New(rand.NewChaCha8(seed))
	rnd.Shuffle(len(ordered), func(i, j int) {
		ordered[i], ordered[j] = ordered[j], ordered[i]
	})

	return func(yield func(T) bool) {
		for _, n := range ordered {
			yield(n)
		}
	}
}

func main() {
	var seed [32]byte
	_, err := crand.Read(seed[:])
	if err != nil {
		panic(err)
	}
	for i := range randomItr(seed, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}) {
		fmt.Println(i)
	}
}
