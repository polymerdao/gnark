package circuits

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
)

type xorCircuit struct {
	Op1, Op2, Res frontend.Variable
}

func (circuit *xorCircuit) Define(curveID ecc.ID, cs frontend.API) error {
	d := cs.Xor(circuit.Op1, circuit.Op2)

	cs.AssertIsEqual(d, circuit.Res)
	return nil
}

func init() {

	good := []frontend.Circuit{
		&xorCircuit{
			Op1: frontend.Value(1),
			Op2: frontend.Value(1),
			Res: frontend.Value(0),
		},
		&xorCircuit{
			Op1: frontend.Value(1),
			Op2: frontend.Value(0),
			Res: frontend.Value(1),
		},
		&xorCircuit{
			Op1: frontend.Value(0),
			Op2: frontend.Value(1),
			Res: frontend.Value(1),
		},
		&xorCircuit{
			Op1: frontend.Value(0),
			Op2: frontend.Value(0),
			Res: frontend.Value(0),
		},
	}

	bad := []frontend.Circuit{
		&xorCircuit{
			Op1: frontend.Value(1),
			Op2: frontend.Value(1),
			Res: frontend.Value(1),
		},
		&xorCircuit{
			Op1: frontend.Value(1),
			Op2: frontend.Value(0),
			Res: frontend.Value(0),
		},
		&xorCircuit{
			Op1: frontend.Value(0),
			Op2: frontend.Value(1),
			Res: frontend.Value(0),
		},
		&xorCircuit{
			Op1: frontend.Value(0),
			Op2: frontend.Value(0),
			Res: frontend.Value(1),
		},
		&xorCircuit{
			Op1: frontend.Value(42),
			Op2: frontend.Value(1),
			Res: frontend.Value(1),
		},
		&xorCircuit{
			Op1: frontend.Value(1),
			Op2: frontend.Value(1),
			Res: frontend.Value(42),
		},
	}

	addNewEntry("xor", &xorCircuit{}, good, bad)
}
