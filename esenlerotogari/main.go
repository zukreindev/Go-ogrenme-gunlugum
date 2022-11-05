package main

import (
	"fmt"
)

func main() {
	type yolculuk struct {
		nereden string
		nereye  string
		saat    float64
	}

	yolculuk1 := yolculuk{
		nereden: "Ankara",
		nereye:  "İstanbul",
		saat:    3.5,
	}

	anlıkSAAT := 1.3

	var name string
	var nereye string
	var saat = 0.0

	fmt.Println("Merhaba Ben esenler otogarı otobüs bileti satan robotum. Sana nasıl yardımcı olabilirim?")
	fmt.Println("Saat:", anlıkSAAT)
	fmt.Println("Hangi şehirdesiniz")
	fmt.Scan(&name)
	fmt.Println("Nereye gidiyorsunuz?")
	fmt.Scan(&nereye)
	fmt.Println("Saat kaçta gidicen abim?")
	fmt.Scan(&saat)

	yolculuk2 := yolculuk{
		nereden: name,
		nereye:  nereye,
		saat:    saat,
	}
	if yolculuk1.nereye == yolculuk2.nereye {
		if yolculuk1.saat < yolculuk2.saat {
			fmt.Printf("Malesef Belirttiğiniz saatlere uygun yolculuğumuz bulunamamıştır :(")
		} else if yolculuk1.nereden != yolculuk2.nereden {
			fmt.Printf(yolculuk2.nereye, " Yolculuğumuz mevcuttur lakin", yolculuk2.nereden, " Şehrinden kalkacaktır")
		} else {
			fmt.Printf("Yolculuk sırasına alındız")

		}
	} else {
		fmt.Println("Malesef paşam başka bir şehirden geliyorsunuz")
	}

}
