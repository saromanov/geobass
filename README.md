# geobass
[![Go Report Card](https://goreportcard.com/badge/github.com/saromanov/geobass)](https://goreportcard.com/report/github.com/saromanov/geobass)
[![Build Status](https://travis-ci.org/saromanov/geobass.svg?branch=master)](https://travis-ci.org/saromanov/geobass)
[![Coverage Status](https://coveralls.io/repos/github/saromanov/geobass/badge.svg?branch=master)](https://coveralls.io/github/saromanov/geobass?branch=master)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/e88a9cb5910743cc829930720195264a)](https://www.codacy.com/app/saromanov/geobass?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=saromanov/geobass&amp;utm_campaign=Badge_Grade)

Very simple db which handle geolocation points by accuracy

## Getting started

```
go get github.com/saromanov/geobass
```

## Examples

In this example, create new db with 11KM of accuracy. It means, that you can
get points which is nearly of the stored point, within 11 km
```go
package main

import (
	"fmt"

	"github.com/saromanov/geobass"
)

func main() {
	g := geobass.New(geobass.Range11KM)
	g.Set(geobass.Point{
		Latitude:  40.9887,
		Longitude: 28.7817,
	}, "data")
	fmt.Println(g.Get(geobass.Point{
		Latitude:  40.9858,
		Longitude: 28.7852,
	}))
}
```