/*
Package slow5 contains slow5 parsers and writers.

Right now, only parsing slow5 files is supported. Support for writing and blow5
coming soon.

slow5 is a file format alternative to fast5, which is the file format outputted
by Oxford Nanopore sequencing devices. fast5 uses hdf5, which is a complex file
format that can only be read and written with a single software library built
in 1998. On the other hand, slow5 uses a .tsv file format, which is easy to
both parse and write.

slow5 files contain both general metadata about the sequencing run and raw
signal reads from the sequencing run. This raw signal can be used directly or
basecalled and used for alignment.

More information on slow5 can be found here: https://github.com/hasindu2008/slow5tools
*/
package slow5

import (
	"bufio"
	"io"
)

/******************************************************************************
Oct 10, 2021

slow5 parser begins here. Specification below:
https://hasindu2008.github.io/slow5specs/slow5-v1.0.0.pdf

slow5 is able to combine multiple sequencing runs into a single file format,
but we read each sequencing run separately. Each sequencing run contains a
header with metadata and a list of reads. However, unlike many other file
formats, a slow5 file should almost never be read into a common struct, since
most runs are large, and will take a ton of memory. Instead, the default way
to parse slow5 files is to produce a list of headers and a channel of raw
reads. In order to connect the two (if needed), use ReadGroupID.

Nanopore changes the attributes found in the header quite often, so we store
most of these attributes in a map for future proofing. Even the binary file
format, blow5, does not have types for these attributes, and just stores them
as a long string.

Reads have 8 required columns, and a few auxiliary. These are typed, since they
are what will probably be used in real software.

Cheers mate,

Keoni

******************************************************************************/

// Header contains metadata about the sequencing run in general.
type Header struct {
	ReadGroupID        uint32
	Slow5Version       string
	Attributes         map[string]string
	EndReasonHeaderMap map[string]int
}

// Read contains metadata and raw signal strengths for a single nanopore read.
type Read struct {
	ReadID       string
	ReadGroupID  uint32
	Digitisation float64
	Offset       float64
	Range        float64
	SamplingRate float64
	LenRawSignal uint64
	RawSignal    []int16

	// Auxiliary fields
	ChannelNumber string
	MedianBefore  float64
	ReadNumber    int32
	StartMux      uint8
	StartTime     uint64
	EndReason     string // enum{unknown,partial,mux_change,unblock_mux_change,data_service_unblock_mux_change,signal_positive,signal_negative}

	Error error // in case there is an error while parsing!
}

var knownEndReasons = map[string]bool{"unknown": true,
	"partial":                         true,
	"mux_change":                      true,
	"unblock_mux_change":              true,
	"data_service_unblock_mux_change": true,
	"signal_positive":                 true,
	"signal_negative":                 true,
}

// Parser is a flexible parser that provides ample
// control over reading slow5 sequences.
// It is initialized with NewParser.
type Parser struct {
	// reader keeps state of current reader.
	reader       bufio.Reader
	line         uint
	headerMap    map[int]string
	endReasonMap map[int]string
}

// NewParser parsers a slow5 file.
func NewParser(r io.Reader, maxLineSize int) (*Parser, []Header, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// First, we need to identify the number of read groups. This number will be the length of our
// ReadGroups output, and we will need it for iteration through the rest of the header.

// Terminate if we hit the beginning of the raw read headers
// Get endReasonEnums. This is simply a string between enum{} that is used for the reasons that a read could have ended.

// Add endReasonMap to each output header. This helps the writer

// Get the read headers and their identifiers. Though the primary read headers are in a defined order, the auxiliary headers are not.

// Check to make sure we have the right amount of information for the num_read_groups

// ParseNext parses the next read from a parser.
func (parser *Parser) ParseNext() (Read, error) { _ = "STUB: not implemented"; return *new(Read), nil }

// Reads have started.
// Once we have the read headers, start to parse the actual reads

// For whatever reason, this is a string.

/******************************************************************************
March 26, 2023

Start of Write functions

Slow5 write takes in a header, a channel of reads, and an io.Writer output.
The intended use case of slow5 write is reading from a location, such as a
database, and then directly writing the output to somewhere (stdout, a file,
etc).

A channel is used here so that reading and writing of slow5 can be done
concurrently. In almost all cases, you do not want to have slow5 files entirely
in memory, because they're freakin' huge.

Cheers,

Keoni

******************************************************************************/

// Write writes a list of headers and a channel of reads to an output.
func Write(headers []Header, reads <-chan Read, output io.Writer) error {
	_ = "STUB: not implemented"
	// First, write the slow5 version number
	return nil
}

// Then, write the number of read groups (ie, the number of headers)

// Next, we need a map of what attribute values are available

// Now that we know what attribute values are possible, lets build a map
// with those values with "." placeholders (as defined in slow5 spec)

// Build a list with all header values

// Sort the header strings

// Write the header attribute strings to the output

// Now we handle endReasons / failureReasons. In the slow5 spec, these are
// enums depending on what is present in the FAST5 files. This means the
// labels may not be consistent and there is no exhaustive list of enum
// labels. So, we have to create this from the endReasonMap each time we
// write the slow5 file (not a const!)
// Invert the endReasonMap

// Build endReasonString

// Remove trailing comma

// Write the read headers
// These are according to the slow5 specifications

// Iterate over reads. This is reading from a channel, and will end
// when the channel is closed.

// converts []int16 to string

// Don't add a comma to last number

// Look at above output.Write("#read_id ... for the values here.
