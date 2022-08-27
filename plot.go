package main

import (
  "fmt"
	"log"
)


const (
  ATTR_AREA							= "area"
	ATTR_OWNER						= "owner"
	ATTR_FARMABLE					= "farmable"
	ATTR_WATER						= "water"
	ATTR_ELEVATION				= "elevation"
	ATTR_RESOURCES				= "resources"
	ATTR_INHABITANTS			= "inhabitants"
)


type Plot struct {
  ID									string			`json:"id"`
	Area								int					`json:"area"`
	Owner								string			`json:"owner"`
	Farmable						int					`json:"farmable"`
	Water								int					`json:"water"`
	Elevation						int					`json:"elevation"`
	Resources						map[string] Resource			`json:"resources"`
	Inhabitants					map[string] User					`json:"inhabitants"`
}


func createPlot(kingdom string, municipal string, plot string) {

	key := fmt.Sprintf("%s:%s:%s:%s:%s", KEY_KINGDOM, kingdom, KEY_MUNICIPAL,
	  municipal, KEY_PLOT, plot)

	err := rds.HSet(ctx, key,
		ATTR_AREA, initRating(DEFAULT_AREA_MAX),
		ATTR_KINGDOM, kingdom,
	)

	if err != nil {
		log.Println(err)
	}
	
} // createPlot
