package parcel_locker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"task_manager/internal/model"
)

const (
	parcelLockersDistanceSearchUrlTpl = "%s/parcel-locker-distance-search?longitude=%f&latitude=%f&radius=%f"
)

type ParcelLockerClient struct {
	LocationServiceEndpoint string
}

type ParcelLockersDistanceSearchResponse struct {
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Distance  float64 `json:"distance"`
}

type ParcelLockersNear struct {
	ParcelLockers []ParcelLockersDistanceSearchResponse
}

func (ls *ParcelLockerClient) FindParcelLockersNear(shipping *model.CustomerShipping, radius float64) (ParcelLockersNear, error) {
	parcels := ParcelLockersNear{}

	if shipping.Address != nil {
		endpoint := fmt.Sprintf(parcelLockersDistanceSearchUrlTpl, ls.LocationServiceEndpoint, shipping.Address.Longitude, shipping.Address.Latitude, radius)
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
