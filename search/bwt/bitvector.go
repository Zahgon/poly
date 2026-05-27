package bwt

const wordSize = 64

// bitvector a sequence of 1's and 0's. You can also think
// of this as an array of bits. This allows us to encode
// data in a memory efficient manner.
type bitvector struct {
	bits         []uint64
	numberOfBits int
}

// newBitVector will return an initialized bitvector with
// the specified number of zeroed bits.
func newBitVector(initialNumberOfBits int) bitvector {
	_ = "STUB: not implemented"
	return *new(bitvector)
}

// getBitSet gets the while word as some offset from the
// bitvector. Useful if you'd prefer to work with the
// word rather than with individual bits.
func (b bitvector) getBitSet(bitSetPos int) uint64 { _ = "STUB: not implemented"; return 0 }

// getBit returns the value of the bit at a given offset
// True represents 1
// False represents 0
func (b bitvector) getBit(i int) bool { _ = "STUB: not implemented"; return false }

// setBit sets the value of the bit at a given offset
// True represents 1
// False represents 0
func (b bitvector) setBit(i int, val bool) { _ = "STUB: not implemented"; return }

func (b bitvector) checkBounds(i int) { _ = "STUB: not implemented"; return }

func (b bitvector) len() int { _ = "STUB: not implemented"; return 0 }

func getNumOfBitSetsNeededForNumOfBits(n int) int { _ = "STUB: not implemented"; return 0 }
