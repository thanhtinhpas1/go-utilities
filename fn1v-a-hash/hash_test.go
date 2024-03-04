package main

import (
	"math/rand"
	"testing"
)

const (
	numWords = 10000
	wordLen  = 7
)

func BenchmarkMostlyUniqueCounter(b *testing.B) {
	words := make([][]byte, numWords)
	for i := 0; i < numWords; i++ {
		words[i] = randomWord(wordLen)
	}

	b.ResetTimer()
	var counts Counter
	for i := 0; i < b.N; i++ {
		for i := 0; i < numWords; i++ {
			counts.Incr(words[i], 1)
		}
	}
}

func BenchmarkNonUniqueCounter(b *testing.B) {
	words := make([][]byte, numWords/10)
	for i := 0; i < numWords/10; i++ {
		words[i] = randomWord(wordLen)
	}
	var counts Counter
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for wi := 0; wi < numWords/10; wi++ {
			counts.Incr(words[wi], 1)
			counts.Incr(words[wi], 1)
			counts.Incr(words[wi], 1)
			counts.Incr(words[wi], 1)
			counts.Incr(words[wi], 1)
			counts.Incr(words[wi], 1)
			counts.Incr(words[wi], 1)
			counts.Incr(words[wi], 1)
			counts.Incr(words[wi], 1)
			counts.Incr(words[wi], 1)
		}
	}
}

func BenchmarkMostlyUniqueMapBytes(b *testing.B) {
	words := make([][]byte, numWords)
	for i := 0; i < numWords; i++ {
		words[i] = randomWord(wordLen)
	}
	counts := make(map[string]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for wi := 0; wi < numWords; wi++ {
			counts[string(words[wi])]++
		}
	}
}

func BenchmarkNonUniqueMapBytes(b *testing.B) {
	words := make([][]byte, numWords/10)
	for i := 0; i < numWords/10; i++ {
		words[i] = randomWord(wordLen)
	}
	counts := make(map[string]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for wi := 0; wi < numWords/10; wi++ {
			counts[string(words[wi])]++
			counts[string(words[wi])]++
			counts[string(words[wi])]++
			counts[string(words[wi])]++
			counts[string(words[wi])]++
			counts[string(words[wi])]++
			counts[string(words[wi])]++
			counts[string(words[wi])]++
			counts[string(words[wi])]++
			counts[string(words[wi])]++
		}
	}
}

func BenchmarkMostlyUniqueMapPointerBytes(b *testing.B) {
	words := make([][]byte, numWords)
	for i := 0; i < numWords; i++ {
		words[i] = randomWord(wordLen)
	}
	counts := make(map[string]*int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for wi := 0; wi < numWords; wi++ {
			word := words[wi]
			p, ok := counts[string(word)]
			if ok {
				*p++
			} else {
				n := 1
				counts[string(word)] = &n
			}
		}
	}
}

func BenchmarkNonUniqueMapPointerBytes(b *testing.B) {
	words := make([][]byte, numWords/10)
	for i := 0; i < numWords/10; i++ {
		words[i] = randomWord(wordLen)
	}
	counts := make(map[string]*int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for wi := 0; wi < numWords/10; wi++ {
			for j := 0; j < 10; j++ {
				word := words[wi]
				p, ok := counts[string(word)]
				if ok {
					*p++
				} else {
					n := 1
					counts[string(word)] = &n
				}
			}
		}
	}
}

func BenchmarkMostlyUniqueMapString(b *testing.B) {
	words := make([]string, numWords)
	for i := 0; i < numWords; i++ {
		words[i] = string(randomWord(wordLen))
	}
	counts := make(map[string]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for wi := 0; wi < numWords; wi++ {
			counts[words[wi]]++
		}
	}
}

func BenchmarkNonUniqueMapString(b *testing.B) {
	words := make([]string, numWords/10)
	for i := 0; i < numWords/10; i++ {
		words[i] = string(randomWord(wordLen))
	}
	counts := make(map[string]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for wi := 0; wi < numWords/10; wi++ {
			counts[words[wi]]++
			counts[words[wi]]++
			counts[words[wi]]++
			counts[words[wi]]++
			counts[words[wi]]++
			counts[words[wi]]++
			counts[words[wi]]++
			counts[words[wi]]++
			counts[words[wi]]++
			counts[words[wi]]++
		}
	}
}

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func randomWord(length int) []byte {
	b := make([]byte, length)

	for i := 0; i < length; i++ {
		b[i] = chars[rand.Intn(len(chars))]
	}

	return b
}
