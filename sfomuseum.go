package tools

import (
	"github.com/sfomuseum/go-sfomuseum-airports/sfomuseum"
	sfomuseum_props "github.com/sfomuseum/go-sfomuseum-geojson/properties/sfomuseum"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/utils"
)

func SFOMuseumAirportFromFeature(f geojson.Feature) (*sfomuseum.Airport, error) {

	pt := sfomuseum_props.Placetype(f)
	
	if pt != "airport" {
		return nil, nil
	}
	
	wof_id := whosonfirst.Id(f)
	name := whosonfirst.Name(f)

	sfom_id := utils.Int64Property(f.Bytes(), []string{"properties.sfomuseum:airport_id"}, -1)
	
	concordances, err := whosonfirst.Concordances(f)
	
	if err != nil {
		return nil, err
	}
	
	a := &sfomuseum.Airport{
		WOFID:       wof_id,
		SFOMuseumID: int(sfom_id),
		Name:        name,
	}
	
	iata_code, ok := concordances["iata:code"]
	
	if ok && iata_code != "" {
		a.IATACode = iata_code
	}
	
	icao_code, ok := concordances["icao:code"]
	
	if ok && icao_code != "" {
		a.ICAOCode = icao_code
	}
	
	return a, nil
}
