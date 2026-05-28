package techpalace

import (
    "strings"
    "fmt"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	//panic("Please implement the WelcomeMessage() function")
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	//panic("Please implement the AddBorder() function")
return fmt.Sprintln(strings.Repeat("*", numStarsPerLine)) +
    fmt.Sprintln(welcomeMsg) + fmt.Sprint(strings.Repeat("*", numStarsPerLine))
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	//panic("Please implement the CleanupMessage() function")
 return  strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}
