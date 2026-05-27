package collatzconjecture

import (
    "errors"
)

func CollatzConjecture(n int) (int, error) {
	count := 0
	//result := n
 	if n < 1 {
        return 0, errors.New("Only positive number")
        }
    for n != 1 {
        if n%2 == 0 {
            n = n / 2
        } else {
            n = 3*n + 1
            // count += 1
        }
        count += 1

    }
    return count, nil
}

