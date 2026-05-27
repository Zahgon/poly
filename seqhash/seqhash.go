/*
Package seqhash contains the seqhash algorithm.

This package contains the reference seqhash algorithm.

There is a big problem with current sequence databases - they all use different
identifiers and accession numbers. This means cross-referencing databases is
a complicated exercise, especially as the quantity of databases increases, or if
you need to compare "wild" DNA sequences.

Seqhash is a simple algorithm to produce consistent identifiers for any genetic sequence. The
basic premise of the Seqhash algorithm is to hash sequences with the hash being a robust
cross-database identifier. Sequences themselves shouldn't be used as a database index
(often, they're too big), so a hash based off of a sequence is the next best thing.

Usability wise, you should be able to Seqhash any rotation of a sequence in any direction and
get a consistent hash.

The Seqhash algorithm makes several opinionated design choices, primarily to make working
with Seqhashes more consistent and nice. The Seqhash algorithm only uses a single hash function,
Blake3, and only operates on DNA, RNA, and Protein sequences. These identifiers will be seen
by human beings, so versioning and metadata is attached to the front of the hashes so that
a human operator can quickly identify problems with hashing.

If the sequence is DNA or RNA, the Seqhash algorithm needs to know whether or not the nucleic
acid is circular and/or double stranded. If circular, the sequence is rotated to a deterministic
point. If double stranded, the sequence is compared to its reverse complement, and the lexicographically
minimal sequence is taken (whether or not the min or max is used doesn't matter, just needs to
be consistent).

If the sequence is RNA, the sequence will be converted to DNA before hashing. While the full Seqhash
will still be different between RNA and DNA (due to the metadata string), the hash afterwards will be the same.
This makes it easy to cross reference DNA and RNA sequences. This fact is important for parts of Poly
store that relate to storing and searching large quantities of sequences - deduplication can easily
be used on those Seqhashes to save a lot of space.

For DNA or RNA sequences, only ATUGCYRSWKMBDHVNZ characters are allowed. For Proteins,
only ACDEFGHIKLMNPQRSTVWYUO*BXZ characters are allowed in sequences. Selenocysteine (Sec; U) and pyrrolysine
(Pyl; O) are included in the protein character set - usually U and O don't occur within protein sequences,
but for certain organisms they do, and it is certainly a relevant amino acid for those particular proteins.

A Seqhash is separated into 3 different elements divided by underscores. It looks like the following:

v1_DCD_4b0616d1b3fc632e42d78521deb38b44fba95cca9fde159e01cd567fa996ceb9

The first element is the version tag (v1 for version 1). If there is ever a Seqhash version 2, this tag
will differentiate seqhashes. The second element is the metadata tag, which has 3 letters. The first letter
codes for the sequenceType (D for DNA, R for RNA, and P for Protein). The second letter codes for whether or
not the sequence is circular (C for Circular, L for Linear). The final letter codes for whether or not the
sequence is double stranded (D for Double stranded, S for Single stranded). The final element is the blake3
hash of the sequence (once rotated and complemented, as stated above).

Seqhash is a simple algorithm that allows for much better indexing of genetic sequences than what is
currently available.
*/
package seqhash

// Seqhash is a struct that contains the Seqhash algorithm sequence types.
type SequenceType string

const (
	DNA     SequenceType = "DNA"
	RNA     SequenceType = "RNA"
	PROTEIN SequenceType = "PROTEIN"
)

// boothLeastRotation gets the least rotation of a circular string.
func boothLeastRotation(sequence string) int {
	_ = "STUB: not implemented"
	// https://en.wikipedia.org/wiki/Lexicographically_minimal_string_rotation
	// this is generally over commented but I'm keeping it this way for now. - Tim
	return 0
}

// first concatenate the sequence to itself to avoid modular arithmetic
// maybe do this as a buffer just for speed? May get annoying with larger sequences.

//initializing failure slice.

// iterate through each character in the doubled over sequence

// get character

// get failure

// while failure does not equal -1 and character does not equal the character found at the least rotation + failure + 1 <- why this?

// if character is lexically less than whatever is at the leastRotationIndex index update leastRotation index

// update failure using previous failure as index?

// if character does not equal whatever character is at leastRotationIndex plus failure.

// if character is lexically less then what is rotated least leastRotationIndex gets value of character index.

// assign -1 to whatever is at the index of difference between character and rotation indices.

// if character does equal whatever character is at leastRotationIndex plus failure.

// assign failure + 1 at the index of difference between character and rotation indices.

// end loop

// RotateSequence rotates circular sequences to deterministic point.
func RotateSequence(sequence string) string { _ = "STUB: not implemented"; return "" }

// writing the same sequence twice. using build incase of very long circular genome.

// Hash is a function to create Seqhashes, a specific kind of identifier.
func Hash(sequence string, sequenceType SequenceType, circular bool, doubleStranded bool) (string, error) {
	_ = "STUB: not implemented"
	// By definition, Seqhashes are of uppercase sequences
	return "", nil
}

// If RNA, convert to a DNA sequence. The hash itself between a DNA and RNA sequence will not
// be different, but their Seqhash will have a different metadata string (R vs D)

// Run checks on the input

// Selenocysteine (Sec; U) and pyrrolysine (Pyl; O) are added
// in accordance with https://www.uniprot.org/help/sequences
// The release notes https://web.expasy.org/docs/relnotes/relstat.html
// also state there are Asx (B), Glx (Z), and Xaa (X) amino acids, so
// these are added in as well.

// There is no check for circular proteins since proteins can be circular

// Gets Deterministic sequence based off of metadata + sequence

// Build 3 letter metadata

// Get first letter. D for DNA, R for RNA, and P for Protein

// Get 2nd letter. C for circular, L for Linear

// Get 3rd letter. D for Double stranded, S for Single stranded
