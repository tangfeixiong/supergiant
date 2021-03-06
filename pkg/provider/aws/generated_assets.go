// Code generated by go-bindata.
// sources:
// config/providers/aws/master_userdata.txt
// config/providers/aws/minion_userdata.txt
// DO NOT EDIT!

package aws

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

var _configProvidersAwsMaster_userdataTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x59\x6f\x73\xdb\x36\xd2\x7f\xcf\x4f\xb1\x65\x34\x8f\xec\xc6\x14\x25\xc7\x4d\x9e\x51\xca\xcc\x29\xb6\x92\x6a\x2a\x4b\x3e\x49\x6e\xda\x26\x1d\x0e\x44\x82\x12\xc6\x24\xc0\x02\xa0\x6c\xc5\x55\x3f\xfb\x0d\x00\x52\xa2\xfe\x50\x4e\xae\x9d\xeb\x2b\x89\xc4\xee\x6f\x7f\x58\xec\x02\x8b\xe5\xb3\x6f\xc0\x9d\x12\xea\x4e\x91\x98\x5b\xc9\x5d\x48\x38\x38\x29\xb8\x0b\xc4\xdd\x00\x05\x73\xec\xde\x65\x53\xcc\x29\x96\x58\x38\x84\x0a\x89\xe2\xd8\x0a\xc2\x27\x04\x38\x46\x21\xa3\xf1\x12\xc6\x9d\xfe\xc4\xbf\xee\x8c\x27\xdd\x91\x57\x7f\x7c\x84\x46\xe7\xc3\xf8\x92\xd1\x88\xcc\x1a\xd7\x48\x48\xcc\x6f\x38\x59\x20\x89\x7b\x37\xb0\x5a\xd5\x37\x7a\xbd\xc1\x78\xd2\x19\x5c\x76\xfd\x9b\x51\xf7\x5d\xef\x67\xa3\x3b\x40\x09\xde\x16\x1b\x0c\xaf\xba\xfe\x31\x59\x27\x21\x94\x30\x5a\x52\xb9\xec\xdf\x2a\x36\x7e\xef\xc6\x1f\x75\x06\xef\xbb\x5e\xbd\xd5\x6c\x9c\x5f\x5c\x34\x9a\x8d\xa6\xdb\x7a\x59\x92\xec\xf4\xfb\xc3\xcb\xce\xa4\xeb\x6b\x2b\x97\xbd\xab\xd1\xd8\xab\x4b\x9e\xe1\x92\xcc\xb8\x3b\xfa\xa9\x3b\xf2\xdf\xf6\x06\x9d\xd1\x2f\xfe\xa4\x33\xf2\x6f\x47\x7d\xaf\x3e\x97\x32\x15\x6d\xd7\x15\x2f\x1a\x28\x41\x9f\x19\x45\xf7\xa2\x11\xb0\xa4\xec\xab\x96\xd3\x72\x5e\x39\x88\x4b\x12\xa1\x40\x0a\x37\xc4\x0b\x1c\x97\x05\x04\xe6\x0b\xcc\x9d\x98\xd0\xec\xc1\x41\x49\xf8\xf2\xa2\x21\x11\x6f\xcc\x3e\xd7\x77\xfc\xfb\x37\x9b\x45\xb1\xdc\x37\xf4\xeb\x70\xd0\xdd\x5d\xc1\xce\x02\x91\x18\x4d\x49\x4c\xe4\xf2\x57\x46\x77\xd6\xe6\xc7\xdb\xb7\x5d\xff\x76\x5c\x2c\xfc\xad\xc0\x9c\xee\x2d\xa0\x16\xba\xe9\x8c\xc7\x1f\x86\xa3\x2b\x23\x78\x83\x84\xb8\x67\x3c\xdc\x16\x54\x8e\xee\x5d\x76\xfd\x83\xcb\xd7\xdc\x5f\xbc\xee\xa0\xf3\xb6\xbf\x11\xbf\x1e\x0e\x7a\x93\xe1\xa8\x37\x78\xef\xd5\x09\x8d\xe2\xec\x21\x9c\x56\x4b\xf7\x87\xef\xdf\x6b\xd1\x08\xc5\x02\xef\xcb\xe9\x80\xa8\x14\xca\x07\xfc\xab\xee\x78\xd2\x1b\x74\x26\xbd\xe1\xc0\xab\xe3\x18\x09\x49\x02\x81\x11\x0f\xe6\x65\xc4\x7e\x67\x3c\xe9\x5d\x8e\xbb\x9d\xd1\xe5\x0f\x05\xa6\x3f\xea\xde\xf4\x7b\x97\x9d\xb1\x57\x6f\x55\xb3\xbc\x1a\xec\xc7\xe3\x8e\xc8\x6d\x6f\x4f\xe2\x6a\x30\xae\x80\x57\x23\x79\x3c\xf7\x6e\xd6\x7e\x6d\x35\x77\x44\xae\x86\xd7\x9d\xde\xc0\xab\x07\x71\xa6\x12\xb8\x11\xb3\x00\xc5\xe5\xb4\xb9\xba\xee\x8d\xc7\xbd\xe1\xc0\xbf\x1c\x0e\x26\xa3\x61\xdf\xab\xab\x64\x14\x29\x0a\x70\x9f\x44\x38\x58\x06\x31\x3e\xeb\x93\x84\xc8\x11\xa2\x33\xcc\xcf\xc6\x38\xc8\x38\x91\xcb\x4b\x46\x25\x7e\x90\x57\x98\x2e\xcf\xc6\x98\x2f\x48\x80\x3b\x41\xc0\x32\x2a\xcf\x46\x58\xb0\x8c\x07\xf8\xdf\x19\x93\xa8\x64\xcc\x6c\x2d\xbb\xc9\xfc\x52\xc7\xc3\xf9\xc5\x4e\xa0\xf5\xbb\x13\x7f\x32\xfc\xb1\x3b\xf0\x6a\x27\x61\x08\x24\xf2\x54\xf4\xbb\x19\x47\x34\x64\x09\x4c\x85\xd7\x3a\xff\x7f\xd0\x06\xbd\x16\x9c\xbf\xd1\xa3\x34\x8b\x63\xf8\x03\xa6\x48\xe0\x97\x17\xf0\x07\x48\x0e\x4e\x08\xb6\xf7\xdc\xb5\xe1\x0f\x08\x43\xa5\xf6\xe2\xfc\x90\xd6\xe9\x6e\x98\x8f\x86\x3f\xff\xf2\x0f\x11\xb8\x1a\x5e\xfe\xd8\x1d\xf9\xe3\xc9\x70\xd4\x51\x6e\x42\x59\x24\xf6\xdd\xd8\xfd\x79\x32\xea\xf8\xe3\x8e\x0a\xac\xde\x4d\xbb\x88\x80\xb3\xab\xc1\xb8\xbd\xd9\x21\x76\x1e\x1b\x21\x8e\x50\x16\xcb\x8a\xd7\x0d\xb1\x08\x8e\x0c\x35\xb6\xc2\x68\x47\xd0\x49\xf4\x19\x51\xde\xf0\x6f\xaf\xfd\xeb\xde\xa0\x37\x1c\x98\xe0\xb5\x2c\x0b\xa5\xd2\x99\x61\x09\x59\x1a\x22\x89\xd7\x8f\xf9\x51\x04\x8e\xb3\xc4\x02\x82\x8c\xc7\x96\x15\xb2\x7b\x1a\x33\x14\x3a\x8c\x3b\xd3\x4c\xc8\x93\x53\x78\xb4\x00\xb4\x69\x70\x38\x64\x3c\xf6\xec\x5a\xcb\x2e\xbf\x8b\x48\x8c\x3d\xbb\xf6\x98\xf1\xf8\xd9\xb3\x6f\xdd\x95\x1a\xe4\x09\x38\x11\xd8\x35\x35\xa6\x9e\x33\x2a\x49\x0c\x1f\x3f\x82\x83\xc1\xae\x3d\xb6\x8c\x20\xfc\xf6\xdb\x6b\x08\x99\x05\x00\x80\x83\x39\x03\xfb\x2a\xb7\x4f\xe8\x4c\xe3\xc2\x49\xad\x75\x6a\x6b\x01\x45\x10\x1c\x87\xa4\x8b\x0b\x70\xfa\xac\x00\x07\xc7\x09\x18\xa5\x38\x90\x8e\x24\x09\x66\x99\x84\xf3\x26\x38\x0e\xc7\x92\x2f\xe1\x65\xf1\xcf\x09\x71\x8c\x96\xd0\x6a\x42\x4e\x1f\x20\x09\xbf\x13\x59\x52\x22\x19\x32\x8a\xad\x95\xf2\x58\xee\x1a\xbd\xd3\x6f\xb9\x40\xbd\xf0\x13\x16\x62\xe3\x05\x0b\x80\x44\x10\xa6\x77\x33\x70\x84\x1e\xcc\x8f\x53\xf8\xbf\x4d\x90\xbd\x06\x39\xc7\xb4\x34\x49\xcf\x83\x31\x8a\xe5\x58\xa2\xe0\x0e\x50\xac\x96\x6e\x59\xac\x06\x0e\xcf\x40\xdc\x91\x34\x55\x1e\x28\x56\x48\x48\x9c\x82\xe7\x19\xda\x1c\xcb\x8c\x2b\xb8\x88\x28\xfb\x6b\xc8\x11\x8e\x38\x16\x73\xa5\x97\xa2\xe0\x0e\xcd\x30\x84\x48\x22\x95\x15\xb9\xae\x59\x84\xed\x68\xd8\xf1\xbf\xe7\xed\x8c\x43\x84\x88\x26\xa5\xbd\xa8\xc0\x0b\x1e\x5a\x41\xc4\x18\xa7\xf0\x5d\xe1\x3c\x0b\xe0\x50\x81\xa4\x1d\x53\x54\x3e\x00\xdb\xc5\xd1\xd6\xa0\x05\x70\xd5\x7d\x3b\xf6\x4e\xb4\x89\x98\x4c\x3f\x27\xbf\xbf\xf0\x5f\x34\xce\x1b\x2f\x9e\x87\x91\x98\x39\xad\x3f\xa7\x29\x7b\xd5\xfc\x33\x14\xf2\x79\xcb\x37\x67\x7e\x88\xa7\x5a\x3c\x5d\xca\x39\xa3\xce\xe7\xe4\x77\xbf\xf5\xa2\xd1\x6a\x34\x8f\x4a\x6b\xbb\x01\x4b\x12\x46\xfd\xf3\x66\xeb\xa2\xd1\x6a\xb4\x5e\x3c\x0f\x45\xa1\xa4\x14\xe2\x38\x17\x3f\x35\x2b\xfd\xf1\xa3\x8a\xde\x75\x10\xac\x6c\xf0\x3c\xb0\x4d\x0e\x9a\x68\x5e\x2f\xb5\x9a\xc6\x73\xef\x24\x0f\x0a\x2d\xf1\x84\x19\x6d\x24\x22\xd6\xae\xae\x0e\xa8\x2f\xd1\xbd\x1d\xf5\xfd\xb7\x9d\x71\xd7\xb3\xd7\x65\x8e\x64\x1c\xcd\x70\x63\xc6\xd8\x2c\xc6\x28\x25\x7b\xc5\x0e\xc7\x31\x46\xc2\xac\x82\x8e\xe7\x88\x71\x50\x80\x84\xaa\xa9\x2a\x26\x1f\xff\xf5\xdb\xca\x5e\xc7\x89\xf2\x02\x7c\x93\xe7\x71\x88\xa7\x2a\x8b\x4b\xd3\x56\x81\xb0\xbd\x85\x28\xb9\x82\xd9\xca\xcd\x75\xb4\xac\x9e\x6b\x11\x37\x4f\xd9\x5d\xc7\x67\xcf\x84\x8a\x0a\x45\x03\x76\x06\x64\x46\x19\xc7\x10\xe2\x14\xd3\x10\xd3\x60\x09\x01\x4b\xd2\x18\x11\x2a\x05\x9c\xdc\x93\x38\x86\x88\x3c\x40\x8c\x24\xe6\xa7\xeb\x00\x36\x69\xeb\xa8\x6c\x73\x04\x4a\xb0\xb3\xc0\x5c\xa8\xdc\x75\x9c\x88\xf1\x00\x3b\x06\x4f\x80\x43\xd6\x73\x2d\x11\x7e\x06\x93\x39\x11\xa0\xd1\x8b\x44\x45\x74\x09\x2c\x52\xde\x80\x8c\x26\x58\x6e\x28\x11\x2c\x20\xe2\x2c\x01\x34\x65\x0b\xdc\xb0\x0e\x4f\xe8\x80\xd2\xa1\xd4\x5d\xef\xdc\x11\x38\xcb\xea\x0c\x2e\xc4\xbe\x36\x85\x9f\x41\x9f\xcd\x00\x81\xda\x4f\x85\x44\x49\x5a\x66\xfb\x8e\x50\x22\xe6\x38\x2c\xd0\x15\x9e\xda\xcd\x34\xa8\xde\x3c\x8d\xe8\x07\x44\xa4\xde\xc5\x19\x07\x13\xfe\x90\x86\x20\x19\x4c\x31\x20\x29\x55\xea\x87\xb6\x85\xa4\xc4\x49\x2a\xbd\xa6\x75\x3f\x57\xdb\xbd\x2a\xc5\xf2\x09\x69\x98\x8e\x19\x07\xbb\x76\x72\x52\xcb\x85\x9f\xb7\x4e\x4f\x6d\x85\x14\xcc\x71\x70\xa7\x0d\xe8\xbd\xf6\x61\x11\x4e\xd7\x69\xea\xe0\xcd\xcb\xed\xc4\x34\xf4\xde\xb1\x8c\x86\x1b\x11\xe3\x8f\x29\xc7\xe8\xae\x48\xc2\x82\xda\x8e\x65\x0b\x72\x8f\xb5\x2c\xe3\x2f\x83\x77\xad\x8a\x0b\x35\x5f\x33\x57\x27\x0d\xed\xd2\x5d\x31\xa1\xd2\x5d\x0f\x58\xc9\x5d\x24\xc0\x91\x80\x1f\xe4\x45\x89\xba\x01\xda\xb0\xde\x51\x03\x23\x0f\x94\x21\xb5\x30\x00\x4d\x68\xda\xf0\xe6\x0d\xb8\x58\x06\x6e\x24\x24\x9a\x5a\x89\x62\xb1\x6b\xae\xe0\x91\xc0\xab\x66\x73\x9f\x8e\xde\x8a\xb1\x0c\xc2\x2a\xbe\xae\xe0\x8b\xd2\x96\x71\x54\x4c\xef\x58\x6c\x81\x79\x8c\x96\xc7\x05\xc5\x3c\xe5\xec\x61\x69\x59\x31\x55\x67\xa7\x13\x55\xd1\x82\x0d\xc1\x2a\xd1\x6d\x82\xb0\x4b\xf8\x98\x5a\xc1\x03\xb6\x59\x1d\x55\x29\xcd\x11\xf6\x67\x6d\x91\x08\xbe\x01\x12\x82\x26\x7f\xa8\x0e\xc8\x04\xe6\x28\x0c\x95\x05\x57\x4c\x09\x75\x29\x8b\xd9\x8c\x50\x55\xb5\x6e\xa6\xad\x67\x1c\x11\x2b\x98\xb3\x7b\x0a\xce\xc8\xe0\x55\xad\x5d\x30\x9f\xf1\xf4\x49\x29\xcb\x2a\x2d\xca\x2e\x73\x37\x25\x71\x8c\xb8\x15\x20\x09\xdf\x7f\xdf\x1d\xbe\x83\x37\x55\x32\x6e\x5e\x93\x3a\x29\xe2\x28\x11\x0d\x11\x0b\x53\x33\xd1\x00\xfb\x29\xc7\x11\x79\x68\x43\xbd\x76\x62\x42\xba\xb6\xd3\x7f\x50\x55\xb9\xc0\xa1\x3e\x43\x84\x5b\x77\xeb\x75\x77\x66\x9f\xd6\x2d\xca\x42\xec\x1f\x81\x39\xd4\xcb\xa8\xc2\xca\x09\xfa\x01\x09\x79\x19\x63\xf7\x76\x5c\xa5\x8f\x62\x55\xf0\x49\xec\x6b\x52\x0a\x45\x94\x61\x0e\x74\x3e\xaa\x90\x84\xb9\xa7\xf9\x05\x23\x92\xfa\x5c\xdd\xea\xca\x70\x55\x77\xf7\x2a\x4c\x4c\xd1\x34\xde\x40\x26\x8c\x12\xc9\x38\xa1\xb3\x32\x68\xe5\x05\xff\x0b\x51\x63\x36\x9b\x1d\x87\xcc\x2f\xe3\x5f\x88\x97\x91\x23\x50\xb7\xbd\x27\x50\xf4\x3a\x54\x53\x2a\x37\x1c\xaa\x90\x72\x6d\x3f\xc4\x42\x12\x8a\x24\x61\xb4\x8c\x74\xa0\x2b\x51\xc9\xa9\xdc\xab\xf0\x39\x4e\x63\x12\xa0\xad\x00\x39\xde\xb5\xf8\x42\x8f\x85\x54\x1c\x71\xd9\xd5\xa0\x12\x27\xa4\xe2\x20\xab\x72\x6b\xe3\x98\xae\x69\xac\xed\x6a\xae\x5b\x1f\xc7\x54\x43\x96\x20\x42\x77\x55\x4d\x4b\xa4\x32\xd9\xc2\x84\x08\x55\x7a\xf9\x01\xa3\x92\xb3\x78\x2b\xd5\x76\xbb\x25\x95\xdb\x47\x96\xe8\x28\x11\x6d\x58\x2b\x3f\x96\xae\xc3\x2b\xfb\xd4\xea\x0e\xdf\x59\x9b\xbb\xf2\xdb\xce\xb8\x77\xe9\x77\x6e\x27\x3f\xf8\xef\x7a\xfd\xae\x67\xef\xef\x77\xea\x41\x9f\x25\x8e\xaa\x9e\xb5\x5b\xdc\x29\x12\x24\xf0\x51\x26\xe7\x8d\x40\x2c\x6c\x6b\xab\x24\xde\xc1\xdc\x2a\x8f\x8f\xec\xbd\x07\xec\x58\x00\x27\x59\x82\xc4\x1d\x34\x5f\xbd\x7a\x5d\xaa\x5d\x6a\x8f\x5b\x4d\xc1\xd5\x59\xfe\xe2\x76\xdc\x1d\xad\xce\x94\x33\xa9\x0d\x6f\x0e\x92\x39\x55\x27\x8a\xa5\xec\xc4\x58\xfa\x92\xdd\x61\xea\xd5\xb6\x3a\x3f\x7a\xd0\xd7\xa7\x60\x79\xbc\xdc\x9b\x39\x7a\x88\x1c\x9a\xc8\xa6\xcb\x33\x18\x7e\x18\x18\x90\xf1\x57\xb9\xfc\x8e\xb2\x7b\x6a\xf8\x08\xe3\xf4\xdc\x35\x99\xc7\xef\xcf\x66\xcc\x83\xd7\x85\x73\xb6\x26\x77\x96\x3f\x15\xbf\xca\x2f\xb5\x3d\x16\xf0\xda\x2a\x29\x97\x27\x7f\xb6\x79\x51\xfa\xab\xcb\xae\x7d\x98\xd3\x2f\x73\x4c\x8c\xe5\x7a\x05\x54\x10\xf9\xa6\x81\x72\x5c\xa3\x54\xcf\x68\xa5\x6a\x07\x3c\x7e\xb2\xdf\x62\xc4\x31\x9f\xa8\x09\x7c\xb2\xdb\xf0\x69\xc7\x29\x9f\xec\x33\xf8\x64\xf7\xa8\xc0\x41\xc6\xb1\x92\x50\x65\x37\xac\xb4\x73\xf6\x98\x7d\xe1\xac\x1c\x53\x38\x95\x1c\xa8\xfe\x06\xba\x0f\xfe\xf4\x0c\x8d\xb6\xbb\x51\xb1\x75\x19\xa2\x83\xb8\x1a\x72\x65\x9b\x42\xc5\x42\x29\xf9\xc9\x5c\xde\xda\xb0\x68\x59\x77\x84\x86\x6d\x30\x3d\x78\x4b\x55\x5b\xa2\x6d\x39\x40\x51\x82\xdb\x50\xe2\x6a\x2a\xb1\xb6\xce\x2c\xed\x98\x36\x6c\x59\xd3\xef\x56\x45\x29\x51\xc2\xd0\xbd\x20\x0b\x20\x1f\x31\x08\xea\x42\xa4\x1d\x6a\x6e\x94\x32\x16\xea\x42\x49\xa2\xa5\x71\xaf\x15\x98\xde\xad\x86\xc9\xff\x1b\xc5\x02\x65\x0d\x9b\xf3\xda\xa6\x6a\x2c\xe7\xc5\x84\x83\x4c\xd7\xd7\xc9\x81\xac\x20\xe3\x1c\x6f\x9e\xab\x05\xf5\x06\xf8\xb5\x51\xfa\x35\x2b\x59\xc4\xea\xc1\x65\x3c\x00\xf6\xdf\xae\xa1\xe2\x57\xb9\x80\xeb\x50\xff\x87\x57\xcf\x90\xfc\x3b\x97\xae\xa8\x26\xf3\x41\xe1\x9d\xd8\x62\x29\x24\x4e\xda\x42\xdd\xa8\xb3\x18\x73\x1b\x8a\x57\xf9\x69\x1a\xab\x12\x11\x51\x34\x2b\x8f\xe5\xc5\xd0\xe6\xc5\xa6\x88\xdc\xbc\x0b\xa9\x50\x47\x06\xe3\x90\xdb\xcb\xdb\x33\xbb\x2c\xca\xad\x9a\xfc\xd0\xf8\xdf\xb4\xef\x37\x47\xa2\x59\xee\xb3\xda\x63\xce\xa9\xfc\x57\xef\xd7\xea\xd8\xdc\xdd\xb1\x57\x76\x7e\x89\x5f\xb7\x2d\xca\x0d\xe8\x29\xa1\x88\x2f\x21\x6f\x91\x81\x44\x1c\x4e\x6a\x07\xbf\x67\x9e\xda\x7b\x9d\xf3\xbc\xaa\xdf\x13\xb5\xbf\xdc\x52\xe9\xcb\x65\x85\x81\x92\xc4\x1a\xf7\x96\xa6\x28\xb8\x5b\xb7\x65\x24\xc7\xd8\xb6\x78\x02\x0e\x8f\xa0\x74\x29\x56\x36\x1e\x3e\x47\xca\x2d\x65\x98\xbc\x83\x9f\x63\x8d\x32\x4a\x15\x52\x41\x6c\xdd\x90\x0e\x38\x49\xa5\x6d\x89\x2c\x64\x25\x50\xbd\x0b\xa8\x75\x74\x73\xc1\x86\x98\x6b\x03\x87\x1c\x51\x58\x2a\xed\x46\x58\x06\x66\x23\x31\x3d\xcf\x46\x98\xd3\x30\x37\xd9\x36\xd4\x4a\x1f\xcb\xd5\x61\xb5\xaf\x90\x5f\x7a\x1b\x6a\x83\xb1\xb6\x6e\xb2\xfb\xa2\x33\x8e\x88\xaa\x26\x94\xa8\xf9\xaf\x32\x9a\xb3\x18\x0b\x93\xda\x0e\xec\x7d\x68\xd1\xdb\x06\xcb\xc2\x36\xa0\x7b\xa1\x1e\xa6\xdc\x31\x57\x4c\xbb\xf6\xb8\xf3\xa9\x6d\x65\x9b\x94\xcd\x1b\x52\xa6\xaf\x69\xbe\x32\x0d\x6f\x26\xe3\xd5\x56\xc3\xb8\x44\xf5\x29\xae\x00\x21\x0b\xee\x30\xf7\x59\x2a\xb7\xab\xfb\x0d\x78\x55\x8d\xac\x08\xa9\x0a\xf0\x00\xa7\xd1\x70\x38\xf9\xeb\x9c\x38\x63\xf2\x00\x27\x05\xfe\x35\x9c\x8a\x92\xf4\x2f\x92\x2a\xce\x82\x5d\x56\x65\xf8\xaf\xa1\xb5\xf7\x15\xf0\x2f\x70\xcb\xbf\x09\xe0\x07\xc9\x91\x2f\xd0\xf6\x55\x6f\xcf\xd0\x53\x2c\x0f\x25\x91\x49\x85\xb0\x22\x0d\xf2\x51\x17\x65\x92\xa9\xb3\x06\xa7\xd2\x50\x53\x2f\x7c\xf3\xa2\x0d\x13\x75\xf8\xe9\x30\x3e\x8e\xc2\x31\x0a\x24\xcb\x13\x2f\x7f\x50\x59\xe4\x40\xbd\xe4\x04\xf7\x5b\x57\x48\xc4\x65\xbd\x48\x30\x5d\x48\xe4\xe2\xee\x9c\xcc\xe6\x42\x22\x89\x1d\x8a\xef\x75\x4f\xc9\xe4\x4f\xe9\x63\x5c\xee\xb5\x62\x7f\x62\x29\xa6\xfa\x4b\x8c\x21\xba\x69\x87\x96\xc8\xe5\xa2\x7b\xb3\x3a\x2c\x5c\x1c\xb0\xe5\xef\x36\xa0\x39\xef\x8c\x98\xcf\x7c\x66\xe4\x3f\x01\x00\x00\xff\xff\x11\xf4\x83\xca\x22\x24\x00\x00")

