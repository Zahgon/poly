/*
Package random provides functions to generate random DNA, RNA, and protein sequences.
*/
package random

// ProteinSequence returns a random protein sequence string of a given length and seed.
// All returned sequences start M (Methionine) and end with * (stop codon).
func ProteinSequence(length int, seed int64) (string, error) {
	_ = "STUB: not implemented"
	//The length needs to be greater than two because the random protein sequenced returned always contain a start and stop codon. You could see more about this stuff here: https://en.wikipedia.org/wiki/Genetic_code#Start_and_stop_codons
	return "", nil
}

// https://en.wikipedia.org/wiki/Amino_acid#Table_of_standard_amino_acid_abbreviations_and_properties

//M is the standard abbreviation for the Methionine aminoacid. A protein sequence start with M because the start codon is translated to Methionine

//* is the standard abbreviation for the stop codon. That's a signal for the ribosome to stop the translation and because of that a protein sequence is finished with *

// DNASequence returns a random DNA sequence string of a given length and seed.
func DNASequence(length int, seed int64) (string, error) { _ = "STUB: not implemented"; return "", nil }

// RNASequence returns a random DNA sequence string of a given length and seed.
func RNASequence(length int, seed int64) (string, error) { _ = "STUB: not implemented"; return "", nil }

func randomNucelotideSequence(length int, seed int64, alphabet []rune) string {
	_ = "STUB: not implemented"
	return ""
}
