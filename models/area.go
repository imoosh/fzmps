package models



type Province struct {
    Id           int
    ProvinceCode int
    ProvinceName string
    ShortName    string
}

func (Province) TableName() string {
    return "province"
}

type City struct {
    Id           int
    CityCode     int
    ProvinceCode int
    CityName     string
    ProvinceId   int
    Orders       int
}

func (City) TableName() string {
    return "city"
}

type County struct {
    Id           int
    CountyCode   int
    CityCode     int
    ProvinceCode int
    CountyName   string
    Orders       int
    ProvinceId   int
    CityId       int
}

func (County) TableName() string {
    return "county"
}
