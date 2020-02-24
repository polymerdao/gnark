package fp6

import (
	"github.com/consensys/gnark/ecc/internal/tower"
	"github.com/consensys/gnark/ecc/internal/tower/fp2"
)

// CodeSource is the aggregated source code
var CodeSource []string

// CodeTest is the aggregated test code
var CodeTest []string

// CodeTestPoints is the aggregated test points code
var CodeTestPoints []string

func init() {
	CodeSource = []string{
		base,
		mul,
		Inline,
		fp2.Inline,
	}

	CodeTest = []string{ // TODO move this to tower
		tower.Tests,
		customTests,
	}

	CodeTestPoints = []string{
		tower.TestPoints,
	}
}

const base = `
// Code generated by internal/fp6 DO NOT EDIT 

package {{.PackageName}}

// {{.Name}} is a degree-three finite field extension of fp2:
// B0 + B1v + B2v^2 where v^3-{{.Fp6NonResidue}} is irrep in fp2

type {{.Name}} struct {
	B0, B1, B2 {{.Fp2Name}}
}

// SetString sets a {{.Name}} elmt from stringf
func (z *{{.Name}}) SetString(s1, s2, s3, s4, s5, s6 string) *{{.Name}} {
	z.B0.SetString(s1, s2)
	z.B1.SetString(s3, s4)
	z.B2.SetString(s5, s6)
	return z
}

// Set Sets a {{.Name}} elmt form another {{.Name}} elmt
func (z *{{.Name}}) Set(x *{{.Name}}) *{{.Name}} {
	z.B0 = x.B0
	z.B1 = x.B1
	z.B2 = x.B2
	return z
}

// Equal compares two elements in {{.Name}}
func (z *{{.Name}}) Equal(x *{{.Name}}) bool {
	return z.B0.Equal(&x.B0) && z.B1.Equal(&x.B1) && z.B2.Equal(&x.B2)
}

// ToMont converts to Mont form
func (z *{{.Name}}) ToMont() *{{.Name}} {
	z.B0.ToMont()
	z.B1.ToMont()
	z.B2.ToMont()
	return z
}

// FromMont converts from Mont form
func (z *{{.Name}}) FromMont() *{{.Name}} {
	z.B0.FromMont()
	z.B1.FromMont()
	z.B2.FromMont()
	return z
}

// Add adds two elements of {{.Name}}
func (z *{{.Name}}) Add(x, y *{{.Name}}) *{{.Name}} {
	z.B0.Add(&x.B0, &y.B0)
	z.B1.Add(&x.B1, &y.B1)
	z.B2.Add(&x.B2, &y.B2)
	return z
}

// Neg negates the {{.Name}} number
func (z *{{.Name}}) Neg(x *{{.Name}}) *{{.Name}} {
	z.B0.Neg(&z.B0)
	z.B1.Neg(&z.B1)
	z.B2.Neg(&z.B2)
	return z
}

// Sub two elements of {{.Name}}
func (z *{{.Name}}) Sub(x, y *{{.Name}}) *{{.Name}} {
	z.B0.Sub(&x.B0, &y.B0)
	z.B1.Sub(&x.B1, &y.B1)
	z.B2.Sub(&x.B2, &y.B2)
	return z
}

// MulByGen Multiplies by v, root of X^3-{{.Fp6NonResidue}}
// TODO deprecate in favor of inlined MulByNonResidue in fp12 package
func (z *{{.Name}}) MulByGen(x *{{.Name}}) *{{.Name}} {
	var result {{.Name}}

	result.B1 = x.B0
	result.B2 = x.B1
	{{- template "fp2InlineMulByNonResidue" dict "all" . "out" "result.B0" "in" "&x.B2" }}

	z.Set(&result)
	return z
}

// Double doubles an element in {{.Name}}
func (z *{{.Name}}) Double(x *{{.Name}}) *{{.Name}} {
	z.B0.Double(&x.B0)
	z.B1.Double(&x.B1)
	z.B2.Double(&x.B2)
	return z
}

// String puts {{.Name}} elmt in string form
func (z *{{.Name}}) String() string {
	return (z.B0.String() + "+(" + z.B1.String() + ")*v+(" + z.B2.String() + ")*v**2")
}
`
