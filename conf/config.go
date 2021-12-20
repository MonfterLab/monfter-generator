package conf

import "fmt"

type AccessoryTypeName string
type PropertyTypeName string

const (
	AccessoryTypeNameCap        AccessoryTypeName = "CAP"
	AccessoryTypeNameEye        AccessoryTypeName = "EYE"
	AccessoryTypeNameNose       AccessoryTypeName = "NOSE"
	AccessoryTypeNameMouth      AccessoryTypeName = "Mouth"
	AccessoryTypeNameClothes    AccessoryTypeName = "Clothes"
	AccessoryTypeNameHorn       AccessoryTypeName = "Horn"
	AccessoryTypeNameFace       AccessoryTypeName = "FACE"
	AccessoryTypeNameBody       AccessoryTypeName = "BODY"
	AccessoryTypeNameBackground AccessoryTypeName = "BACKGROUND"
)

const (
	PropertyTypeNameCap        PropertyTypeName = "CAP"
	PropertyTypeNameEye        PropertyTypeName = "EYE"
	PropertyTypeNameNose       PropertyTypeName = "NOSE"
	PropertyTypeNameMouth      PropertyTypeName = "Mouth"
	PropertyTypeNameClothes    PropertyTypeName = "Clothes"
	PropertyTypeNameHorn       PropertyTypeName = "Horn"
	PropertyTypeNameFace       PropertyTypeName = "FACE"
	PropertyTypeNameBody       PropertyTypeName = "BODY"
	PropertyTypeNameBackground PropertyTypeName = "BACKGROUND"
)

type AccessoryType struct {
	Name  AccessoryTypeName
	Index int
}

type AccessoryTypeConfig struct {
	Name          AccessoryTypeName
	Index         int
	Necessary     bool
	AccessoryList []Accessory
}

type Accessory struct {
	Id                      int
	Name                    string  `json:"name"`
	Chance                  float64 `json:"chance"`
	PropertyList            []Property
	IncompatibleAccessories map[AccessoryTypeName][]int // 不兼容的
}

type Property struct {
	Type PropertyTypeName
	Name string
}

type AccessoryResult struct {
	AccessoryType AccessoryTypeName
	Index         int
	Accessory     Accessory
}

type ChanceType struct {
	Index int
	Val   float64
	Min   int
	Max   int
}

const GenerateTotal = 10

const ChanceRangeMax = 100000

var AccessoryConfigList = []AccessoryTypeConfig{
	{
		Index:         1,
		Name:          AccessoryTypeNameBackground,
		Necessary:     true,
		AccessoryList: AccessoryBackground,
	},
	{
		Index:         2,
		Name:          AccessoryTypeNameBody,
		Necessary:     true,
		AccessoryList: AccessoryBodyList,
	},
	{
		Index:         3,
		Name:          AccessoryTypeNameHorn,
		Necessary:     true,
		AccessoryList: AccessoryHornList,
	},
	{
		Index:         4,
		Name:          AccessoryTypeNameClothes,
		Necessary:     false,
		AccessoryList: AccessoryClothesList,
	},
	{
		Index:         5,
		Name:          AccessoryTypeNameMouth,
		Necessary:     true,
		AccessoryList: AccessoryMouthList,
	},
	{
		Index:         6,
		Name:          AccessoryTypeNameNose,
		Necessary:     true,
		AccessoryList: AccessoryNoseList,
	},
	{
		Index:         7,
		Name:          AccessoryTypeNameEye,
		Necessary:     true,
		AccessoryList: AccessoryEyeList,
	},
	{
		Index:         8,
		Name:          AccessoryTypeNameCap,
		Necessary:     false,
		AccessoryList: AccessoryCapList,
	},
}

func ShowConfig() {
	for _, accessoryConfig := range AccessoryConfigList {
		fmt.Println("trait:", accessoryConfig.Name)
		fmt.Println("ID", ",", "Value", ",", "Chance")
		for _, accessory := range accessoryConfig.AccessoryList {
			fmt.Println(accessory.Id, ",", accessory.Name, ",", accessory.Chance)
		}
		fmt.Println("")
		fmt.Println("")
	}
}
