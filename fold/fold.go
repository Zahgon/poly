/*
Package fold is a package for folding DNA and RNA sequences into secondary structures.

This package provides everything you need to fold a DNA or RNA sequence into a secondary structure
and get the minimum free energy of the structure. Most of the code was ported from the
python SeqFold package by Lattice Automation and Joshua Timmons but we hope to have a
linear fold algorithm in the near future.

Biological context:

DNA, RNA, and proteins all fold. Protein is a particularly tricky thing to predict
partially because there are so many more amino acids than there are nucleotides.

ACG(T/U) vs. ACDEFGHIKLMNPQRSTVWY (20 amino acids)

These folding predictions help us understand how to design primers, guide RNAs, and
other nucleic acid sequences that fold into a particular structure.

Fortunately for us, DNA and RNA are much easier to predict because there are only 4 nucleotides
and the rules for folding are much more well defined.

Each function has citations to the original papers that describe the algorithms used.
Most of the algorithms used in this package are based on the work of Zuker and Stiegler, 1981
but we're hoping to add more algorithms in the near future such as linear fold.

TTFN,
Tim
*/
package fold

// Zuker folds the DNA sequence and return the lowest free energy score.
//
// Based on the approach described in:
// Zuker and Stiegler, 1981
// https://www.ncbi.nlm.nih.gov/pmc/articles/PMC326673/pdf/nar00394-0137.pdf
//
// If the sequence is 50 or more bp long, "isolated" matching bp
// are ignored in pairedMinimumFreeEnergyV(start,end). This is based on an approach described in:
// Mathews, Sabina, Zuker and Turner, 1999
// https://www.ncbi.nlm.nih.gov/pubmed/10329189
// Args:
//
//	seq: The sequence to Fold
//	temp: The temperature the Fold takes place in, in Celsius
//
// Returns a slice of NucleicAcidStructure with the energy and description,
// i.e. stacks, bulges, hairpins, etc.
func Zuker(seq string, temp float64) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

// get the minimum free energy structure out of the cache

// unpairedMinimumFreeEnergyW returns the minimum free energy of a subsequence
// at start and terminating at end.
//
// From Zuker and Stiegler, 1981: let W(i,j) be the minimum free energy of all
// possible admissible structures formed from the subsequence Sij.
//
// Figure 2B in Zuker and Stiegler, 1981
// Args:
//
//		seq: The sequence being folded
//		start: The start index
//		end: The end index (inclusive)
//	 foldContext: The context for this sequence
//
// Returns the free energy for the subsequence from start to end
func unpairedMinimumFreeEnergyW(start, end int, foldContext context) (nucleicAcidStructure, error) {
	_ = "STUB: not implemented"
	return *new(nucleicAcidStructure), nil
}

// pairedMinimumFreeEnergyV returns the minimum free energy of a subsequence of
// paired bases and end.
// From Figure 2B of Zuker, 1981: let V(i,j) be the minimum free energy of all
// possible admissible structures formed from Sij in which Si and Sj base pair
// with each other. If Si and Sj cannot base pair, then V(i,j) = infinity
//
// If start and end don't bp, store and return INF.
// See: Figure 2B of Zuker, 1981
// Args:
//
//		start: The start index
//		end: The end index (inclusive)
//	 foldContext: The context for this sequence
//
// Returns the minimum energy folding structure possible between start and end on seq
func pairedMinimumFreeEnergyV(start, end int, foldContext context) (nucleicAcidStructure, error) {
	_ = "STUB: not implemented"
	return *new(nucleicAcidStructure), nil
}

// the ends must basepair for pairedMinimumFreeEnergyV(start,end)

// if the basepair is isolated, and the seq large, penalize at 1,600 kcal/mol
// heuristic for speeding this up
// from https://www.ncbi.nlm.nih.gov/pubmed/10329189

// small hairpin; 4bp

// rightOfStart and leftOfEnd must match

// it's a neighboring/stacking pair in a helix

// there's a dangling end

// it's an interior loop

// technically an interior loop of 1. really 1bp mismatch

// it's a bulge on the left side

// it's a bulge on the right side

// it's basically a hairpin, only outside bp match

