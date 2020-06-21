package main

import (
	"log"
	"sort"
	"time"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud/openstack/imageservice/v2/images"
	"github.com/rackspace/gophercloud/pagination"
)

var (
	allImages          map[int64]string
	imageServiceClient *gophercloud.ServiceClient
)

// Init setups Openstack provider
func Init(authData AuthData) *gophercloud.ProviderClient {

	allImages = make(map[int64]string)

	opts := gophercloud.AuthOptions{
		IdentityEndpoint: authData.Endpoint,
		Username:         authData.Username,
		Password:         authData.Password,
		TenantName:       authData.TenantName,
		DomainName:       authData.DomainName,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		panic(err)
	}

	return provider
}

// GetImageList retrieves the images that match the pattern
func GetImageList(
	provider *gophercloud.ProviderClient,
	imageName string,
	region string) {

	cl, err := openstack.NewImageServiceV2(
		provider,
		gophercloud.EndpointOpts{Region: region},
	)
	imageServiceClient = cl

	if err != nil {
		panic(err)
	}

	opts := images.ListOpts{
		Name:    imageName,
		Status:  "active",
		SortKey: "created_at",
		SortDir: "desc",
		Limit:   1000,
	}

	pager := images.List(imageServiceClient, opts)
	pager.EachPage(extractImages)
}

func extractImages(page pagination.Page) (bool, error) {
	imageList, _ := images.ExtractImages(page)
	for _, image := range imageList {
		t, _ := time.Parse(
			time.RFC3339,
			image.CreatedDate)

		allImages[t.Unix()] = image.ID
	}
	return true, nil
}

// ProcessImages sorts the images by creating date
func ProcessImages(numOfImagesToKeep int, check bool) {

	if numOfImagesToKeep >= len(allImages) {
		return
	}

	var keys []int64
	for k := range allImages {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] > keys[j] })
	for i := 0; i < numOfImagesToKeep; i++ {
		delete(allImages, keys[i])
	}
	deleteImages(allImages, check)
}

func deleteImages(imagesToDelete map[int64]string, check bool) {
	for _, v := range imagesToDelete {
		if check {
			log.Printf("CHECK MODE: Image ID: [%s] will be deleted\n", v)
		} else {
			log.Printf("Image ID: [%s] is being deleted\n", v)
			res := images.Delete(imageServiceClient, v)
			if res.Err != nil {
				log.Printf("Error while deleting image: [%s]\n", v)
			}
		}
	}

}
