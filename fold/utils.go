package fold

import (
	"golang.org/x/exp/constraints"
)

func max[T constraints.Ordered](a, b T) T { _ = "STUB: not implemented"; return *new(T) }

func abs(x int) int { _ = "STUB: not implemented"; return 0 }
