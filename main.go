//Maintener : Yuri

package main

//go run main.go genKey  crypt rsa
//change import formats
import (
	"fmt"
	"math/big"
	"math/rand"
	"os"
	str "strings"
	"time"
)

func main() {
	// adsda.exe  genkey
	// adsda.exe filename  sign/crypt if crypt= rsa/elgamal
	//			0			1				2					3

	// go run main.go try.txt sign
	// @TODO: check qtde of args

	// go run main.go genKey sign/crypt

	if str.Compare(os.Args[1], "genKey") == 0 {
		fmt.Print(" /--* Starting generating Key *--\\ \n")

		r := generateKeys()
		fmt.Print(r)
	}
	/*
		if str.Compare(os.Args[2], "sign") == 0 {
			fmt.Print(" /--* Starting digital assing *--\\ \n")
			digitalSign()
		}
		//(args[2], "sign") == 0 {
	*/
}

func generateKeys() *big.Int {

	var f big.Int

	if str.Compare(os.Args[2], "crypt") == 0 {
		if str.Compare(os.Args[3], "rsa") == 0 {
			f.MulRange(1, 11)
			fmt.Print(f, "\n\n")
			var res bool
			for tega := 1; tega < 45; tega++ {
				//@TODO: check if n is greater than 3
				res = millerRabin(big.NewInt(7))
				//	fmt.Println(res)
				if res {
					fmt.Println("Composto")
					//	break
				}
			}

		}
	}
	return f.MulRange(1, 2)
}

func millerRabin(n *big.Int) bool {
	// returns true  if the number is composite and false if it's probably a prime number
	//
	var q, p, tmp big.Int
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	q.Add(n, big.NewInt(-1)) // q = n - 1
	orig := q
	k := big.NewInt(0) // k = 0
	//generate rand b in range [2,n-1]
	// (max - min) + min

	var rng big.Int
	rng.Add(&q, big.NewInt(-2)) //	(max - min) + min

	var b big.Int
	b.Rand(r, &rng).Add(&b, big.NewInt(2))

	// while q is odd (par)
	for tmp.QuoRem(&q, big.NewInt(2), &p); p.Cmp(big.NewInt(0)) == 0; tmp.QuoRem(&q, big.NewInt(2), &p) {
		k.Add(k, big.NewInt(1))
		q.Quo(&q, big.NewInt(2))

	}
	//find miller rabbin exp sequency
	var pots []big.Int // exps slice
	for i := big.NewInt(0); i.Cmp(k) != +1; i.Add(i, big.NewInt(1)) {
		var fake big.Int
		fake.Exp(big.NewInt(2), i, nil)
		pots = append(pots, fake)
	}

	for _, i := range pots {

		var t big.Int
		// q = i * teste
		q.Mul(&i, &orig)
		// t = b**q mod n
		t.Exp(&b, &q, n)

		if (t.Cmp(big.NewInt(1)) == 0) || (t.Cmp(&orig) == 0) {
			fmt.Println("INCONCLUSIVE")
			return false // inconclusive
		}
	}
	return true //composite
}

//func genRSA() {

//}

//func digitalSign() {

//}
