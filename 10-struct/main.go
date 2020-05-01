package main

import "fmt"

type Sepeda struct {
	jumlahBan, jumlahGigi int
	warna string
}	

func (s *Sepeda) waktuTempuh(jarak float32) float32 {
	waktu := jarak * 2.5 
	return waktu
}

func main() {

	var sepeda = make(map[int]Sepeda)

	sepeda[0] = Sepeda{jumlahBan:2, jumlahGigi:1, warna:"Hitam"} 
	sepeda[1] = Sepeda{jumlahBan:2, jumlahGigi:2, warna:"Putih"} 
	sepeda[2] = Sepeda{jumlahBan:2, jumlahGigi:3, warna:"Biru"} 
	sepeda[3] = Sepeda{jumlahBan:2, jumlahGigi:4, warna:"Merah"} 
	sepeda[4] = Sepeda{jumlahBan:2, jumlahGigi:5, warna:"Abu-abu"} 

	// s := Sepeda{}
	// fmt.Println("Sepeda")
	// fmt.Println(s.waktuTempuh(20))
	

	for _, item := range sepeda {
		fmt.Println(
					"Sepeda", item.warna, 
					"Beroda", item.jumlahBan, 
					"dengan Gigi", item.jumlahGigi,
					"Menempuh jarak 20km dengan waktu", item.waktuTempuh(20), "Menit",
				  )
	}
}