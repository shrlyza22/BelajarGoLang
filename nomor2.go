package main

import ( 
	"fmt"
	"strings"
)


func lowToUpper(H string) string {
	return strings.ToUpper(H)
}

func main() {
	var huruf string

	
	fmt.Print("Masukkan huruf kecil: ")
	fmt.Scan(&huruf)

	
	fmt.Printf("Huruf besar: %s\n", lowToUpper(huruf))
}
