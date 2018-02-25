package functions

import (
	"math/big"
	"math/rand"
	"time"
)

//MillerRabin functions  returns true if a number ,type big.Int, is probably prime
// or false if it is not, a composite number
func MillerRabin(n *big.Int, accuracy int) bool {
	//variables that are used as constants
	var (
		two                      = big.NewInt(2)
		minusOne                 = big.NewInt(-1)
		one                      = big.NewInt(1)
		s                        = 0
		t, ck, number, max, x, a big.Int
	)

	// if number  is even and greather than 0 is composite
	if n.Cmp(two) > 0 && n.Bit(0) == 0 {
		return false
	}

	t.Add(n, minusOne) // n-1
	ck.Add(n, minusOne)
	for t.Bit(0) == 0 {
		t.Rsh(&t, 1) //t >>=1
		s++
		//t.Quo(&t, TWO)
	}

	//certo ate aqui
	number = *n
	//rand.Intn(max - min) + min
	max.Add(n, minusOne)

LOOP:
	for vez := 0; vez < accuracy; vez++ {

		ao := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
		a.Rand(ao, &max).Add(&a, one)

		x.Exp(&a, &t, n)

		if x.Cmp(one) == 0 || x.Cmp(&ck) == 0 {
			continue LOOP
		}
		for r := 1; r < s; r++ {

			x.Exp(&x, two, &number)

			if x.Cmp(one) == 0 {
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
