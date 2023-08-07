package main

import (
	"fmt"
	"time"
)
func main(){
	duration := 5 * time.Second

	fmt.Println("Bekleme başladı: ", time.Now())

	<-time.After(duration)
	// Belirtilen süre sonra bir kanalda değer gönderilir

	fmt.Println("Bekleme Sona Erdi: ", time.Now())
}