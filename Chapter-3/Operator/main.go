// Degisken atama
package main

import (
	"fmt"
)
func main(){
	// string int yazmak zorunda degiliz, golang bunu anliyor ve ona gore islem yapiyor.
	var name string = "Beyza"
	var surname string = "Alkis"
	var age int = 23
	var student bool = false
	fmt.Println(name, surname, age, student)


	// var yazmadan da degisken atamasi yapabiliyoruz
	school_name := "Bursa Uludag Universitesi"
	education_name := "Yonetim Bilisim Sistemleri"
	grad_year := 2023
	fmt.Println(school_name, education_name, grad_year)


	// bunlara ek olarak sabit degsikenlere(const) deger atadiktan sonra degistiremeyiz.
	// const ile atama degeri olan ":=" bu deger kullanilamaz.
	const Pi float64 = 3.141
	fmt.Println("Pi: ", Pi)


	// kodlarimizi gruplamak istersek & isareti ile bellekteki degerini ogrenerek destek aliyoruz.
	// bu iki degiskenin ayri degiskenler oldugunu anliyoruz.
	variables := "cat"
	{
		variables := "dog"
		fmt.Println(variables)
		fmt.Println(&variables)
	}
	fmt.Println(variables)
	fmt.Println(&variables)
}