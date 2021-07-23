package main

// Given a name, return a string with the message: One for X, one for me.Where X is the given name.
// However, if the name is missing, return the string: One for you, one for me.
// HINT: The string "One for X, one for me.Where X is the given name." is returned if the name is missing.

import "fmt"

func main() {
	fmt.Println(twofer("Bob"))
	fmt.Println(twofer("Alice"))
	fmt.Println(twofer(""))
	fmt.Println(hammingDistance("GGGCCGTTGGT", "GGACCGTTGAC"))
}

func twofer(name string) string {
	if len(name) == 0 {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}

// Calculate the Hamming Distance between two DNA strands.
// A mutation is simply a mistake that occurs during the creation or
// copying of a nucleic acid, in particular DNA. Because nucleic acids are
// vital to cellular functions, mutations tend to cause a ripple effect
// throughout the cell. Although mutations are technically mistakes, a very
// rare mutation may equip the cell with a beneficial attribute. In fact,
// the macro effects of evolution are attributable by the accumulated
// result of beneficial mutations throughout history.
// Each nucleotide in DNA contains a base pairing pairing,
// which is a double helix that consists of two purines (A or G) and
// two pyrimidines (C or T).
// The base pairing is essentially the structure of a nucleotide in
// the DNA molecule.
// Given two strings, return the Hamming Distance between them.
// For example, the Hamming Distance between “GGGCCGTTGGT” and “GGACCGTTGAC”
// is 3.

func hammingDistance(s1, s2 string) int {
	if len(s1) != len(s2) {
		panic("Strings are not of equal length")
	}
	count := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
		}
	}
	return count
}