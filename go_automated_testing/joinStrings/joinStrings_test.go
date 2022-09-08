// This test code will be part of the same package as the code we are testing
/*
To run tests, you use the go test command. The command
takes the import paths of one or more packages, just like go
install or go doc . It will ﬁnd all ﬁles in those package
directories whose names end in _test.go, and run every
function contained in those ﬁles whose name starts with
Test .
*/
package joinstrings

import "testing"

// Note: Func name should begin with the word Test Capital letter
// Name after Test can be whatever you want
// Note: Func will be passed a pointer to a testing T value
func TestTwoElements(t *testing.T) {
	// Call a method on the testing T to fail the test
	t.Error("No test written yet")
}

func TestThreeElements(t *testing.T) {
	// Call a method on the testing T to fail the test
	t.Error("No test written yet")
}
