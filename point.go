package geobass

import "github.com/shopspring/decimal"

// Point defines point on store
type Point struct {
	Latitude  float64
	Longitude float64
}

func (g Point) truncate(precision Accuracy) Point {
	fLon := decimal.NewFromFloat(g.Longitude)
	fLat := decimal.NewFromFloat(g.Latitude)
	lat, _ := fLat.Truncate(int32(precision)).Float64()
	long, _ := fLon.Truncate(int32(precision)).Float64()

	return Point{
		Latitude:  lat,
		Longitude: long,
	}
}
