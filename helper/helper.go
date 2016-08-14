package helper

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
	"strings"
)

func Resizing(f string) string {
	file, err := os.Open(f)
	if err != nil {
		log.Println("File open:", f, err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Println("Image decode:", img, err)
	}

	m := resize.Thumbnail(400, 400, img, resize.Lanczos3)

	fileNameSlice := strings.Split(f, "/")
	filename := fileNameSlice[len(fileNameSlice)-1]
	parentFolder := fileNameSlice[len(fileNameSlice)-2]
	path := StringConcat("/", ".", "upload", parentFolder, "thumbnail")
	FindOrCreatePath("folder", path)
	out, err := os.Create(StringConcat("/", path, filename))
	if err != nil {
		log.Println("File create:", err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)

	return path
}

func FindOrCreatePath(pathType, path string) bool {

	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		switch {
		case pathType == "folder":
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				panic(err)
				return false
			}
		case pathType == "file":
			_, err := os.Create(path)
			if err != nil {
				panic(err)
				return false
			}
		}
	}
	return true
}

func StringConcat(separator string, str ...string) string {
	if len(str) != 0 {
		// s = append([]string{".", "upload"}, s...)
		return strings.Join(str, separator)
	}
	return ""
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Warning(err error, msg string) {
	if err != nil {
		log.Println(msg, err)
	}
}
