package geobass

import "github.com/shopspring/decimal"

// Point defines point on store
type Point struct {
	Latitude  float64
	Longitude float64
}

func (g Point) truncate(precision int32) Point {
	fLon := decimal.NewFromFloat(g.Longitude)
	fLat := decimal.NewFromFloat(g.Latitude)
	lat, _ := fLat.Truncate(precision).Float64()
	long, _ := fLon.Truncate(precision).Float64()

	return Point{
		Latitude:  lat,
		Longitude: long,
	}
}
