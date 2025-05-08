package main
import (
	"fmt"
)

func main(){
	var time float64
	var hasil int
	fmt.Scan(&time)
	hasil = litters(time)
	fmt.Print(hasil)
}
func litters(time float64)int{
	var i int
	var liter float64
	for i = 1; i <= int(time); i++{
		liter += 0.5
	}
	return int(liter)
}
