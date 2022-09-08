package gadget

import "fmt"

// Declares a struct type
// We will make it satisfy an interface
type TapePlayer struct {
	Batteries string
}

// Declares a struct type
// We will make it satisfy an interface
type TapeRecorder struct {
	Microphones int
}

// Notice how the method's signature of the TapePlayer struct matches 
// the signature of the Player interface.
func (t TapePlayer) Play(song string) {
	fmt.Println("Playing ", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Playlist stopped")
}

func (r TapeRecorder) Play(song string) {
	fmt.Println("Recording ", song)
}

// A type can still satisfy an interface even if it has methods of its own
// Note: method Recording() not defined in the interface
func (r TapeRecorder) Record() {
	fmt.Println("Starts recording ")
}

func (r TapeRecorder) Stop() {
	fmt.Println("Recording stopped")
}

/*
If a type declares methods with pointer receivers,
then you’ll only be able to use pointers to that type
when assigning to interface variables.
The toggle method on the Switch type below has to use a
pointer receiver so it can modify the receiver.
*/

type Switch string

func (s *Switch) Toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}
