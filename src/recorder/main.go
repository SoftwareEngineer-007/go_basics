package main

import gadget "command-line-argumentsD:\\Projects\\go_basics\\src\\gadget\\tape.go"

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs { // перебираем все песни в цикле
		device.Play(song) // воспроизведение текущей песни
	}
	device.Stop() // плеер останавливается после завершения
}

func main() {
	player := gadget.TapePlayer{}                             // создание TapePlayer
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"} // сегмент с названиями песен
	playList(player, mixtape)                                 // песни воспроизводятся при помощи TapePlayer
}
