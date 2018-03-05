package functions

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

//	Constants values used a lot.
var (
	ZERO = big.NewInt(0)
	ONE  = big.NewInt(1)
	TWO  = big.NewInt(2)
)

//GenPrime generates a prime number with n bits of lenght
func GenPrime(n int) *big.Int {
	var b []byte
	b = make([]byte, n, n)
	key := new(big.Int)
	for {
		rand.Read(b[:])
		if int64(binary.LittleEndian.Uint64(b[:])) > 0 {
			key.SetBytes(b[:])

			isPrime := MillerRabin(key, 2)

			if isPrime {
				break
			}

		}
	}
	fmt.Println("genkey : ", key)
	return key

}

//GetPrimitiveRoot gets the primitive root of U(p)  for the el gamal algo
//p = prime number  and return the g (primitive root )
func GetPrimitiveRoot(p *big.Int) big.Int {
	fmt.Println("GetPrimRoot")

	ao := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	if p.Cmp(TWO) == 0 {
		return *ONE
	}
	var (
		g, p2, tmp, tmp3, mod, mod2 big.Int
		p1                          = TWO
	)
	fmt.Println("p:", &p)
	p2.Sub(p, ONE)
	fmt.Println(&p2)
	tmp.Sub(p, ONE).Quo(&tmp, p1)
	tmp3.Sub(p, ONE).Quo(&tmp3, &p2)
	//fmt.Println("p-1: ", p2)

	for {
		g.Rand(ao, &p2).Add(&g, TWO) // g gets a random integer btween 2  -- p-1
		// fmt.Println("random: ", g)

		if mod.Exp(&g, &tmp, p).Cmp(ONE) != 0 {
			if mod2.Exp(&g, &tmp3, p).Cmp(ONE) != 0 {
				fmt.Println("ge:", g)
				return g
			}
		}
	}
}

//ExtEuclidian returns the result of the extended  euclidian algorithm
//takes as parametes
//@TODO: explain parameters
func ExtEuclidian(n1, n2, x1, y1, x2, y2 big.Int) big.Int {
	var (
		quo, //quotient
		rem, //remainder
		tmp1,
		tmp2 big.Int
	)
	if n2.Cmp(ZERO) != 0 {
		quo.QuoRem(&n1, &n2, &rem)

		tmp1.Mul(&x2, &quo).Sub(&x1, &tmp1) //tmp1  = x1  - (x2*q)
		tmp2.Mul(&y2, &quo).Sub(&y1, &tmp2) //tmp1  = x1  - (x2*q)

		if rem.Cmp(ZERO) != 0 {
			return ExtEuclidian(n2, rem, x2, y2, tmp1, tmp2)
		}
		return x2

	}
	return *ONE

}

//MillerRabin functions  returns true if a number ,type big.Int, is probably prime
// or false if it is not, a composite number
func MillerRabin(n *big.Int, accuracy int) bool {
	//variables that are used as constants
	var (
		s                        = 0
		t, ck, number, max, x, a big.Int
	)

	// if number  is even and greather than 0 is composite
	if n.Cmp(TWO) > 0 && n.Bit(0) == 0 {
		return false
	}

	t.Sub(n, ONE) // n-1
	ck.Sub(n, ONE)
	for t.Bit(0) == 0 {
		t.Rsh(&t, 1) //t >>=1
		s++
		//t.Quo(&t, TWO)
	}
	number = *n
	//rand.Intn(max - min) + min
	max.Sub(n, ONE)

LOOP:
	for vez := 0; vez < accuracy; vez++ {

		ao := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
		a.Rand(ao, &max).Add(&a, ONE)

		x.Exp(&a, &t, n)

		if x.Cmp(ONE) == 0 || x.Cmp(&ck) == 0 {
			continue LOOP
		}
		for r := 1; r < s; r++ {

			x.Exp(&x, TWO, &number)

			if x.Cmp(ONE) == 0 {
				return false
			}
			if x.Cmp(&ck) == 0 {
				continue LOOP
			}
		}
		return false //composto
	}

	return true
}
