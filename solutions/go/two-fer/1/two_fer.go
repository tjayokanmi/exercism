
package twofer


func ShareWith(name string) string {
    result := ""
if name == ""{
result = "One for you, one for me."
} else {
    result = "One for " + name+ ", one for me."
}
    
	
	return result
}
