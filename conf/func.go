package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"monfter-generator/models"
	"monfter-generator/utils"
)

var ExistingResultMap = make(map[string]interface{})

func SelectAccessory(accessoryTypeName AccessoryTypeName, accessoryList []Accessory) (Accessory, bool) {

	var pro Accessory

	chanceRangeList := make([]ChanceType, 0)

	var chanceSum float64
	for _, pro := range accessoryList {
		chanceSum += pro.Chance
	}
	cursor := 0
	for key, pro := range accessoryList {
		chanceRange := ChanceType{}
		chanceRange.Index = key
		if chanceSum > 1 {
			chanceRange.Val = pro.Chance / chanceSum
		} else {
			chanceRange.Val = pro.Chance
		}
		chanceRange.Min = cursor
		cursor = cursor + int(math.Ceil(chanceRange.Val*float64(ChanceRangeMax)))
		chanceRange.Max = cursor - 1
		chanceRangeList = append(chanceRangeList, chanceRange)
	}

	number := int(getRandomNumber(0, ChanceRangeMax))

	selectIndex := -1
	for _, changeRange := range chanceRangeList {
		if changeRange.Min <= number && number < changeRange.Max {
			selectIndex = changeRange.Index
			break
		}
	}

	if selectIndex >= 0 {
		return accessoryList[selectIndex], true
	}

	return pro, false
}

func getRandomNumber(min, max int32) int32 {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	number := rnd.Int31n(max) + min

	return number
}

func saveImage(img image.Image, filePath string) {
	fp, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("save dest image:%v err:%v\n", filePath, err)
		return
	}

	png.Encode(fp, img)

	defer fp.Close()
}

func GenerateNFT(tokenId interface{}, accessoryResultList []AccessoryResult) error {
	currPath, _ := os.Getwd()

	sort.Slice(accessoryResultList, func(i, j int) bool {
		return accessoryResultList[i].Index < accessoryResultList[j].Index
	})

	accessoryImgList := make([]image.Image, 0)

	flag := true
	for _, result := range accessoryResultList {
		fileName := fmt.Sprintf("%v/assets/%v/%v.png", currPath, strings.ToLower(string(result.AccessoryType)), result.Accessory.Id)
		itemFile, err := os.Open(fileName)
		if err != nil {
			flag = false

			fmt.Printf("open source file:%v err:%v\n", fileName, err)
			break
		}

		img, imgErr := png.Decode(itemFile)
		if imgErr != nil {
			flag = false
			fmt.Printf("decode source file:%v err:%v\n", fileName, err)
			break
		}

		accessoryImgList = append(accessoryImgList, img)
	}

	if !flag {
		fmt.Println(tokenId, ": Failed to load resource file.")
		return errors.New("Failed to load resource file.")
	}

	b := accessoryImgList[0].Bounds()
	canvas := image.NewNRGBA64(b)
	for _, itemImg := range accessoryImgList {
		draw.Draw(canvas, b, itemImg, image.Point{}, draw.Over)
	}

	// rotate
	//dstImg := RotateFlip(canvas)

	// save file
	dstFilePath := fmt.Sprintf("%v/dist/images/%v.png", currPath, tokenId)
	saveImage(canvas, dstFilePath)
	return nil
}

func isCompatible(accessoryResult AccessoryResult, results []AccessoryResult) (bool, error) {
	for _, r := range results {
		if r.AccessoryType == accessoryResult.AccessoryType {
			continue
		}
		if incompatibleAccessories, ok := r.Accessory.IncompatibleAccessories[accessoryResult.AccessoryType]; !ok {
			if utils.FindByIntArray(accessoryResult.Accessory.Id, incompatibleAccessories) >= 0 {
				return false, errors.New(
					fmt.Sprintf("%v:%v and %v:%v",
						r.AccessoryType, r.Accessory.Name,
						accessoryResult.AccessoryType, accessoryResult.Accessory.Name))
			}
		}
		if incompatibleAccessories, ok := accessoryResult.Accessory.IncompatibleAccessories[r.AccessoryType]; !ok {
			if utils.FindByIntArray(r.Accessory.Id, incompatibleAccessories) >= 0 {
				return false, errors.New(
					fmt.Sprintf("%v:%v and %v:%v",
						accessoryResult.AccessoryType, accessoryResult.Accessory.Name,
						r.AccessoryType, r.Accessory.Name))
			}
		}
	}

	return true, nil
}

