- [Chapter-2]()
  **Int(Integer Veri Tipi)**

Birçok veri tipinin olması sistem hafızasında yer sorunu oluşturmaktadır. Hafızada kapladığı yere göre minik bir özet geçerek integer veri tipleri aşağıda yer almaktadır.

| Veri Tipi |        Açıklama         |                                   Veri Aralığı |
| --------- | :---------------------: | ---------------------------------------------: |
| int8      |  8-bit signed integer   |        -127 ve 128 değerleri arasında yer alır |
| int16     |  16-bit signed integer  | -32,768 ve +32,767 değerleri arasında yer alır |
| uint8     | 8-bit unsigned integer  |           0 ve 255 değerleri arasında yer alır |
| uint16    | 16-bit unsigned integer |         0 ve 65535 değerleri arasında yer alır |

```
package main

import ("fmt")
func main()  {
	var mathGrade  = 90
	var physicsGrade  = 70
	var chemistryGrade  = 83
	fmt.Println(mathGrade, physicsGrade, chemistryGrade)
	//average of grades
	fmt.Println((mathGrade+physicsGrade+chemistryGrade)/3)
}
```

**Float(Ondalıklı) Veri Tipi**

- Float virgül ya da nokta ile ayrılmış ondalıklı değerleri ifade eder.
- Pozitif ve negatif sayıları küsüratları ile tutar ve küsüratlar nokta (.) ile gösterilir.
