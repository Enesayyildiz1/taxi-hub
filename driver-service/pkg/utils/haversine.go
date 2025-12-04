package utils

import (
	"math"
)

const (
	EarthRadiusKm = 6371.0
)

func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	lat1Rad := degreesToRadians(lat1)
	lon1Rad := degreesToRadians(lon1)
	lat2Rad := degreesToRadians(lat2)
	lon2Rad := degreesToRadians(lon2)

	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := EarthRadiusKm * c

	return distance
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func IsWithinRadius(centerLat, centerLon, pointLat, pointLon, radiusKm float64) bool {
	distance := CalculateDistance(centerLat, centerLon, pointLat, pointLon)
	return distance <= radiusKm
}
