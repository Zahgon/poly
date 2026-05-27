/*
Package clone provides functions for cloning DNA sequences.

Since 1973, the most common way to make recombinant DNA has been restriction
enzyme cloning (though lately, homologous recombination based methods like
Gibson assembly have attracted a lot of use). The cloning functions here allow
for simulation of restriction enzyme cloning.

For a historical review leading up to the discovery:
https://doi.org/10.1073/pnas.1313397110

The idea of restriction enzyme cloning is that you can cut DNA at specific
locations with restriction enzyme and then glue them back together in different
patterns using ligase. The final product is (99.9% of the time) a circular plasmid
that you can transform into a bacterial cell for propagation.

While simulation is simple for simple cases, there are a lot of edge cases to handle, for example:
  - Which input sequences are circular? How do we handle their rotations?
  - Is the enzyme that is cutting directional? How do we handle that directionality?
  - Are there multiple possible outputs of our ligation reaction? For example, ligations may be
    able to create a "library" of plasmids, in which there are millions of valid combinations.
  - How do we handle sequences that get ligated in multiple orientations?

These cloning functions handle all those problems so that they appear simple to the end user.

In particular, there is a focus here on GoldenGate Assembly:
https://en.wikipedia.org/wiki/Golden_Gate_Cloning
https://www.neb.com/applications/cloning-and-synthetic-biology/dna-assembly-and-cloning/golden-gate-assembly

GoldenGate is a particular kind of restriction enzyme cloning reaction that you can do
in a single tube and that is extraordinarily efficient (up to 50 parts) and is popular
for new modular DNA part toolkits. Users can easily simulate GoldenGate assembly reactions
with just their input fragments + the enzyme name.

Let's build some DNA!

# Keoni

PS: We do NOT (yet) handle restriction enzymes which recognize one site but cut
in multiple places (Type IIG enzymes) such as BcgI.
*/
package clone

import (
	"regexp"
)

// Part is a simple struct that can carry a circular or linear DNA sequence.
// In the field of synthetic biology, the term "DNA Part" was popularized by
// the iGEM competition http://parts.igem.org/Main_Page , so we use that term
// here.
type Part struct {
	Sequence string
	Circular bool
}

// Overhang is a struct that represents the ends of a linearized sequence where Enzymes had cut.
type Overhang struct {
	Length                        int
	Position                      int
	Forward                       bool
	RecognitionSitePlusSkipLength int
}

// Fragment is a struct that represents linear DNA sequences with sticky ends.
type Fragment struct {
	Sequence        string
	ForwardOverhang string
	ReverseOverhang string
}

// Enzyme is a struct that represents restriction enzymes.
type Enzyme struct {
	Name            string
	RegexpFor       *regexp.Regexp
	RegexpRev       *regexp.Regexp
	Skip            int
	OverheadLength  int
	RecognitionSite string
}

// EnzymeManager manager for Enzymes. Allows for management of enzymes throughout the lifecyle of your
// program. EnzymeManager is not safe for concurrent use.
type EnzymeManager struct {
	// enzymeMap Map of enzymes that exist for the lifetime of the manager. Not safe for concurrent use.
	enzymeMap map[string]Enzyme
}

// NewEnzymeManager creates a new EnzymeManager given some enzymes.
func NewEnzymeManager(enzymes []Enzyme) EnzymeManager {
	_ = "STUB: not implemented"
	return *new(EnzymeManager)
}

/******************************************************************************

Base cloning functions begin here.

******************************************************************************/

// CutWithEnzymeByName cuts a given sequence with an enzyme represented by the
// enzyme's name. It is a convenience wrapper around CutWithEnzyme that
// allows us to specify the enzyme by name.
func (enzymeManager EnzymeManager) CutWithEnzymeByName(part Part, directional bool, name string) ([]Fragment, error) {
	_ = "STUB: not implemented"
	// Get the enzyme from the enzyme map
	return nil, nil
}

// Return an error if there was an error

// Cut the sequence with the enzyme

