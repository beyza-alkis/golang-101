# golang-101
Learn Golang Zero To Mastery With Me 🦋

![Image of Gopher](https://res.cloudinary.com/practicaldev/image/fetch/s--8-yHySDD--/c_imagga_scale,f_auto,fl_progressive,h_900,q_auto,w_1600/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/db3m4zatbzjqhkb0v05c.jpg)

Golang öğrenirken aldığım notlarla bir kaynakça oluşturmaya karar verdim. Repo'mu tamamen bitirdikten sonra bu satırlara tekrar döneceğim. :)

--------

* [Chapter-1](https://github.com/beyza-alkis/golang-101/blob/master/chapter-1/main.go)

 **"Hello, World!" Örneği**
 
Bir programlama dilini öğrenirken geleneği bozmadım ve ilk çıktımızı "Hello, World!" ile oluşturmak istedim. Kodlarımız bu şekilde, peki bu kodlar ne işe yarıyor?
Gelin birlikte öğrenelim. 😊


```
package main

import "fmt"
func main(){
	fmt.Println("Hello, World!")
}
```

- Package: Go programlama dili packagelar ile birlikte çalışıyor bu yüzden dosyada package olmak zorunda. İçerisinde package olan kod sayfalarıyla iletişim halinde olmamızı sağlar. Go programları paketler halinde çalıştığı için bu zorunlu bir ifadedir. 'main' ifadesi ise bizim dosyamızın adıdır.
- import "fmt": 'import' terimi yazıldığı pakete başka bir paketten özellik eklemeye yarar. 
