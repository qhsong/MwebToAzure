package main

import (
	"fmt"
	"log"

	"strings"
	"time"

	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func UploadImgToAzureHandler(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	filename := header.Filename
	azureKey := c.PostForm("azureKey")
	azureName := c.PostForm("azureName")
	azureContainer := c.PostForm("azureContainer")
	appendDate := c.PostForm("appendDate")
	fmt.Println(azureContainer, azureKey, azureName)

	if strings.ToLower(appendDate) == "true" {
		filename = time.Now().Format("2006/01/") + filename
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	azureURL, err := UploadFileToAzure(content, filename, azureName, azureKey, azureContainer)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{
		"url":      filename,
		"azureURL": azureURL,
	})
}

func main() {
	r := gin.Default()
	r.POST("/api/upload/azure", UploadImgToAzureHandler)
	r.Run()
}