// add pairedMinimumFreeEnergyV(start', end')

// Bulge calculates the free energy associated with a bulge.
//
// Args:
//
//		start: The start index of the bulge
//		rightOfStart: The index to the right of start
//		end: The end index of the bulge
//		leftOfEnd: The index to the left of end
//	 foldContext: The FoldingContext for this sequence
//
// Returns the increment in free energy from the bulge
func Bulge(start, rightOfStart, end, leftOfEnd int, foldContext context) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// add penalty based on size

// it's too large for pre-calculated list, extrapolate

// if len 1, include the delta G of intervening nearestNeighbors (SantaLucia 2004)

// penalize AT terminal bonds

func addBranch(structure nucleicAcidStructure, branches *[]subsequence, foldContext context) error {
	_ = "STUB: not implemented"
	return nil
}

// multibranch calculates a multi-branch foldEnergy penalty using a linear formula.
//
// From Jaeger, Turner, and Zuker, 1989.
// Found to be better than logarithmic in Ward, et al. 2017
// Args:
//
//		start: The left starting index
//		mid: The mid-point in the search
//		end: The right ending index
//	 foldContext: The FoldingContext for this sequence
//		helix: Whether this multibranch is enclosed by a helix
//		helix: Whether pairedMinimumFreeEnergyV(start, end) bond with one another in a helix
//
// Returns a multi-branch structure
func multibranch(start, mid, end int, foldContext context, helix bool) (nucleicAcidStructure, error) {
	_ = "STUB: not implemented"
	return *new(nucleicAcidStructure), nil
}

// gather all branches of this multi-branch structure

// in python this was a recursive closure, in Go this is not possible so
// we pull it out and pass all the parameters

// this isn't multi-branched

// if there's a helix, start,end counts as well

// count up unpaired bp and asymmetry

// add foldEnergy from unpaired bp to the right
// of the helix as though it was a dangling end
// if there's only one bp, it goes to whichever
// helix (this or the next) has the more favorable energy

// pass

// add energy

// this is just for readability of the formulas below

// penalty for unmatched bp and multi-branch

// energy of min-energy neighbors

// pointer to next structures

// branches.pop()

// internalLoop calculates the free energy of an internal loop.
//
// The first and last bp of both left and right sequences
// are not themselves parts of the loop, but are the terminal
// bp on either side of it. They are needed for when there's
// a single internal looping bp (where just the mismatching
// free energies are used)
// Note that both left and right sequences are in 5' to 3' direction
// This is adapted from the "Internal Loops" section of SantaLucia/Hicks, 2004
// Args:
//
//		start:  The index of the start of structure on left side
//	rightOfStart: The index to the right of start
//		end:  The index of the end of structure on right side
//		leftOfEnd: The index to the left of end
//	 foldContext: The FoldingContext for this sequence
//
// Returns the free energy associated with the internal loop
func internalLoop(start, rightOfStart, end, leftOfEnd int, foldContext context) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// single bp mismatch, sum up the two single mismatch pairs

// apply a penalty based on loop size

// it's too large an internal loop, extrapolate

// apply an asymmetry penalty

// apply penalty based on the mismatching pairs on either side of the loop

// stack returns the free energy of a stack.
//
// Using the indexes start and end, check whether it's at the end of
// the sequence or internal. Then check whether it's a match
// or mismatch, and return.
// Two edge-cases are terminal mismatches and dangling ends.
// The energy of a dangling end is added to the energy of a pair
// where start XOR end is at the sequence's end.
// Args:
//
//		start: The start index on left side of the pair/stack
//	rightOfStart: The index to the right of start
//		end: The end index on right side of the pair/stack
//		leftOfEnd: The index to the left of end
//	 foldContext: The FoldingContext for this sequence
//
// Returns the free energy of the nearestNeighbors pairing
func stack(start, rightOfStart, end, leftOfEnd int, foldContext context) float64 {
	_ = "STUB: not implemented"
	// if any(x >= len(seq) for x in [start,rightOfStart, end, leftOfEnd]):
	//    return 0.0
	return 0
}

// if any(x == -1 for x in [start,rightOfStart, end, leftOfEnd]):