func configProvidersAwsMaster_userdataTxtBytes() ([]byte, error) {
	return bindataRead(
		_configProvidersAwsMaster_userdataTxt,
		"config/providers/aws/master_userdata.txt",
	)
}

func configProvidersAwsMaster_userdataTxt() (*asset, error) {
	bytes, err := configProvidersAwsMaster_userdataTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/providers/aws/master_userdata.txt", size: 9250, mode: os.FileMode(420), modTime: time.Unix(1473994417, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configProvidersAwsMinion_userdataTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe4\x5a\x5f\x73\xdb\x38\x92\x7f\xe7\xa7\xe8\xd0\x9a\xb5\x7d\x31\x44\xc9\x8e\xb3\x3b\xce\x2a\xbb\x4e\xac\xcc\xba\xc6\x8e\x53\x96\x93\x5c\x5d\x2e\xa7\x40\x64\x4b\x42\x19\x04\xb8\x00\x28\x59\x71\x3c\x9f\xfd\x0a\x00\x49\x91\xb2\x64\x27\x99\x7b\xbb\x3c\xc4\x24\xd1\xe8\x6e\xf4\x9f\x5f\x37\x00\x6d\x3d\x81\x68\xc4\x44\x34\xa2\x7a\x1a\x0c\x8e\xcf\xae\x86\xe7\xc7\x83\xab\xfe\x65\x6f\xfb\xf6\x16\xda\xc7\x1f\x07\xaf\xa5\x18\xb3\x49\xfb\x9c\x6a\x83\xea\x9d\x62\x33\x6a\xf0\xf4\x1d\xdc\xdd\x6d\x07\x27\x17\xaf\x7f\xef\x5f\x0e\x2f\xde\x5d\x0d\x7a\xdb\xdb\x81\x42\x9a\x48\xc1\x17\x50\x7c\x1f\x5c\x5d\x5c\x1e\xff\xd6\xef\x6d\xd3\x7c\xac\xb7\x83\x20\x08\x68\x66\xc8\x04\x0d\xe4\x59\x42\x0d\x56\xaf\x4c\x68\x43\x39\x07\x42\x16\xa8\x21\xce\x15\x0f\x82\x44\xce\x05\x97\x34\x21\x52\x91\x51\xae\xcd\xce\x2e\xdc\x06\x00\x5c\xc6\x94\x03\x51\x90\x2b\xde\x0b\x5b\xdd\xb0\xfe\x6d\xcc\x38\xf6\xc2\xd6\x6d\xae\xf8\xd6\xd6\x7f\x44\x77\x76\x50\xa5\x40\xc6\x10\xb6\xec\x98\x7d\xcf\x85\x61\x1c\x3e\x7d\x02\x82\x10\xb6\x6e\xbb\x9e\x10\x3e\x7f\x7e\x01\x89\x0c\x00\x00\x30\x9e\x4a\x08\x4f\x0a\xf9\x4c\x4c\x1c\x5f\xd8\x69\x75\x77\x43\x47\x60\x15\x04\x42\x58\x36\x7b\x06\xe4\x4c\x96\xcc\x81\x90\x58\x0a\x81\xb1\x21\x86\xa5\x28\x73\x03\xfb\x1d\x20\x44\xa1\x51\x0b\x78\x5e\x3e\x91\x04\x39\x5d\x40\xb7\x03\x85\xfa\x00\x69\x72\xa8\xf3\xb4\xa6\x64\x22\x05\x06\x77\xd6\x62\x85\x69\x88\xa6\xbc\x69\x02\xfb\x61\x98\xca\x04\xbd\x15\x02\x00\x36\x86\x24\xbb\x9e\x00\xd1\x6e\x90\xa4\x4c\x30\x29\xe0\x2f\x2f\xa3\x04\x67\x91\xc8\x39\x7f\x01\x66\x8a\xa2\xb6\xc8\x5e\x0f\x06\x94\x9b\x81\xa1\xf1\x35\x50\x6e\x1d\xb8\x28\xbd\x81\xc9\x1e\xe8\x6b\x96\x65\xd6\x02\xa5\x87\xb4\xc1\x0c\x7a\x3d\xaf\xb6\x42\x93\x2b\xcb\x6e\xcc\xac\xfc\x8a\xe5\x25\x8e\x15\xea\xa9\x9d\x97\xd1\xf8\x9a\x4e\x10\x12\x6a\xe8\x88\x6a\x2c\xe6\x7a\x27\x34\xa3\x61\xc5\xfe\xbd\xde\xca\x38\x8c\x29\x73\x4a\x39\x2b\x5a\xe6\xa5\x1e\x6e\x82\xe6\x88\x19\x1c\x96\xc6\x0b\x00\xd2\xeb\x84\x29\x20\x19\x44\x33\xaa\xa2\x98\xc6\x53\x8c\x9c\x61\x8a\xc5\x04\x00\x71\xb2\x71\x30\x00\x38\xe9\xbf\x1a\xf4\x76\x9c\x08\xce\x46\x5f\xd3\x7f\x1f\x0c\x0f\xda\xfb\xed\x83\xa7\xc9\x58\x4f\x48\xf7\x8f\x51\x26\xff\xda\xf9\x23\xd1\xe6\x69\x77\x48\xd3\xe4\xf9\xb3\x76\x82\x23\x47\x9e\x2d\xcc\x54\x0a\xf2\x35\xfd\xf7\xb0\x7b\xd0\xee\xb6\x3b\x0f\x52\x3b\xb9\xb1\x4c\x53\x29\x86\xfb\x9d\xee\xb3\x76\xb7\xdd\x3d\x78\x9a\xe8\x72\x92\x9d\xc0\x79\x41\xbe\xeb\x3d\xfd\xe9\x93\x8d\xde\x2a\x08\xee\x42\xe8\xf5\x20\x4c\x5d\x92\xfa\x68\xae\x5c\x6d\x97\xf1\xb4\xb7\x53\x04\x85\xa3\x78\x44\x8c\x13\x32\x66\xc1\xea\x5c\x17\x50\xdf\x33\xf7\xfd\xe5\xd9\xf0\xd5\xf1\xa0\xdf\x0b\xa7\xc6\x64\xfa\x28\x8a\xb4\x91\x8a\x4e\xb0\x3d\x91\x72\xc2\x91\x66\x4c\xb7\x63\x99\x46\xd7\xf9\x08\x95\x40\x83\x9a\x28\xe4\x48\xb5\xf7\x82\x8b\xe7\xb1\x54\x60\x19\x32\x61\x97\x6a\x35\xf9\xf4\xcf\xcf\x77\x61\x15\x27\xd6\x0a\xf0\xa4\xc8\xe3\x04\x47\x36\x8b\x6b\xcb\xb6\x81\xd0\x84\x10\x4b\x57\x6a\x76\x17\x15\x73\x1c\xad\x5b\x6b\x19\x37\x8f\xc9\xad\xe2\xf3\xd4\x87\x8a\x0d\x45\xcf\x6c\x0f\xd8\x44\x48\x85\x90\x60\x86\x22\x41\x11\x2f\x20\x96\x69\xc6\x29\x13\x46\xc3\xce\x9c\x71\x0e\x63\x76\x03\x9c\x1a\x54\xbb\x55\x00\xfb\xb4\x25\x36\xdb\x88\xa6\x29\x92\x19\x2a\x6d\x73\x97\x90\xb1\x54\x31\x12\xcf\x4f\x03\x61\xd5\x5a\x6b\x0a\x6f\xc1\xd5\x94\x69\x70\xdc\xcb\x44\xa5\x62\x01\x72\x6c\xad\x01\xb9\x48\xd1\x2c\x55\x62\xa8\x61\xac\x64\x0a\x74\x24\x67\xd8\x0e\xd6\x2f\x68\xcd\xa4\x75\xa9\x5b\x21\xf7\x18\xc8\x62\x73\x06\x97\x64\x3f\x9a\xc2\x5b\x70\x26\x27\x40\xc1\xe2\xa9\x36\x34\xcd\xea\xda\xbe\x61\x82\xe9\x29\x26\x25\x77\xcb\xcf\xa2\x99\x63\xea\xc0\x73\xc4\x65\x7c\x3d\x4c\x70\xc6\x62\xd4\xbd\x9d\xdd\x20\xc0\x6c\x8a\x29\x2a\xca\xab\xaf\xad\x9d\x02\xce\x35\xe3\x28\x0c\xd8\x80\x3d\x8a\xa2\xee\xf3\x5f\xdb\xfb\x87\xcf\xda\xc5\xdf\xc8\x06\x3d\xe9\x76\x49\xe7\x30\x4a\xd1\x50\x62\x01\x2d\x72\xfc\x89\xe7\x44\x52\xea\xa0\x32\x82\x6f\x30\x51\x98\x41\x25\x6a\x37\xb0\x11\xb5\x2a\xd9\x86\x57\xeb\x9e\x3a\x85\x09\xfd\x1a\x5f\x4f\x31\xbe\xb6\xab\xaa\xc8\xc0\x93\x1d\x41\xeb\x76\x75\xaa\x8b\x08\x3a\xd7\xc5\xeb\xff\xfd\xba\xd6\x88\xdc\xb5\x3e\xf2\xcf\xc3\x8c\x9a\x69\x2f\x0c\x0b\x74\x02\x32\x02\x57\x76\x5a\x4b\x95\x1a\xd9\xd9\x98\xb5\x4a\x69\xb9\x20\xd7\xe8\x28\xb7\xc0\xd9\xc1\xa5\xa5\x0d\xe7\x9b\x59\x42\xb4\x59\x70\x04\x41\x53\x4f\x72\x33\x4b\x86\xee\x53\xaf\xb5\xe3\x4c\x57\x97\xfa\x0d\x34\x26\x10\xea\x48\x27\xd1\xcd\x2c\x89\xc2\xdd\x25\x7a\x54\x5a\x56\x1c\x56\x21\xe4\xbe\x9a\x15\x69\x0d\x3a\x7c\x01\xf4\xb0\x4c\xbe\x3a\x34\xa8\xe6\xdd\x35\xd1\xd8\xbb\x16\xe0\xb5\xcc\x79\x02\x42\x1a\x18\x33\x91\x40\xc2\xf4\xf5\x5a\xb7\xfe\xb3\x75\xbb\x5c\xcc\x5d\xc3\x32\x25\xab\x13\x34\x18\x1b\x4c\xea\x71\xf2\x00\xbb\xba\x6e\x7e\x0d\x8d\x3c\x79\xda\xdb\x69\xd2\x14\xc5\xc0\xe7\x64\x2a\x67\x38\x4c\x64\x7c\x8d\xca\x7a\xdb\xbd\x5a\x10\xe7\x68\xec\xfb\x6a\x87\x17\x78\xd2\x61\x01\xff\xbd\xd6\x6d\xb3\x37\x3c\x22\xb6\x37\xbc\x0b\x02\x6f\xbb\xd6\xed\x56\x43\x17\x0b\xbc\xb6\xb0\x75\xea\x36\xf4\xcb\x7e\x2b\x6b\xcb\x75\x93\x0a\x67\x69\x18\xcb\x5c\x24\x2f\x3c\x1e\xe6\x1a\xc1\x8a\x00\x29\x40\x49\x69\x3c\x6e\x36\x74\x0a\xed\x78\x18\x14\x66\xf5\xdc\x5f\xd5\x19\x5a\x43\xde\x53\x2b\xf4\xf8\x74\x89\xd6\x04\x0e\x6e\xf1\x86\x69\x63\xd3\x35\x95\xb9\x30\xba\xa8\x24\xf5\x89\x2e\xe7\x9b\xac\xee\x56\x40\xf3\xbd\x70\xb3\x7d\x4d\xa9\x53\x16\xae\x72\xbd\x7a\xee\x68\x56\x09\x7c\x2f\x81\x89\xad\x12\xb6\x28\xfe\xf7\xb7\xff\x59\xa1\xf8\x96\x84\x10\xa1\x89\xa3\xb1\x36\x74\x54\xc3\xd9\xd2\xfa\x4d\xcb\x38\xd3\x87\x23\xa3\xc6\x7a\xa5\xa5\x58\xdf\xb8\x3b\x4a\x62\xa4\xe4\x3a\xa8\x72\x6c\xb3\x53\xbb\x4d\x9e\xa5\x05\x2e\x04\x6e\xf0\x6c\xe9\xd8\xb1\x54\x29\x35\xce\x46\x73\x66\xa6\x5e\x6e\x58\x30\x49\xaf\xc7\xba\xed\xbe\xd8\xa2\xb4\xea\xb8\xce\x67\x6f\xa6\x2a\x87\x4a\xa9\x6f\x2c\x6b\x48\x73\x6e\x58\xc6\x37\x29\xa0\xf7\xd6\xcb\x06\xaa\xe1\xf2\xf8\xf4\x84\x74\xd6\x6b\x41\x1c\xac\x82\xa2\x2c\xe9\xac\x8b\xa5\x25\x90\x94\xea\xac\xd1\x1b\x20\x4a\x85\x89\x96\x9a\x15\xb2\x41\x48\x6a\xeb\x23\x40\x07\x3a\x21\xbc\x7c\xd9\x74\x71\xbd\x0f\x6e\xcc\xf7\x63\x2e\x90\x56\x06\x1e\x9a\x55\xeb\xd8\x0a\xba\x3a\x20\x34\x69\xc3\x25\x41\x05\x11\x1b\xb9\x79\x64\x7b\x28\x10\x6d\x9a\x12\x21\xf9\x2c\x5d\x89\xc6\x07\xc2\xec\xc9\xc6\x30\x5b\x72\x03\x8d\xdc\xe1\xe7\x1e\x8c\x72\xb3\x2e\x06\x4a\x60\x99\xa3\xaa\x82\xd0\x6d\x74\x6d\x3d\x1a\x33\xa5\x8d\x07\x9b\x11\x02\x9d\x51\xc6\xe9\xa8\x56\x1d\x82\x32\x1e\x80\x18\xc0\x1b\xf3\x6c\x73\x50\x7e\xbf\xeb\x1d\x1f\xfb\xef\xff\x8b\xf3\xbd\x35\x6c\x2f\x62\xf7\x37\xdf\xbe\x3d\x10\x23\x2b\xd1\xb1\x05\x1f\x11\x28\x9f\xd3\x85\x76\xd5\xe0\xec\xc3\xf9\x1e\xe0\x0c\x85\x4f\x60\x29\xb0\xf0\x6f\x41\x7d\x2a\xa0\x2e\x0c\xec\xf6\x6a\x0f\x4e\x9c\x30\x88\xa9\x28\x99\x40\xc2\x14\xc6\x86\x2f\x8a\x79\xc7\x5c\xcb\x3d\x18\xe3\xdc\x92\xc9\x04\xc1\x16\x4f\x0d\x54\x21\x4c\xa4\x4c\xea\x18\xaf\x2d\x7c\x9c\x7d\x38\xdf\x07\x2a\x12\xe7\xcb\xf0\x01\x58\xe5\xb3\x74\x3f\x28\x84\x9c\x48\xb1\x6d\x40\xe6\x26\xcb\x0d\xe8\x2c\x57\x4c\xe6\x1a\xc2\x37\x8c\xdb\x55\xe8\x58\xb1\xcc\x48\x05\xff\x09\x1c\xe9\x35\x26\xb6\xea\xcd\x26\xb1\x42\xbb\x7f\x66\x62\x26\x63\x6a\x98\x14\xed\xb0\x60\xf7\xbb\x90\x73\x01\xa3\x7c\x72\x04\xd8\x9e\xb4\xe1\xfd\x28\x17\x26\x87\xad\xc3\x5f\xbb\x7f\xdb\x3f\xf0\x2a\xdf\x64\x52\x19\xab\xed\x70\xf0\xfe\xdd\xbb\xcb\xfe\x60\x30\x7c\x73\x32\xfc\x78\x7c\xf9\xf6\xf4\xed\x6f\x83\x5e\xd7\xab\xf6\x43\xa5\x0e\x20\x9b\x15\x4a\xad\xab\x61\xae\x2a\xd9\x87\x4a\xf5\xd9\x84\x2c\x83\x7f\x1d\x80\x36\x70\xe0\xf1\xf8\x59\x41\x84\xad\xa6\xc7\xcd\x94\x09\xc8\x94\x9c\x31\xbb\x09\x63\x62\xb2\x07\x29\x15\x74\x82\x09\x8c\x16\x45\xf7\x50\xcd\x74\x5b\x2f\xa6\x1d\x14\x8c\x50\x1b\x90\x99\x35\xb1\xc7\x12\x66\xec\x90\xa6\x09\x5f\xc0\x48\xc9\x6b\x14\xd6\x23\xa9\xd4\xc6\x76\x68\x46\x49\x5d\xf1\x79\x65\x9d\x50\xee\x96\x27\xcc\x4c\xf3\x91\xdb\x21\x7b\x71\xe5\x1f\xa6\x75\x8e\x3a\x7a\xd6\x39\x78\x1e\x54\x53\xff\xd6\xf9\x05\x26\x12\x35\x18\xe9\xd4\xf0\xb4\x6e\x19\x24\x93\x92\xbf\x80\x39\xc2\x9c\x0a\x63\x09\x38\xd2\x19\x82\x96\x29\x82\xce\xa8\x2b\xaa\x0a\xa6\x52\x1b\x32\x93\x3c\x4f\xb1\xd4\x88\x97\x1e\x22\xdc\x0a\xf8\xf0\x1b\x10\x62\x39\x5a\x86\x85\x84\xe5\x7b\xdd\x3f\xa5\x5e\xf5\xc3\xc1\xb0\xea\xfb\xec\xeb\x9d\xdd\x95\x78\xe7\x10\x99\x19\x48\xd2\x76\xc9\x29\xc1\x59\xcf\xb5\xd9\xde\x15\xd1\x6c\x42\x96\x9c\x49\x21\xb6\x92\x1b\x56\x26\x78\x2b\x0d\x82\x99\x52\x63\x97\x9a\xb8\x24\x71\x7d\x99\x9f\xf1\xa2\x34\x89\xb3\x92\xcf\xdb\xd2\x58\x25\xaf\xa5\x39\x2f\x31\xa5\xcc\xba\xbd\x30\xd0\xce\x7e\xe7\x97\x5d\xeb\x47\x6b\xa9\x25\x66\xb9\x33\xad\x65\x18\x5c\x9c\x5c\x1c\xc1\x60\xea\xba\x7a\x63\x63\xc2\x16\x83\x22\x94\xa4\xe4\xff\x00\x9f\x61\x73\x47\x30\x47\x8b\x40\xaa\x72\x8a\x16\x34\xd3\x53\x69\xfc\x4c\xcb\xf8\x1f\x6b\xdc\xd0\xed\x74\x7e\x79\x73\xd9\xef\x03\x11\x75\x3d\x1a\xc6\x5f\xb6\x1f\x55\xb9\x71\xf6\xac\x13\xd5\x61\xbc\xa4\x7f\x14\xef\xeb\xb5\xf3\x21\x8e\xab\xb5\xaa\x31\xe4\xcb\xd6\x63\x35\x6b\x7d\x65\xba\x57\x7c\x7e\xa8\xba\x34\x7b\xbe\x2d\xb7\x27\x58\xba\xfc\x23\x5a\x54\x4f\x9c\x2b\x32\xce\x4c\x19\x2f\x7f\xa9\xd9\xd9\x67\xf4\x1c\x41\x48\xe0\x52\x4c\x50\x41\x22\x5d\xcc\xed\xc1\x08\x63\x9a\xd7\xb8\xdb\x7c\x82\x22\x9f\x60\x22\x81\x89\x22\xda\x6a\xc6\xa0\x0a\xe9\x9e\x2b\x00\x1e\x25\xa4\x99\xa2\x9a\x33\x8d\x30\x43\xb5\x00\xa4\x7a\x51\xf1\x33\x12\xc6\x6e\x37\x93\x81\x4e\x6d\x65\x28\x78\xb7\xcb\x25\x14\x87\x67\xbd\x2f\x5c\x8f\x86\xc5\x0b\x90\x18\x88\xfe\x52\x50\x54\x67\x86\xc5\xe8\x5d\x68\x9b\xa3\x70\x3e\x45\xfc\xba\xb0\x88\x08\x0d\x48\xdc\x14\x79\xcb\x9c\xb7\xff\xd5\x50\x79\x1d\x04\x00\x5c\xfd\xeb\xf4\xed\xbb\x8b\x8b\xb3\xe1\xe0\xf4\xbf\xfa\xbd\xd6\x0e\x9f\x35\xe3\x35\x5a\xe1\x42\x24\x9c\x7d\x70\xc4\x40\x88\x90\x53\x74\x07\xf0\x1a\x08\xc9\x05\x33\x1a\xce\xdd\x67\x9d\x8f\xc7\xec\x66\x77\x8d\xaa\x1f\x5a\xb7\x0d\x91\x77\xe7\x40\xae\x1e\x96\x28\x60\x35\x7b\x6a\x81\x52\x20\x7c\xb3\x1c\xb8\x5d\xbb\xce\x33\x5b\x18\x7d\x4d\xf8\xe8\xac\x58\xcd\x29\x6e\x10\xca\x6d\xb9\xb7\xf1\x0b\x98\x3b\x5c\x2a\xdb\x87\x7b\x55\x26\xdc\x68\x7a\x0b\xbe\x35\x35\xd7\x65\x7c\xd9\x68\x3e\x9a\xfb\xf7\xa1\x62\x73\x83\xb8\x39\xdf\x6b\xad\xe8\xfa\xd6\xf4\xe7\x72\x3c\xf8\x7e\x3c\xaa\x43\xc0\x43\x1d\xe8\x8f\xa2\x84\xdd\x86\xad\x9c\xb0\x9c\x4e\x84\x54\xfe\x50\xf4\xda\x35\x4c\x2b\x67\x18\xf7\x5b\x8e\xd0\x1f\x99\x58\xa7\x04\x3f\xb6\xb7\x7e\xb0\x6a\xea\x72\xab\xfb\xfd\x1b\xa5\xef\xef\x94\xb7\xca\xe3\x5f\x7f\x5e\x72\x6d\xcd\xc2\x6d\xf7\x9b\x73\x6b\x8c\x0d\xbd\x29\x13\xf9\x0d\x61\xa9\x2d\xe5\x78\x63\x14\x25\xad\x9d\x5c\xd0\x14\x81\xa8\xdd\xe0\x3e\x57\x7f\x38\xb0\x89\x9b\x53\xbc\x3a\x3f\x78\xc4\x14\xe5\xa9\xcd\x4f\x75\x7d\x8f\xf0\x6e\xcc\x6b\x9e\x0c\xfd\x4c\x30\xd8\x38\x28\x0e\x07\xdd\x9d\x42\x2d\x68\xef\x56\x7c\x70\xbe\xec\x5a\x2c\xea\xbb\x9e\xc1\x86\xec\xf2\x78\xb1\xb8\xb2\xe2\x6c\x54\x34\x85\xcd\x6e\x36\x9d\xdd\x1b\x6f\x0a\x8c\xca\xbb\x9d\x2a\xc9\x56\xc6\xab\x06\x97\x0b\x6b\x8b\xb5\xa3\xab\x32\x96\x26\xbd\xbc\xb8\xb8\xea\x85\x6b\x27\x85\x8f\x19\x7e\x02\xd5\x17\xcb\x66\x93\xe5\x8a\x5c\x5e\x6b\xba\x4a\xaf\x82\x68\xd5\x88\x5b\xb0\xe3\x3a\x75\xea\x4f\xd6\xdd\xc9\x7c\x9a\x99\x05\xb1\xb6\x28\x2b\x76\xb9\xcb\xa3\xc0\xa5\x01\x39\xf6\x8d\xe0\x93\xdd\xb5\x4e\x28\x05\x6d\xf4\x42\x49\xb0\xa2\xfd\x46\x3f\x54\x04\xc5\xc3\xaa\x27\x56\xc7\xef\x49\x0a\x00\x7e\x7f\xff\xaa\x7f\xd6\xbf\x6a\x7a\x63\x75\xa2\x37\x6f\x10\x2c\xd1\xd6\x62\xb4\xa6\xdc\x44\xfe\xaa\xaf\x9d\x04\x3e\xe8\xfd\xad\xe1\x11\xb4\x6a\x3f\x10\x08\xe1\xe5\x9a\x09\x91\x27\x6d\xc7\x52\x8c\x83\x20\x88\xa9\x81\xbf\xff\xbd\x7f\xf1\x06\x5e\xae\xa1\x9d\x28\xca\x84\xf6\xb4\xfe\xf9\x28\x00\x50\x92\xa3\x7b\x00\x20\xb5\x86\xc9\x6d\x65\x02\x80\x78\xa4\x48\xcc\x12\x75\x04\xdd\x4e\xbb\xbb\x7f\xd0\x7e\x76\xd8\xee\x44\x07\x1d\x3b\xc4\x65\x9e\x1c\x01\x9d\xeb\xa0\x7f\xf1\xa6\x8a\x9c\xaf\x36\x72\xfe\x75\x31\xb8\x7a\x7b\x7c\xde\x1f\x5e\x7c\xe8\x5f\x5e\x9e\x9e\xf4\x9b\xe1\x73\x6f\xb8\xf7\xa5\x79\x7f\xe2\xde\x36\x5e\xa2\x74\xfe\x4a\x3a\x5d\xd2\xfd\xb5\x76\x89\xe2\xae\xe9\x89\xed\x04\x2d\x1e\x7e\xb9\x17\xca\x8f\x28\x54\xb3\xdc\x63\xa6\x03\x28\xa5\x0c\xe5\x0c\x95\x62\x09\x1e\x6d\x90\xe0\xec\xb2\xaa\x49\x3d\x09\x7f\x5e\x87\x02\xf5\x64\x66\xf4\x11\x6c\x17\x17\x31\x61\xab\xc6\x3c\x2c\xee\x62\x08\x42\xa8\xa3\xed\x68\x7b\x3b\x9a\x84\xbb\xdb\x0f\xe9\xe4\x61\xe0\xcf\xea\xa4\xa4\x34\x6b\x74\xb2\xcc\x7f\x44\xa7\x7a\x52\xfd\x09\xa5\x8a\xec\xbb\xa7\x55\x9d\xfd\xa3\x6a\xd5\x7e\x16\x12\x04\x1a\x95\x3b\x73\xa9\xff\xec\x43\x1b\xaa\x4c\x90\x73\x96\x32\x63\xfb\xc8\x5c\xb8\x47\x4c\xaa\x6f\x02\x9e\x1f\x1e\x1e\x1c\xfa\x14\xdf\x9e\xa5\x6d\x3d\x77\x57\x7d\xa8\x35\xf4\xa0\xdb\xd9\xae\xda\x36\xbd\xd0\xb1\xe1\x5e\x7f\xff\x0c\x0d\xf2\x5e\x37\xf8\xdf\x00\x00\x00\xff\xff\x07\x81\x61\xc9\x47\x24\x00\x00")

func configProvidersAwsMinion_userdataTxtBytes() ([]byte, error) {
	return bindataRead(
		_configProvidersAwsMinion_userdataTxt,
		"config/providers/aws/minion_userdata.txt",
	)
}

func configProvidersAwsMinion_userdataTxt() (*asset, error) {
	bytes, err := configProvidersAwsMinion_userdataTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/providers/aws/minion_userdata.txt", size: 9287, mode: os.FileMode(420), modTime: time.Unix(1473994417, 0)}
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
	"config/providers/aws/master_userdata.txt": configProvidersAwsMaster_userdataTxt,
	"config/providers/aws/minion_userdata.txt": configProvidersAwsMinion_userdataTxt,
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
	"config": &bintree{nil, map[string]*bintree{
		"providers": &bintree{nil, map[string]*bintree{
			"aws": &bintree{nil, map[string]*bintree{
				"master_userdata.txt": &bintree{configProvidersAwsMaster_userdataTxt, map[string]*bintree{}},
				"minion_userdata.txt": &bintree{configProvidersAwsMinion_userdataTxt, map[string]*bintree{}},
			}},
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

