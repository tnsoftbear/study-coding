package location

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"task_manager/internal/types"
)

const (
	// getLocationEndpoint     = "%s/location/%s"
	// locationEndpoint        = "%s/location"
	searchLocationsEndpoint = "%s/location/search?longitude=%f&latitude=%f&radius=%f"
)

type LocationService struct {
	LocationServiceEndpoint string
}

// type Coordinates struct {
// 	Longitude 	float64 	`json:"longitude"`
// 	Latitude 	float64 	`json:"latitude"`
// }

type SearchLocationResponse struct {
	Name string `json:"name"`
	//	Coord 		*Coordinates 	`json:"coord"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Distance  float64 `json:"distance"`
}

type ParcelLockersNear struct {
	Locations []SearchLocationResponse
}

func (ls *LocationService) FindParcelLockersNear(shipping *types.CustomerShipping, radius float64) (ParcelLockersNear, error) {
	parcels := ParcelLockersNear{}

	if shipping.Address != nil {
		endpoint := fmt.Sprintf(searchLocationsEndpoint, ls.LocationServiceEndpoint, shipping.Address.Longitude, shipping.Address.Latitude, radius)
		respose, err := http.Get(endpoint)
		if err != nil {
			return parcels, err
		}

		responseData, err := io.ReadAll(respose.Body)
		if err != nil {
			return parcels, err
		}

		if err := json.Unmarshal(responseData, &parcels); err != nil {
			return parcels, err
		}

		return parcels, nil
	}

	return parcels, nil
}
