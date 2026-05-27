/*
Package pcr designs and simulates simple PCR reactions.

PCR, or polymerase chain reaction, is a method developed in 1983 to copy DNA
templates using small fragments of synthesized single-stranded DNA, amplifying
those DNA templates to ~x1,000,000,000 their starting concentration. These small
fragments, referred to as "primers" or "oligos", can be designed on a computer
and then synthesized for amplifying a variety of different templates.

This package allows users to simulate a PCR reaction or design new primers to
amplify a given template. This package assumes perfect annealing to template at
a target temperature, so should only be used for PCR reactions where this is a
reasonable assumption.

If you are trying to simulate amplification out of a large pool, such as an
oligo pool, use the `Simulate` rather than `SimulateSimple` function to detect
if there is concatemerization happening in your multiplex reaction. In most
other cases, use `SimulateSimple`.

IMPORTANT! The targetTm in all functions is specifically for Taq polymerase.
*/
package pcr

// https://doi.org/10.1089/dna.1994.13.75
const minimalPrimerLength int = 7

// what we want for designs
const designedMinimalPrimerLength int = 15

// DesignPrimersWithOverhangs designs two primers to amplify a target sequence,
// adding on an overhang to the forward and reverse strand. This overhang can
// contain additional DNA needed for assembly, like Gibson assembly overhangs
// or GoldenGate restriction enzyme sites.
func DesignPrimersWithOverhangs(sequence, forwardOverhang, reverseOverhang string, targetTm float64) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// Add overhangs to primer

// DesignPrimers designs two primers to amplify a target sequence and only that
// target sequence (no overhangs).
func DesignPrimers(sequence string, targetTm float64) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// SimulateSimple simulates a PCR reaction. It takes in a list of sequences and
// a list of primers, with support for complex multiplex reactions, produces
// a list of all possible PCR fragments from such a reaction. It does not
// detect concatemerization, which could be useful or very detrimental to
// your reactions. The variable `circular` is for if the target template is
// circular, like a plasmid.
func SimulateSimple(sequences []string, targetTm float64, circular bool, primerList []string) []string {
	_ = "STUB: not implemented"
	// Set all primers to uppercase.
	return nil
}

// Suffix array construction allows function to operate on
// very large sequences without being worried about exeuction
// time. For small sequences, it doesn't really matter.
// https://eli.thegreenplace.net/2016/suffix-arrays-in-the-go-standard-library/

// Use the minimal binding sites of the primer to find positions in the template

// For each primer, we want to look for all possible binding sites in our gene.
// We then append this to a list of binding sites for that primer.

// Next, iterate through the forwardLocations list

// First, make sure that this isn't the last element in forwardLocations

// If this isn't the last element in forwardLocations, then we can select the first reverseLocation that is less than the next forwardLocation

// If both are true, we have found the sequence we are aiming to PCR! Now, we get all primers from that forwardLocation and then
// build PCR fragments with each one.

// If the sequence is circular and we haven't found a fragment yet, check the other side of the origin

// If either one of these are true, create a new pcrFragment and append to pcrFragments

// Simulate simulates a PCR reaction, including concatemerization analysis. It
// takes in a list of sequences and list of primers, produces all possible PCR
// fragments in a given reaction, and then attempts to see if the output
// fragments can amplify themselves. If they can, concatemerization is occurring
// in your reaction, which can lead to confusing results. The variable
// `circular` is for if the target template is circular, like a plasmid.
func Simulate(sequences []string, targetTm float64, circular bool, primerList []string) ([]string, error) {
	_ = "STUB: not implemented"
	// make sure no primers are too short
	return nil, nil
}

func generatePcrFragments(sequence string, forwardLocation int, reverseLocation int, forwardPrimerIndxs []int, reversePrimerIndxs []int, minimalPrimers []string, primerList []string) []string {
	_ = "STUB: not implemented"
	return nil
}
