/*
Package genbank provides genbank parsers and writers.

GenBank is a flat text file format developed in the 1980s to annotate genetic
sequences, and has since become the standard for sharing annotated genetic
sequences.

This package provides a parser and writer to convert between the GenBank file
format and the more general Genbank struct.
*/
package genbank

import (
	"io"
	"os"
	"regexp"
	"strings"
)

/******************************************************************************

GBK specific IO related things begin here.

******************************************************************************/

var (
	readFileFn        = os.ReadFile
	parseMultiNthFn   = ParseMultiNth
	parseReferencesFn = parseReferences
)

// Genbank is the main struct for the Genbank file format.
type Genbank struct {
	Meta     Meta
	Features []Feature
	Sequence string // will be changed and include reader, writer, and byte slice.
}

// Meta holds the meta data for Genbank and other annotated sequence files.
type Meta struct {
	Date                 string            `json:"date"`
	Definition           string            `json:"definition"`
	Accession            string            `json:"accession"`
	Version              string            `json:"version"`
	Keywords             string            `json:"keywords"`
	Organism             string            `json:"organism"`
	Source               string            `json:"source"`
	Taxonomy             []string          `json:"taxonomy"`
	Origin               string            `json:"origin"`
	Locus                Locus             `json:"locus"`
	References           []Reference       `json:"references"`
	BaseCount            []BaseCount       `json:"base_count"`
	Other                map[string]string `json:"other"`
	Name                 string            `json:"name"`
	SequenceHash         string            `json:"sequence_hash"`
	SequenceHashFunction string            `json:"hash_function"`
}

// Feature holds the information for a feature in a Genbank file and other annotated sequence files.
type Feature struct {
	Type                 string            `json:"type"`
	Description          string            `json:"description"`
	Attributes           map[string]string `json:"attributes"`
	SequenceHash         string            `json:"sequence_hash"`
	SequenceHashFunction string            `json:"hash_function"`
	Sequence             string            `json:"sequence"`
	Location             Location          `json:"location"`
	ParentSequence       *Genbank          `json:"-"`
}

// Reference holds information for one reference in a Meta struct.
type Reference struct {
	Authors    string `json:"authors"`
	Title      string `json:"title"`
	Journal    string `json:"journal"`
	PubMed     string `json:"pub_med"`
	Remark     string `json:"remark"`
	Range      string `json:"range"`
	Consortium string `json:"consortium"`
}

// Locus holds Locus information in a Meta struct.
type Locus struct {
	Name             string `json:"name"`
	SequenceLength   string `json:"sequence_length"`
	MoleculeType     string `json:"molecule_type"`
	GenbankDivision  string `json:"genbank_division"`
	ModificationDate string `json:"modification_date"`
	SequenceCoding   string `json:"sequence_coding"`
	Circular         bool   `json:"circular"`
}

// Location is a struct that holds the location of a feature.
type Location struct {
	Start             int        `json:"start"`
	End               int        `json:"end"`
	Complement        bool       `json:"complement"`
	Join              bool       `json:"join"`
	FivePrimePartial  bool       `json:"five_prime_partial"`
	ThreePrimePartial bool       `json:"three_prime_partial"`
	GbkLocationString string     `json:"gbk_location_string"`
	SubLocations      []Location `json:"sub_locations"`
}

// BaseCount is a struct that holds the base counts for a sequence.
type BaseCount struct {
	Base  string
	Count int
}

// Precompiled regular expressions:
var (
	basePairRegex         = regexp.MustCompile(` \d* \w{2} `)
	circularRegex         = regexp.MustCompile(` circular `)
	modificationDateRegex = regexp.MustCompile(`\d{2}-[A-Z]{3}-\d{4}`)
	partialRegex          = regexp.MustCompile("<|>")
	sequenceRegex         = regexp.MustCompile("[^a-zA-Z]+")
)

// AddFeature adds a feature to a Genbank struct.
func (sequence *Genbank) AddFeature(feature *Feature) error { _ = "STUB: not implemented"; return nil }

// GetSequence returns the sequence of a feature.
func (feature Feature) GetSequence() (string, error) { _ = "STUB: not implemented"; return "", nil }

// getFeatureSequence takes a feature and location object and returns a sequence string.
func getFeatureSequence(feature Feature, location Location) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// reverse complements resulting string if needed.

