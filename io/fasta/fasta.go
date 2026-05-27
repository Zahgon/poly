/*
Package fasta contains fasta parsers and writers.

Fasta is a flat text file format developed in 1985 to store nucleotide and
amino acid sequences. It is extremely simple and well-supported across many
languages. However, this simplicity means that annotation of genetic objects
is not supported.

This package provides a parser and writer for working with Fasta formatted
genetic sequences.
*/
package fasta

import (
	"bufio"
	"compress/gzip"
	"io"
	"os"
)

/******************************************************************************
Apr 25, 2021

Fasta Parser begins here

Many thanks to Jordan Campbell (https://github.com/0x106) for building the first
parser for Poly and thanks to Tim Stiles (https://github.com/TimothyStiles)
for helping complete that PR. This work expands on the previous work by allowing
for concurrent  parsing and giving Poly a specific  parser subpackage,
as well as few bug fixes.

Fasta is a very simple file format for working with DNA, RNA, or protein sequences.
It was first released in 1985 and is still widely used in bioinformatics.

https://en.wikipedia.org/wiki/_format

One interesting use of the concurrent  parser is working with the Uniprot
fasta dump files, which are far too large to fit into RAM. This parser is able
to easily handle those files by doing computation actively while the data dump
is getting parsed.

https://www.uniprot.org/downloads

I have removed the  Parsers from the io.go file and moved them into this
subpackage.

Hack the Planet,

Keoni

******************************************************************************/

var (
	gzipReaderFn = gzip.NewReader
	openFn       = os.Open
	buildFn      = Build
)

// Fasta is a struct representing a single Fasta file element with a Name and its corresponding Sequence.
type Fasta struct {
	Name     string `json:"name"`
	Sequence string `json:"sequence"`
}

// Parse parses a given Fasta file into an array of Fasta structs. Internally, it uses ParseFastaConcurrent.
func Parse(r io.Reader) ([]Fasta, error) {
	_ = "STUB: not implemented"
	// 32kB is a magic number often used by the Go stdlib for parsing. We multiply it by two.
	return nil, nil
}

// Parser is a flexible parser that provides ample
// control over reading fasta-formatted sequences.
// It is initialized with NewParser.
type Parser struct {
	// reader keeps state of current reader.
	reader bufio.Reader
	line   uint
}

// NewParser returns a Parser that uses r as the source
// from which to parse fasta formatted sequences.
func NewParser(r io.Reader, maxLineSize int) *Parser { _ = "STUB: not implemented"; return nil }

// ParseAll parses all sequences in underlying reader only returning non-EOF errors.
// It returns all valid fasta sequences up to error if encountered.
func (parser *Parser) ParseAll() ([]Fasta, error) { _ = "STUB: not implemented"; return nil, nil }

// ParseN parses up to maxSequences fasta sequences from the Parser's underlying reader.
// ParseN does not return EOF if encountered.
// If an non-EOF error is encountered it returns it and all correctly parsed sequences up to then.
func (parser *Parser) ParseN(maxSequences int) (fastas []Fasta, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EOF not treated as parsing error.

// ParseByteLimited parses fastas until byte limit is reached.
// This is NOT a hard limit. To set a hard limit on bytes read use a
// io.LimitReader to wrap the reader passed to the Parser.
func (parser *Parser) ParseByteLimited(byteLimit int64) (fastas []Fasta, bytesRead int64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// EOF not treated as parsing error.

// ParseNext reads next fasta genome in underlying reader and returns the result
// and the amount of bytes read during the call.
// ParseNext only returns an error if it:
//   - Attempts to read and fails to find a valid fasta sequence.
//   - Returns reader's EOF if called after reader has been exhausted.
//   - If a EOF is encountered immediately after a sequence with no newline ending.
//     In this case the Fasta up to that point is returned with an EOF error.
//
// It is worth noting the amount of bytes read are always right up to before
// the next fasta starts which means this function can effectively be used
// to index where fastas start in a file or string.
func (parser *Parser) ParseNext() (Fasta, int64, error) {
	_ = "STUB: not implemented"
	return *new(Fasta), 0, nil
}

// Early return on error. Probably will be EOF.

// Initialization of parser state variables.

// Parser looks for a line starting with '>' (U+003E)
// that contains the next fasta sequence name.

// parse loop begins here.

// OR short circuits so no panic here.

// More general case of error handling.

// got EOF on a empty or commented line.

// Buffer size too small to read fasta line.

// Unexpected error.

// So got to this point the line is probably OK, we will return a Fasta.
// with the EOF error.

// Exclude newline delimiter.

// We are currently parsing a fasta and next line contains a new fasta.
// We handle this situation by appending current line to sequence if not a comment
// and ending the current fasta parsing.

// We got the start of a fasta.

// This continue will also skip line if we are looking for name
// and the current line does not contain the name.

// If we got to this point we are currently inside of the fasta
// sequence contents. We append line to what we found of sequence so far.

// parse loop ends here.

// Parsing ended. Check for inconsistencies.

// We found a fasta name but no sequence to go with it.

// Stdlib strings.Builder.String() does this so it *should* be safe.

// Gotten to this point err is non-nil only in EOF case.
// We report this error to note the fasta may be incomplete/corrupt
// like in the case of using an io.LimitReader wrapping the underlying reader.
// We return the fasta as well since some libraries generate fastas with no
// ending newline i.e Zymo. It is up to the user to decide whether they want
// an EOF-ended fasta or not, the rest of this library discards EOF-ended fastas.

// Reset discards all data in buffer and resets state.
func (parser *Parser) Reset(r io.Reader) { _ = "STUB: not implemented"; return }

// ParseConcurrent concurrently parses a given Fasta file in an io.Reader into a channel of Fasta structs.
func ParseConcurrent(r io.Reader, sequences chan<- Fasta) {
	_ = "STUB: not implemented"
	// Initialize necessary variables
	return
}

// Start the scanner

// if there's nothing on this line skip this iteration of the loop

// if it's a comment skip this line

// start of a fasta line

// Process normal new lines

// Reset sequence lines

// New name

// Process first line of file

// Add final sequence in file to channel

/******************************************************************************

Start of  Read functions

******************************************************************************/

// ReadGzConcurrent concurrently reads a gzipped Fasta file into a Fasta channel.
// Deprecated: Use Parser.ParseNext() instead.
func ReadGzConcurrent(path string, sequences chan<- Fasta) { _ = "STUB: not implemented"; return }

// TODO: these errors need to be handled/logged

// ReadConcurrent concurrently reads a flat Fasta file into a Fasta channel.
func ReadConcurrent(path string, sequences chan<- Fasta) { _ = "STUB: not implemented"; return }

// TODO: these errors need to be handled/logged

// ReadGz reads a gzipped  file into an array of Fasta structs.
func ReadGz(path string) ([]Fasta, error) { _ = "STUB: not implemented"; return nil, nil }

// Read reads a  file into an array of Fasta structs
func Read(path string) ([]Fasta, error) { _ = "STUB: not implemented"; return nil, nil }

/******************************************************************************

Start of  Write functions

******************************************************************************/

// Build converts a Fastas array into a byte array to be written to a file.
func Build(fastas []Fasta) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// write the fasta sequence 80 characters at a time

// Write writes a fasta array to a file.
func Write(fastas []Fasta, path string) error { _ = "STUB: not implemented"; return nil }

//  fasta.Build returns only nil errors.
