package main

import (
	"encoding/base64"
	"log"
	"strings"

	"github.com/Azure/azure-sdk-for-go/storage"
)

/*
	Upload file to azure
*/
func UploadFileToAzure(file []byte, filename, azureName, azureKey, azureContainer string) (string, error) {
	client, err := storage.NewBasicClient(azureName, azureKey)
	if err != nil {
		log.Println("Connection Error!")
		return "", err
	}
	blobService := client.GetBlobService()
	_, err = blobService.CreateContainerIfNotExists(azureContainer, storage.ContainerAccessTypeContainer)
	if err != nil {
		return "", err
	}
	result, err := blobService.BlobExists(azureContainer, filename)
	if err != nil {
		return "", err
	}
	if result {
		// File duplicate
		if strings.Contains(filename, ".") {
			fileArray := strings.Split(filename, ".")
			fileArray[len(fileArray)-2] += "_1"
			filename = strings.Join(fileArray, ".")
		} else {
			filename += "_1"
		}
	}
	blockID := base64.StdEncoding.EncodeToString([]byte(filename))
	err = blobService.PutBlock(azureContainer, filename, blockID, file)
	if err != nil {
		return "", err
	}

	blobService.PutBlockList(azureContainer, filename, []storage.Block{{blockID, storage.BlockStatusUncommitted}})

	url := blobService.GetBlobURL(azureContainer, filename)
	return url, nil
}
