package country

import (
	"github.com/go-pprof-practices/country/asia/china"
	"github.com/go-pprof-practices/country/asia/singapore"
	"github.com/go-pprof-practices/country/europe/england"
	"github.com/go-pprof-practices/country/europe/germany"
	"github.com/go-pprof-practices/country/north_america/america"
)

var (
	// AllCountries 各个国家有各个国家的问题
	AllCountries = []Country{
		//亚洲
		&china.China{},
		&singapore.Singapore{},
		//欧洲
		&england.England{},
		&germany.Germany{},
		// 北美洲
		&america.America{},
	}
)

type Country interface {
	Name() string
	Exist()

	Trade()
	Develop()
	War()
	Legislate()
}
