// A Channel-based implementation of the concurrent prime sieve using the sive of atkin

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func GenerateCandidates(ch chan<- int) {
	ch <- 2
	ch <- 3
	for i, t := 5, 2; ; i, t = i+t, 6-t {
		ch <- i
	}
}

func FilterPrimes(in <-chan int, out chan<- int, prime int) {
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
	go GenerateCandidates(ch)
	for i := 0; i < numPrimes; i++ {
		prime := <-ch
		primes[i] = prime
		ch1 := make(chan int)
		go FilterPrimes(ch, ch1, prime)
		ch = ch1
	}
	return primes
}