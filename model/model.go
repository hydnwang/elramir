package model

import (
	// "bytes"
	"database/sql"
	"github.com/elramir/helper"
	"github.com/hydnwang/goexif/exif"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
	"log"
	"os"
	"time"
)

type Photo struct {
	Id      int64
	Name    string `db:"name"`
	Path    string `db:"file_path"`
	Created string `db:"created_at"`
}

func newPhoto(name, path string) Photo {
	layout := "2006-01-02 15:04:05"
	return Photo{
		Name:    name,
		Path:    path,
		Created: time.Now().Format(layout),
	}
}

var (
	dbmap         *gorp.DbMap
	dbCollect     = make(map[string]*gorp.DbMap)
	defaultDbDir  = "./db"
	defaultDbName = "default.db"
)

func InitDB() {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	dir := defaultDbDir
	helper.FindOrCreatePath("folder", dir)

	fname := defaultDbName
	path := helper.StringConcat("/", dir, fname)
	helper.FindOrCreatePath("file", path)

	db, err := sql.Open("sqlite3", path)
	helper.CheckErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(Photo{}, "photos").SetKeys(true, "Id")
	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	helper.CheckErr(err, "Create tables failed")

	dbCollect["default"] = dbmap
}

func GetPhoto(cid string, id int) Photo {
	var photo Photo
	// err := dbmap.SelectOne(&photo, "SELECT * FROM photo_"+cid+" WHERE id=?", id)
	_ = createDBIfNotExist(cid)
	_, ok := dbCollect[cid]
	if ok {
		err := dbCollect[cid].SelectOne(&photo, "SELECT * FROM photos WHERE id=?", id)
		// helper.CheckErr(err, "Search failed:")
		helper.Warning(err, "Search failed:")
		log.Println("File found:", photo)
		return photo
	}
	return Photo{}
}

func GetPhotos(cid string) []Photo {
	var photos []Photo

	_ = createDBIfNotExist(cid)
	_, ok := dbCollect[cid]
	if ok {
		_, err := dbCollect[cid].Select(&photos, "SELECT * FROM photos")
		// helper.CheckErr(err, "Get photos failed")
		helper.Warning(err, "Get photos failed")
		log.Println("Get Photos")
		return photos
	}
	return nil
}

func PostPhoto(name, path, cid string) bool {
	// create new photo
	p := newPhoto(name, path)

	// Create table if not exist
	_ = createDBIfNotExist(cid)

	err := dbCollect[cid].Insert(&p)
	// err := dbmap.Insert(&p1, &p2)

	helper.Warning(err, "Insert failed")
	log.Println("Insert completed!")
	log.Println(dbCollect)
	return true
}

func ReadPhotoEXIF(file string) {
	f, err := os.Open(file)
	helper.Warning(err, "read EXIF failed")
	defer f.Close()

	x, err := exif.Decode(f)
	helper.Warning(err, "EXIF decoding failed")

	// camModel, _ := x.Get(exif.Model)
	// normally, don't ignore errors!
	// log.Println(camModel.StringVal())
	// log.Println(x.String)
	// buf := bytes.NewBufferString("")
	// log.Println(buf.Read(x.Raw))
	for k, v := range x.Main {
		log.Println(k, "=", v)
	}

}

func createDBIfNotExist(cid string) bool {

	fname := "/test_" + cid + ".db"
	path := helper.StringConcat("/", defaultDbDir, fname)

	if _, ok := dbCollect[cid]; !ok {
		helper.FindOrCreatePath("file", path)
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbMap.AddTableWithName(Photo{}, "photos").SetKeys(true, "Id")
	err = dbMap.CreateTablesIfNotExists()
	helper.CheckErr(err, "Create tables failed")

	dbCollect[cid] = dbMap

	return true
}
