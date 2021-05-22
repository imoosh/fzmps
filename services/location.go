package services

import (
    "centnet-fzmps/common/log"
    "centnet-fzmps/conf"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

//https://lbs.qq.com/service/webService/webServiceGuide/webServiceGcoder

const (
    locationUrl = "https://apis.map.qq.com/ws/geocoder/v1/?location="
)

type Loc struct {
    Nation       string
    Province     string
    City         string
    District     string
    Street       string
    StreetNumber string
}

type LocationResp struct {
    Status    int    `json:"status"`
    Message   string `json:"message"`
    RequestId string `json:"request_id"`
    Result    struct {
        AddressComponent Loc `json:"address_component"`
    } `json:"result"`
}

func GetLocation(lat, lng float32) *Loc {
    key := conf.Conf.Services.Location.Key
    if len(key) == 0 {
        return nil
    }
    url := fmt.Sprintf("%s%v,%v&key=%s&get_poi=0", locationUrl, lat, lng, key)

    resp, err := http.Get(url)
    if err != nil {
        log.Error(err)
        return nil
    }
    defer resp.Body.Close()

    bs, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Error(err)
        return nil
    }

    var loc LocationResp
    err = json.Unmarshal(bs, &loc)
    if err != nil {
        log.Error(err)
        return nil
    }

    if loc.Status != 0 {
        return nil
    }

    log.Debug(loc.Result.AddressComponent)

    return &loc.Result.AddressComponent
}
