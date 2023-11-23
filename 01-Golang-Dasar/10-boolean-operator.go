package main

import "fmt"

func main() {
	var ujian = 80
	var absesnsi = 80

	var lulusUjian = ujian > 80
	var lulusAbsensi = absesnsi > 80

	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println(lulus)

	fmt.Println(ujian >= 80 && absesnsi >= 80)
}
