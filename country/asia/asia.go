package asia

import "github.com/go-pprof-practices/country"

type Asia interface {
	country.Country
	Expand()
	PreserveCulture()
}
