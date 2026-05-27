package raindrops

import (
	//"fmt"
	"strconv"
	//    "fmt"
)

func Convert(number int) string {
	var result string
 //    if number%3 == 0 && number%5 == 0 && number%7 == 0 {
	// 	result += "PlingPlangPlong" 
 //    } else if number%3 == 0 && number%5 == 0 {
	// 	result += "PlingPlang"
 //        }	else if number%3 == 0 && number%7 == 0 {
	// 	result += "PlingPlong"
	// } else if number%5 == 0 && number%7 == 0 {
	// 	result += "PlangPlong" 
 if number%3 == 0 {
		result += "Pling"
     }
if number%5 == 0 {
		result += "Plang"
	}  
if number%7 == 0 {
		result += "Plong"
	} 
if result == "" {
		result += strconv.Itoa(number)
	}

return result

	// panic("Please implement the Convert function")
}