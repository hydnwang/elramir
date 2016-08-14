# Elramir

This web API server, which is built by Golang, is intended to be somewhere we can store and preview  our photos.

It's a demo project that constructed a Go web API stucture template for my further use.

## Environment

* Golang v1.6.2
* Sqlite3
* [Glide](https://github.com/Masterminds/glide)
* Packages
  * [gin](https://github.com/gin-gonic/gin) (Web API framework)
  * [gorp](https://github.com/go-gorp/gorp) (Database ORM framework)
  * [goexif](https://github.com/hydnwang/goexif) (Decode photo EXIF meta data, use my own fork version)
  * [Resize](https://github.com/nfnt/resize) (Resize photos to smaller ones, as preview thumbnails.)

## Folder Structure

```
elramir
|_ config
    |_ config.go
|_ db
    |_ fake_data.sql
    |_ default.db
|_ handler
    |_ application.go
    |_ photo.go
|_ helper
    |_ helper.go
|_ model
    |_ model.go
|_ server
    |_ init.go
    |_ routes.go
|_ upload
    |_ (store uploaded files if needed)
|_ vendor
    |_ (packages and dependencies installed by glide)
|_ main.go
|_ glide.yml
|_ glide.lock
|_ README.md

```

## Setup & Run

> Important: Please use `"go get"` rather than `"git clone"` to download this project, "go get" not only works the same as `git clone` but also will save files to `$GOPATH/src/`

### Download source code

``` sh
$ go get github.com/hydnwang/elramir && echo Success!
Success!
```

### Install `glide` and packages

> `glide` is a local Go package manager to help gethering and organizing all packages you need in one place. Be aware that I use `Homebrew` to install it for my Mac OS.

``` sh
$ brew install glide
$ ...(installing)
$ cd $GOPATH/src/github.com/hydnwang/elramir
$ glide up
$ ...(installing packages)
```

### Database and fake data

There's nothing to be worried about when it comes to the database, since our server is smart enough to bring up a whole new database while there's no existing one, the only thing we should care about is how we generate a bunch of fake data for testing and here's how:

first, check what's inside the file: `db/fake_data.sql`, unmark SQL commands you need and then execute it: 

```sh
cat ./db/fake_data.sql | sqlite3 ./db/default.db
```

### Go Run!!

```sh
$ cd $GOPATH/src/github.com/hydnwang/elramir

// Use debug mode and listen to port 3000
$ go run main.go -m debug -p 3000
```

Then, just wait a moment (it could be a while since Go is compiling our code) til the server up, and test API by either the browser or the API tools like Postman and DHC.

## Run Test

There are few ways to run testing: 

> **Important:** please put your own photo named `photo.jpg` in folder `tmp` before starting the test.

### 1. Run testing right in the package directory

``` sh
$ cd $GOPATH/src/YOUR_PACKAGE_PATH
$ go test -v -cover
```

### 2. Run package testing somewhere else

``` sh
# Package path: $GOPATH/src/YOUR_PACKAGE_PATH
$ go test YOUR_PACKAGE_PATH -v -cover
```

### 3. Run package testing somewhere else with directory path

``` sh
# Package path: $GOPATH/src/YOUR_PACKAGE_PATH
$ go test $GOPATH/src/YOUR_PACKAGE_PATH -v -cover
```

### 4. Try GoConvey (recommended)

``` sh
# Don't use glide, we need its bin exec file.
$ go get github.com/smartystreets/goconvey
$ cd $GOPATH/src/YOUR_PACKAGE_PATH
$ $GOPATH/bin/goconvey
```
GoConvey will automatically open default browser for you and start to run test on your behalf, just press enter, sit back and wait!

\** If you only want to run one or some of your test, just pass specific test case name (regex pattern) after option "run" being given: 

``` sh
$ go test PATH -run YOUR_TEST_NAME
# let's say the test function name is TestGetPhotoHandler()
$ go test PATH -run GetPhoto
```

This way, your test wiil find itself in the up one level directory, then there you go!

## Reference

...