// Read reads a GBK file from path and returns a Genbank struct.
func Read(path string) (Genbank, error) { _ = "STUB: not implemented"; return *new(Genbank), nil }

// ReadMulti reads a multi Gbk from path and parses it into a slice of Genbank structs.
func ReadMulti(path string) ([]Genbank, error) { _ = "STUB: not implemented"; return nil, nil }

// ReadMultiNth reads a multi Gbk from path and parses N entries into a slice of Genbank structs.
func ReadMultiNth(path string, count int) ([]Genbank, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Write takes an Genbank list and a path string and writes out a genbank record to that path.
func Write(sequences Genbank, path string) error {
	_ = "STUB: not implemented"
	// build function always returns nil error.
	// This is for API consistency in case we need to
	// add error handling in the future.
	return nil
}

// WriteMulti takes a slice of Genbank structs and a path string and writes out a multi genbank record to that path.
func WriteMulti(sequences []Genbank, path string) error {
	_ = "STUB: not implemented"
	// buildmulti function always returns nil error.
	// This is for API consistency in case we need to
	// add error handling in the future.
	return nil
}

// Build builds a GBK byte slice to be written out to db or file.
func Build(gbk Genbank) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// BuildMulti builds a MultiGBK byte slice to be written out to db or file.
func BuildMulti(sequences []Genbank) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// building locus

// building other standard meta features

// building references
// TODO: could use reflection to get keys and make more general.

// building other meta fields that are catch all

// start writing features section.

// start writing sequence section.

// iterate over every character in sequence range.

// if 60th character add newline then whitespace and index number and space before adding next base.

// genbank indexes at 1 for some reason
// <- I wish I was kidding

// if base index is divisible by ten add a space (genbank convention)

// else just add the base.

// finish genbank file with "//" on newline (again a genbank convention)

// Parse takes in a reader representing a single gbk/gb/genbank file and parses it into a Genbank struct.
func Parse(r io.Reader) (Genbank, error) { _ = "STUB: not implemented"; return *new(Genbank), nil }

// ParseMulti takes in a reader representing a multi gbk/gb/genbank file and parses it into a slice of Genbank structs.
func ParseMulti(r io.Reader) ([]Genbank, error) { _ = "STUB: not implemented"; return nil, nil }

type parseLoopParameters struct {
	newLocation      bool
	quoteActive      bool
	attribute        string
	attributeValue   string
	emptyAttribute   bool
	sequenceBuilder  strings.Builder
	parseStep        string
	genbank          Genbank // since we are scanning lines we need a Genbank struct to store the data outside the loop.
	feature          Feature
	features         []Feature
	metadataTag      string
	metadataData     []string //this stutters but will remain to make it easier to batch rename variables when compared to parameters.metadataTag.
	genbankStarted   bool
	currentLine      string
	prevline         string
	multiLineFeature bool
}

// method to init loop parameters
func (params *parseLoopParameters) init() {
	params.newLocation = true
	params.feature.Attributes = make(map[string]string)
	params.parseStep = "metadata"
	params.genbankStarted = false
	params.genbank.Meta.Other = make(map[string]string)
}

// ParseMultiNth takes in a reader representing a multi gbk/gb/genbank file and parses the first n records into a slice of Genbank structs.
func ParseMultiNth(r io.Reader, count int) ([]Genbank, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sequence setup

// Loop through each line of the file

// get line from scanner and split it

// keep scanning until we find the start of the first record

// We detect the beginning of a new genbank file with "LOCUS"

// Handle empty lines

// If we are currently reading a line, we need to figure out if it is a new meta line.

// If this is true, it means we are beginning a new meta tag. In that case, let's save
// the older data, and then continue along.

// We know that we are now parsing features, so lets initialize our first feature

// example string for BASE COUNT: "BASE COUNT    67070277 a   48055043 c   48111528 g   67244164 t   18475410 n"

// starts at two because we don't want to include "BASE COUNT" in our fields

// Switch to sequence parsing
// we detect the beginning of the sequence with "ORIGIN"

// save our completed attribute / qualifier string to the current feature

// add our features to the genbank

// end sequence parsing flag logic

// check if current line contains anything but whitespace

// determine if current line is a new top level feature

// save our completed attribute / qualifier string to the current feature

// }
// checks for empty types

// An initial feature line looks like this: `source          1..2686` with a type separated by its location

// without this we can't tell if something is a multiline feature or multiline qualifier
// current line is continuation of a feature or qualifier (sub-constituent of a feature)
// if it's a continuation of the current feature, add it to the location

// without this we can't tell if something is a multiline feature or multiline qualifier
// it's a continued line of a qualifier

// current line is a new qualifier

// if we have an exception case, like (adenine(1518)-N(6)/adenine(1519)-N(6))-

// save our completed attribute / qualifier string to the current feature

// handle case of ` /pseudo `, which has no text

// this is normally triggered

// without this we can't tell if something is a multiline feature or multiline qualifier

// throw error if line is malformed

// end of sequence

// add line to total sequence

func countLeadingSpaces(line string) int { _ = "STUB: not implemented"; return 0 }

func parseMetadata(metadataData []string) string { _ = "STUB: not implemented"; return "" }

// Remove trailing whitespace

func parseReferences(metadataData []string) (Reference, error) {
	_ = "STUB: not implemented"
	return *new(Reference), nil
}

// Otherwise, simply append the next metadata.

func (reference *Reference) addKey(referenceKey string, referenceValue string) error {
	_ = "STUB: not implemented"
	return nil
}

var genBankMoleculeTypes = []string{
	"DNA",
	"genomic DNA",
	"genomic RNA",
	"mRNA",
	"tRNA",
	"rRNA",
	"other RNA",
	"other DNA",
	"transcribed RNA",
	"viral cRNA",
	"unassigned DNA",
	"unassigned RNA",
}

// used in parseLocus function though it could be useful elsewhere.
var genbankDivisions = []string{
	"PRI", //primate sequences
	"ROD", //rodent sequences
	"MAM", //other mamallian sequences
	"VRT", //other vertebrate sequences
	"INV", //invertebrate sequences
	"PLN", //plant, fungal, and algal sequences
	"BCT", //bacterial sequences
	"VRL", //viral sequences
	"PHG", //bacteriophage sequences
	"SYN", //synthetic sequences
	"UNA", //unannotated sequences
	"EST", //EST sequences (expressed sequence tags)
	"PAT", //patent sequences
	"STS", //STS sequences (sequence tagged sites)
	"GSS", //GSS sequences (genome survey sequences)
	"HTG", //HTG sequences (high-throughput genomic sequences)
	"HTC", //unfinished high-throughput cDNA sequencing
	"ENV", //environmental sampling sequences
}

// TODO rewrite with proper error handling.
// parses locus from provided string.
func parseLocus(locusString string) Locus { _ = "STUB: not implemented"; return *new(Locus) }

// sequence length and coding

// molecule type

// circularity flag

// genbank division

// ModificationDate

// indices for random points of interests on a gbk line.
const subMetaIndex = 5
const qualifierIndex = 21

func getSourceOrganism(metadataData []string) (string, string, []string) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// Taxonomy ends with a ".", which we check for here

func parseLocation(locationString string) (Location, error) {
	_ = "STUB: not implemented"
	return *new(Location), nil
}

// Case checks for simple expression of x..x
//Case checks for simple expression x

// to remove FivePrimePartial and ThreePrimePartial indicators from start and end before converting to int.

// This case checks for join(complement(x..x),complement(x..x)), or any more complicated derivatives

// "(" is at 0, so we start at 1

// This is the default join(x..x,x..x)

// location.Complement = true

// if excess root node then trim node. Maybe should just be handled with second arg?

// buildMetaString is a helper function to build the meta section of genbank files.
func buildMetaString(name string, data string) string { _ = "STUB: not implemented"; return "" }

// I wish I was kidding.

// BuildLocationString is a recursive function that takes a location object and creates a gbk location string for Build()
func BuildLocationString(location Location) string { _ = "STUB: not implemented"; return "" }

// BuildFeatureString is a helper function to build gbk feature strings for Build()
func BuildFeatureString(feature Feature) string { _ = "STUB: not implemented"; return "" }

// I wish I was kidding.

func generateWhiteSpace(length int) string { _ = "STUB: not implemented"; return "" }

/******************************************************************************

GBK specific IO related things end here.

******************************************************************************/
