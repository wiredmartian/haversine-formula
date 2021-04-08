package main

/* HAVERSINE FORMULA
	a = sin²(φB - φA/2) + cos φA * cos φB * sin²(λB - λA/2)
	c = 2 * atan2( √a, √(1−a) )
	d = R ⋅ c
*/
import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println("Hello from go")

	// ==== Test formula

	// point 1 GPS coordinates
	point1 := Coordinates {
		latitude: 51.510357,
		longitude: -0.116773,
	}

	// point 2 GPS coordinates
	point2 := Coordinates {
		latitude: 38.889931,
		longitude: -77.009003,
	}

	// distance in km
	distance := haversine(point1, point2)

	fmt.Printf("Distance is %.3fkm", distance)
}

// Coordinates coordinates struct
type Coordinates struct {
	latitude float64
	longitude float64
}

// calculate distance between 2 point on a sphere
func haversine(point1 Coordinates, point2 Coordinates) float64 {
	R := 6371.0 // Earth's radius in km

	dlat := degToRadians(point2.latitude - point1.latitude)
	dlon := degToRadians(point2.longitude - point1.longitude)

	phi1 := degToRadians(point1.latitude)
	phi2 := degToRadians(point2.latitude)

	a := math.Sin(dlat / 2) * math.Sin(dlat / 2) +
		 math.Cos(phi1) * math.Cos(phi2) *
		 math.Sin(dlon / 2) * math.Sin(dlon / 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1 - a))
	return c * R
}

// convert degrees to radians
func degToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}