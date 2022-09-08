package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

// The fmt.Printf function uses this interface to print out values
type Stringer interface {
	String() string
}

/*
 *** The first thing you need to implement an interface is to create a custom type. ***
 */

// for Square struct to be said to have implemented the Shape interface
// it must implement all methods defined in the Shape interface
type Square struct {
	size float64
}

func (sq Square) Area() float64 {
	return sq.size * sq.size
}

func (sq Square) Perimeter() float64 {
	return sq.size * 4
}

// for Circle struct to be said to have implemented the Shape interface
// it must implement all methods defined in the Shape interface
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// printInformation prints out the type of the object it receives, along with its area and perimeter
// Notice how the printInformation function has Shape as a parameter. This means that -
// the printInformation function expects an object that implements the methods that are defined in the Shape interface.
// In our case - you could send to this function a Square or a Circle object, and it will work,
// although the output will be different. output varies depending on the object type it receives.
func printInformation(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println()
}

type Person struct {
	Name, Country string
}

// Person struct implements the stringer interface method string()
// method to print a custom string
func (p Person) String() string {
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}

func main() {
	var s Shape = Square{3}
	printInformation(s)

	// Notice that for the c object, we don't specify that it's a Shape object.
	// However, the printInformation function expects an object that implements
	// the methods that are defined in the Shape interface.
	c := Circle{6}
	printInformation(c)

	rs := Person{"John Doe", "USA"}
	ab := Person{"Mark Collins", "United Kingdom"}
	// The fmt.Printf function uses this interface to print out values
	fmt.Printf("%s\n%s\n", rs, ab)

	//
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Notice that the io.Copy(os.Stdout, resp.Body) call is the one that prints to the terminal
	// the content you've obtained from the call to the GitHub API.
	/*
			When you look at the source of the io.Copy function, you see the following:

			func Copy(dst Writer, src Reader) (written int64, err error)

			This mean io.Copy expects any object that implements the methods defined in the Writer interface
			since it receives a variable of type Writer interface - see below

			If you go deeper into the details of the first parameter, dst Writer, you'll notice that Writer is an
			interface:

			type Writer interface {
		    Write(p []byte) (n int, err error)

			Because Writer is an interface and it's an object that the Copy function is expecting, you could write your custom implementation of the Write method. Therefore, you can customize the content that you print to the terminal.

			Recall that the first thing you need to implement an interface is to create a custom type. In this case, you can create an empty struct, because you simply need to write your custom Write method.
		}


	*/

	//fmt.Println("Output using the inbuilt Writer method")
	//io.Copy(os.Stdout, resp.Body)

	// Finally, you have to modify the main() function to use your custom object
	writer := customWriter{}

	fmt.Println("Output using our custom Writer method")
	io.Copy(writer, resp.Body)

	// Declare and initialize a database struct type instance
	db := database{"Go T-Shirt": 25, "Go Jacket": 55}

	fmt.Printf("db type: %v", db)

	// ListenAndServe function expects any object that implements the methods defined in the Handler interface
	log.Fatal(http.ListenAndServe("localhost:8000", db))

	/*
			   *** Output ***
			   Go T-Shirt: $25.00
		       Go Jacket: $55.00
	*/
}

// CustomWriter implements the Writer interface
type customWriter struct{}

//  GitHubResponse struct parses the API response in JSON format to a Golang object.
type GitHubResponse []struct {
	FullName string `json:"full_name"`
}

// Now our customeWriter struct implements the Writer interface
// Notice how the Write method signature matches the method declared in the Writer interface
func (w customWriter) Write(p []byte) (n int, err error) {
	var resp GitHubResponse
	json.Unmarshal(p, &resp)
	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	return len(p), nil
}

/*
   ***** Another Example - Using the Handler interface *****

   type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(address string, h Handler) error

Notice how the ListenAndServe function is expecting a server address, such as http://localhost:8000, and an instance of the Handler that will dispatch the response from the call to the server address.
*/

// dollars is a custom type for a float32 type, with the idea of writing a custom implementation of the String() method
type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

// Now our database struct implements the Handler interface
// Notice how the ServerHTTP method signature matches the method declared in the Handler interface
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
