package authsvc

import (
	"math"
	"tauth/entities"
)

// Calculate the Euclidean distance between two slices of floats
func euclideanDistance(a, b []float64) float64 {
	if len(a) != len(b) {
		return math.Inf(1) // Return infinity if lengths are not equal
	}
	var sum float64
	for i := 0; i < len(a); i++ {
		diff := a[i] - b[i]
		sum += diff * diff
	}
	return math.Sqrt(sum)
}

// Function to compare stored profile with the new input
func compareKeystrokeProfiles(storedProfile, loginProfile entities.KeystrokeProfile) (bool, float64) {
	// Compare average dwell and flight times
	dwellDistance := math.Abs(storedProfile.AverageDwellTime - loginProfile.AverageDwellTime)
	flightDistance := math.Abs(storedProfile.AverageFlightTime - loginProfile.AverageFlightTime)

	// Set thresholds based on experimentation or statistical analysis
	const dwellThreshold = 20.0
	const flightThreshold = 20.0

	// Check if the dwell and flight distances are within acceptable ranges
	if dwellDistance > dwellThreshold || flightDistance > flightThreshold {
		return false, 0
	}

	// Compare entire dwell times and flight times if more accuracy needed
	dwellTimesDistance := euclideanDistance(storedProfile.DwellTimes, loginProfile.DwellTimes)
	flightTimesDistance := euclideanDistance(storedProfile.FlightTimes, loginProfile.FlightTimes)

	// Set threshold for overall similarity
	const similarityThreshold = 50.0

	// Total distance can be used to determine how close the typing patterns are
	totalDistance := dwellTimesDistance + flightTimesDistance
	isMatch := totalDistance < similarityThreshold
	return isMatch, totalDistance
}

func average(items []float64) float64 {
	var sum float64

	if len(items) == 0 {
		return 0
	}

	for _, item := range items {
		sum += item
	}

	return sum / float64(len(items))
}
