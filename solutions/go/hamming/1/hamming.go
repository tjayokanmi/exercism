package hamming

import (
    "errors"
)
func Distance(a, b string) (int, error) {
	var distance int
    if len(a) != len(b) {
        return 0, errors.New("Sequence of different length")
    }
    for i := 0; i < len(a); i++ {
    if a[i] != b[i]{
        distance++
    }
        }
    return distance, nil
    
	panic("Implement the Distance function")
}
