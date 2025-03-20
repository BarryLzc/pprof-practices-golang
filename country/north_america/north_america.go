package north_america

import "github.com/go-pprof-practices/country"

type NorthAmerica interface {
	country.Country
	Lead()
	Rebel()
}
