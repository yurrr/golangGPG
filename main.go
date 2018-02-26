//Maintener : Yuri

package main

//go run main.go genKey  crypt rsa
//change import formats
import (
	"encoding/binary"
	"fmt"
	fcns "golangGPG/functions"
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
	rand.Seed(time.Now().UTC().UnixNano())

	/*
		if str.Compare(os.Args[2], "sign") == 0 {
			fmt.Print(" /--* Starting digital assing *--\\ \n")
			digitalSign()
		}
		//(args[2], "sign") == 0 {
	*/
	if str.Compare(os.Args[1], "genKey") == 0 {
		fmt.Print(" /--* Starting generating Key *--\\ \n")
		/*
			fmt.Println("1234 - 1 0")
			fmt.Println("54 - 0 1")

			res := fcns.ExtEuclidian(*big.NewInt(1234), *big.NewInt(54), *big.NewInt(1), *big.NewInt(0), *big.NewInt(0), *big.NewInt(1))
			fmt.Println(res.Add(&res, big.NewInt(1)))
		*/
		generateKeys()
		//	fmt.Print(r)
	}
}

func generateKeys() *big.Int {

	var f big.Int

	if str.Compare(os.Args[2], "crypt") == 0 {
		if str.Compare(os.Args[3], "rsa") == 0 {
			// generate keys to rsa
			var (
				keys []big.Int
				cont = 0
			)
			for cont < 2 {
				var b [16]byte
				rand.Read(b[:])

				if int64(binary.LittleEndian.Uint64(b[:])) > 0 {
					key := new(big.Int)
					key.SetBytes(b[:])

					isPrime := fcns.MillerRabin(key, 2)

					if isPrime {
						fmt.Println("atual1:", key)
						keys = append(keys, *key)
						cont++
					}

				}
			}

			n, e := genPubRSA(&keys[0], &keys[1])
			_, d := genPrivRSA(&keys[0], &keys[1], &e)

			fmt.Println("n:", n)
			fmt.Println("e:", e)
			fmt.Println("d:", d)

		} else if str.Compare(os.Args[3], "elGamal") == 0 {
			var (
				b       [32]byte
				isPrime = false
			)
			for !isPrime {
				rand.Read(b[:])
				if int64(binary.LittleEndian.Uint64(b[:])) > 0 {
					key := new(big.Int)
					key.SetBytes(b[:])
					isPrime = fcns.MillerRabin(key, 2)

					if isPrime {
						fmt.Println("atual1:", key)
					}

				}
			}

		}
	}
	return f.MulRange(1, 2)
}
func genPrivRSA(p, q, e *big.Int) (big.Int, big.Int) {
	// /Privada: (n, d), onde n é como acima e ed ≡ 1 (mod φ(n)).
	/*
			fi = (p-1)*(q-1)
		    print fi
		    d = headerEuc(e,fi) % fi
		    print d
	*/

	var (
		n, fi, p2, q2 big.Int
	)
	p2.Add(p, big.NewInt(-1))
	q2.Add(q, big.NewInt(-1))
	fi.Mul(&p2, &q2)

	n.Mul(p, q)

	d := fcns.ExtEuclidian(*e, fi, *big.NewInt(1), *big.NewInt(0), *big.NewInt(0), *big.NewInt(1))
	d.Rem(&d, &fi)
	fmt.Println("n : ", n)
	fmt.Println("d : ", d)
	return n, d
}
func genPubGamal(p, q *big.Int) (big.Int, big.Int, big.Int) {

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
	fmt.Println("n : ", n)
	fmt.Println("e : ", e)
	return n, e
	//d.Exp

}

//func digitalSign() {

//}
