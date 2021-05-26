package models

type FzmpsProvince struct {
    Id           int
    ProvinceCode int
    ProvinceName string
    ShortName    string
}

func (FzmpsProvince) TableName() string {
    return "fzmps_province"
}

type FzmpsCity struct {
    Id           int
    CityCode     int
    ProvinceCode int
    CityName     string
    ProvinceId   int
    Orders       int
}

func (FzmpsCity) TableName() string {
    return "fzmps_city"
}

type FzmpsCounty struct {
    Id           int
    CountyCode   int
    CityCode     int
    ProvinceCode int
    CountyName   string
    Orders       int
    ProvinceId   int
    CityId       int
}

func (FzmpsCounty) TableName() string {
    return "fzmps_county"
}
