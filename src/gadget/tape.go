package gadget

import "fmt"

// ...определение типа TapePlayer...
type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

// ...определение типа TapeRecorder...
type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) { // содержит метод Play, как у TapePlayer
	fmt.Println("Playing", song)
}
func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}
func (t TapeRecorder) Stop() { // содержит метод Stop, как у TapePlayer
	fmt.Println("Stopped!")
}
