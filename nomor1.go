package main
import "fmt"

func Reamur(C float64) float64 {
	return 4.0/5.0* C 
}
func Fahrenheit(C float64) float64{
	return 9.0/5.0 * C + 32
}
func Kelvin(C float64) float64 {
	return C + 273
}

func main () {
	var awal, akhir, langkah float64 
	fmt.Print("Masukkan suhu awal, suhu akhir, dan langkah:")
	fmt.Scan(&awal,&akhir,&langkah)
	
	fmt.Println("Celcius Reamur Fahrenheit Kelvin")
		for C := awal; C <= akhir; C += langkah {
		fmt.Printf("%.2f   %.2f   %.2f   %.2f\n", C, Reamur(C), Fahrenheit(C), Kelvin(C))
		}	
}