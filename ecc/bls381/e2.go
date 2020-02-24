// Code generated by internal/fp2 DO NOT EDIT

package bls381

import (
	"github.com/consensys/gnark/ecc/bls381/fp"
)

// e2 is a degree-two finite field extension of fp.Element:
// A0 + A1u where u^2 == -1 is a quadratic nonresidue in fp

type e2 struct {
	A0, A1 fp.Element
}

// SetString sets a e2 element from strings
func (z *e2) SetString(s1, s2 string) *e2 {
	z.A0.SetString(s1)
	z.A1.SetString(s2)
	return z
}

func (z *e2) SetZero() *e2 {
	z.A0.SetZero()
	z.A1.SetZero()
	return z
}

// Clone returns a copy of self
func (z *e2) Clone() *e2 {
	return &e2{
		A0: z.A0,
		A1: z.A1,
	}
}

// Set sets an e2 from x
func (z *e2) Set(x *e2) *e2 {
	z.A0.Set(&x.A0)
	z.A1.Set(&x.A1)
	return z
}

// Set sets z to 1
func (z *e2) SetOne() *e2 {
	z.A0.SetOne()
	z.A1.SetZero()
	return z
}

// SetRandom sets a0 and a1 to random values
func (z *e2) SetRandom() *e2 {
	z.A0.SetRandom()
	z.A1.SetRandom()
	return z
}

// Equal returns true if the two elements are equal, fasle otherwise
func (z *e2) Equal(x *e2) bool {
	return z.A0.Equal(&x.A0) && z.A1.Equal(&x.A1)
}

// Equal returns true if the two elements are equal, fasle otherwise
func (z *e2) IsZero() bool {
	return z.A0.IsZero() && z.A1.IsZero()
}

// Neg negates an e2 element
func (z *e2) Neg(x *e2) *e2 {
	z.A0.Neg(&x.A0)
	z.A1.Neg(&x.A1)
	return z
}

// String implements Stringer interface for fancy printing
func (z *e2) String() string {
	return (z.A0.String() + "+" + z.A1.String() + "*u")
}

// ToMont converts to mont form
func (z *e2) ToMont() *e2 {
	z.A0.ToMont()
	z.A1.ToMont()
	return z
}

// FromMont converts from mont form
func (z *e2) FromMont() *e2 {
	z.A0.FromMont()
	z.A1.FromMont()
	return z
}

// Add adds two elements of e2
func (z *e2) Add(x, y *e2) *e2 {
	z.A0.Add(&x.A0, &y.A0)
	z.A1.Add(&x.A1, &y.A1)
	return z
}

// AddAssign adds x to z
func (z *e2) AddAssign(x *e2) *e2 {
	z.A0.AddAssign(&x.A0)
	z.A1.AddAssign(&x.A1)
	return z
}

// Sub two elements of e2
func (z *e2) Sub(x, y *e2) *e2 {
	z.A0.Sub(&x.A0, &y.A0)
	z.A1.Sub(&x.A1, &y.A1)
	return z
}

// SubAssign subs x from z
func (z *e2) SubAssign(x *e2) *e2 {
	z.A0.SubAssign(&x.A0)
	z.A1.SubAssign(&x.A1)
	return z
}

// Double doubles an e2 element
func (z *e2) Double(x *e2) *e2 {
	z.A0.Double(&x.A0)
	z.A1.Double(&x.A1)
	return z
}

// Mul sets z to the e2-product of x,y, returns z
func (z *e2) Mul(x, y *e2) *e2 {
	// (a+bu)*(c+du) == (ac+(-1)*bd) + (ad+bc)u where u^2 == -1
	// Karatsuba: 3 fp multiplications instead of 4
	// [1]: ac
	// [2]: bd
	// [3]: (a+b)*(c+d)
	// Then z.A0: [1] + (-1)*[2]
	// Then z.A1: [3] - [2] - [1]
	var ac, bd, cplusd, aplusbcplusd fp.Element

	ac.Mul(&x.A0, &y.A0)            // [1]: ac
	bd.Mul(&x.A1, &y.A1)            // [2]: bd
	cplusd.Add(&y.A0, &y.A1)        // c+d
	aplusbcplusd.Add(&x.A0, &x.A1)  // a+b
	aplusbcplusd.MulAssign(&cplusd) // [3]: (a+b)*(c+d)
	z.A1.Add(&ac, &bd)              // ad+bc, [2] + [1]
	z.A1.Sub(&aplusbcplusd, &z.A1)  // z.A1: [3] - [2] - [1]
	z.A0.Sub(&ac, &bd)              // z.A0: [1] - [2]
	return z
}

