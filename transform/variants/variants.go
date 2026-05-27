/*
Package variants contains a function for generating all variants of a sequence.

Sometimes sequencers will only give you an *estimate* of what the basepair at
a given position is. This package provides a function for generating all
possible deterministic variants of a sequence given a sequence
with ambiguous bases.
*/
package variants

// AllVariantsIUPAC takes a string as input
// and returns all iupac variants as output
func AllVariantsIUPAC(seq string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// rune map of all iupac nucleotide variants

func cartRune(inLists ...[]rune) [][]rune {
	_ = "STUB: not implemented"
	// An iterative approach to calculate Cartesian product of two or more lists
	// Adapted from https://rosettacode.org/wiki/Cartesian_product_of_two_or_more_lists
	// supposedly "minimizes allocations and computes and fills the result sequentially"
	return nil
}

// a counter used to determine the possible number of variants

// in the future this could be part of an error return?

// this is the 2D slice where all variants will be stored
// this is an empty slice with a length totaling the size of all input characters
// these will be all the possible variants

// define end point

// start at end point
