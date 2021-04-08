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

func main() {
	fmt.Println("---------Hello from go---------")

	// ==== Test formula
	runTest()
}

// Coordinates coordinates struct
type Coordinates struct {
	latitude  float64
	longitude float64
}

// calculate distance between 2 point on a sphere
func haversine(point1 Coordinates, point2 Coordinates) float64 {
	R := 6371.0 // Earth's radius in km

	dlat := degToRadians(point2.latitude - point1.latitude)
	dlon := degToRadians(point2.longitude - point1.longitude)

	phi1 := degToRadians(point1.latitude)
	phi2 := degToRadians(point2.latitude)

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(phi1)*math.Cos(phi2)*
			math.Sin(dlon/2)*math.Sin(dlon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return c * R
}

// convert degrees to radians
func degToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// test formula
func runTest() {
	// point 1 GPS coordinates
	point1 := Coordinates{
		latitude:  51.510357,
		longitude: -0.116773,
	}

	// point 2 GPS coordinates
	point2 := Coordinates{
		latitude:  38.889931,
		longitude: -77.009003,
	}

	// point 3 GPS coordinates (Morningside)
	point3 := Coordinates{
		latitude:  31.0118578,
		longitude: -29.8173733,
	}

	// point 4 GPS coordinates (DUT Ritson Campus)
	point4 := Coordinates{
		latitude:  31.0056468,
		longitude: -29.8518299,
	}
	// distance in km
	distance1 := haversine(point1, point2)
	distance2 := haversine(point3, point4)

	fmt.Printf("\nDistance is %.3fkm", distance1)
	fmt.Printf("\n\nStraight line distance between Morningside and DUT: %.2fkm\n", distance2)
}
