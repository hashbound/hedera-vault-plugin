package key

import "fmt"

type Algorithm uint32

const (
	ED25519 = 0
	ECDSA   = 1
)

func AlgorithmFromString(algo string) Algorithm {
	switch algo {
	case Algorithm(ED25519).String():
		return ED25519
	case Algorithm(ECDSA).String():
		return ECDSA
	}

	panic(fmt.Sprintf("unreachable: Algorithm.FromString() switch statement is non-exhaustive. Status: %s", algo))
}

func (algo Algorithm) String() string {
	switch algo {
	case ED25519:
		return "ED25519"
	case ECDSA:
		return "ECDSA"
	}

	panic(fmt.Sprintf("unreachable: Algorithm.String() switch statement is non-exhaustive. Status: %v", uint32(algo)))
}
