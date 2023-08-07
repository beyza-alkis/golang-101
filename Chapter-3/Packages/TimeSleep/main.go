package main

import (
	"fmt"
	"time"
)
func main() {
	fmt.Println("Başlangıç zamanı: ", time.Now())

	time.Sleep(5 * time.Second)
	//10 dakika boyunca işlemi duraklatır ve sonrasında bitiş zamanını ekrana yazdırır.

	fmt.Println("Bitiş zamanı: ", time.Now())
}