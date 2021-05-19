.PHONY: default release init dist webapp clean
export GOPATH:=$(shell pwd)
export GOBIN:=$(GOPATH)/tools

APP_NAME:=shopmanagement
APP_BUNDLE_NAME:="Shop Management"
APP_DIR:=Release
APP_VERSION:=1.0.0
APP_VERSION_SHORT:=1.0
BUNDLE_ID:=com.przon.shopmanagement
COPYRIGHT_YEARS:=2020-2021
COPYRIGHT_OWNER:="Md Delwar Hossain"
TARGET_EXE:=$(APP_DIR)/$(APP_BUNDLE_NAME).exe
TOOLS_DIR:=$(GOPATH)/tools
export PATH:=$(TOOLS_DIR):$(PATH)
export PKG_CONFIG_PATH:=$(PKG_CONFIG_PATH):$(GOPATH)

default: release
init:
	go get -d app
	go build -o $(TOOLS_DIR)/cef.exe tools/cef-installer
	go build -o $(TOOLS_DIR)/bindata.exe tools/go-bindata
	cef install
dist:
	cef dist \
	--dir $(APP_DIR) \
	--bundle $(APP_BUNDLE_NAME) \
	--executable $(APP_NAME) \
	--release $(APP_VERSION) \
	--short-release $(APP_VERSION_SHORT) \
	--year $(COPYRIGHT_YEARS) \
	--owner $(COPYRIGHT_OWNER) \
	--id $(BUNDLE_ID)
	cp -rf public $(APP_DIR)
	windres -o src/app/resource.syso res/main.rc
webapp:
	bindata -nomemcopy -pkg=resources -o=src/app/views/resources.go assets/... templates/...
	go build -o $(TARGET_EXE) app
release: dist webapp
clean:
	rm -rf tools pkg $(APP_DIR) cef cef.pc src/github.com src/golang.org src/gopkg.in