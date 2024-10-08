package main

import (
	"fmt"
	"log"
	"infectious"

)

func main() {
	fmt.Println("Welcome to playing around with infectious!")

	// Example usage of the infectious package
	// You can add your code to work with the infectious package here
	const (
		required = 4
		total    = 8
	)
	// Create a *FEC, which will require 'required' pieces
	// for reconstruction at minimum, and generate 'total'
	// total pieces.
	f, err := infectious.NewFEC(required, total)
	if err != nil {
		panic(err)
	}

	// Prepare to receive the shares of encoded data.
	shares := make([]infectious.Share, total)
	output := func(s infectious.Share) {
		// we need to make a copy of the data. The share data
		// memory gets reused when we return.
		shares[s.Number] = s.DeepCopy()
	}
	// the data to encode must be padded to a multiple of 'required',
	// hence the underscores.
	err = f.Encode([]byte("hello, world! __"), output)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	for _, share := range shares {
		fmt.Printf("Share Number: %d, Data: ", share.Number)
        for _, b := range share.Data {
            fmt.Printf("%d, ", b)
        }
        fmt.Printf("\n")
	}

}
