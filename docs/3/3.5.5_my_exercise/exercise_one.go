package go_unit_test_bootcamp

/*
import (
    "fmt"
)
*/

// Takes an array of ints, returns an int
func FindMissingDrone(droneIds []int) int {
    // TODO: implement this

	for _, i := range droneIds {

		var counter int = 0
		var target_drone int

		for _, j := range droneIds {
			if i == j {
				counter = counter + 1
				target_drone = i
			}
		}

		if counter == 1 {
			return target_drone
		}

    }

    return -1
}