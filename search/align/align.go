/*
Package align is a package for aligning (comparing) DNA, RNA, and protein sequences.

Biology is fickle and full of quirks that make it hard to do even the most basic of tasks
which we would normally take for granted when working with other kinds of data.

Comparing two biological sequences to see if they're roughly equivalent is one of those tasks.

Essentially two almost identical sequences with almost identical functionality can contain
small insertions or deletions that shift the entire string such that a meaningful comparison via
hamming distance or levenshtein distance becomes impossible.

For example:

Timothy Stiles
||||||| ||||||
Timothy Stiles

is an easy match with hamming or levenshtein distance.

However, say we introduce a new character to the beginning of the sequence.

Timothy Stiles
xxxxxxxxxxxxxxxx
A Timothy Stiles

Now our edit distance via levenshtein is maximized at 16 and we wouldn't
be able to tell that semantically these strings are almost identical.

This frame shifting seen above is incredibly common within biological sequences and alignment
algorithms are designed in part to deal with these shifts so that when we compare two sequences
like the two below we can get a more useful edit distance.

GAAAAAAT
GAA----T

As of writing this package includes the two most basic algorithms for alignment,
Needleman-Wunsch and Smith-Waterman. Needleman-Wunsch is used when you are
looking for global alignment between two full-length sequences, while
Smith-Waterman is better at smaller sequences with local similarities and
handling sequences with long non-homologous regions. BLAST, on the other hand,
takes advantage of more heuristic techniques to speed up alignment, and is better
at finding similar sequences in large database, sacrificing precision for faster
results.

Both are "dynamic programming algorithms" which is a fancy 1980's term for they use
matrices. If you're familiar with kernel operations, linear filters, or whatever term
ML researchers are using nowadays for, "slide a window over a matrix and determine that
entry's values using its neighbor's values", then this should be pretty easy to grok.

If not these algorithms essentially compare every character in one sequence with another
sequence and create an edit distance along with human readable string to show gaps like the
previous example.

I'm not really an expert on alignment so if you want to learn more about this class of algorithms
wikipedia has a decent overview.

https://en.wikipedia.org/wiki/Sequence_alignment

Even if I may not know the answer to your alignment questions please ask and I'll do my best
to help!

TTFN,
Tim
*/
package align

import (
	"github.com/bebop/poly/search/align/matrix"
)

// Scoring is a struct that holds the scoring matrix for match, mismatch, and gap penalties.
type Scoring struct {
	SubstitutionMatrix *matrix.SubstitutionMatrix
	GapPenalty         int
}

// NewScoring returns a new Scoring struct with default values for DNA.
func NewScoring(substitutionMatrix *matrix.SubstitutionMatrix, gapPenalty int) (Scoring, error) {
	_ = "STUB: not implemented"
	return *new(Scoring), nil
}

func (s Scoring) Score(a, b byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// NeedlemanWunsch performs global alignment between two strings using the Needleman-Wunsch algorithm.
// It returns the final score and the optimal alignments of the two strings in O(nm) time and O(nm) space.
// https://en.wikipedia.org/wiki/Needleman-Wunsch_algorithm
func NeedlemanWunsch(stringA string, stringB string, scoring Scoring) (int, string, string, error) {
	_ = "STUB: not implemented"
	// Get the M and N dimensions of the matrix. The M x N matrix is standard linear algebra notation.
	// But I added columns and rows to the variable name to make it more clear what the dimensions are.
	return 0, "", "", nil
}

// Initialize the matrix columns
// matrix is a 2D slice of ints.

// Initialize the matrix rows.

// Fill in the first column with gap penalties.

// Fill in the first row with gap penalties.

// Fill in the rest of the matrix.

// Calculate the scores for scoring.Match/mismatch and gap.

// Traceback to find the optimal alignment.

// Reverse the alignments to get the optimal alignment.

// SmithWaterman performs local alignment between two strings using the Smith-Waterman algorithm.
// It returns the max score and optimal local alignments between two strings alignments of the two strings in O(nm) time and O(nm) space.
// https://en.wikipedia.org/wiki/Smith-Waterman_algorithm
func SmithWaterman(stringA string, stringB string, scoring Scoring) (int, string, string, error) {
	_ = "STUB: not implemented"
	return 0, "", "", nil
}

// Initialize the alignment matrix

// Initialize variables to keep track of the maximum score and its position

// Fill the alignment matrix

// Traceback to construct the aligned strings

func reverseRuneArray(runes []rune) []rune {
	_ = "STUB: not implemented" // wasn't able to find a built-in reverse function for runes
	return nil
}

func max(a, b int) int {
	_ = "STUB: not implemented" // funny enough Go's built-in max only handles floats?
	return 0
}
