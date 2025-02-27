package circuits

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
)

type checkAssertEqualCircuit struct {
	X frontend.Variable
	Y frontend.Variable `gnark:",public"`
}

func (circuit *checkAssertEqualCircuit) Define(curveID ecc.ID, cs frontend.API) error {
	cs.AssertIsEqual(circuit.X, circuit.Y)
	return nil
}

func init() {

	var circuit, good, bad checkAssertEqualCircuit

	good.X.Assign(3)
	good.Y.Assign(3)

	bad.X.Assign(5)
	bad.Y.Assign(2)

	addEntry("assert_equal", &circuit, &good, &bad)
}