func CheckIsCompatibleForResults(results []AccessoryResult) (bool, error) {
	// check compatible
	for _, r := range results {
		if ok, err := isCompatible(r, results); !ok {
			return false, err
		}
	}

	return true, nil
}

func isKeyExist(key string) bool {
	if _, err := models.MonfterKeyIns().FindByKey(key); err == nil {
		fmt.Printf("the key:%v exist\n", key)
		return true
	}

	return false
}

func GenerateNFTByResults(tokenId, key string, results []AccessoryResult) bool {
	// del current token traits
	deleteNum, deleteErr := models.MonfterTraitsIns().DeleteByTokenId(tokenId)
	if deleteErr != nil {
		fmt.Println("main", tokenId, "is no traits to delete", deleteErr)
	} else {
		fmt.Println("main", tokenId, "delete traits:", deleteNum)
	}

	// del the key from special token
	deleteKeyNum, deleteKeyErr := models.MonfterKeyIns().DeleteByTokenId(tokenId)
	if deleteKeyErr != nil {
		fmt.Println("main", tokenId, "is no key to delete", deleteKeyErr)
	} else {
		fmt.Println("main", tokenId, "delete key:", deleteKeyNum)
	}

	keyModel := models.MonfterKeyIns()
	keyModel.Key = key
	keyModel.TokenId = tokenId
	if _, err := keyModel.Add(); err != nil {
		fmt.Println("add key failed", key)
	}

	fmt.Println("generate...")
	err := GenerateNFT(tokenId, results)
	if err != nil {
		fmt.Println("generate failed", err)
		return false
	} else {
		fmt.Println("generate succeed.")
	}

	// save meta
	fmt.Println("save meta info...")
	metadata := make(map[string]interface{})
	traits := make([]map[string]interface{}, 0)
	for _, result := range results {
		if result.AccessoryType == AccessoryTypeNameNose {
			continue
		}

		// save to db
		traitInfo := models.MonfterTraits{}
		traitInfo, err = traitInfo.FindByTokenIdAndType(tokenId, string(result.AccessoryType))
		if err != nil {
			traitModel := models.MonfterTraits{}
			traitModel.TokenId = tokenId
			traitModel.DisplayType = "string"
			traitModel.TraitType = string(result.AccessoryType)
			traitModel.Value = result.Accessory.Name
			if _, err := traitModel.Add(); err != nil {
				fmt.Println(fmt.Sprintf("save %v trait is failed", result.AccessoryType), "result:", result)
			}
		} else {
			traitInfo.Value = result.Accessory.Name
			if _, err := traitInfo.Save(); err != nil {
				fmt.Println(fmt.Sprintf("save %v trait is failed", result.AccessoryType), "result:", result)
			}
		}

		trait := make(map[string]interface{})
		trait["display_type"] = "string"
		trait["trait_type"] = string(result.AccessoryType)
		trait["value"] = result.Accessory.Name

		traits = append(traits, trait)
	}

	metadata["description"] = "description"
	metadata["external_url"] = "external_url"
	metadata["image"] = "image"
	metadata["name"] = "name"
	metadata["attributes"] = traits

	if metadataJson, err := json.Marshal(metadata); err == nil {
		fileErr := ioutil.WriteFile(fmt.Sprintf("./dist/metadata/%v", tokenId), metadataJson, 0666)
		if fileErr != nil {
			fmt.Println("fileErr:", fileErr)
		}
	}

	return true
}

