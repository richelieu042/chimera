package gaodeKit

import (
	"context"

	"github.com/richelieu042/chimera/v3/src/component/web/http_client/reqKit"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
)

const (
	weatherUrl = "https://restapi.amap.com/v3/weather/weatherInfo"
)

// GetLive 获取"实况"天气.
/*
@param city 城市编码
*/
func (client *Client) GetLive(city string) (*Live, error) {
	wResp := &WeatherResponse{}

	err := reqKit.GetAndInto(context.TODO(), weatherUrl, map[string][]string{
		"key":        {client.key},
		"city":       {city},
		"extensions": {"base"},
	}, wResp)
	if err != nil {
		return nil, err
	}

	if err := wResp.IsSuccess(); err != nil {
		return nil, err
	}
	if len(wResp.Lives) == 0 {
		return nil, errorKit.Newf("len(wResp.Lives) == 0")
	}
	return wResp.Lives[0], nil
}

// GetTodayCast 获取今天的"预报"天气.
func (client *Client) GetTodayCast(city string) (*Cast, error) {
	forecast, err := client.GetForecast(city)
	if err != nil {
		return nil, err
	}

	if len(forecast.Casts) == 0 {
		return nil, errorKit.Newf("len(forecast.Casts) == 0")
	}

	return forecast.Casts[0], nil
}

// GetForecast 获取"预报"天气.
/*
@param city 城市编码
*/
func (client *Client) GetForecast(city string) (*Forecast, error) {
	wResp := &WeatherResponse{}

	err := reqKit.GetAndInto(context.TODO(), weatherUrl, map[string][]string{
		"key":        {client.key},
		"city":       {city},
		"extensions": {"all"},
	}, wResp)
	if err != nil {
		return nil, err
	}

	if err := wResp.IsSuccess(); err != nil {
		return nil, err
	}
	if len(wResp.Forecasts) == 0 {
		return nil, errorKit.Newf("len(wResp.Forecasts) == 0")
	}
	return wResp.Forecasts[0], nil
}
