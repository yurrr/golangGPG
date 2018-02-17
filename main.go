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
			for tega := 1; tega < 40; tega++ {
				res = millerRabin(big.NewInt(67280421310721))
				fmt.Println(res)
				if res {
					break
				}
			}
			if res == false {
				fmt.Println("COMPOSTO\n")
			}

		}
	}

	return f.MulRange(1, 2)
}

func millerRabin(n *big.Int) bool {
	//@TODO: check if n is greater than 3
	var q, p, tmp big.Int
	q.Add(n, big.NewInt(-1)) // q = n - 1
	k := big.NewInt(0)       // k = 0

	//generate rand b in range [2,n-1]
	// (max - min) + min
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	var rng, b big.Int
	rng.Add(&q, big.NewInt(-2)) //	(max - min) + min
	b.Rand(r, &rng).Add(&b, big.NewInt(2))

	// while q is odd (par)
	for tmp.QuoRem(&q, big.NewInt(2), &p); p.Cmp(big.NewInt(0)) == 0; {
		k.Add(k, big.NewInt(1))
		q.Quo(&q, big.NewInt(2))
		tmp.QuoRem(&q, big.NewInt(2), &p)
	}
	// t = b**q mod n
	var t big.Int
	t.Exp(&b, &q, n)

	if t.Cmp(big.NewInt(1)) == 0 || t.Cmp(&q) == 0 {
		return false // inconclusive
	}

	for i := big.NewInt(1); i.Cmp(k) == -1; i.Add(i, big.NewInt(1)) {
		t.Exp(&t, big.NewInt(2), n)
		if t.Cmp(&q) == 0 {
			return false //inconclusive
		}
	}

	return true //composite
}

//func genRSA() {

//}

//func digitalSign() {

//}