func BatchGenerateNFT(startTokenId, number int) {
	// sort
	sort.Slice(AccessoryConfigList, func(i, j int) bool {
		return AccessoryConfigList[i].Index < AccessoryConfigList[j].Index
	})

	var (
		incompatibleCount = 0
		repeatCount       = 0
		endTokenId        = startTokenId + number
	)

	for i := startTokenId; i < endTokenId; {
		if i >= endTokenId {
			break
		}

		var (
			results   = make([]AccessoryResult, 0)
			n         = i + 1
			status    = true
			numberArr = make([]string, 0)
		)
		fmt.Printf("generate:%v\n", n)

		for _, accessoryType := range AccessoryConfigList {
			pro, ok := SelectAccessory(accessoryType.Name, accessoryType.AccessoryList)
			if !ok {
				if accessoryType.Necessary {
					status = false
					break
				} else {
					numberArr = append(numberArr, "N")
					continue
				}
			}

			numberArr = append(numberArr, strconv.Itoa(pro.Id))
			results = append(results, AccessoryResult{AccessoryType: accessoryType.Name, Index: accessoryType.Index, Accessory: pro})
		}

		// check compatible
		if ok, _ := CheckIsCompatibleForResults(results); !ok {
			status = false
			incompatibleCount += 1
		}

		key := strings.Join(numberArr, "-")
		if isExist := isKeyExist(key); isExist {
			fmt.Printf("the key:%v exist retry:%v\n", key, n)
			status = false
		}

		if !status {
			// regenerate
			repeatCount++
			continue
		}

		if ok := GenerateNFTByResults(strconv.Itoa(i), key, results); !ok {
			repeatCount++
			continue
		}

		fmt.Println("------ FINISHED ------")
		i++
	}

	fmt.Println("stats:")
	fmt.Println("startTokenId:", startTokenId)
	fmt.Println("endTokenId:", endTokenId)
	fmt.Println("number:", number)
	fmt.Println("incompatibleCount:", incompatibleCount)
	fmt.Println("repeatCount:", repeatCount)
}

func ReplaceNFT(tokenId, key string) {
	accessoryIndexes := strings.Split(key, "-")
	results := make([]AccessoryResult, 0)
	for key, index := range accessoryIndexes {
		if strings.ToUpper(index) == "N" {
			continue
		}

		i, _ := strconv.Atoi(index)
		accessoryType := AccessoryConfigList[key]
		var accessory Accessory
		for _, row := range accessoryType.AccessoryList {
			if row.Id == i {
				accessory = row
				break
			}
		}

		if accessory.Id <= 0 {
			fmt.Printf("accessory:%v with index:%v not find\n", accessoryType.Name, i)
			return
		}

		results = append(results, AccessoryResult{
			AccessoryType: accessoryType.Name,
			Index:         accessoryType.Index,
			Accessory:     accessory,
		})
	}

	// check compatible
	if ok, err := CheckIsCompatibleForResults(results); !ok {
		fmt.Println("check compatible is failed, err:", err)
		return
	}

	if ok := GenerateNFTByResults(tokenId, key, results); ok {
		fmt.Println("generate succeed")
	} else {
		fmt.Println("generate failed")
	}
}

func RotateFlip(m image.Image) image.Image {
	newImg := image.NewRGBA(image.Rect(0, 0, m.Bounds().Dy(), m.Bounds().Dx()))
	for x := m.Bounds().Min.Y; x < m.Bounds().Max.Y; x++ {
		for y := m.Bounds().Max.X - 1; y >= m.Bounds().Min.X; y-- {
			newImg.Set(m.Bounds().Max.X-x, y, m.At(x, y))
		}
	}
	return newImg
}

func ShowTotalChance() {
	for _, accessoryConf := range AccessoryConfigList {
		var chanceSum float64
		for _, accessory := range accessoryConf.AccessoryList {
			chanceSum += accessory.Chance
		}
		fmt.Println(accessoryConf.Name, "chanceSum:", chanceSum)
	}
}
