/*
Package fragment optimally fragments DNA for GoldenGate systems.

Optimal fragmentation is accomplished by using empirical fidelity data derived
by NEB in the paper "Enabling one-pot Golden Gate assemblies of unprecedented
complexity using data-optimized assembly design". We use the BsaI-T4 ligase
data provided in table S1.

Paper link: https://doi.org/10.1371/journal.pone.0238592
Data link: https://doi.org/10.1371/journal.pone.0238592.s001
*/
package fragment

// SetEfficiency gets the estimated fidelity rate of a given set of
// GoldenGate overhangs.
func SetEfficiency(overhangs []string) float64 { _ = "STUB: not implemented"; return 0 }

// NextOverhangs gets a list of possible next overhangs to use for an overhang
// list, along with their efficiencies. This can be used for more optimal
// fragmentation of sequences with potential degeneracy.
func NextOverhangs(currentOverhangs []string) ([]string, []float64) {
	_ = "STUB: not implemented"
	return nil, nil
}

// These 4 for loops generate all combinations of 4 base pairs
// checking all 256 4mer combinations for palindromes. Palindromes
// can cause problems in large combinatorial reactions, so are
// removed here.

// NextOverhang gets next most efficient overhang to use for a given set of
// overhangs. This is useful for when developing a new set of standard
// overhangs. Note: NextOverhang is biased towards high AT overhangs, but this
// will not affect fidelity at all.
func NextOverhang(currentOverhangs []string) string { _ = "STUB: not implemented"; return "" }

// optimizeOverhangIteration takes in a sequence and optimally fragments it.
func optimizeOverhangIteration(sequence string, minFragmentSize int, maxFragmentSize int, existingFragments []string, excludeOverhangs []string, includeOverhangs []string) ([]string, float64, error) {
	_ = "STUB: not implemented"
	// If the sequence is smaller than maxFragment size, stop iteration.
	return nil, 0, nil
}

// Make sure minFragmentSize > maxFragmentSize

// Minimum lengths (given oligos) for assembly is 8 base pairs
// https://doi.org/10.1186/1756-0500-3-291
// For GoldenGate, 2 8bp oligos create 12 base pairs (4bp overhangs on two sides of 4bp),
// so we check for minimal size of 12 base pairs.

// If our iteration is approaching the end of the sequence, that means we need to gracefully handle
// the end so we aren't left with a tiny fragment that cannot be synthesized. For example, if our goal
// is fragments of 100bp, and we have 110 base pairs left, we want each final fragment to be 55bp, not
// 100 and 10bp

// buffer is needed equations above pass.

// Get all sets of 4 between the min and max FragmentSize

// We go from max -> min, so we can maximize the size of our fragments

// Make sure overhang isn't already in set

// Make sure overhang is in set of includeOverhangs. If includeOverhangs is
// blank, skip this check.

// See if this overhang is a palindrome

// Get this overhang set's efficiency

// If this overhang is more efficient than any other found so far, set it as the best!

// Set variables

// Fragment fragments a sequence into fragments between the min and max size,
// choosing fragment ends for optimal assembly efficiency. Since fragments will
// be inserted into either a vector or primer binding sites, the first 4 and
// last 4 base pairs are the initial overhang set.
func Fragment(sequence string, minFragmentSize int, maxFragmentSize int, excludeOverhangs []string) ([]string, float64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// FragmentWithOverhangs fragments a sequence with only a certain overhang set.
// This is useful if you are constraining the set of possible overhangs when
// doing more advanced forms of cloning.
func FragmentWithOverhangs(sequence string, minFragmentSize int, maxFragmentSize int, excludeOverhangs []string, includeOverhangs []string) ([]string, float64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}
