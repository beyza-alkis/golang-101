package main

import "fmt"
func SumNumbers(numbers ...int)int{
	// dogru deger vermesi icin total degeri 0 
	total := 0
	// for dongu baslatmak icin kullaniliyor.
	// _ degisken olarak kullanilmayacak degeri temsil eder, hata almamak icin kullanilabilir.
	// range dongu olusturmak icin kullanilir yani fonksiyonun icerisine ne kadar sayi eklersek islem ona gore sekillenir.
	for _, num := range numbers{
		total += num
		total *= num
	}
	return total
}
func main(){
	result1 := SumNumbers(1,2,3,4,5)
	result2 := SumNumbers(10,20,30,40,50)
	result3 := SumNumbers(100,200,300,400,500)
	fmt.Println("Result 1: ", result1)
	fmt.Println("Result 2: ", result2)
	fmt.Println("Result 3: ", result3)
}