package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/astaxie/beego"
	. "monfter-generator/conf"
	"monfter-generator/models"
)

var (
	configFilePath string
	n              int
	startTokenId   int
	option         string
	targetKey      string
)

const (
	OptionBatch   = "batch"
	OptionReplace = "replace"
	OptionConfig  = "config"
)

func init() {
	env := os.Getenv("ENV")

	// init
	defConfigFilePath := ""
	if env == "DEV" {
		defConfigFilePath = "conf/app.conf.dev"
	} else {
		defConfigFilePath = "conf/app.conf"
	}

	flag.StringVar(&configFilePath, "c", defConfigFilePath, "config file")
	flag.StringVar(&option, "o", OptionBatch, "option: batch, replace")
	flag.IntVar(&n, "n", 0, "nft number")
	flag.IntVar(&startTokenId, "tokenId", 0, "start token_id")
	flag.StringVar(&targetKey, "key", "", "target key")
	flag.Parse()

	err := beego.LoadAppConfig("ini", configFilePath)
	if err != nil {
		panic(err)
	}

	// Init Model
	models.Init()
}

func main() {
	// show total chance
	ShowTotalChance()
	switch option {
	case OptionBatch:
		if n > 0 {
			BatchGenerateNFT(startTokenId, n)
		}
		break
	case OptionReplace:
		ReplaceNFT(strconv.Itoa(startTokenId), targetKey)
		break
	case OptionConfig:
		ShowConfig()
		break
	}
}
