package authsvc

import (
	"errors"
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
func compareKeystrokeProfiles(storedProfile, loginProfile entities.KeystrokeProfile) (bool, float64, error) {
	// Validate input profiles
	if err := validateProfiles(storedProfile, loginProfile); err != nil {
		return false, 0, err
	}

	// Compare sequence lengths
	if len(storedProfile.DwellTimes) != len(loginProfile.DwellTimes) ||
		len(storedProfile.FlightTimes) != len(loginProfile.FlightTimes) {
		return false, 0, errors.New("keystroke sequence length mismatch")
	}

	// Calculate normalized differences for average times
	dwellTimeDiff := math.Abs(storedProfile.AverageDwellTime-loginProfile.AverageDwellTime) / storedProfile.AverageDwellTime
	flightTimeDiff := math.Abs(storedProfile.AverageFlightTime-loginProfile.AverageFlightTime) / storedProfile.AverageFlightTime

	// Stricter thresholds for average time differences (20% variation)
	const avgTimeThreshold = 0.20
	if dwellTimeDiff > avgTimeThreshold || flightTimeDiff > avgTimeThreshold {
		return false, math.Max(dwellTimeDiff, flightTimeDiff) * 100, nil
	}

	// Calculate detailed pattern similarity using dynamic time warping
	dwellSimilarity := calculateDTWDistance(storedProfile.DwellTimes, loginProfile.DwellTimes)
	flightSimilarity := calculateDTWDistance(storedProfile.FlightTimes, loginProfile.FlightTimes)

	// Calculate weighted total similarity score (0-100 scale)
	const dwellWeight = 0.6
	const flightWeight = 0.4
	totalSimilarity := (dwellSimilarity*dwellWeight + flightSimilarity*flightWeight) * 100

	// Threshold for accepting the login attempt
	const similarityThreshold = 75.0 // Requires 75% similarity
	isMatch := totalSimilarity >= similarityThreshold

	return isMatch, totalSimilarity, nil
}

// validateProfiles checks if the profiles contain valid data
func validateProfiles(stored, login entities.KeystrokeProfile) error {
	if len(stored.DwellTimes) == 0 || len(stored.FlightTimes) == 0 ||
		len(login.DwellTimes) == 0 || len(login.FlightTimes) == 0 {
		return errors.New("empty keystroke profiles")
	}

	// Check for negative times
	for _, t := range append(stored.DwellTimes, stored.FlightTimes...) {
		if t < 0 {
			return errors.New("negative time values in stored profile")
		}
	}
	for _, t := range append(login.DwellTimes, login.FlightTimes...) {
		if t < 0 {
			return errors.New("negative time values in login profile")
		}
	}

	return nil
}

// calculateDTWDistance implements Dynamic Time Warping algorithm to compare sequences
func calculateDTWDistance(s1, s2 []float64) float64 {
	n, m := len(s1), len(s2)
	dtw := make([][]float64, n+1)
	for i := range dtw {
		dtw[i] = make([]float64, m+1)
		for j := range dtw[i] {
			dtw[i][j] = math.Inf(1)
		}
	}
	dtw[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			cost := math.Abs(s1[i-1] - s2[j-1])
			dtw[i][j] = cost + min(dtw[i-1][j], dtw[i][j-1], dtw[i-1][j-1])
		}
	}

	// Normalize the distance to a similarity score between 0 and 1
	maxPossibleDist := float64(n+m) * math.Max(maxSlice(s1), maxSlice(s2))
	similarity := 1 - (dtw[n][m] / maxPossibleDist)
	return similarity
}

// Helper functions
func min(numbers ...float64) float64 {
	result := numbers[0]
	for _, num := range numbers[1:] {
		if num < result {
			result = num
		}
	}
	return result
}

func maxSlice(s []float64) float64 {
	if len(s) == 0 {
		return 0
	}
	max := s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}
	return max
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
