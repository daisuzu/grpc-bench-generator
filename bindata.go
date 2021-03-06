// Code generated by go-bindata.
// sources:
// templates/main.go.tmpl
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesMainGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x57\xdd\x6f\xdb\xc8\x11\x7f\x26\xff\x8a\x29\xd1\x04\xa4\x42\xd3\xb9\xeb\xc7\x83\x2f\x7e\xb8\xc8\x97\x22\x45\xce\x35\xe2\x04\x79\x48\x03\x63\xb5\x1a\x52\x5b\x93\xbb\xec\xee\xd2\xb2\x41\xf0\x7f\x2f\x66\x96\xa4\x68\xcb\xf5\x05\x88\x24\x72\xbe\xbf\x7e\x33\x6e\x85\xbc\x15\x15\x42\x23\x94\x8e\x63\xd5\xb4\xc6\x7a\x48\xe3\x28\xd9\x3c\x78\x74\x49\x1c\x25\xa8\xa5\xd9\x2a\x5d\x9d\xfe\xc7\x19\x4d\x2f\xca\x5a\x54\xfc\xdd\x78\xfa\x52\x86\x3e\x1b\xe1\x77\xf4\xad\xd1\x8f\x5f\xa7\x3b\xef\xdb\x24\x8e\x6e\x60\x7e\x3a\x6d\x5b\x6b\x4a\xa2\xbb\x07\x2d\xe9\xdb\xab\x06\x93\x38\x8e\x92\xca\xd4\x42\x57\x85\xb1\xd5\xe9\xfd\x29\xf1\x4b\xa3\x3d\xde\xb3\xb2\xca\x98\xaa\xc6\x62\xc1\x52\xd9\x56\xfe\x7f\xca\xe9\x06\xb5\xdc\x35\xc2\xde\x9e\x3a\x2f\xbc\x7b\x81\x53\x9a\x2d\xbe\x44\xa7\x8f\xda\x54\xe4\x62\xbb\x81\xa4\xef\x8b\xab\xdb\xea\x52\x34\x38\x0c\x49\x9c\xc5\xb1\x34\xda\x79\xe8\x9c\xa8\xf0\x1a\xed\x9d\x92\x08\xe7\xd0\xf7\xc5\xd7\xc5\x9b\x61\x58\xb2\x7d\xbe\x5a\x2f\x58\x3e\x5f\xad\x87\x21\x8e\xef\x84\xa5\xa4\x3b\xb4\x77\x68\x61\xf1\xef\x1c\x28\xdb\xc5\xb5\xb7\x4a\x57\x69\x12\x18\x92\x1c\x12\xfa\xff\x65\x87\x30\x8a\x88\xed\xd6\xa2\x73\x49\x16\x47\x8d\xb8\x5f\x1b\x2d\x3b\x6b\x51\xfb\xcf\x57\x6b\x37\x29\xf9\xa8\x7d\x9a\x34\xe2\xfe\x46\xce\xe4\x1b\xdb\x4a\x97\xe4\xf0\xd3\xa8\xad\x11\xf7\xa0\xbb\x66\x83\x16\x4c\x09\x07\x3e\x20\x3d\xa4\x7c\xdb\x59\xe1\x95\xd1\x47\x1e\xb2\xf2\x89\x9a\xe4\x40\xfd\x50\xfc\x2e\xee\x3f\x6a\xff\x97\x9f\x47\xed\xb3\xb0\xd2\xe0\x50\x1a\xbd\x75\xe0\x0d\xd8\x4e\x83\xdf\x21\xcc\x55\x03\x59\x2b\xd4\x9e\xec\x79\x2b\x24\xc2\x33\x19\x79\x6f\x4c\x9d\x26\x4c\x4e\x72\xf0\xb6\xc3\x1c\x92\x6f\x3b\xf4\x3b\xb4\x40\xaf\x95\xae\x40\x39\x30\x9a\xd4\xa0\x16\x9b\x1a\xaf\xaf\x3f\x3d\xaf\x26\x90\x6f\x9c\xab\x9f\xd1\x65\xa0\x73\x08\xbe\xe6\xf8\xdd\x58\xe4\x17\x2b\xa4\xd8\x27\x6a\x96\x0b\x2c\x45\x57\xfb\xb9\x11\x92\xfc\x51\xaf\x64\x71\x64\x5b\x09\xf0\x5c\x80\x93\x42\xea\xf4\x47\xca\xb8\x65\x26\x45\x9f\xaf\xd6\x54\x15\xe1\xc5\x8b\x4a\x88\x81\xb5\x90\x60\x72\xdd\xa2\x54\xa5\x42\xc7\x69\x67\x61\x6f\x60\x43\xcd\xa4\x3d\x15\xe7\x9f\xd7\xff\xba\x4c\x32\xea\x6f\xff\xd0\x22\x78\x74\x7e\xcd\x25\x01\xe7\x6d\x27\x3d\xf4\x71\xd4\x69\x61\x1f\xd6\xa2\xae\xd1\x42\xd9\x69\x99\x66\x80\xd6\x1a\x1b\x0f\xa3\x14\x8d\x5e\xe7\xd6\x34\x60\x0b\x31\x69\x3a\xed\xa1\x11\xed\x77\x1e\xbd\x82\xe8\x3f\x94\xf6\x71\x44\x90\x50\xfc\xde\x79\xbc\x27\x15\xa4\x12\x52\x27\x61\xb5\xd0\x93\xc1\xaf\xdb\x6d\x8a\xd6\x06\x53\x19\x29\x74\xb2\xf8\x64\xe4\x6d\x4a\x69\xc0\x12\x2d\x38\x59\x7c\xd5\xf5\xf8\xca\xc9\x82\x2d\x7e\xa7\x41\x66\x63\x24\x9e\xfd\x78\xf3\xe6\x05\x23\x57\x56\x69\x9f\xee\x41\x99\xe2\x9b\x55\x1e\x83\xa1\xb2\xf1\xc5\x87\x96\x48\xb5\x4e\xf7\x94\x46\x96\x01\x16\xa2\xde\x28\x8d\x85\xdb\x1c\xee\xe0\xec\x1c\xac\xd0\x15\xc2\x64\x9d\xc4\x17\xf2\x25\x8b\xbf\x72\x67\xf0\x6a\xfb\x6f\x1a\x16\x92\xca\xe2\x68\x78\xc1\xa7\xb1\x92\x19\xa5\x92\x5a\xbb\x8f\x23\x02\x8d\x0d\x30\x4e\x17\xef\xbb\xb2\x44\xcb\xf1\x06\xef\x5f\x6f\xa8\xb9\xd0\x77\x56\xc3\x66\xea\x83\x6c\x36\xd0\x0a\xeb\xf0\x42\x78\x91\x6e\xe0\xfb\x0f\xd2\x41\x8e\x2b\xed\xd1\x96\x42\x62\x3f\x70\xc8\xaa\xa4\x44\x53\x3c\x84\xfd\xc5\x57\xdd\x08\xeb\x76\xa2\x4e\x37\xe4\xf0\x2f\x4c\xfc\xd3\x39\x68\x55\x73\x84\x23\x58\x16\x1f\x84\x17\x75\x99\x26\xa5\x50\x35\x6e\x79\x82\x26\xc9\xb9\xdb\x1c\xea\xed\x19\xbc\xba\x4b\x72\xd2\xf2\x28\x78\x59\x1b\x87\x9f\x8c\x69\xbf\x52\x8b\xa5\x2e\x64\xc2\x15\x94\x6f\x97\x13\x2a\x69\x58\x85\x7a\x72\x53\xae\x8d\xd6\x39\x78\x09\xab\x43\xa3\xe6\x70\x94\xc1\x3e\x54\x48\x51\x38\x6f\x7f\x01\x05\xef\xe0\xa7\xb7\xf4\xe3\xcd\x1b\xf6\xde\xcb\x62\xd1\xd3\x29\xbb\x14\xc9\x1d\xb1\x37\xe2\x16\x53\xb9\x13\x84\x5c\x3e\x87\xd5\x11\xc4\xae\xfe\x9a\x85\x7a\xa4\x71\x14\x35\x1d\x2c\x7a\x39\x8a\xf6\x55\x78\xfe\x26\x94\xff\x87\x35\x5d\x1b\x47\x59\x1c\xed\xab\x82\x9a\xf9\x58\x57\x16\x1f\x3b\x7a\xcc\x75\xf0\xbb\x32\xd3\xfc\xd1\x13\x8b\x86\xee\x93\xbb\xf0\x26\x72\x5e\x58\x4f\xda\x68\xdb\x16\x97\x66\x9f\x66\xf1\x13\xd0\x01\xdb\xca\xdf\x42\xa9\x8f\xf3\x10\x45\x11\xd6\xa2\x75\x38\xeb\xb8\x56\x5a\x62\xca\x7a\x03\xbd\xe9\xe6\x29\x24\x83\x1c\x5a\x90\x99\xe9\x87\xa1\x24\x0e\xc9\x2c\xc1\x2a\xbf\x1a\xe8\x63\x5f\x15\x17\x46\x23\x33\x0d\x63\x09\x4e\x4f\xe1\xda\x9b\x96\xb1\x2a\x2c\x06\xd8\xef\x50\xb3\x23\x84\xf0\x5d\x5b\xc4\xd1\xd6\x68\x7c\x5c\xa9\x00\x39\xfd\x90\xc5\x4f\x32\xf4\xee\x84\x43\xf8\xb5\xf4\x68\x53\xfe\x79\x31\xee\xa5\x74\x35\x6d\xa8\x0c\x56\x63\xa0\xbc\xa6\xc8\x1d\xee\xca\x94\xec\x90\x5b\xe4\x9b\xb9\xe5\x7c\xd8\x0e\x43\xc1\xcc\x2d\xeb\x77\x58\x63\xc0\xba\x48\x0a\xc7\x65\x78\x77\x02\x6f\xcf\xa6\xe7\x77\x27\xa4\x84\x1e\x49\xc3\x39\x94\xa2\x76\x18\x73\x02\xa8\xe1\xd8\x8c\xdc\x85\x0e\xa1\x96\x21\x53\xd4\xf3\xc5\x9a\x49\x14\xcf\x38\x67\x57\x23\x18\xb9\x79\xbe\x9f\x23\xca\x05\x75\x9e\x31\x3a\xf7\x42\x3e\x78\x4f\x5c\x11\x1c\x4c\xaa\x8b\xdf\x78\x1d\x7e\x19\x97\xe8\x39\xac\x78\xcd\x3e\xcd\x63\xad\x5c\x3e\xc1\x83\x46\x5f\x7c\x52\xce\xa3\x4e\x13\x2f\x5b\x5a\x32\x67\x6f\x09\x13\x27\x08\x59\xa0\xc4\x11\x4c\x7c\x98\x61\xa2\x66\x15\x8f\x60\x81\xfb\xe2\x69\x4c\xc9\xb8\x8d\xe8\x9a\x54\x35\x39\x39\xde\x3f\x67\x90\xe4\xa4\x85\x9a\xcb\xa6\xd9\x32\x2f\x0b\x30\xa3\x53\xb4\xa0\x15\x8c\x29\x07\xa1\x55\x7d\x0c\x66\x2f\xb8\xc9\x37\xd7\x91\x97\xd4\x13\x01\x04\xa4\x3c\xc2\xa7\x40\x98\x37\x17\x43\xeb\xea\x70\x95\x70\xb3\xc8\x90\xce\x73\x60\xd9\x0b\x25\xea\x74\x15\xce\x3b\xea\x38\xc0\xda\xe1\x1f\x31\xe6\xe1\xd5\x37\xe5\x77\x1f\xb5\x43\xd9\x59\xe4\xe0\x87\xf8\x99\x42\x3c\x0d\xf0\xa0\xed\xd5\x7f\x33\x38\x80\xb7\xb4\x28\x3c\x82\x80\x43\x38\x21\xf6\xd9\xe8\x84\xe0\x71\xe4\x28\xbf\x01\xae\x2f\x71\xcf\x88\x9d\xfe\xfc\xb7\xbf\xf3\x26\x26\xd2\xeb\x05\x24\xf7\xbc\x1b\xcf\xc2\xcc\x1e\x5f\x04\xd9\x10\x47\x3e\x08\x1d\x70\xbd\x67\x23\x7b\xe5\xe5\x2e\x98\xa7\x63\xac\x8f\xa3\xbe\x0f\x98\xf7\x67\x77\x27\xe9\x3c\xcf\xf9\x17\x09\x17\xe3\xad\xe5\xe0\x64\xa0\xf9\xa2\x11\x4c\xfa\x7e\x62\x1c\x86\x84\x27\xb3\x66\xb8\x6d\x37\x45\x20\x15\x1f\x84\xf4\xc6\x3e\x0c\x43\x2a\x25\x15\x78\xb2\x49\xf7\x1a\x25\x6f\x36\x68\xdb\xc9\x20\x91\xce\xce\xd9\x70\xc1\x77\x37\x1b\x3c\x58\x1c\x39\x47\x8b\x51\xdf\xab\x92\x85\x0a\xde\x73\x23\x77\xa4\x74\x18\xa8\x7d\x1a\x9c\x21\x86\x8f\x7a\x18\x18\x24\x0f\x6b\x3b\x2c\xed\x74\x45\xfb\x34\xcb\xe1\xb5\xd2\xcc\xf0\x18\xbc\x09\x5d\x16\xa7\xd9\xb8\x0e\x6e\xe6\xb1\x95\xb5\x2a\x96\x8e\xa5\xe3\x1f\x5d\xc5\x7b\x21\x6f\x2b\x6b\x3a\xbd\x4d\xb3\x1c\x46\xdd\xd3\x39\x81\xd6\xce\x80\xdd\xf7\x27\x80\x7a\x3b\xcc\x0f\x53\x4c\xd7\xde\xa2\x68\xa6\xa0\x1e\xb5\x1a\x0d\xf0\xa5\xf1\xa0\x9a\xb6\xc6\x06\xb5\xc7\x2d\x43\xc5\x23\x65\x7d\x8f\x7a\x3b\x8a\x0f\xf1\xf2\x91\x3a\x40\x95\x4f\xd6\xd4\xb2\xaf\x9f\xde\x0f\x39\xd0\xc4\x78\x49\x37\xc1\x78\x68\xfc\x2f\x00\x00\xff\xff\x83\x3d\x30\x8c\xf6\x0e\x00\x00")

func templatesMainGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesMainGoTmpl,
		"templates/main.go.tmpl",
	)
}

func templatesMainGoTmpl() (*asset, error) {
	bytes, err := templatesMainGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/main.go.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/main.go.tmpl": templatesMainGoTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"main.go.tmpl": &bintree{templatesMainGoTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

