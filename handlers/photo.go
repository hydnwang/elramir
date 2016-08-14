package handlers

import (
	"github.com/elramir/helper"
	"github.com/elramir/model"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func PhotosIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": "ok",
	})
}

func PostPhoto(c *gin.Context) {
	// Required fields
	var m string
	cid := c.PostForm("cid")
	if cid == "" {
		m = "Invalid user"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "unauthorized",
			"message": m,
		})
	} else {
		m = "Valid"
		f, h, err := c.Request.FormFile("upload")
		if err != nil {
			log.Fatal(err)
		}

		// File processing
		dir := helper.StringConcat("/", ".", "upload", cid)
		helper.FindOrCreatePath("folder", dir)
		filePath := helper.StringConcat("/", ".", "upload", cid, h.Filename)
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = io.Copy(file, f)
		if err != nil {
			log.Fatal(err)
		}

		resizePath := helper.Resizing(filePath)
		ext := filepath.Ext(filePath)
		extToLower := strings.Trim(strings.ToLower(ext), ".")
		if extToLower == "jpg" || extToLower == "jpeg" {
			model.ReadPhotoEXIF(filePath)
		}
		result := model.PostPhoto(h.Filename, dir, cid)

		// Respond when done
		c.JSON(http.StatusOK, gin.H{
			"status":   "uploaded",
			"message":  result,
			"header":   h,
			"filename": filePath,
			"resize":   resizePath,
		})
	}
}

func GetPhotos(c *gin.Context) {
	cid := c.Param("cid")
	result := model.GetPhotos(cid)
	// Respond when done
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

func GetPhoto(c *gin.Context) {
	cid := c.Param("cid")
	id, _ := strconv.Atoi(c.Param("id"))
	photo := model.GetPhoto(cid, id)
	file := helper.StringConcat("/", photo.Path, photo.Name)
	log.Println("File path:", file)
	c.File(file)
}

func GetThumbnail(c *gin.Context) {
	cid := c.Param("cid")
	id, _ := strconv.Atoi(c.Param("id"))
	photo := model.GetPhoto(cid, id)
	file := helper.StringConcat("/", photo.Path, "thumbnail", photo.Name)
	log.Println("File path:", file)
	c.File(file)
}

// func GetXML(c *gin.Context) {
// 	c.XML(http.StatusOK, gin.H{
// 		"message": "hey",
// 		"status":  http.StatusOK,
// 	})
// }
