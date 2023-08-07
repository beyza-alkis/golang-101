package main

import "fmt"

// 1 - add = fonksiyon ismimizin adi, parantez icerisine fonksiyonun alacagi parametreler icin degisken ve tip tanimlamasi yapiyoruz.
func add(a int, b int)int{
	return a + b // a + b toplamini donduruyoruz.
}

// 2 - fonksiyonlar parametresiz de olabilir.
func print(){
	fmt.Println("Write down")
}

// 3 - fonksiyonlarin tum degisken tipleri ayni ise sadece en sagdaki degiskenin tipini belirtmek yeterlidir.
func operation(number int)(a, b, c int){ // return yapacak degiskenleri tanimliyoruz
	a = number * 5
	b = number / 2
	c = number + 7
	return 
}


func main(){
	// 1 numarali islemin ciktisi icin;
	fmt.Println(add(10, 7)) // sonucu ekrana bastiriyoruz. 

	// 2 numarali islemin ciktisi icin;
	print()
	
	// 3 numarali islemin ciktisi icin;
	fmt.Println(operation(20))
}