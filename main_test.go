package main

import (
	"bytes"
	"flag"
	"github.com/appleboy/gofight"
	"github.com/buger/jsonparser"
	"github.com/hydnwang/elramir/config"
	"github.com/hydnwang/elramir/server"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

var (
	cwd_arg = flag.String("cwd", "", "set cwd")
	cid     = "0"
)

func init() {

	log.Println("Test started!")
	config.SetDefault()
	config.Mode = "test"

	clearTestData()

	// flagParse()
	// if err := os.Chdir("../"); err != nil {
	// 	log.Println("Chdir error:", err)
	// }
}

func TestRootHandler(t *testing.T) {

	message := make(map[int]interface{})
	r := gofight.New()

	r.GET("/").
		// turn on the debug mode.
		SetDebug(true).
		Run(server.RoutersEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			data := []byte(r.Body.String())
			message[0] = r.Code
			message[1], _ = jsonparser.GetString(data, "message")
			log.Println("Header: ", r.HeaderMap)
			log.Println("Body: ", r.Body.String())
			assert.Equal(t, "Welcome to Elramir", message[1])
			assert.Equal(t, http.StatusOK, r.Code)
		})

	Convey("When visit the root path", t, func() {
		Convey("Should see welcome message", func() {
			So(message[1], ShouldEqual, "Welcome to Elramir")
			So(message[0], ShouldEqual, 200)
			log.Println("welcome message:", message[1])
		})
	})

}

func TestPostPhotoHandler(t *testing.T) {

	path := "./tmp/photo.jpg"
	paramName := "upload"

	file, err := os.Open(path)
	checkErr(err)

	fileContents, err := ioutil.ReadAll(file)
	checkErr(err)

	fi, err := file.Stat()
	checkErr(err)
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	checkErr(err)
	part.Write(fileContents)

	_ = writer.WriteField("cid", cid)
	err = writer.Close()
	checkErr(err)

	contentType := writer.FormDataContentType()

	r := gofight.New()

	r.POST("/api/v1/photos").
		SetHeader(gofight.H{
			"Content-Type": contentType,
		}).
		SetBody(body.String()).
		Run(server.RoutersEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			log.Println("Header: ", r.HeaderMap)
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestGetPhotosHandler(t *testing.T) {

	r := gofight.New()

	r.GET("/api/v1/photos/"+cid).
		SetDebug(true).
		Run(server.RoutersEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			log.Println("Header: ", r.HeaderMap)
			log.Println("Body: ", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestGetPhotoHandler(t *testing.T) {

	id := "1"
	r := gofight.New()

	r.GET("/api/v1/photos/"+cid+"/"+id).
		SetDebug(true).
		Run(server.RoutersEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			log.Println("Header: ", r.HeaderMap)
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestGetThumbHandler(t *testing.T) {

	id := "1"
	r := gofight.New()

	r.GET("/api/v1/thumb/"+cid+"/"+id).
		SetDebug(true).
		Run(server.RoutersEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			log.Println("Header: ", r.HeaderMap)
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func flagParse() {
	flag.Parse()
	if *cwd_arg != "" {
		if err := os.Chdir(*cwd_arg); err != nil {
			log.Println("Chdir error:", err)
		}
	}
}

func clearTestData() {

	dbPath := "./db/test_" + cid + ".db"
	filePath := "./upload/" + cid

	var err error

	err = os.RemoveAll(dbPath)
	checkErr(err)

	err = os.RemoveAll(filePath)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Println("Error occurred:", err)
	}
}
