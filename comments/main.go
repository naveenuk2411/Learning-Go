// A documentation comment should be a complete sentence that starts with the name of the thing being described and ends with a period.

// Package comments should be written directly before a package clause (package x) and begin with Package x ...
// Package weather provides tool to tell the forecast conditions of a city.
package weather

// CurrentCondition represents the current forecast condition.
var CurrentCondition string

// CurrentLocation represents the current location of the city.
var CurrentLocation string

// A function comment should be written directly before the function declaration. It should be a full sentence that starts with the
// function name.
// For example, an exported comment for the function Calculate should take the form Calculate ....
// Forecast takes in the location of the city and the forecase condition and returns the weather forecast string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
