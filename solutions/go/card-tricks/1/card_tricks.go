package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	
    return []int{2,6,9}
//    panic("Please implement the FavoriteCards function")
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
   var result int
    if len(slice) - 1 >= index && index >= 0 {
    result = slice[index]
    } else {
        result = -1
    }
    return result
	//panic("Please implement the GetItem function")
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    if len(slice) - 1 >= index && index >= 0 {
        slice[index] = value
    } else {
        slice = append(slice, value)
    }
    return slice
	//panic("Please implement the SetItem function")
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
    var slices []int
    if len(values) == 0 {
        slices = slice
    } else {
   // for i, v := range values {
        slices = append(values, slice...)
     //  i++
   // }
    
    }
    return slices 
	//panic("Please implement the PrependItems function")
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if len(slice) >= index && index >= 0 {
       // slice = append(slice[0:index],slice[index+1:])
        slice = append(slice[:index], slice[index+1:]...)
    } 
    return slice
    }
    //panic("Please implement the RemoveItem function")
