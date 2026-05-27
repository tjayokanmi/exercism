// Package weather provides tool for Goblinocus weather forecast, it tools help predict the weather.
package weather 

// CurrentCondition represent the Weather current condition either it's windy, rainy, etc.
var CurrentCondition string 

// CurrentLocation represents a location in Goblinocus that the weather is being forecasted.
var CurrentLocation  string 

// Forecast function return the weather condition for a specified city.
func Forecast(city, condition string) string { 
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
