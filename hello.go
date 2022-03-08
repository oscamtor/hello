
// Declare a main package. In Go, code executed as an application must be in a main package.
package main


// Import two packages: example.com/greetings and the fmt package. This gives your code access to functions in those packages. Importing oscamtor/grettings/v2 (the package contained in the module you created earlier) gives you access to the Hello() function. You also import fmt, with functions for handling input and output text (such as printing text to the console)
import (
    "fmt"

    "oscamtor/grettings/v2"
)

// Get a greeting by calling the greetings package’s Hello function.
func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Oscar")
    fmt.Println(message)
}