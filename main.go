//Maintener : Yuri  

package main

//change import formats
import( "fmt"
        "os"
        str "strings"
        "math/big"
        "math/rand"
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
          


          r :=  generateKeys()
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
      if str.Compare(os.Args[3],"rsa") == 0    
       rsss := millerRabin( 55555 )
  
   }
   
   return f.MulRange(1, 10000000000000000)
}
func  millerRabin  ( n *big.Int) bool {
    //@TODO: check if n is greater than 3
    q := n - 1
    var k big.Int
    k= 1
    qt := true

        return qt
}
//func genRSA() {

//}

//func digitalSign() {

//}
