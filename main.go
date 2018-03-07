//Maintener : Yuri
package main

import (
	"fmt"
	fn "golangGPG/functions"
	"math/big"
	"math/rand"
	"os"
	str "strings"
	"time"
)

var (
	two = big.NewInt(2)
)

func main() {
	// @TODO: check qtde of args
	// @TODO: change functions names
	rand.Seed(time.Now().UTC().UnixNano())

	if str.Compare(os.Args[1], "genKey") == 0 {
		fmt.Print(" /--* Starting generating Key *--\\ \n")

		generateKeys()
	}
}
func generateKeys() *big.Int {
	var f big.Int
	//fmt.Sprintf("%x", n) convert big to hex
	if str.Compare(os.Args[2], "crypt") == 0 {
		// generate key-par (pub,priv) RSA
		if str.Compare(os.Args[3], "rsa") == 0 {
			n, e, d := buildRSA(fn.GenPrime(16), fn.GenPrime(16))

			fmt.Println("n:", &n)
			fmt.Println("e:", &e)
			fmt.Println("d:", &d)
		} else if str.Compare(os.Args[3], "elGamal") == 0 {
			p, g, c, a := buildElGamal(fn.GenPrime(32))

			fmt.Println("primo: ", p)
			fmt.Println("gerador: ", g)
			fmt.Println("c: ", c)
			fmt.Println("a: ", a)

		}
	} else if str.Compare(os.Args[2], "sign") == 0 {
		pub, priv := digitalSign(fn.GenPrime(32))

		fmt.Println("pub:  ", pub)
		fmt.Println("priv:  ", priv)

	}
	return f.MulRange(1, 2)
}
func digitalSign(p *big.Int) (big.Int, big.Int) {
	_, pub, priv := genElGamal(p) //ra stands for random number used to generate ...
	return pub, priv
}
func buildRSA(p1, p2 *big.Int) (big.Int, big.Int, big.Int) {
	n, e := genPubRSA(p1, p2)
	_, d := genPrivRSA(p1, p2, &e)

	return n, e, d
}
func buildElGamal(p *big.Int) (big.Int, big.Int, big.Int, big.Int) {
	g, c, ra := genElGamal(p) //ra stands for random number used to generate ...
	return *p, g, c, ra
}
func genPrivRSA(p, q, e *big.Int) (big.Int, big.Int) {
	var (
		n, fi, p2, q2 big.Int
	)
	p2.Sub(p, big.NewInt(1))
	q2.Sub(q, big.NewInt(1))
	fi.Mul(&p2, &q2)
	fmt.Println("fi:", &fi)

	n.Mul(p, q)
	fmt.Println("e2:", e)
	// pra ser negativo o euclidiano gera um numero negativo
	d := fn.ExtEuclidian(*e, fi, *big.NewInt(1), *big.NewInt(0), *big.NewInt(0), *big.NewInt(1))
	fmt.Println("d1:::", &d)
	d.Mod(&d, &fi)
	fmt.Println("d:::", &d)
	return n, d
}
func genPubRSA(p, q *big.Int) (big.Int, big.Int) {
	var (
		n, fi, p2, q2, e, gc big.Int
	)
	n.Mul(p, q)

	p2.Add(p, big.NewInt(-1))
	q2.Add(q, big.NewInt(-1))
	fi.Mul(&p2, &q2)
	ao := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	e.Rand(ao, &fi).Add(&e, big.NewInt(1))

	for gc.GCD(nil, nil, &e, &fi); gc.Cmp(big.NewInt(1)) != 0; {

		ao := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
		e.Rand(ao, &fi).Add(&e, big.NewInt(1))
		gc.GCD(nil, nil, &e, &fi)

	}
	return n, e
	//d.Exp

}

func genElGamal(p *big.Int) (big.Int, big.Int, big.Int) {
	var tmp444, ra, c big.Int

	g := fn.GetPrimitiveRoot(p)
	g.Exp(&g, two, p)
	ao := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	tmp444.Sub(p, big.NewInt(1)).Quo(&tmp444, big.NewInt(2))
	ra.Rand(ao, &tmp444)
	c.Exp(&g, &ra, p)

	return g, c, ra

}
