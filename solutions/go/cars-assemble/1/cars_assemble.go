package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return float64(productionRate) * (successRate / 100)
	panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    minutes := (float64(productionRate) * (successRate/100))/ 60.0
    return int(minutes)
    
	panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    tens := carsCount/10 
    remainder := carsCount - (tens * 10)
  	cost :=  (tens * 95000) + (remainder * 10000)
    return uint(cost)
	panic("CalculateCost not implemented")
}
