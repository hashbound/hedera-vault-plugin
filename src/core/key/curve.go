package key

import "fmt"

type Curve int32

const (
	None      = -1
	secp256k1 = 0
)

func CurveFromString(curve string) Curve {
	switch curve {
	case Curve(secp256k1).String():
		return secp256k1
	}

	panic(fmt.Sprintf("unreachable: Curve.CurveFromString() switch statement is non-exhaustive. Status: %v", curve))
}

func (curve Curve) String() string {
	switch curve {
	case secp256k1:
		return "secp256k1"
	}

	panic(fmt.Sprintf("unreachable: Curve.FromString() switch statement is non-exhaustive. Status: %v", uint32(curve)))
}
