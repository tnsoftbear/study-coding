package parcel_locker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"task_manager/internal/model"
	"task_manager/internal/config"
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

func NewParcelLockerClient() *ParcelLockerClient {
	return &ParcelLockerClient{
		LocationServiceEndpoint: config.GetStrEnv("PARCEL_LOCKER_SERVICE_ADDR", "http://localhost:8081"),
	}
}

func (ls *ParcelLockerClient) FindParcelLockersNear(shipping *model.CustomerShipping, distance float64) (ParcelLockersNear, error) {
	parcels := ParcelLockersNear{}

	if shipping.Address != nil {
		endpoint := fmt.Sprintf(parcelLockersDistanceSearchUrlTpl, ls.LocationServiceEndpoint, shipping.Address.Longitude, shipping.Address.Latitude, distance)
		respose, err := http.Get(endpoint)
		if err != nil {
			return parcels, err
		}

		responseData, err := io.ReadAll(respose.Body)
		if err != nil {
			return parcels, err
		}

		if err := json.Unmarshal(responseData, &parcels.ParcelLockers); err != nil {
			return parcels, err
		}

		return parcels, nil
	}

	return parcels, nil
}
