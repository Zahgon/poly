/*
Package pileup contains pileup parsers and writers.

The pileup format is a text-based bioinformatics format to summarize aligned
reads against a reference sequence. In comparison to simply getting a consensus
sequence from sequencing data, pileup files can contain more context about the
mutations in a sequencing run, which is especially useful when analyzing
plasmid sequencing data from Nanopore sequencing runs.

Pileup files are basically tsv files with 6 columns: Sequence Identifier, Position,
Reference Base, Read Count, Read Results, and Quality. An example from
wikipedia (https://en.wikipedia.org/wiki/Pileup_format) is shown below:

	```
	seq1 	272 	T 	24 	,.$.....,,.,.,...,,,.,..^+. 	<<<+;<<<<<<<<<<<=<;<;7<&
	seq1 	273 	T 	23 	,.....,,.,.,...,,,.,..A 	<<<;<<<<<<<<<3<=<<<;<<+
	seq1 	274 	T 	23 	,.$....,,.,.,...,,,.,... 	7<7;<;<<<<<<<<<=<;<;<<6
	seq1 	275 	A 	23 	,$....,,.,.,...,,,.,...^l. 	<+;9*<<<<<<<<<=<<:;<<<<
	seq1 	276 	G 	22 	...T,,.,.,...,,,.,.... 	33;+<<7=7<<7<&<<1;<<6<
	seq1 	277 	T 	22 	....,,.,.,.C.,,,.,..G. 	+7<;<<<<<<<&<=<<:;<<&<
	seq1 	278 	G 	23 	....,,.,.,...,,,.,....^k. 	%38*<<;<7<<7<=<<<;<<<<<
	seq1 	279 	C 	23 	A..T,,.,.,...,,,.,..... 	75&<<<<<<<<<=<<<9<<:<<<
	```

	1. Sequence Identifier: The sequence identifier of the reference sequence
	2. Position: Position of row in the reference sequence (indexed at 1)
	3. Reference Base: Base pair in reference sequence
	4. Read Count: Number of aligned reads to this particular base pair
	5. Read Results: The resultant alignments
	6. Quality: Phred quality scores associated with each base

This package provides a parser and writer for working with pileup files.
*/
package pileup

import (
	"bufio"
	"io"
)

// https://en.wikipedia.org/wiki/Pileup_format

// Pileup struct is a single position in a pileup file. Pileup files "pile"
// a bunch of separate bam/sam alignments into something more readable at a per
// base pair level, so are only useful as a grouping.
type Pileup struct {
	Sequence      string   `json:"sequence"`
	Position      uint     `json:"position"`
	ReferenceBase string   `json:"reference_base"`
	ReadCount     uint     `json:"read_count"`
	ReadResults   []string `json:"read_results"`
	Quality       string   `json:"quality"`
}

// Parse parses a given Pileup file into an array of Pileup structs.
func Parse(r io.Reader) ([]Pileup, error) {
	_ = "STUB: not implemented"
	// 32kB is a magic number often used by the Go stdlib for parsing. We multiply it by two.
	return nil, nil
}

// Parser is a pileup parser.
type Parser struct {
	reader bufio.Reader
	line   uint
}

// NewParser creates a parser from an io.Reader for pileup data.
func NewParser(r io.Reader, maxLineSize int) *Parser { _ = "STUB: not implemented"; return nil }

// ParseAll parses all sequences in underlying reader only returning non-EOF errors.
// It returns all valid pileup sequences up to error if encountered.
func (parser *Parser) ParseAll() ([]Pileup, error) { _ = "STUB: not implemented"; return nil, nil }

// ParseN parses up to maxRows pileup sequences from the Parser's underlying reader.
// ParseN does not return EOF if encountered.
// If an non-EOF error is encountered it returns it and all correctly parsed sequences up to then.
func (parser *Parser) ParseN(maxRows int) (pileups []Pileup, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EOF not treated as parsing error.

// ParseNext parses the next pileup row in a pileup file.
// ParseNext returns an EOF if encountered.
func (parser *Parser) ParseNext() (Pileup, error) {
	_ = "STUB: not implemented"
	return *new(Pileup), nil
}

// Early return on error. Probably will be EOF.

// Parse out a single line

// Exclude newline delimiter.

// Check that there are 6 values, as defined by the pileup format

// Convert Position and ReadCount to integers

// Parse ReadResults

// This applies to the last read segement

// formatted in `+4ATGC` format. We need to know the number of jumps
// because you can have +10AAAAAAAAAA

// Because of the above check, this will never err

// The 1 makes sure to include the regularExpressionInt in readResult string

// Reset discards all data in buffer and resets state.
func (parser *Parser) Reset(r io.Reader) { _ = "STUB: not implemented"; return }

/******************************************************************************

Start of  Read functions

******************************************************************************/

// Read reads a  file into an array of Pileup structs
func Read(path string) ([]Pileup, error) { _ = "STUB: not implemented"; return nil, nil }

/******************************************************************************

Start of  Write functions

******************************************************************************/

// WritePileups writes a pileup array to an io.Writer
func WritePileups(pileups []Pileup, w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Write writes a pileup array to a file
func Write(pileups []Pileup, path string) error { _ = "STUB: not implemented"; return nil }
