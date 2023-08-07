package main

import (
	"fmt"
	"time"
)
func main(){
	// 1 Saat 15 Dakika 25 Saniye Süre
	duration := 1 * time.Hour + 15 * time.Minute + 25 * time.Second

	// Süreyi saniye cinsinden hesaplamak
	seconds := duration.Seconds()
	fmt.Printf("Toplam süre: %v\n", duration)
	fmt.Printf("Toplam süre saniye: %f\n", seconds)


	oneHourLater := time.Now().Add(1 * time.Hour)
	fmt.Printf("1 saat sonra: %v\n", oneHourLater)


	oneHourAgo := time.Now().Add(-1 * time.Hour)
	fmt.Printf("1 saat önce: %v\n", oneHourAgo)
}