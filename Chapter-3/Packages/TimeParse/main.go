package main

import (
	"fmt"
	"time"
)
func main() {
	layout := "2006-01-02 15:04:05" //layout değişkeni, ayrıştırılacak dizenin beklenen zaman düzenini belirtir. 
	dateStr := "2023-07-25 10:30:45" //dateStr değişkeni, ayrıştırılacak olan tarih ve saati içeren dizedir.

	t, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Hata", err)
		return
	}
	fmt.Println("Zaman", t)
}


