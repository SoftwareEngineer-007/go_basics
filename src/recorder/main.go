package main

import "src/gadget"

type Player interface { // определяет тип интерфейса
	Play(string) // должен содержать метод Play с параметром string
	Stop()       // также необходим метод Stop
}

func playList(device Player, songs []string) {
	for _, song := range songs { // перебираем все песни в цикле
		device.Play(song) // воспроизведение текущей песни
	}
	device.Stop() // плеер останавливается после завершения
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"} // сегмент с названиями песен
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape) // песни воспроизводятся при помощи TapePlayer
	player = gadget.TapeRecorder{}
	playList(player, mixtape)
}
