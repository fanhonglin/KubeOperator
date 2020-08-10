// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package i18n generated by go-bindata.// sources:
// locales/en-US/home.yml
// locales/zh-CN/home.yml
package i18n

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _localesEnUsHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x54\x41\x8e\xe3\x36\x10\xbc\xf3\x15\x6d\xf9\x1a\xe4\x01\xbe\x31\x12\x67\xac\xac\x2d\x09\x94\xb4\xc9\xe4\x22\xd0\x64\x7b\xcc\x8c\x4c\x0a\x24\x85\xc9\xe4\x96\x7f\xe5\x4f\xf9\x42\x40\x59\xf6\x7a\x33\x1e\xac\x83\xf5\xc5\x10\xc0\xaa\xee\xaa\xee\xae\xa5\xb4\xc7\xa3\x35\x84\x2c\x65\x3f\xfa\x80\x8e\xa4\x9b\xb6\x6e\x18\xef\x32\xb6\x61\x0d\xeb\x1e\x68\xbe\x61\xd9\x0a\x12\x29\x0c\x18\x1b\x40\x61\x8f\x01\x61\x7e\x0e\xda\x80\x1c\x9d\x43\x13\xc0\x07\x11\x30\x21\x4b\xe9\x50\xa1\x09\x5a\xf4\xe4\x2b\x92\x8e\xb3\xba\x6c\x79\xca\x56\x90\x3c\x08\xdd\xa3\x82\x60\x67\xbe\x05\x34\x07\x74\x08\xda\x83\x30\x20\xbc\xb7\x52\x8b\x80\x0a\x0e\xd6\x07\x18\x8d\x42\x07\xe1\xa0\x3d\xbc\xe0\x5b\xf2\x01\x6d\xf7\x5b\x59\xfc\x2f\xee\x3f\xad\xc1\x1b\xdc\x0f\xb4\xdd\x34\x5d\xca\x59\xc6\x8a\x26\xa7\x9b\x2e\xa5\x45\x57\x94\xcd\x6c\xc9\x0a\x92\x0c\xf7\x62\xec\x03\x7c\x51\x0a\x52\x98\xe8\xce\x0e\xe7\xa2\x2a\x21\x64\x19\x9b\x27\xeb\xb2\x6e\x3a\xba\xe1\x8c\x66\x4f\x1d\xfb\x35\xaf\x9b\x7a\x05\xc9\x3a\xea\x12\xbd\x43\xa1\xde\x00\xff\xd0\x3e\xf8\x8b\xae\x09\x71\x31\xfe\xa6\x9c\x93\x2f\x51\xd1\x4c\x71\x25\xeb\x55\x87\x03\x84\xc3\x65\x46\xb7\x78\xbb\x9f\x9e\xba\x8a\x97\x3f\xb3\xb4\xf9\xae\x12\x83\xb3\xbf\xa3\x0c\x51\xec\xe8\xd1\x91\x8a\xd6\xf5\x2f\x25\xcf\x26\xbf\xb6\xb4\x49\xd7\x2b\x48\x22\xd9\x20\xbc\x7f\xb5\x4e\x45\x42\x6d\xa4\x75\x6e\x82\x95\x3c\x7f\xcc\x0b\xba\x79\xf7\xde\x3a\xfd\xac\x8d\xe8\x3f\x02\xb6\x35\xe3\x5d\x5e\x4f\x38\x9a\x36\xf9\x67\x36\x03\x63\x1b\xf1\x6d\x9c\x06\x1a\xb1\xeb\x51\x2d\xa0\xea\x51\x78\x04\x69\x4d\x10\x32\x4c\x9d\x0b\x75\xd4\x46\xfb\xe0\x44\xb0\x6e\x31\x13\x46\xb6\x87\xb2\x2d\xa2\xef\x85\x05\x3f\xca\xc3\x44\xf8\xcf\xdf\x7f\x25\x84\x66\xdb\xbc\x78\xbf\x0c\xb1\xa8\x9a\x17\x62\x22\x3d\xb5\xf0\x6e\x21\x16\x5f\x37\xcd\xd9\x86\x36\x2c\xbb\x9a\x42\x1b\x61\x07\x11\x5b\xbf\xf6\x7a\xb6\xf8\x46\x0b\x6d\x95\xd1\xfb\x5a\x40\xa5\x4f\x1d\x10\xb2\x74\xf8\xac\xad\x39\xef\x04\x67\x8f\x79\x59\xdc\x7b\xa1\x70\x02\x7f\x6b\x2b\xe2\x61\xc5\x95\x88\xff\xe7\x42\xf1\x38\xef\x2e\x33\x5d\xe6\xb7\x56\xaf\x17\x26\x16\xb1\x03\x1a\x1f\x84\x7c\x21\x8f\xac\x39\xeb\x61\x9c\x97\xfc\x34\xc4\xb9\xe5\xbd\x1d\xcd\x74\x94\xb3\x9f\x5b\x3c\xee\xd0\x5d\x46\x42\xb3\xec\x7a\x04\x3b\x44\x03\x42\x29\xbc\x86\x5c\xb2\x61\x9e\xd9\xc7\xc1\x30\x03\x6e\xa5\xc2\x19\xbb\xa6\x75\x37\x67\xed\x0a\x92\x6a\x06\x5c\x09\x9d\xaf\x77\x01\xe9\x8d\x4d\x22\x64\x69\xac\x42\x52\x94\x19\xbb\x84\x0b\x6f\x8b\x22\x2f\x1e\xbb\x86\xd6\x9f\x56\x90\x50\xa5\x20\x3e\x02\xeb\xce\xa9\x3d\x7d\x9e\x4d\x75\xa3\x31\xda\x3c\xff\x30\x9c\x8e\xe3\x55\xe8\x00\x3a\x80\xb2\x06\x7f\x8c\xaa\x77\x42\xbe\x8c\x03\x95\xd2\x8e\x26\x90\x8a\x72\xba\xed\xd8\xb6\x6a\x9e\x56\x90\xe4\xc6\x8f\xfb\xbd\x96\x3a\x06\xff\x20\x9c\x38\x62\x40\xe7\x13\x12\xed\xa8\xdb\xaa\x2a\xf9\xb4\xd2\xc6\x8f\xc3\x60\x5d\xd4\x13\xde\x06\x4c\x48\xba\x66\xe9\xa7\x2f\xe1\xf6\x19\x9d\xde\x6b\x29\xc2\x34\xa2\x69\x17\x26\xc3\x7b\x61\xfe\x93\xf2\x77\x04\x56\x44\xdd\x1f\x58\xf1\xf7\x6f\x00\x00\x00\xff\xff\x9a\xc1\x4e\xbd\xff\x06\x00\x00")

func localesEnUsHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesEnUsHomeYml,
		"locales/en-US/home.yml",
	)
}

func localesEnUsHomeYml() (*asset, error) {
	bytes, err := localesEnUsHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/en-US/home.yml", size: 1791, mode: os.FileMode(420), modTime: time.Unix(1596627123, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesZhCnHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xd1\x52\xdb\x46\x14\x7d\xd7\x57\x78\xc4\x6b\xa7\x1f\xc0\xdb\x56\x5a\x40\x8d\x2c\x79\x56\x52\x5b\xfa\xa2\x71\x8c\xa6\xa5\x01\x89\x31\xe6\xa5\x4f\x35\x14\x82\x89\x5d\xd2\xc4\xb4\x4e\xe3\x29\x76\x6b\x26\x9e\xb4\x96\x08\xee\x60\x13\x61\xfa\x33\xda\x5d\xf9\x89\x5f\xe8\xac\x64\x6b\xe4\xb8\x4c\xfc\xba\x73\xcf\xb9\x67\xcf\xb9\xf7\x2e\x15\x9c\xed\x6d\xc7\xe6\xb8\xa5\xc2\xd6\xde\x6e\xc9\x2a\x72\x82\x6c\x68\x3a\x44\xa6\x08\x65\xa8\x43\x73\x05\x48\x32\x14\x97\x33\x3c\xf9\xb5\x45\xfa\x67\xf8\xb8\x35\x7e\xd5\xc1\xa3\x97\xb8\x52\xa3\x27\xd7\xe4\x87\x32\xfd\xed\xc7\xf1\xeb\x23\x7a\xd7\xe1\x19\x49\xd1\xda\xb0\xec\xd2\x66\x7e\x8b\x9b\xc1\x9b\x08\x6a\xaa\x81\x04\xb8\x9c\xe1\x27\x14\x9d\x77\xe1\x3f\x17\xf7\xb7\xe5\xd0\xbb\xa0\x6f\x1a\xe3\x17\x17\xc1\xf0\x19\x69\x56\xf0\x61\x3f\x2c\xd7\x83\xa1\x4f\x9a\xef\xf9\x07\x48\xcc\xaf\x55\x65\x51\x26\x7c\xea\xd1\x7a\x17\x57\x23\xb2\x15\x60\xc8\xba\x29\x20\x28\x42\x45\x97\x80\x6c\x0a\x40\x31\x15\x55\x9f\x7c\x76\x39\xc3\x8f\xfd\x46\xe8\x76\xf0\xd3\x1e\xa9\xb9\xc1\xb0\x16\x1e\x8c\xe2\x26\xf7\xb7\x65\xf6\xbf\x6f\x9d\xdd\x12\xb7\xa6\x6a\xba\x09\x64\x04\x81\xb8\x6e\xc2\xaf\x24\x4d\xd7\x96\x33\x7c\x2c\x19\x0f\xae\x70\xaf\x81\x9b\xdd\x44\x7b\x54\x9d\xb8\x38\x2f\x39\xc1\x51\xff\x34\x96\x3c\xb5\x73\x9e\xc0\xfc\x6c\xdd\xcc\x21\xf5\x73\x28\xe8\x8b\x72\xb5\x6f\xe8\x6b\x77\xa2\x7e\x6f\xd7\x2a\x72\x39\xa0\x69\x5f\xaa\x48\x8c\xfe\x9d\x05\xba\xb0\xc6\xa8\xbc\x23\xda\x2a\x8f\xeb\xaf\x42\xcf\xe3\x39\x15\x49\xab\x92\x02\xe4\xd9\x92\x9f\xce\x67\xab\x0c\x0d\x22\x53\xd2\xa2\x22\x20\xe8\xd2\x17\xcc\x3f\x5a\xef\x92\xe3\x01\x69\xbe\xc5\xcf\x99\xef\x91\xac\x41\x58\xae\xd3\xbe\x4f\xdd\x36\x7d\x7e\x84\x7f\x6e\x44\x6a\x22\x34\x83\xae\xa8\x86\x32\x9d\xaf\x5e\x27\xc6\x47\x15\x40\xcc\x4a\xca\x43\x09\x65\xf2\x1b\xdb\x9b\x76\x26\x2e\x8f\x83\x0a\xff\xf8\x2b\x95\x55\x5a\x1d\x82\x32\xd0\xa1\x98\xb2\x6e\x22\xf3\xaa\x9d\xcc\x49\x6c\xd4\x87\x5d\x8d\x9c\x08\xa2\xae\x60\xae\x5d\xf0\xaf\x4b\xea\x37\xb1\xb3\xdc\x52\xd1\xfa\x66\xd3\xb1\xa7\x91\x21\xb8\x2a\xa9\xca\x42\xb3\x8f\xab\xef\xf1\xf9\x79\x3a\xb2\xd4\xc4\x72\x4b\xdf\x3b\xb6\x35\x65\x65\x53\xbf\x18\xe7\x94\x61\x66\x12\x0e\xba\x74\x74\x15\xba\x6d\x7c\xfc\x82\x31\x3b\x3b\x96\xbd\x5b\xca\x17\x9e\x70\xab\x50\x9f\x2a\x86\x08\xa9\x88\x85\x51\xb9\x0b\x86\x35\x7c\x7c\x19\xcb\x63\xf5\x3b\x45\xe7\x3b\xab\x50\xca\x5a\xdb\x8f\xad\x62\x62\x2f\x10\xc5\xc4\xce\xb8\x1b\x19\xf8\xf8\xa4\x95\x42\x24\x7b\x37\xb1\xff\xa1\x48\xe3\x04\xe6\x96\x6e\x8a\x5a\x03\x9a\x39\xb9\x4d\x0c\x12\x15\xa7\xf7\xe5\xfe\xb6\xfc\x3f\x0b\x6b\x3b\x1b\x16\xa7\xa8\x22\x4c\x16\x16\x19\x8a\x22\x29\xab\xa6\x0e\xb4\x47\xcc\xbd\xc3\xeb\xc0\xff\x25\x3c\xd9\xa7\xfb\x37\xe4\xec\x72\xfc\xf4\x94\xbc\xac\x05\xa3\x26\xe9\xfd\x89\x9b\x5d\x52\x79\x13\xb6\xab\x9f\x64\x42\x6f\x40\x7b\x15\x7c\x77\x88\xdd\x83\xc0\xff\x3b\x7e\xc6\x6e\x95\x78\x67\x9f\xb2\x36\x8f\xf3\x85\x27\x7b\x3b\xa0\x50\x70\xf6\xec\x12\x97\x03\x08\x64\x4d\x98\xcd\xe9\xeb\xac\xc3\xe9\x3e\x39\xbb\x64\xda\xae\xfb\x3c\xc7\x3e\xae\x19\xb9\x9c\x8a\xf4\xe8\x68\xd4\x48\xdd\x23\x55\x76\x45\xe9\x3b\x1f\xff\xfe\x8c\xe7\x84\x35\x28\x3c\x4a\xdd\xdd\x56\x7b\xfc\xb6\x9a\xc4\x1b\x19\xbb\x95\xb7\x3f\x38\x8d\x1f\x39\x0c\xe9\xf0\xe7\xcf\x03\x9b\xe0\xff\x02\x00\x00\xff\xff\xd7\xad\x88\x2b\x09\x06\x00\x00")

func localesZhCnHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesZhCnHomeYml,
		"locales/zh-CN/home.yml",
	)
}

func localesZhCnHomeYml() (*asset, error) {
	bytes, err := localesZhCnHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/zh-CN/home.yml", size: 1545, mode: os.FileMode(420), modTime: time.Unix(1597032031, 0)}
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
	"locales/en-US/home.yml": localesEnUsHomeYml,
	"locales/zh-CN/home.yml": localesZhCnHomeYml,
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
	"locales": &bintree{nil, map[string]*bintree{
		"en-US": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesEnUsHomeYml, map[string]*bintree{}},
		}},
		"zh-CN": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesZhCnHomeYml, map[string]*bintree{}},
		}},
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
