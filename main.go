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
		} else if str.Compare(os.Args[3], "el gamal") == 0 {
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

func genPubRSA(p, q *big.Int) {
	var (
		n big.Int
	)
	n.Mul(p, q)

}

//func digitalSign() {

//}
