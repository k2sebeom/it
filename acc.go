package it

import (
	"cmp"
	"iter"
)

// Max returns the maximum element of seq.
// The second return value reports whether the element exists (seq is not empty).
func Max[V cmp.Ordered](seq iter.Seq[V]) (V, bool) {
	var acc V
	ok := false
	for v := range seq {
		if !ok || v > acc {
			acc, ok = v, true
		}
	}
	return acc, ok
}

// Min returns the minimum element of seq.
// The second return value reports whether the element exists (seq is not empty).
func Min[V cmp.Ordered](seq iter.Seq[V]) (V, bool) {
	var acc V
	ok := false
	for v := range seq {
		if !ok || v < acc {
			acc, ok = v, true
		}
	}
	return acc, ok
}

// Fold applies f to each element of seq and returns the accumulated result.
//
// init is the initial value of the accumulator.
// f takes the accumulator and an element of seq as arguments.
// The result of each call to f becomes the accumulator for the next call.
func Fold[V, R any](seq iter.Seq[V], init R, f func(R, V) R) R {
	acc := init
	for v := range seq {
		acc = f(acc, v)
	}
	return acc
}
