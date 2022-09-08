package main

import (
	"fmt"
	"git_repo/go_playground/playground/go_interfaces/gadget"
)

/*
Interfaces in Go are a type of data that's used to represent the behavior of other types.
An interface is like a blueprint or a contract that an object should satisfy.
Interfaces in Go are satisfied implicitly. Go doesn't offer keywords to implement an interface

An interface in Go is an abstract type that includes only the methods that a concrete type must possess or implement. That's why we say that an interface is like a blueprint.
*/

// Declare a Player Interface
// Note: A type satisfies this interface if it implements the two methods defined
// in the interface
// However, another interface might have a different name but the same methods. 
// How or when does Go know which interface a concrete type is implementing? 
// Go knows it when you're using it, at runtime.
// When you have a value in a variable  with an interface type, you can only call
// methods deﬁned as part of that interface, regardless of what methods
// the concrete type had.
type Player interface {
	Play(string) // Require a Play() method with a string parameter
	Stop()       // Also requires a Stop() method
}

// Defines a normal function that accepts any player
// If and only if it satisifies the interface - Implements everything defined in the interface

func playlist(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song) // satisfies the interface Player

	}

	device.Stop() // satifies the Interface Player

}

type Toggleable interface {
	Toggle()
}

func main() {
	// Defines a slice of playlist songs
	mixtape := []string{"Just a dream", "Tapout", "Godly", "Roju"}

	// Declare a variable using the interface

	// Declare a variable using the interface
	// Since TapePlayer struct satisfies the interface
	// We can assign the variable player of type Player interface to it
	var player Player = gadget.TapePlayer{}

	// Then we can call any method that is part of the interface
	// Pass a TapePlayer to playlist
	playlist(player, mixtape)

	// Note: player can only call methods defined in the interface
	// We cannot call Record() since it is not defined in the interface
	player = gadget.TapeRecorder{}

	// player.Record() -> not allawed

	// Define a variable of type TapeRecorder to access Record()
	recorder := gadget.TapeRecorder{}

	recorder.Record()

	playlist(player, mixtape)

	// The Toggle method on the Switch type has to use a
	// pointer receiver so it can modify the receiver.

	s := gadget.Switch("Off")

	// assign a pointer to a Switch to the Toggleable variable,
	// instead of a direct Switch value:
	var t Toggleable = &s // Pointer receiver

	t.Toggle()
	t.Toggle()

	// Type Asssertion
	/*
		When you have a value of a concrete type eg - recorder above assigned to a
		variable with an interface type eg player above, a type assertion lets you
		get the concrete type back. It’s kind of like a type conversion.

		Once you’ve used a type assertion to get a value of a
		concrete type back, you can call methods on it that are
		deﬁned on that type, but aren’t part of the interface.
	*/

	// Converts back to Concrete type using Type Asssertion
	var myRecorder gadget.TapeRecorder = player.(gadget.TapeRecorder)

	// Can now call a method that is defined on the Concrete type(Not the interface)
	myRecorder.Record()

	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
}

//Testing Type Asssertion
func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()

	// Converts back to Concrete type using Type Asssertion
	// Assign the second return value to a variable ok - comma ok idiom
	// This helps you avoid run time panics
	myRec, ok := player.(gadget.TapeRecorder)

	// Call the Record method if the original value was a TapeRecorder

	if ok {
		myRec.Record()
	} else {
		fmt.Println("Player is not a TapeRecorder")
	}

}
