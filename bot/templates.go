// Code generated by go-bindata. DO NOT EDIT.
// sources:
// data/users_export_template.html
package bot

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

var _dataUsers_export_templateHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\x51\xab\x9b\x30\x14\x7e\xf7\x57\x1c\xbc\x5c\xd8\xe0\x5a\x6b\xef\xba\x75\xd6\xca\xe0\x3e\xed\x65\x8c\x3d\xee\x2d\x9a\xa3\x0d\xc4\x24\x24\x69\xab\x13\xff\xfb\x88\x96\xb6\x7a\x2b\x8c\xb2\x08\x42\x92\xef\xfb\xce\x77\x92\x9c\x93\xec\x6d\xc5\x81\x13\x51\xee\x7c\x65\x83\x4c\xfb\xa9\x97\xec\x91\xd0\xd4\x03\x00\x48\x2c\xb3\x1c\xd3\x37\x79\x64\xf4\x37\x6a\x09\x01\xbc\x49\x61\x35\xcb\x0e\x8c\x4a\x8d\x26\x09\x07\xc4\x80\x36\xb6\xe1\x08\xb6\x51\xb8\xf3\x2d\xd6\x36\xcc\x8d\xf1\x53\xaf\xdf\x74\x23\x93\xb4\x81\xf6\x32\x75\xa3\x90\xc2\x06\x05\xa9\x18\x6f\x62\x08\x88\x52\x1c\x03\xd3\x18\x8b\xd5\x0b\xe4\x84\xb3\x4c\xb3\x17\x30\x44\x98\xc0\xa0\x66\xc5\xf6\x42\xee\xae\xb2\x8b\x3c\xd0\xf2\x34\x11\xa6\xcc\x28\x4e\x9a\x18\x0a\x8e\xf5\x76\x1c\x93\x63\x1d\x14\x5c\x9e\x62\x70\xbc\x93\x26\x6a\x4e\x37\x97\xfc\x50\x89\xa9\x67\x8e\x75\x0c\x11\x44\xb0\x5e\x3e\xdf\x65\x7e\xab\x90\x32\x02\x52\xf0\x06\x4c\xae\x11\x05\x10\x41\xe1\x43\xc5\x44\x70\x62\xd4\xee\x63\xf8\xf2\x79\xa3\xea\x8f\x13\xe5\xb9\x88\xe3\xa8\xab\xf5\xf3\x38\xa1\xee\x41\x0f\xd1\x6a\xb3\x7c\xd8\xc4\xf2\x7f\x99\x58\x2f\x1f\x36\x11\x7d\x5a\xac\x36\xff\xe2\xa3\x97\x13\x16\x85\x9d\xe8\x55\x44\x97\x4c\xc4\x10\x4d\xd3\x71\xef\x37\x20\x9c\x95\x22\x86\x1c\x85\x45\x3d\xff\x46\x06\x61\x56\x95\x13\xf1\x4c\x6a\x8a\x3a\xd0\x84\xb2\x83\x71\x31\x5c\xa6\xdb\x3b\x90\x18\x22\x55\x83\x91\x9c\x51\x78\xc2\xc2\x7d\x63\xd8\xf9\xac\x36\x33\xef\xed\xc6\x85\x9a\x78\xc8\x25\x97\x3a\x86\xa7\xd7\xd7\xd7\xed\xfb\xca\x33\xec\x0f\xc6\x10\x69\xac\xee\x6c\x9e\x90\x95\x7b\x1b\x43\x26\x39\xbd\x0d\xeb\xfe\x49\xd8\x17\x7b\xea\x25\xe1\xd0\x2e\x12\x57\xdc\xa9\x97\x50\x76\x84\x9c\x13\x63\x76\x7e\x5f\x96\xfe\xd0\x1b\xda\x56\x13\x51\x22\x2c\xba\xeb\xfd\x8c\xb1\xc3\x75\x9f\xe1\x73\x90\x3e\xc7\x09\xa6\xc7\xb9\xc3\x37\x3a\xdf\xf9\x6d\xbb\xf8\xa9\x65\xc1\x38\x2e\xbe\x57\xa4\xc4\xe8\xeb\xaa\xeb\xfc\x77\x84\xe1\x8e\x5d\xef\x1a\x51\x7e\x21\xe1\x3f\x48\x85\xb3\x14\xc2\xed\x1c\xe1\xa6\xd1\x5d\x6c\xa9\xf4\x2e\x36\x09\xd5\x24\xcf\x90\xb2\xe3\x75\xe9\x66\xda\xb6\x28\x68\xd7\x79\xe7\xa5\x24\x3c\x9f\x73\xe8\x3a\x77\xfa\x37\x00\x00\xff\xff\x1c\x29\x6d\xab\xc0\x05\x00\x00")

func dataUsers_export_templateHtmlBytes() ([]byte, error) {
	return bindataRead(
		_dataUsers_export_templateHtml,
		"data/users_export_template.html",
	)
}

func dataUsers_export_templateHtml() (*asset, error) {
	bytes, err := dataUsers_export_templateHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/users_export_template.html", size: 1472, mode: os.FileMode(420), modTime: time.Unix(1584742285, 0)}
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
	"data/users_export_template.html": dataUsers_export_templateHtml,
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
	"data": &bintree{nil, map[string]*bintree{
		"users_export_template.html": &bintree{dataUsers_export_templateHtml, map[string]*bintree{}},
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