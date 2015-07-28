// A Channel-based implementation of a concurrent prime sieve using the sieve of eratosthenes with wheel factorization optimization

package main

import (
	"container/ring"
	"fmt"
	"log"
	"os"
	"strconv"
)

func wheel2357() *ring.Ring {

	var gaps2357 = []int{2, 4, 2, 4, 6, 2, 6, 4, 2, 4, 6, 6, 2, 6, 4, 2, 6, 4, 6, 8, 4, 2, 4, 2, 4, 8, 6, 4, 6, 2, 4, 6, 2, 6, 6, 4, 2, 4, 6, 2, 6, 4, 2, 4, 2, 10, 2, 10}
	r := ring.New(len(gaps2357))
	for _, i := range gaps2357 {
		r.Value = i
		r = r.Next()
	}
	return r
}

func generateCandidates(ch chan<- int) {
	ch <- 2
	ch <- 3
	ch <- 5
	ch <- 7
	nextGap := wheel2357()

	var gap int
	for i := 11; ; nextGap = nextGap.Next() {

		ch <- i
		gap = nextGap.Value.(int)
		i += gap
	}
}

func filterPrimes(in <-chan int, out chan<- int, prime int) {
	for {

		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("usage: primesieve <number-primes>")
		}
	}()

	numPrimes, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	primes := GenPrimes(numPrimes)
	for _, p := range primes {
		fmt.Println(p)
	}
}

func GenPrimes(numPrimes int) []int {
	ch := make(chan int)
	primes := make([]int, numPrimes)
	go generateCandidates(ch)
	for i := 0; i < numPrimes; i++ {
		prime := <-ch
		primes[i] = prime
		ch1 := make(chan int)
		go filterPrimes(ch, ch1, prime)
		ch = ch1
	}
	return primes
}
