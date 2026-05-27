/*
Package alphabet provides structs for defining biological sequence alphabets.
*/
package alphabet

// Alphabet is a struct that holds a list of symbols and a map of symbols to their index in the list.
type Alphabet struct {
	symbols  []string
	encoding map[interface{}]int
}

// Error is an error type that is returned when a symbol is not in the alphabet.
type Error struct {
	message string
}

// Error returns the error message for AlphabetError.
func (e *Error) Error() string {
	_ = "STUB: not implemented"

	// NewAlphabet creates a new alphabet from a list of symbols.
	return ""
}

func NewAlphabet(symbols []string) *Alphabet { _ = "STUB: not implemented"; return nil }

// Encode returns the index of a symbol in the alphabet.
func (alphabet *Alphabet) Encode(symbol interface{}) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Decode returns the symbol at a given index in the alphabet.
func (alphabet *Alphabet) Decode(code interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Extend returns a new alphabet that is the original alphabet extended with a list of symbols.
func (alphabet *Alphabet) Extend(symbols []string) *Alphabet { _ = "STUB: not implemented"; return nil }

// Symbols returns the list of symbols in the alphabet.
func (alphabet *Alphabet) Symbols() []string { _ = "STUB: not implemented"; return nil }

var (
	DNA     = NewAlphabet([]string{"A", "C", "G", "T"})
	RNA     = NewAlphabet([]string{"A", "C", "G", "U"})
	Protein = NewAlphabet([]string{"A", "C", "D", "E", "F", "G", "H", "I", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "V", "W", "Y"})
)
