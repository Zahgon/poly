/*
Package gff provides gff parsers and writers.

GFF stands for "general feature format". It is an alternative to GenBank for
storing data about genomic sequences. While not often used in synthetic biology
research, it is more commonly used in bioinformatics for digesting features of
genomic sequences.

This package provides a parser and writer to convert between the gff file
format and the more general poly.Sequence struct.
*/
package gff

import (
	"io"
	"os"
	"strconv"
)

var (
	readAllFn = io.ReadAll
	atoiFn    = strconv.Atoi
	openFn    = os.Open
)

// Gff is a struct that represents a gff file.
type Gff struct {
	Meta     Meta
	Features []Feature // will need a GetFeatures interface to standardize
	Sequence string
}

// Meta holds meta information about a gff file.
type Meta struct {
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	Version              string   `json:"gff_version"`
	RegionStart          int      `json:"region_start"`
	RegionEnd            int      `json:"region_end"`
	Size                 int      `json:"size"`
	SequenceHash         string   `json:"sequence_hash"`
	SequenceHashFunction string   `json:"hash_function"`
	CheckSum             [32]byte `json:"checkSum"` // blake3 checksum of the parsed file itself. Useful for if you want to check if incoming genbank/gff files are different.
}

// Feature is a struct that represents a feature in a gff file.
type Feature struct {
	Name           string            `json:"name"`
	Source         string            `json:"source"`
	Type           string            `json:"type"`
	Score          string            `json:"score"`
	Strand         string            `json:"strand"`
	Phase          string            `json:"phase"`
	Attributes     map[string]string `json:"attributes"`
	Location       Location          `json:"location"`
	ParentSequence *Gff              `json:"-"`
}

// Location is a struct that represents a location in a gff file.
type Location struct {
	Start             int        `json:"start"`
	End               int        `json:"end"`
	Complement        bool       `json:"complement"`
	Join              bool       `json:"join"`
	FivePrimePartial  bool       `json:"five_prime_partial"`
	ThreePrimePartial bool       `json:"three_prime_partial"`
	SubLocations      []Location `json:"sub_locations"`
}

// AddFeature takes a feature and adds it to the Gff struct.
func (sequence *Gff) AddFeature(feature *Feature) error { _ = "STUB: not implemented"; return nil }

// GetSequence takes a feature and returns a sequence string for that feature.
func (feature Feature) GetSequence() (string, error) { _ = "STUB: not implemented"; return "", nil }

// getFeatureSequence takes a feature and location object and returns a sequence string.
func getFeatureSequence(feature Feature, location Location) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// reverse complements resulting string if needed.

// Parse Takes in a string representing a gffv3 file and parses it into an Sequence object.
func Parse(file io.Reader) (Gff, error) { _ = "STUB: not implemented"; return *new(Gff), nil }

// Add the CheckSum to sequence (blake3)

// get name for general meta

// Formally region name, but changed to name here for generality/interoperability.

// get meta info only specific to GFF files

// sequence.Sequence = sequence.Sequence + line

// Indexing starts at 1 for gff so we need to shift down for Sequence 0 index.

// var eqIndex int

// regionString takes in the lines array,fieldName that is needed in gff file, and
// returns the region containing fieldName if found
// throws error if not found
func extractInfoFromField(lines []string, fieldName string) ([]string, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Build takes an Annotated sequence and returns a byte array representing a gff to be written out.
func Build(sequence Gff) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Indexing starts at 1 for gff so we need to shift up from Sequence 0 index.

// Read takes in a filepath for a .gffv3 file and parses it into an Annotated poly.Sequence struct.
func Read(path string) (Gff, error) { _ = "STUB: not implemented"; return *new(Gff), nil }

// Write takes an poly.Sequence struct and a path string and writes out a gff to that path.
func Write(sequence Gff, path string) error { _ = "STUB: not implemented"; return nil }
