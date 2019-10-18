package main

import "fmt"

const NumLockers = 138

type LockerManager struct {
	isLockerRented   [NumLockers]bool
	numLockersRented int
}

// Find an empty locker, mark it as rented, return its number...
func (m *LockerManager) Rent(lockerNumber int) int {
	if !m.isLockerRented[lockerNumber] {
		m.isLockerRented[lockerNumber] = true
	}
	return lockerNumber
}

// Mark a locker as no longer rented
func (m *LockerManager) Release(lockerNumber int) {
	if m.isLockerRented[lockerNumber] {
		m.isLockerRented[lockerNumber] = false
	}
	return
}

// Say whether a locker is available for rent
func (m *LockerManager) IsFree(lockerNumber int) bool {
	if !m.isLockerRented[lockerNumber] {
		return true
	}
	return false
}

// Say whether any lockers are left to rent
//func (m *LockerManager) IsAvailable() bool {
//}

func main() {
	test := LockerManager{}
	test2 := LockerManager{}
	test2.Rent(2)
	fmt.Println(test.Rent(23))
	fmt.Println(test.IsFree(24))
}
