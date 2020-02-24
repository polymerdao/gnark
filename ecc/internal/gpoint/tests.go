package gpoint

const tests = `
// Code generated by internal/gpoint DO NOT EDIT 
package {{.PackageName}}

import (
	"testing"
	"fmt"
	"github.com/consensys/gnark/internal/pool"
	"github.com/consensys/gnark/ecc/{{.PackageName}}/fr"
)

{{- if eq .Name "G2"}} 
func Test{{.Name}}NotReallyHere(t *testing.T) {
	t.Skip("testPoints{{.Name}}() not available?")
}
{{- else}}

func Test{{.Name}}JacToAffineFromJac(t *testing.T) {

	p := testPoints{{.Name}}()

	_p := {{.Name}}Affine{}
	p[0].ToAffineFromJac(&_p)
	if !_p.X.Equal(&p[1].X) || !_p.Y.Equal(&p[1].Y){
		t.Fatal("ToAffineFromJac failed")
	}
	
}

func Test{{.Name}}Conv(t *testing.T) {
	p := testPoints{{.Name}}()

	for i := 0 ; i < len(p) ; i++ {
		var pJac {{.Name}}Jac
		var pAff {{.Name}}Affine
		p[i].ToAffineFromJac(&pAff)
		pAff.ToJacobian(&pJac)
		if !pJac.Equal(&p[i]) {
			t.Fatal("jacobian to affine to jacobian fails")
		}
	}
}


func Test{{.Name}}JacAdd(t *testing.T) {

	curve := {{toUpper .PackageName}}()
	p := testPoints{{.Name}}()

	// p3 = p1 + p2
	p1 := p[1].Clone()
	_p2 := {{.Name}}Affine{}
	p[2].ToAffineFromJac(&_p2)
	p[1].AddMixed(&_p2)
	p[2].Add(curve, p1)

	if !p[3].Equal(&p[1]) {
		t.Fatal("Add failed")
	}

	// test commutativity
	if !p[3].Equal(&p[2]) {
		t.Fatal("Add failed")
	}
}

func Test{{.Name}}JacSub(t *testing.T) {

	curve := {{toUpper .PackageName}}()
	p := testPoints{{.Name}}()

	// p4 = p1 - p2
	p[1].Sub(curve, p[2])

	if !p[4].Equal(&p[1]) {
		t.Fatal("Sub failed")
	}
}

func Test{{.Name}}JacDouble(t *testing.T) {

	curve := {{toUpper .PackageName}}()
	p := testPoints{{.Name}}()

	// p5 = 2 * p1
	p[1].Double()
	if !p[5].Equal(&p[1]) {
		t.Fatal("Double failed")
	}

	G := curve.{{toLower .Name}}Infinity.Clone()
	R := curve.{{toLower .Name}}Infinity.Clone()
	G.Double()

	if !G.Equal(R) {
		t.Fatal("Double failed (infinity case)")
	}
}

func Test{{.Name}}JacScalarMul(t *testing.T) {

	curve := {{toUpper .PackageName}}()
	p := testPoints{{.Name}}()

	// p6 = [p1]32394 (scalar mul)
	scalar := fr.Element{32394}
	p[1].ScalarMul(curve, &p[1], scalar)

	if !p[1].Equal(&p[6]) {
		t.Error("ScalarMul failed")
	}
}



func Test{{.Name}}JacMultiExp(t *testing.T) {
	curve := {{toUpper .PackageName}}()
	// var points []{{.Name}}Jac
	var scalars []fr.Element
	var got {{.Name}}Jac

	//
	// Test 1: testPoints{{.Name}}multiExp
	// 
	// TODO why is this commented?
	// numPoints, wants := testPoints{{.Name}}MultiExpResults()

	// for i := range numPoints {
	// 	if numPoints[i] > 10000 {
	// 		continue
	// 	}
	// 	points, scalars = testPoints{{.Name}}MultiExp(numPoints[i])

	// 	got.multiExp(curve, points, scalars)
	// 	if !got.Equal(&wants[i]) {
	// 		t.Error("multiExp {{.Name}}Jac fail for points:", numPoints[i])
	// 	}
	// }

	//
	// Test 2: testPoints{{.Name}}()
	//
	p := testPoints{{.Name}}()

	// scalars
	s1 := fr.Element{23872983, 238203802, 9827897384, 2372}
	s2 := fr.Element{128923, 2878236, 398478, 187970707}
	s3 := fr.Element{9038947, 3947970, 29080823, 282739}

	scalars = []fr.Element{
		s1,
		s2,
		s3,
	}

	got.multiExp(curve, p[17:20], scalars)
	if !got.Equal(&p[20]) {
		t.Error("multiExp {{.Name}}Jac failed")
	}

	//
	// Test 3: edge cases
	//

	// one input point p[1]
	scalars[0] = fr.Element{32394, 0, 0, 0} // single-word scalar
	got.multiExp(curve, p[1:2], scalars[:1])
	if !got.Equal(&p[6]) {
		t.Error("multiExp {{.Name}}Jac failed, scalar:", scalars[0])
	}

	scalars[0] = fr.Element{2, 0, 0, 0} // scalar = 2
	got.multiExp(curve, p[1:2], scalars[:1])
	if !got.Equal(&p[5]) {
		t.Error("multiExp {{.Name}}Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{1, 0, 0, 0} // scalar = 1
	got.multiExp(curve, p[1:2], scalars[:1])
	if !got.Equal(&p[1]) {
		t.Error("multiExp {{.Name}}Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{0, 0, 0, 0} // scalar = 0
	got.multiExp(curve, p[1:2], scalars[:1])
	if !got.Equal(&curve.{{toLower .Name}}Infinity) {
		t.Error("multiExp {{.Name}}Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{^uint64(0), ^uint64(0), ^uint64(0), ^uint64(0)} // scalar == (4-word maxuint)
	got.multiExp(curve, p[1:2], scalars[:1])
	if !got.Equal(&p[21]) {
		t.Error("multiExp {{.Name}}Jac failed, scalar:", scalars[0])
	}

	// one input point curve.{{toLower .Name}}Infinity
	infinity := []{{.Name}}Jac{curve.{{toLower .Name}}Infinity}

	scalars[0] = fr.Element{32394, 0, 0, 0} // single-word scalar
	got.multiExp(curve, infinity, scalars[:1])
	if !got.Equal(&curve.{{toLower .Name}}Infinity) {
		t.Error("multiExp {{.Name}}Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{2, 0, 0, 0} // scalar = 2
	got.multiExp(curve, infinity, scalars[:1])
	if !got.Equal(&curve.{{toLower .Name}}Infinity) {
		t.Error("multiExp {{.Name}}Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{1, 0, 0, 0} // scalar = 1
	got.multiExp(curve, infinity, scalars[:1])
	if !got.Equal(&curve.{{toLower .Name}}Infinity) {
		t.Error("multiExp {{.Name}}Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{0, 0, 0, 0} // scalar = 0
	got.multiExp(curve, infinity, scalars[:1])
	if !got.Equal(&curve.{{toLower .Name}}Infinity) {
		t.Error("multiExp {{.Name}}Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{^uint64(0), ^uint64(0), ^uint64(0), ^uint64(0)} // scalar == (4-word maxuint)
	got.multiExp(curve, infinity, scalars[:1])
	if !got.Equal(&curve.{{toLower .Name}}Infinity) {
		t.Error("multiExp {{.Name}}Jac failed, scalar:", scalars[0])
	}

	// two input points: p[1], curve.{{toLower .Name}}Infinity
	twoPoints := []{{.Name}}Jac{p[1], curve.{{toLower .Name}}Infinity}

	scalars[0] = fr.Element{32394, 0, 0, 0} // single-word scalar
	scalars[1] = fr.Element{2, 0, 0, 0}     // scalar = 2
	got.multiExp(curve, twoPoints, scalars[:2])
	if !got.Equal(&p[6]) {
		t.Error("multiExp {{.Name}}Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{2, 0, 0, 0} // scalar = 2
	scalars[1] = fr.Element{1, 0, 0, 0} // scalar = 1
	got.multiExp(curve, twoPoints, scalars[:2])
	if !got.Equal(&p[5]) {
		t.Error("multiExp {{.Name}}Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{1, 0, 0, 0} // scalar = 1
	scalars[1] = fr.Element{0, 0, 0, 0} // scalar = 0
	got.multiExp(curve, twoPoints, scalars[:2])
	if !got.Equal(&p[1]) {
		t.Error("multiExp {{.Name}}Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{0, 0, 0, 0}                                     // scalar = 0
	scalars[1] = fr.Element{^uint64(0), ^uint64(0), ^uint64(0), ^uint64(0)} // scalar == (4-word maxuint)
	got.multiExp(curve, twoPoints, scalars[:2])
	if !got.Equal(&curve.{{toLower .Name}}Infinity) {
		t.Error("multiExp {{.Name}}Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{^uint64(0), ^uint64(0), ^uint64(0), ^uint64(0)} // scalar == (4-word maxuint)
	scalars[1] = fr.Element{32394, 0, 0, 0}                                 // single-word scalar
	got.multiExp(curve, twoPoints, scalars[:2])
	if !got.Equal(&p[21]) {
		t.Error("multiExp {{.Name}}Jac failed, scalar:", scalars[0])
	}

	// TODO: Jacobian points with nontrivial Z coord?
}

{{- end}}

func TestMultiExp{{ .Name}}(t *testing.T) {

	curve := {{toUpper .PackageName}}()

	pointsJac := make([]{{ .Name}}Jac, 5)
	pointsAff := make([]{{ .Name}}Affine, 5)
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

	pointsJac[0].ScalarMul(curve, &curve.{{toLower .Name}}Gen, gens[0])
	pointsJac[1].ScalarMul(curve, &curve.{{toLower .Name}}Gen, gens[1])
	pointsJac[2].ScalarMul(curve, &curve.{{toLower .Name}}Gen, gens[2])
	pointsJac[3].ScalarMul(curve, &curve.{{toLower .Name}}Gen, gens[3])
	pointsJac[4].ScalarMul(curve, &curve.{{toLower .Name}}Gen, gens[4])
	for i := 0; i < 5; i++ {
		pointsJac[i].ToAffineFromJac(&pointsAff[i])
	}

	pointsRes := make([]{{ .Name}}Jac, 5)
	pointsRes[0].ScalarMul(curve, &pointsJac[0], scalars[0])
	pointsRes[1].ScalarMul(curve, &pointsJac[1], scalars[1])
	pointsRes[2].ScalarMul(curve, &pointsJac[2], scalars[2])
	pointsRes[3].ScalarMul(curve, &pointsJac[3], scalars[3])
	pointsRes[4].ScalarMul(curve, &pointsJac[4], scalars[4])

	res := curve.{{toLower .Name}}Infinity

	for i := 0; i < 5; i++ {
		res.Add(curve, &pointsRes[i])
	}

	var multiExpRes {{ .Name}}Jac
	<-multiExpRes.MultiExp(curve, pointsAff, scalars)

	if !multiExpRes.Equal(&res) {
		fmt.Println("multiExp failed")
	}
}

func TestMultiExp{{.Name}}LotOfPoints(t *testing.T) {

	curve := {{toUpper .PackageName}}()

	var G {{.Name}}Jac

	samplePoints := make([]{{.Name}}Affine, 1000)
	sampleScalars := make([]fr.Element, 1000)

	G.Set(&curve.{{toLower .Name}}Gen)

	for i := 1; i <= 1000; i++ {
		sampleScalars[i-1].SetUint64(uint64(i)).FromMont()
		G.ToAffineFromJac(&samplePoints[i-1])
	}

	var testPoint {{.Name}}Jac

	<-testPoint.MultiExp(curve, samplePoints, sampleScalars)

	var finalScalar fr.Element
	finalScalar.SetUint64(500500).FromMont()
	var finalPoint {{.Name}}Jac
	finalPoint.ScalarMul(curve, &G, finalScalar)

	if !finalPoint.Equal(&testPoint) {
		t.Fatal("error multi exp")
	}

}


func testPoints{{.Name}}MultiExp(n int) (points []{{.Name}}Jac, scalars []fr.Element) {

	curve := {{toUpper .PackageName}}()

	// points
	points = make([]{{.Name}}Jac, n)
	points[0].Set(&curve.{{toLower .Name}}Gen)
	points[1].Set(&points[0]).Double() // can't call p.Add(a) when p equals a
	for i := 2; i < len(points); i++ {
		points[i].Set(&points[i-1]).Add(curve, &points[0]) // points[i] = i*{{toLower .Name}}Gen
	}

	// scalars
	// non-Montgomery form
	// cardinality of {{.Name}} is the fr modulus, so scalars should be fr.Elements
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
`
