package main

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	numOfImagesToKeep int
	imageName         string
	region            string
	check             bool
	authFile          string
)

// AuthData   contains authentication data
type AuthData struct {
	Endpoint   string `json:"endpoint"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	TenantName string `json:"tenantName"`
	DomainName string `json:"domainName"`
}

func main() {
	var (
		argCheck             = kingpin.Flag("check", "Enable check mode, don't actually delete anything").Default("false").Bool()
		argImageName         = kingpin.Flag("imageName", "Name of the image to save").Required().String()
		argNumOfImagesToKeep = kingpin.Flag("numImages", "Number of images with same name to keep").Default("2").Int()
		argRegion            = kingpin.Flag("region", "Region").String()
		argAuthFile          = kingpin.Flag("authFile", "Absolute path of a JSON file that contains the authentication information").Required().String()
	)

	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	numOfImagesToKeep = *argNumOfImagesToKeep
	imageName = *argImageName
	region = *argRegion
	check = *argCheck
	authFile = *argAuthFile

	if numOfImagesToKeep < 1 {
		log.Fatal("numImages value cannot be less than 1")
	}

	dir := filepath.Dir(authFile)
	base := filepath.Base(authFile)
	ext := filepath.Ext(authFile)

	viper.SetConfigName(strings.TrimSuffix(base, ext))
	viper.SetConfigType("json")
	viper.AddConfigPath(dir)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	authData := AuthData{
		Endpoint:   viper.GetString("endpoint"),
		Username:   viper.GetString("username"),
		Password:   viper.GetString("password"),
		TenantName: viper.GetString("tenantName"),
		DomainName: viper.GetString("domainName"),
	}

	GetImageList(
		Init(authData),
		imageName,
		region,
	)
	ProcessImages(numOfImagesToKeep, check)
}
