package annalyn
// CanFastAttack can be executed only when the knight is sleeping.

var knightIsAwake = false
var archerIsAwake = true
var prisonerIsAwake = false
var petDogIsPresent = false

func CanFastAttack(knightIsAwake bool) bool {
   // var knightIsAwake = true
    if knightIsAwake == false {
        return true
    } else {
        return false
    }
	panic("Please implement the CanFastAttack() function")
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
  if knightIsAwake == true || archerIsAwake == true || prisonerIsAwake == true {
        return true 
    } else {
        return false
    }
	panic("Please implement the CanSpy() function")
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
    if prisonerIsAwake == true && archerIsAwake == false {
        return true
    } else {
        return false
    }
	panic("Please implement the CanSignalPrisoner() function")
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
     if (petDogIsPresent == true && archerIsAwake == false) || (prisonerIsAwake == true && (knightIsAwake == false && archerIsAwake ==false)) {
        return true 
    } else {
       return false
    }
	panic("Please implement the CanFreePrisoner() function")
}
