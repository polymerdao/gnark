// Code generated by internal/gpoint DO NOT EDIT
package bls381

import (
	"fmt"
	"testing"

	"github.com/consensys/gnark/ecc/bls381/fr"
)

func TestG2NotReallyHere(t *testing.T) {
	t.Skip("testPointsG2() not available?")
}

func TestMultiExpG2(t *testing.T) {

	curve := BLS381()

	pointsJac := make([]G2Jac, 5)
	pointsAff := make([]G2Affine, 5)
	scalars := make([]fr.Element, 5)
	scalars[0].SetString("6833313093782752447774533032379533859360921141590695983").FromMont()
	scalars[1].SetString("6833313093782695347774533032379422124572402138975338593590695983").FromMont()
	scalars[2].SetString("683331309378269530181623992840859250215771777453309360695983").FromMont()
	scalars[3].SetString("6833353018162399284042212457240213897533859360921141590695983").FromMont()
	scalars[4].SetString("68333130937826953018162385525244777453303237942212485936983").FromMont()

	gens := make([]fr.Element, 5)
	gens[0].SetUint64(1).FromMont()
	gens[1].SetUint64(5).FromMont()
	gens[2].SetUint64(7).FromMont()
	gens[3].SetUint64(11).FromMont()
	gens[4].SetUint64(13).FromMont()

	pointsJac[0].ScalarMul(curve, &curve.g2Gen, gens[0])
	pointsJac[1].ScalarMul(curve, &curve.g2Gen, gens[1])
	pointsJac[2].ScalarMul(curve, &curve.g2Gen, gens[2])
	pointsJac[3].ScalarMul(curve, &curve.g2Gen, gens[3])
	pointsJac[4].ScalarMul(curve, &curve.g2Gen, gens[4])
	for i := 0; i < 5; i++ {
		pointsJac[i].ToAffineFromJac(&pointsAff[i])
	}

	pointsRes := make([]G2Jac, 5)
	pointsRes[0].ScalarMul(curve, &pointsJac[0], scalars[0])
	pointsRes[1].ScalarMul(curve, &pointsJac[1], scalars[1])
	pointsRes[2].ScalarMul(curve, &pointsJac[2], scalars[2])
	pointsRes[3].ScalarMul(curve, &pointsJac[3], scalars[3])
	pointsRes[4].ScalarMul(curve, &pointsJac[4], scalars[4])

	res := curve.g2Infinity

	for i := 0; i < 5; i++ {
		res.Add(curve, &pointsRes[i])
	}

	var multiExpRes G2Jac
	<-multiExpRes.MultiExp(curve, pointsAff, scalars)

	if !multiExpRes.Equal(&res) {
		fmt.Println("multiExp failed")
	}
}

func TestMultiExpG2LotOfPoints(t *testing.T) {

	curve := BLS381()

	var G G2Jac

	samplePoints := make([]G2Affine, 1000)
	sampleScalars := make([]fr.Element, 1000)

	G.Set(&curve.g2Gen)

	for i := 1; i <= 1000; i++ {
		sampleScalars[i-1].SetUint64(uint64(i)).FromMont()
		G.ToAffineFromJac(&samplePoints[i-1])
	}

	var testPoint G2Jac

	<-testPoint.MultiExp(curve, samplePoints, sampleScalars)

	var finalScalar fr.Element
	finalScalar.SetUint64(500500).FromMont()
	var finalPoint G2Jac
	finalPoint.ScalarMul(curve, &G, finalScalar)

	if !finalPoint.Equal(&testPoint) {
		t.Fatal("error multi exp")
	}

}

func testPointsG2MultiExp(n int) (points []G2Jac, scalars []fr.Element) {

	curve := BLS381()

	// points
	points = make([]G2Jac, n)
	points[0].Set(&curve.g2Gen)
	points[1].Set(&points[0]).Double() // can't call p.Add(a) when p equals a
	for i := 2; i < len(points); i++ {
		points[i].Set(&points[i-1]).Add(curve, &points[0]) // points[i] = i*g2Gen
	}

	// scalars
	// non-Montgomery form
	// cardinality of G2 is the fr modulus, so scalars should be fr.Elements
	// non-Montgomery form
	scalars = make([]fr.Element, n)

	// To ensure a diverse selection of scalars that use all words of an fr.Element,
	// each scalar should be a power of a large generator of fr.
	// 22 is a small generator of fr for bls377.
	// 2^{31}-1 is prime, so 22^{2^31}-1} is a large generator of fr for bls377
	// generator in Montgomery form
	var scalarGenMont fr.Element
	scalarGenMont.SetString("7716837800905789770901243404444209691916730933998574719964609384059111546487")

	scalars[0].Set(&scalarGenMont).FromMont()

	var curScalarMont fr.Element // Montgomery form
	curScalarMont.Set(&scalarGenMont)
	for i := 1; i < len(scalars); i++ {
		curScalarMont.MulAssign(&scalarGenMont)
		scalars[i].Set(&curScalarMont).FromMont() // scalars[i] = scalars[0]^i
	}

	return points, scalars
}

//--------------------//
//     benches		  //
//--------------------//

var benchResG2 G2Jac

func BenchmarkG2WindowedMultiExp(b *testing.B) {
	curve := BLS381()

	var numPoints []int
	for n := 5; n < 400000; n *= 2 {
		numPoints = append(numPoints, n)
	}

	for j := range numPoints {
		points, scalars := testPointsG2MultiExp(numPoints[j])

		b.Run(fmt.Sprintf("%d points", numPoints[j]), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				benchResG2.WindowedMultiExp(curve, points, scalars)
			}
		})
	}
}

func BenchmarkG2MultiExp(b *testing.B) {
	curve := BLS381()

	var numPoints []int
	for n := 5; n < 400000; n *= 2 {
		numPoints = append(numPoints, n)
	}

	for j := range numPoints {
		_points, scalars := testPointsG2MultiExp(numPoints[j])
		points := make([]G2Affine, len(_points))
		for i := 0; i < len(_points); i++ {
			_points[i].ToAffineFromJac(&points[i])
		}

		b.Run(fmt.Sprintf("%d points", numPoints[j]), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				benchResG2.MultiExp(curve, points, scalars)
			}
		})
	}
}