// MulAssign sets z to the e2-product of z,x returns z
func (z *e2) MulAssign(x *e2) *e2 {
	// (a+bu)*(c+du) == (ac+(-1)*bd) + (ad+bc)u where u^2 == -1
	// Karatsuba: 3 fp multiplications instead of 4
	// [1]: ac
	// [2]: bd
	// [3]: (a+b)*(c+d)
	// Then z.A0: [1] + (-1)*[2]
	// Then z.A1: [3] - [2] - [1]
	var ac, bd, cplusd, aplusbcplusd fp.Element

	ac.Mul(&z.A0, &x.A0)            // [1]: ac
	bd.Mul(&z.A1, &x.A1)            // [2]: bd
	cplusd.Add(&x.A0, &x.A1)        // c+d
	aplusbcplusd.Add(&z.A0, &z.A1)  // a+b
	aplusbcplusd.MulAssign(&cplusd) // [3]: (a+b)*(c+d)
	z.A1.Add(&ac, &bd)              // ad+bc, [2] + [1]
	z.A1.Sub(&aplusbcplusd, &z.A1)  // z.A1: [3] - [2] - [1]
	z.A0.Sub(&ac, &bd)              // z.A0: [1] - [2]
	return z
}

// Square sets z to the e2-product of x,x returns z
func (z *e2) Square(x *e2) *e2 {
	// (a+bu)^2 == (a^2+(-1)*b^2) + (2ab)u where u^2 == -1
	// Complex method: 2 fp multiplications instead of 3
	// [1]: ab
	// [2]: (a+b)*(a+(-1)*b)
	// Then z.A0: [2] - (-1+1)*[1]
	// Then z.A1: 2[1]
	// optimize for quadratic nonresidue -1
	var aplusb fp.Element
	var result e2

	aplusb.Add(&x.A0, &x.A1)                       // a+b
	result.A0.Sub(&x.A0, &x.A1)                    // a-b
	result.A0.MulAssign(&aplusb)                   // [2]: (a+b)*(a-b)
	result.A1.Mul(&x.A0, &x.A1).Double(&result.A1) // [1]: ab

	z.Set(&result)

	return z
}

// MulByNonSquare multiplies an element by (0,1)
// TODO deprecate in favor of inlined MulByNonResidue in fp6 package
func (z *e2) MulByNonSquare(x *e2) *e2 {
	a := x.A0
	MulByNonResidue(&z.A0, &x.A1)
	z.A1 = a
	return z
}

// Inverse sets z to the e2-inverse of x, returns z
func (z *e2) Inverse(x *e2) *e2 {
	// Algorithm 8 from https://eprint.iacr.org/2010/354.pdf
	var a0, a1, t0, t1 fp.Element

	a0 = x.A0 // = is slightly faster than Set()
	a1 = x.A1 // = is slightly faster than Set()

	t0.Square(&a0)               // step 1
	t1.Square(&a1)               // step 2
	t0.Add(&t0, &t1)             // step 3
	t1.Inverse(&t0)              // step 4
	z.A0.Mul(&a0, &t1)           // step 5
	z.A1.Neg(&a1).MulAssign(&t1) // step 6

	return z
}

// MulByElement multiplies an element in e2 by an element in fp
func (z *e2) MulByElement(x *e2, y *fp.Element) *e2 {
	var yCopy fp.Element
	yCopy.Set(y)
	z.A0.Mul(&x.A0, &yCopy)
	z.A1.Mul(&x.A1, &yCopy)
	return z
}

// Conjugate conjugates an element in e2
func (z *e2) Conjugate(x *e2) *e2 {
	z.A0.Set(&x.A0)
	z.A1.Neg(&x.A1)
	return z
}

// MulByNonResidue multiplies a fp.Element by -1
// It would be nice to make this a method of fp.Element but fp.Element is outside this package
func MulByNonResidue(out, in *fp.Element) *fp.Element {
	(out).Neg(in)
	return out
}

// MulByNonResidueInv multiplies a fp.Element by -1^{-1}
// It would be nice to make this a method of fp.Element but fp.Element is outside this package
func MulByNonResidueInv(out, in *fp.Element) *fp.Element {
	// TODO this should be a no-op when out==in
	(out).Set(in)
	return out
}
