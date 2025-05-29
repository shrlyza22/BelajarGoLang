package main
import (
 "fmt"
 "strings"
 "strconv"
 )

func adaBilanganM(n, m int) string {
	
	s1 := strconv.Itoa(n)
	s2 := strconv.Itoa(m)
		
		if strings.Contains(s1,s2) { 
			return "YA" 
	} else {
		return "TIDAK"
	}
}
func main() {
	var N, M int

	fmt.Print("Masukkan N dan M: ")
	fmt.Scan(&N, &M)

	fmt.Println(adaBilanganM(N, M))
}