// GetEnzymeByName gets the enzyme by it's name. If the enzyme manager does not
// contain an enzyme with the provided name, an error will be returned
func (enzymeManager EnzymeManager) GetEnzymeByName(name string) (Enzyme, error) {
	_ = "STUB: not implemented"
	return *new(Enzyme), nil
}

// CutWithEnzyme cuts a given sequence with an enzyme represented by an Enzyme struct.
func CutWithEnzyme(part Part, directional bool, enzyme Enzyme) []Fragment {
	_ = "STUB: not implemented"
	return nil
}

// Check for palindromes

// Find and define overhangs

// Palindromic enzymes won't need reverseCuts

// If, on a linear sequence, the last overhang's position + EnzymeSkip + EnzymeOverhangLength is over the length of the sequence, remove that overhang.

// Sort overhangs

// Convert Overhangs into Fragments

// Linear fragments with 1 cut that are no directional will always give a
// 2 fragments
// Check the case of a single cut
// In the case of a single cut in a linear sequence, we get two fragments with only 1 stick end

// Circular fragments with 1 cut will always have 2 overhangs (because of the
// concat earlier). If we don't require directionality, this will always get
// cut into a single fragment

// In the case of a single cut in a circular sequence, we get one fragment out with sticky overhangs

// The following will iterate over the overhangs list to turn them into fragments
// There are two important variables: if the sequence is circular, and if the enzyme cutting is directional. All Type IIS enzymes
// are directional, and in normal GoldenGate reactions these fragments would be constantly cut with enzyme as the reaction runs,
// so are removed from the output sequences. If the enzyme is not directional, all fragments are valid.
// If the sequence is circular, there is a chance that the nextOverhang's position will be greater than the length of the original sequence.
// This is ok, and represents a valid cut/fragmentation of a rotation of the sequence. However, everything after will be a repeat fragment
// of current fragments, so the iteration is terminated.

// If we want directional cutting and the enzyme is not palindromic, we
// can remove fragments that are continuously cut by the enzyme. This is
// the basis of GoldenGate assembly.

// We have to subtract RecognitionSitePlusSkipLength in case we have a recognition site on
// one side of the origin of a circular sequence and the cut site on the other side of the origin

// Convert fragment sequences into fragments

// Minimum lengths (given oligos) for assembly is 8 base pairs
// https://doi.org/10.1186/1756-0500-3-291

func recurseLigate(seedFragment Fragment, fragmentList []Fragment, usedFragments []Fragment, existingSeqhashes map[string]struct{}) (openConstructs []string, infiniteConstructs []string) {
	_ = "STUB: not implemented"
	// Recurse ligate simulates all possible ligations of a series of fragments. Each possible combination begins with a "seed" that fragments from the pool can be added to.
	// If the seed ligates to itself, we can call it done with a successful circularization!
	return nil, nil
}

// If the seed ligates to another fragment, we can recurse and add that fragment to the seed

// If the seedFragment's reverse overhang is ligates to a fragment's forward overhang, we can ligate those together and seed another ligation reaction

// This checks if we can ligate the next fragment in its reverse direction. We have to be careful though - if our seed has a palindrome, it will ligate to itself
// like [-> <- -> <- -> ...] infinitely. We check for that case here as well.
// If the second statement isn't there, program will crash on palindromes

// If fragment is actually attached, move to some checks

// If the newFragment's reverse complement already exists in the used fragment list, we need to cancel the recursion.

// If everything is clear, append fragment to usedFragments and recurse.

// CircularLigate simulates ligation of all possible fragment combinations into circular plasmids.
func CircularLigate(fragments []Fragment) ([]string, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

/******************************************************************************

Specific cloning functions begin here.

******************************************************************************/

// GoldenGate simulates a GoldenGate cloning reaction. As of right now, we only
// support BsaI, BbsI, BtgZI, and BsmBI.
func GoldenGate(sequences []Part, cuttingEnzyme Enzyme) (openConstructs []string, infiniteLoops []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBaseRestrictionEnzymes return a basic slice of common enzymes used in Golden Gate Assembly. Eventually, we want to get the data for this map from ftp://ftp.neb.com/pub/rebase
func GetBaseRestrictionEnzymes() []Enzyme { _ = "STUB: not implemented"; return nil }