// it's a dangling end

// it's internal

// it's terminal

// it's dangling on left

// it's dangling on right

// hairpin calculates the free energy of a hairpin.
// Args:
//
//		start:  The index of start of hairpin
//		end:  The index of end of hairpin
//	 foldContext: The FoldingContext for this sequence
//
// Returns the free energy increment from the hairpin structure
func hairpin(start, end int, foldContext context) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// not known terminal pair, nothing to close "hairpin"

// it's a pre-known hairpin with known value

// add penalty based on size

// it's too large, extrapolate

// add penalty for a terminal mismatch

// add penalty if length 3 and AT closing, formula 8 from SantaLucia, 2004

// convert to entropy

// Find the free energy given delta h, s and temp
// Args:
//
//	enthalpyHDifference: The enthalpy increment in kcal / mol
//	entropySDifference: The entropy increment in cal / mol
//	temp: The temperature in Kelvin
//
// Returns the free energy increment in kcal / (mol x K)
func deltaG(enthalpyHDifference, entropySDifference, temp float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

// jacobsonStockmayer entropy extrapolation formula is used for bulges,
// hairpins, etc that fall outside the maxLenPreCalulated upper limit for
// pre-calculated free-energies. See SantaLucia and Hicks (2004)
//
// Args:
//
//	queryLen: Length of element without known free energy value
//	knownLen: Length of element with known free energy value (dGx)
//	dGx: The free energy of the element knownLen
//	temp: Temperature in Kelvin
//
// Returns the free energy for a structure of length queryLen
// See SantaLucia and Hicks (2004).
// NOTE: the coefficient 2.44 is based on recent kinetics
// measurements in DNA (Goddard NL, Bonnet G, Krichevsky O, Libchaber A. 2000),
// and thus it is used in preference to the older theoretically derived value
// of 1.75.
func jacobsonStockmayer(queryLen, knownLen int, dGx, temp float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

// pair Returns a stack representation, a key for the nearestNeighbors maps
// Args:
//
//	s: Sequence being folded
//	start: leftmost index
//	rightOfStart: index to right of start
//	end: rightmost index
//	leftOfEnd: index to left of end
//
// Returns string representation of the pair
func pair(s string, start, rightOfStart, end, leftOfEnd int) string {
	_ = "STUB: not implemented"
	return ""
}

// Traceback thru the pairedMinimumFreeEnergyV(start,end) and unpairedMinimumFreeEnergyW(start,end) caches to find the structure
// For each step, get to the lowest energy unpairedMinimumFreeEnergyW(start,end) within that block
// Store the structure in unpairedMinimumFreeEnergyW(start,end)
// Inc start and end
// If the next structure is viable according to pairedMinimumFreeEnergyV(start,end), store as well
// Repeat
// Args:
//
//		start: The leftmost index to start searching in
//		end: The rightmost index to start searching in
//	 foldContext: The FoldingContext for this sequence
//
// Returns a list of NucleicAcidStructure in the final secondary structure
func traceback(start, end int, foldContext context) []nucleicAcidStructure {
	_ = "STUB: not implemented"
	// move start,end down-left to start coordinates
	return nil
}

// it's a hairpin, end of structure

// set the energy of everything relative to the hairpin

// it's a stack, bulge, etc
// there's another single structure beyond this

// subsequence = structure.IJs[0]

// it's a multibranch

// Return the struct with the lowest free energy that isn't -inf
// Args:
//
//	structures: NucleicAcidStructure being compared
//
// Returns the min free energy structure
func minimumStructure(structures ...nucleicAcidStructure) nucleicAcidStructure {
	_ = "STUB: not implemented"
	return *new(nucleicAcidStructure)
}

// trackbackEnergy add energy to each structure, based on how it's
// unpairedMinimumFreeEnergyW(start,end) differs from the one after
// Args:
//
//	structures: The NucleicAcidStructure for whom energy is being calculated
//
// Returns a slice of NucleicAcidStructure in the folded DNA with energy
func trackbackEnergy(structures []nucleicAcidStructure) []nucleicAcidStructure {
	_ = "STUB: not implemented"
	return nil
}
