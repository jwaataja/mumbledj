// Code generated by go-bindata.
// sources:
// config.yaml
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

var _configYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3b\x5d\x8f\xdc\x36\x92\xef\xf3\x2b\x2a\xed\x33\xce\x06\x26\x9d\x19\x67\xb3\xbb\x68\x78\x1d\x4c\x6c\xef\xc5\x07\x3b\x09\x6c\x67\x81\x7d\x12\xd8\x52\xa9\xc5\x1d\x89\xd4\x92\x54\xf7\x74\x7e\xfd\xa1\x8a\x1f\x92\x5a\x9a\xe9\xee\xb1\x81\xc3\x01\x37\x2f\x76\x93\xc5\x2a\xb2\xbe\xab\x48\x3d\x81\x0f\x5d\xb3\xae\xf1\xcd\x7f\x5f\x3c\x81\x9f\xf6\xf0\x41\x38\x57\x49\xec\xe0\xbf\x8c\xc4\x0d\x9a\x8b\x27\xf0\x5a\xb7\x7b\x23\x37\x95\x83\x67\xf9\x73\x78\x71\x75\xfd\xe7\x09\x14\x3c\xfb\xf0\xee\x33\xbc\x97\x39\x2a\x8b\xcf\x2f\x9e\x40\xae\x55\x29\x37\xcb\xbd\x68\xea\x8b\x0b\xd1\xca\xec\x16\xf7\x76\x75\x71\x01\x00\xf0\x04\xfe\xa9\xbb\xcf\xdd\x1a\xe1\xe6\xb7\x77\x70\x8b\xfb\x25\x0f\xef\x75\xe7\xba\x35\xae\x60\xb1\x88\x70\x9f\x74\xa7\x8a\xd7\xb5\xee\x8a\x31\xe8\x13\xf8\xe5\xd7\xcf\x6f\x57\xf0\xb9\x4a\x38\x40\x5a\xc2\x60\x20\xaf\x25\x2a\x07\xef\xde\x78\x50\x4b\x28\x72\x42\xe1\x11\x5f\x14\x58\x8a\xae\x76\xfd\x66\xde\xf8\x01\xc8\x75\xd3\xd0\x4a\xa7\x61\x8d\x20\xda\xb6\x96\x58\xf0\x2f\xed\xc6\x64\xdf\x95\x44\x0a\x0a\x0d\x4a\x3b\xd8\x09\xe5\x40\xa4\xe5\xeb\x3d\x04\x12\x97\x60\x91\xd1\x61\xd3\xba\x3d\x58\x67\xa4\xda\xc0\xb3\xc5\xe2\xb9\x47\x17\x56\xac\x60\xf1\x33\xd6\xb5\xfe\x06\xde\x81\x68\x40\x30\x3d\xf8\xbc\x6f\x11\xbe\xa9\xb0\x6e\xa1\xd4\x06\x04\xd4\xd2\x3a\xd0\x25\xaf\x12\xaa\xb0\xcb\xc5\xe4\x00\x95\x50\x0a\x6b\x86\x77\x15\x12\x1e\xa6\xae\x1c\x1a\xe8\x5a\xad\x48\x2a\x0a\x73\x27\xb5\x9a\x3d\xd0\x4e\xda\xea\x70\x75\x58\x42\xff\xa5\x51\xa3\x75\x22\x74\xf4\x7c\x1e\x6c\x28\xd0\xd7\x7e\xf3\xb4\xa8\xb3\x48\xff\xb4\xb5\xd8\x83\xe8\x0a\xa9\xa1\x94\x35\xda\x25\x0b\xd5\xed\x34\xd8\xae\x6d\xb5\x71\x58\x40\x5e\x69\x99\xa3\x05\x61\x10\x16\x65\xd9\xb4\xb8\x59\x00\xa1\x59\x88\x6d\xae\xd5\x76\xe1\xe9\x11\x2a\x34\x59\x60\xd0\x2a\x81\x5e\x5c\x5c\xfc\xbb\xc3\x0e\x93\xc4\x3f\x0a\x27\xe9\x38\xc2\x41\xd3\x59\x47\xe2\x6e\xd0\x81\x36\x80\x77\x39\x62\xe1\xc5\xee\x8c\xdc\x90\x6a\x0b\x70\x46\xe4\xb7\x60\x6f\x65\xeb\x09\xf1\xef\x8c\x7e\x67\x86\x50\xad\xe0\x6a\xf9\xc3\x63\x91\xd3\xae\x59\xb6\x3d\xfe\x38\x74\x1f\x89\x0f\xe2\x4e\x36\x5d\x13\xf6\x55\x74\x0c\xa1\x40\x2a\xb0\x98\x6b\xd2\x0d\xf8\xe4\x25\x73\xc5\xe2\xec\x94\x41\x92\x4e\x4e\xcc\x8c\xe0\x9e\x54\x23\xee\x32\x7f\x9c\x38\xbe\x82\xab\x59\x3a\x16\x5a\x34\x69\x6b\x0f\x51\x88\x30\xf6\x80\x84\xcd\x5a\x34\x59\x9c\x5d\xc1\x0f\x89\xd0\x3b\x0b\xb6\xea\xca\xb2\x26\x05\x42\x25\xd6\x35\x16\xb0\xab\x50\x25\x4d\xb4\x4e\x18\x67\x7f\x64\x78\xd1\x39\xdd\x08\x27\xf3\xcc\x2f\xc2\x8c\x76\x5d\x8a\xda\x62\x44\x78\xa3\x94\xee\x54\x8e\x81\x45\x52\x95\xda\x34\x9e\x4b\xc2\x79\xa4\xb8\x91\x4a\x11\x3d\x5d\x06\xfd\xa3\x9d\xad\x45\x7e\x1b\xa8\x04\x14\x99\xc2\x5d\xd8\xff\x0a\x9c\xe9\x12\x8d\xf7\x5a\xb7\xde\x59\x0c\x30\xb1\xa2\xf9\xad\x97\x52\x49\x5b\x61\xe1\xb9\x50\x6b\xdd\x66\xc1\x31\xa4\xcd\x5e\xf4\x06\x99\x94\xf3\xa6\x28\x0c\x5a\xeb\x4f\x5d\xe9\xae\x2e\x40\x38\x47\x26\x36\x36\x47\x8f\x56\x78\xe8\x15\x2c\xae\x5f\xfc\x65\x79\xb5\xbc\x5a\x5e\x27\x63\xfb\x4d\x1b\x77\x22\x1a\x32\xb4\x15\x2c\xfe\xfc\xa7\xbf\x7c\xff\xd7\x7e\xbd\xb0\x76\xa7\x4d\xc1\x12\x8e\x3b\x55\x1b\x5a\x6f\xd1\x6c\xd1\x4c\x9c\x88\xd2\xd0\x86\x45\xc7\x9c\x43\x84\x1b\x7a\x87\xdf\x2d\x1a\x25\x1a\x64\x82\x31\x2c\x79\xf0\x2e\x4c\xad\x60\x11\x27\xd2\xb2\xbf\xcb\x1a\x5b\xe1\xaa\xe0\x55\x0c\xb4\xd7\x2f\xd8\x99\x78\xcf\xd9\xb9\x0a\x95\x93\xb9\xe0\xcd\x0b\x0b\x02\x0c\x6e\xa4\x75\x68\xb0\xe0\x05\xb3\xe7\x88\x38\xa4\x05\xc5\x66\x7b\xec\x44\x84\x29\x6b\xaf\x5f\x8c\x02\x98\xe7\x7c\xd4\xe2\x28\x01\x41\xc6\x6a\x31\xef\x0c\x0e\x7c\xf2\x8f\x49\x79\xe7\x66\xa1\xd0\x68\x39\xe0\x6c\xd1\xc8\x72\xcf\x48\x73\x34\x4e\x96\x74\x36\x24\xed\xa3\x21\x2f\x1a\x3a\x7a\x40\x97\x6b\x65\xe9\xb4\x2a\xdf\x2f\xe1\x9d\xa3\x03\xad\xd1\xf2\x49\x6a\x14\x5b\x04\x57\x49\x0b\x5a\x5d\xc2\xba\x73\x50\x48\x4b\xe6\x07\xd2\x81\xf4\x51\x81\xbc\x6e\x25\xb6\x52\x6d\x02\x42\x69\x6d\x87\xf6\x40\x23\x44\x24\x4c\x2c\x37\x08\xa6\xf3\x36\xd1\x74\xb5\x93\x2d\x21\x54\xd6\x09\x45\x6e\x5c\x97\x07\xc2\x8d\xa7\x3d\x30\xe2\xa1\x5c\x87\x07\x25\xb1\xcc\x89\xec\x10\xe6\x74\xd1\xd1\xca\xa1\xd8\xee\xa3\x4c\x79\xc6\x7d\xd4\x43\x0e\x72\x1a\xc1\x5b\xdc\x0f\xe9\xdd\xe4\x39\x99\xbc\xd3\xb7\xa8\xe8\x1f\x90\x4a\x3a\x29\x6a\xf9\x07\x26\xdd\xd9\x49\x57\x11\xda\x56\x18\x41\x5e\x76\xbd\xf7\xa9\x80\x9d\xdb\x8c\x18\x21\x24\x79\x9c\xb6\x2f\xbf\x2e\xf3\xeb\x1e\x52\xe4\xe8\x82\x45\x5d\xef\x87\x8e\xc5\xa0\x33\xfb\xa1\xd6\x0e\x55\x43\x94\x94\x89\x14\xd2\xf6\xaa\xe3\x75\x9e\x57\x65\xc1\xf1\x8f\xbd\xec\xcf\x7a\x07\x8d\x50\x7b\x70\xb2\x41\x1b\x5d\xd9\xa1\x41\x31\xe5\x83\x54\xc5\x13\x1d\x12\x08\xd0\x76\x05\xd7\x57\x13\xfc\x21\x74\x1e\x52\xd8\x09\xb2\x04\xf5\xed\x1a\xdd\x0e\x71\x98\x42\x85\xb3\x46\xa4\x43\x42\x92\x52\xae\xad\xa8\x57\xf0\x03\x39\x79\x91\x57\x7d\xf2\xf1\x9a\x7e\x81\xd5\x6a\x63\xc9\x19\xb9\x0a\xf7\x2c\xa0\x42\xef\x54\xad\x45\x81\x85\xc7\x94\xb8\x31\xb2\x89\x14\x92\xb5\x13\xb5\xd7\x72\x4b\x5a\x42\x89\x21\x23\x2e\xa4\xc1\xdc\x69\xb3\xa7\x74\xe0\x83\xfc\x29\xc5\x60\x5a\x96\x11\xec\x0a\x7e\xb8\x7e\x91\x7c\x3c\x1a\xa9\x0b\xf6\x1d\xb2\x41\x9f\xb6\x04\x0e\x60\x2d\x5a\x4b\x81\xb2\xd4\x06\x49\x88\x5a\x6d\x48\xc3\xf3\x1a\x05\x79\xce\xd2\xe8\xc6\x3b\x21\x22\x7c\x49\xf4\x2a\xdd\x99\xa0\x8f\x78\xd7\x4a\x83\x19\x61\x5d\xc1\x8b\x3f\xdd\x43\x2f\x72\x15\x45\x5e\x41\x5e\x61\x7e\x1b\xdd\x98\x3f\x4d\xc9\x69\x13\x61\x2a\x40\x3a\x6c\x2c\x93\x69\xa4\xea\x1c\xda\x98\x65\x62\x7e\x3b\xe6\x78\x48\x8b\x13\x27\x28\x60\x39\x3a\x04\x23\x0d\x98\x96\xf0\x56\x6d\xa5\xd1\x8a\xb3\xf6\xad\x30\x92\xf8\xed\x8d\x85\x3d\xa0\xaf\x03\x3a\x8b\x05\x54\x68\x82\xcd\x27\xf6\xae\x60\xf1\x1f\x3f\xff\xfa\xe1\xed\x77\x4b\x46\xfa\x5d\xc3\x1e\xad\xf8\x17\x65\x9b\x5b\x5d\x77\x0d\x4e\x0a\x0c\x3f\x1c\xf0\xc4\x4c\xe0\x6a\xf9\xa2\x4f\x28\x76\xe4\x97\x3d\x18\x88\xba\xd6\xbb\x3e\x7b\xa0\x29\x82\xbe\xba\x4e\x9a\x2b\x37\xd5\x7d\xf0\x95\x9f\xa3\x05\x7f\xbd\xb8\xb8\x10\x45\x23\x55\x5f\xf1\xbc\x65\xd5\x02\x3f\xfa\xe3\xa1\xfb\xe0\x70\x20\x6d\xf4\x14\xac\x7e\x97\x40\x26\x12\x32\x6b\xc8\x85\x22\xd6\xe0\x1d\xe6\x5d\x70\x45\x34\xdd\x87\xd2\x59\x4b\x7e\x1f\x0a\x18\x26\x0b\x14\xcc\x0f\x5d\x17\xc7\x26\xb2\x63\xaa\x8b\x38\x51\xae\x42\xb6\xc6\xd0\x24\x7a\xde\x1c\xa7\xcb\x1c\x64\xfa\x38\xae\x39\x57\x0c\xf8\x82\xbf\xb1\x21\x0f\x97\x4d\xab\x09\xcc\xd2\xce\x29\x82\x86\x9d\x87\xad\xa4\x8a\x8a\x57\x33\xa9\x15\xff\x97\xfe\xbe\x85\xc5\xa7\xae\x45\x43\xb9\x09\xc9\x96\x8b\x94\xd5\xbc\x8e\xa1\x30\xb9\x2f\xa0\x9a\xce\xca\x1c\xa4\x5a\xc2\x8d\x17\x0b\xcd\x57\x14\x6c\x6d\x85\x75\xdd\xab\x5b\x30\x4f\x02\xcf\xa6\xba\xf5\x81\xc6\x17\x1e\x04\xad\x15\x9b\xe1\xce\x94\xce\x68\x33\x59\x49\x15\x6e\x86\xc6\x68\xb3\x82\xc5\x2f\x1a\x1a\xe1\xf2\x8a\xfc\x7a\xc8\x81\x3a\x55\xb0\xd5\xf0\x9e\x12\x91\xe5\x62\x80\x49\x65\x69\x0b\x59\x6b\xb0\x94\x77\x09\xe1\x8d\xf7\x6f\x21\xcc\x53\xd5\xc6\x68\x75\xe7\xac\x2c\xd8\xf1\x3c\x8c\x38\x15\x33\x11\x21\xd5\x77\x06\xff\xdd\xa1\x25\xdd\x49\x81\x9a\xbc\x39\x2c\x9b\xef\xbb\xbe\xc6\xe8\x11\xf1\x41\x75\x8b\x6a\x84\x85\xd7\xe6\xec\xac\x68\xf9\x1a\x81\x40\xb0\xa0\xd2\xf8\x22\x4a\xb5\x77\xbc\x95\x30\x22\xe7\x5a\xd8\x7a\x81\x14\x68\xe5\x46\x51\x60\x8f\xc0\xde\xa9\x29\x2a\x16\x6a\x70\x78\xe7\x12\xdf\xc7\xaa\xfa\xab\xaa\xf7\xa0\x15\x52\x8d\x1b\x90\x3e\x73\xbc\x23\x63\xdd\x73\x3a\x0f\xd1\x08\x99\x2e\xf3\x73\x05\x8b\x6f\x42\x34\x25\x62\xc4\xf1\x39\x89\x86\x42\x69\x70\x4a\x72\x47\x14\xbf\x75\x2c\xc3\xa4\xaf\x8a\xb8\xd2\x18\xb0\x88\xdc\x90\x54\x9b\x2c\x24\x70\xbd\x46\xbc\xf6\x13\xec\xba\x3b\x63\x50\xb9\x7a\x1f\xd3\xbc\xa2\xef\x22\xfc\x84\xb5\xde\x11\x50\xdf\x6a\x60\x3f\x1c\x39\xd3\x97\xe3\xeb\x7d\x9f\xbf\xc1\x5b\xf6\xdc\xc1\x31\x54\xc2\x06\x6c\xae\x32\x88\xa1\x0b\xd4\x19\x76\x37\xba\xa5\xa8\x19\x8e\xfb\x04\x44\x2d\x85\x45\xbb\x82\x9b\x44\xcf\x5b\x39\x9b\x6c\x70\x31\x51\x52\xd1\x60\x07\x3b\x5a\xa6\x6c\x34\x63\x33\xf6\xce\x06\xfe\x06\x9a\x64\xe3\x7d\x1b\xa3\x99\x59\x7b\xe9\xbd\x1a\xfc\x8d\xfc\x16\x8b\x71\x1e\x2e\xd2\x28\xd0\xe6\x46\xb6\xbe\x3c\x7e\xd3\xff\xa0\x50\xb9\x53\xa9\x65\x12\xd9\x90\x2a\x57\x6e\xdf\xc4\x51\x69\x93\xc7\x8c\x78\x93\x0a\xc0\x3f\x84\x91\xba\xb3\x69\x24\x34\x10\xc4\x9e\xb8\x60\x29\x42\x71\xe1\x32\x54\xc9\x41\x00\x0e\xbb\xa5\x02\xaa\xec\x42\x03\xc8\x08\x65\x6b\xae\x79\x02\xb1\xfe\xcf\xa7\x7d\x9c\x68\x6a\x57\xa1\x81\x5a\xa8\x4d\xc7\x5a\x0e\x6f\x34\xa9\x38\x18\x6c\x34\x79\xab\x08\x49\xbb\xe1\x42\x97\xf3\x50\x58\x3c\x5d\xc0\x33\xdb\xe5\x15\x6d\x6b\xf1\xd4\x2e\x2e\x61\xf1\xb4\x58\x5c\x02\xba\x7c\xf9\x7c\x42\x30\xe6\x39\xb6\x5b\x5b\x27\x1d\x07\x0d\xc6\x63\x3a\xc5\x79\x40\x21\x9c\x58\xc2\x47\x22\xca\x8e\xa6\x42\xdb\x13\xdf\xc9\xba\x86\x5c\x70\xc3\xa8\x6f\x4c\x35\xd2\xae\x91\x5c\x6a\x2a\x84\x7b\x43\x8a\xba\x75\x31\xd8\x03\x79\x72\x51\x14\x8b\xc9\x58\x3f\xd2\xab\x92\xcf\xb9\xe2\xf8\x48\xfc\x8b\x9b\xa2\xb0\xa9\x29\xa4\xfb\x96\x88\x97\x87\x80\x06\x0b\x29\xc0\x4a\x87\x31\x13\x3d\x34\xd5\xa9\xe5\x07\xeb\xef\x4c\xdd\xfb\x5d\xf8\xfd\xe3\xfb\xd4\x42\x22\xeb\xe3\x7e\x24\xb3\x8d\x90\x8a\xa2\x48\x82\x5f\x1c\x22\xda\x8a\x5a\x16\x87\xce\xe4\x17\x0d\x3c\x1e\x1d\xc9\x8e\x7c\x8b\x0f\x0f\x09\x6b\x6b\xf4\x56\x52\xe8\xfd\xfd\xe3\xfb\x67\xf6\xf9\x01\xe6\x80\xd0\x69\x9d\xd5\x5a\x6d\x12\xe6\x7f\xea\xce\xf8\xc9\x67\xf6\xb9\xc7\x8b\x92\x35\xcb\x69\x0d\x04\xca\x65\xa1\x02\x5e\x00\x3a\x67\x47\x44\x86\x42\x1e\xbc\x35\x9a\xaa\x8e\x20\xf8\x66\x09\xbf\xe8\x1e\x19\x07\xcd\x35\xe5\x88\xa2\x28\xf0\xf0\xa8\x5a\x61\x68\x5f\xf1\xec\x0a\x16\x2f\xd7\xaf\x9e\xda\x97\xdf\xad\x5f\x79\x78\x78\xb9\x7e\x75\xcd\x3f\xbd\xbc\x86\x12\x59\xbd\x5c\x9b\x57\x2f\x25\xc3\xcb\x57\x5e\x7c\x4f\xed\x98\x00\x95\x06\x91\x8f\x0f\x90\x78\x5a\xf4\x34\xec\x7d\x62\x67\xd9\x74\x4d\x76\xc0\x45\xc6\x68\x5e\x4d\xb0\x8c\x02\x9b\xa7\x54\x74\xac\x53\x81\x8b\x06\xd6\x98\xcc\xc2\x17\x11\x91\xdd\xd1\xad\x8b\xa2\xa8\x75\x2e\xea\x93\x4c\x83\x21\xa7\xf6\x51\xcf\x19\x08\x27\x76\x8f\xb1\x0f\x62\x0c\x13\xf2\xed\xde\xbd\x75\xd8\x8c\x18\x76\x82\x99\x08\xb3\xe9\x28\x6b\x1f\x6a\x76\x1c\xeb\x35\x98\x2a\xcf\xd9\x43\x29\x9d\xc5\x14\x29\xb3\x43\x35\x1e\xe6\x4e\x34\x31\x59\xd6\x6f\x39\xa3\x14\x53\xe6\x38\x5c\xfb\xf7\xfe\x40\x61\x16\xc4\x56\xc8\x9a\xa2\xdf\x18\x55\xa7\x28\xcb\xdd\x28\xf9\x07\x16\x99\xdb\xb7\x3d\x1a\xc2\x41\x03\x7d\xb7\x87\x77\x04\x4d\xfb\x3d\xf1\xb2\xf9\xbe\xfb\x62\x5b\x8f\xd1\x2a\x31\x8a\x4f\xf5\xff\x86\xfe\x7f\xdd\xd0\x15\xde\xb9\x93\xec\x9c\x00\xa7\x66\xae\xbe\xaa\x99\x8f\xc2\xa0\x0f\xff\x40\x74\xb9\xf6\xbe\x2f\x85\x7d\x12\x8f\x41\x79\x9f\x5f\x93\xf2\xa1\x02\x4b\xa9\xd0\x6b\xaf\x28\x8a\x65\x48\xa5\xa9\xf6\xe6\xa6\xc6\xd1\x83\x27\xd0\xc9\xd1\x73\x7b\xe6\xd1\x7f\xed\x5c\xdb\x39\xbf\xc1\x51\x0b\xa6\x6f\x5c\xf8\xe6\x0b\xc8\x32\xe6\xe5\x9c\xfe\xa9\x98\x73\x27\xad\x9b\x75\x71\x21\x4d\x0f\xdd\x1a\x2e\x75\xc2\xd0\x1c\x25\xcb\x7a\xb9\x7c\xb1\x25\x8a\xa4\x57\x51\x27\xc2\x1a\x96\xd0\x09\xfc\x19\x40\x4f\x59\xe4\x27\xa7\x4e\xb1\x9f\x3b\x37\x8d\x8a\x4c\x1c\xdd\xcd\xac\x75\xe7\x0b\xfe\x78\xde\x78\x7f\xd3\xeb\x0b\xf1\x94\x52\x76\xbc\xe3\x2b\xa6\x53\x79\xe9\xb9\x30\x66\x66\x40\x6e\x21\xf9\x86\xcb\x60\x7f\xeb\x3d\x24\xe3\x8f\xec\x2c\xb5\xc9\xd1\xde\xca\xf6\x38\x2f\x13\xe8\x84\x59\xe5\xb9\xba\xf6\xae\x61\x43\x72\x58\xef\xf9\x86\xd0\x4e\xd9\x73\x94\x07\xfd\x7d\x65\xcb\x7e\x6d\xca\x83\x4a\x58\xef\x7b\x69\xe7\x72\x1d\x68\xb5\xc7\x38\x91\xee\xf2\x4e\xe7\x48\x5c\x32\xc3\x99\xf6\xab\xb2\x66\xa6\x8b\x70\x6f\x42\x31\xdb\xa3\x98\x68\x09\x79\xe8\x56\x18\xdf\xc4\x9a\xc3\x0f\x87\x77\xb7\x53\x76\x27\x2f\x79\x1e\xc7\xa9\x90\x3c\xce\x64\x82\x9a\xf0\xb5\x7a\xac\x61\xf6\xad\xb6\xf1\xab\x83\x87\xb9\x19\x01\xb3\x0a\x45\x81\xa6\x8f\x79\xaf\x63\x3f\x86\xce\x45\x63\xe3\x9d\xf2\xc6\xb2\x7b\x57\xdf\x70\xa7\x6e\x06\x07\x23\xf9\x97\x96\xaa\x39\x21\x06\x78\xb8\x09\x8b\x68\xf8\x4c\xdd\xfb\xa0\xb7\x68\x53\x1b\x04\xa4\x72\x3a\x3c\x3f\x09\x82\x8e\x8f\x31\x64\xe9\xf5\xa6\x16\x7b\xbe\x7e\xe4\x9b\x66\xa7\xc1\xea\x06\xd9\x8d\xd5\xf6\x78\x69\xc8\x55\xba\xcd\x84\xc1\xac\xe6\xab\x3c\x39\xc8\xc9\x7e\xb7\x68\xb8\x2b\x24\x94\xaf\xe6\x23\x69\xca\x13\x12\x38\x97\xcc\x87\x49\x8a\x54\x19\x6d\x3a\xeb\x5f\x6a\xf0\x13\x14\xa5\x77\x84\xcf\x9f\xc7\x4f\xc5\xde\xd5\xad\xac\x4f\x28\x26\x08\x6a\xc2\xe5\xdb\x33\x59\xfc\xc9\xe9\x60\xd2\x7c\xd3\xa3\x0a\xbe\x79\x50\x16\xa4\xb3\x87\x97\x1d\xd1\x4e\xe8\xb8\x27\x16\x3c\x09\x74\xb2\xd1\x7a\xb6\xe2\x79\xc8\x64\xde\x53\x28\xea\x93\xfd\x70\xbb\xe3\x5b\xd2\x93\x6a\xe7\xa8\xb4\x0d\x8a\x22\x2b\x85\xac\x3b\x33\x28\x0c\x84\xac\x7d\x51\x43\xd3\x3e\xc9\xda\xc8\x2d\xbf\x33\xa8\x91\x12\xb0\xb9\x06\x2c\xfd\xd9\x5c\xa8\x07\xd0\xd1\xf4\xc3\x6b\x95\x56\x7d\x6f\xfa\xa0\x2f\x1c\xdf\x03\x10\x12\xa5\xd5\xb7\xfd\xf5\x13\x5f\xdd\x28\x67\x24\xda\xa1\x74\xc2\x4b\x8a\x93\xc4\xe3\x61\xa7\xf2\x21\x6e\x13\x87\x67\x67\xa6\x83\x8f\x75\x80\xe3\x46\x68\x4c\xd6\x53\x0b\xf5\x9e\x24\x76\x5e\xa6\x52\xf9\x32\x4d\x2a\x87\x1b\x34\x3d\x17\x55\x9c\x82\x30\x05\x3b\x61\x53\xbb\x67\xae\xff\xc2\x2e\x40\x86\x72\x22\x94\x12\xab\x23\x29\xcc\xc0\x57\x52\x7e\x7d\x62\x4a\x98\x40\x27\x3c\xa5\x99\xd9\x64\x70\x5c\x5b\x7c\x8d\x4c\x90\xeb\x81\xaf\x9b\x06\x66\x5a\xd5\xfb\x87\x63\x3d\xd1\xe1\x46\xf2\x94\xf2\x61\xa1\x87\x77\xe3\xec\x72\xb8\xe1\x13\x53\x4b\xd5\x35\xfe\x6a\xf2\x04\x99\x44\xd0\x29\xeb\xf3\x2f\x28\x63\x54\xd7\xac\xd1\xb0\xae\x07\x3d\xf7\x57\xa5\x9a\x7c\x83\xbd\x7d\x64\x21\x43\x45\x70\x38\xd8\xf0\x2e\xa3\xb7\xa1\xbe\x16\xe6\x3b\x59\x7f\x4d\x5b\x44\x76\xf3\xd2\x01\x8f\x4e\xf5\x1d\x09\x74\xca\xa3\xae\x99\xf7\x1c\x8f\xaf\x5f\xe6\xb9\xf7\x38\x2f\x91\xba\x1c\x89\x5d\xa3\x4b\x9b\x83\x16\xc7\x03\x4a\xd9\xd6\x9d\x11\x75\x7a\xb6\x76\x84\xf7\xf3\x17\x4b\x8c\xb0\x15\x9d\x3d\x21\xb1\x62\xb0\x73\x39\xf8\x9b\xe0\x32\x7f\xfc\xf8\xee\x94\x74\x9d\x57\x24\xfb\x7d\x1b\x1a\x50\x15\x06\x54\xd2\x82\xa8\x29\x4a\xee\xfd\xf6\x8b\x4b\xf0\x2d\xaf\x53\xaf\xd2\xd2\xc1\xc7\x4d\x20\x4a\xd9\xfd\xf0\x74\xcf\xbc\x36\x5e\x47\x1f\xe7\x57\x84\x9c\xe8\xa1\xc1\xcd\x99\x56\xfc\x31\xa0\xea\x53\xa5\x90\x77\x84\x37\x7a\xc7\xf8\x19\x58\x95\xf5\x77\xe9\x89\xb3\xfe\x3d\x72\x60\xe5\xe4\xae\x7d\x4a\x60\xc8\x03\xe6\x5d\xca\x26\x1f\x58\x1c\x38\x57\x6b\x71\x82\xf7\xf3\x70\x53\xae\x9d\xcd\x33\x42\x13\xea\xc5\x78\x47\xc9\x71\x87\x5f\x77\x1d\x4f\xd0\xfc\x43\x9d\x18\x3c\x26\x18\xfa\xea\xce\x76\xfc\xa2\xaa\xec\xea\x7a\x9f\xd6\xf5\xa7\xb6\x78\x42\xed\xcc\x60\x33\x9a\x72\xf6\xa1\x2d\x06\x7f\xe5\x23\xe8\x7a\xef\xaf\xf9\xb8\x2e\xa9\xeb\x18\x57\xf9\x71\xcc\x31\x16\x30\x6c\xe6\x0f\x70\x68\x23\x3c\x3a\x75\x25\x06\x6d\x77\x4a\x91\xe6\xe1\xce\x75\x26\x1f\x79\xd5\xd9\xde\xe4\x0c\x57\xe2\x2b\xb8\xc7\xf8\x12\x7f\xa2\xa9\x33\x09\xe3\xf7\x78\x13\x8b\x2e\x7e\x21\x70\x94\x67\x3d\xec\xb4\x3d\x77\xcf\xb8\x3d\x37\x5d\xf8\x14\xb5\x27\x7e\xe9\x50\x48\xcb\x4f\xee\x8b\x90\xf2\xe8\x54\x0f\xff\xa7\x4d\x0f\x76\xb9\x13\xca\xc3\x27\xb5\x0e\x28\x47\xf3\x77\xcf\xbd\x75\x79\x6a\xc3\xef\x12\xee\x33\x2f\x5e\x77\xc0\xfc\x88\x95\x0a\xd9\xcd\x23\xb0\x86\x75\xb1\xd7\x5f\xea\xba\xd6\x3b\x4e\xbf\x9f\xda\x28\x29\xff\x0c\xfd\x04\x31\x79\xc0\xa9\x2c\xaa\xae\x9c\x19\x3c\xd7\xc0\x85\x2a\x74\x23\xff\x08\xad\xf3\x2f\x4b\x45\x94\x76\x19\x2a\xdd\x6d\xaa\x87\xde\xa4\x38\xf0\x30\x73\x46\x30\x7c\xb7\x21\x22\x8f\x0e\xeb\x4b\x3f\x1a\xa5\xe2\x0d\xc1\xaf\xee\xa5\x11\x60\x92\x5d\x9c\xd4\x82\x9d\xed\xbe\xce\x36\x5f\x1f\x4c\x51\x6a\xc1\x1f\x9e\xc0\x56\xfb\xeb\x7c\x42\xfb\x88\x0e\x6c\x0c\xb2\x84\xa6\x18\xde\xa6\xf9\x6b\xaf\xe8\x63\x78\x7a\x40\x86\x0a\x91\x03\xfc\xf4\xc7\x60\x13\x6f\x72\xb8\xf8\xfe\x3d\x32\x17\xbb\x75\x23\x9d\x43\x93\x4d\xb0\x5d\xfa\x00\x1d\x01\x7c\x01\x1c\xb7\x72\x39\xa5\xb5\x84\x4f\xb7\xb2\xe5\x56\x80\xec\x5b\xb2\x43\x71\x9d\xde\x27\x7e\xb0\x45\x3c\xdf\x21\xfe\x32\xf9\xfd\x2f\xb5\x89\x1f\xaf\x10\xf7\x20\x3c\x57\x27\xee\x41\xf3\x08\xb5\x88\x98\xce\xd6\x0c\xa7\x37\x9b\x1a\x6b\xad\x4f\x30\xe7\x1e\x76\xa2\x15\x6e\xd0\xe0\x19\x4b\xff\x33\x2f\xb2\xb0\xab\x90\xe3\xba\x36\x2c\x1c\xa7\xf9\x6b\x9c\xe8\xd1\x47\xdf\xef\xf4\xde\x8b\x3f\x98\xf3\xdf\xf0\x48\xbe\x67\x9d\x6a\xde\xc8\x17\xdf\x73\xf5\xc2\x5b\x28\x32\x5d\x96\x2b\x58\xbc\xd7\x9a\x99\x91\x7c\x5b\x98\x06\x5d\x96\x87\xad\x9e\xb8\x50\x3d\xb8\x4e\x8d\x99\x79\x72\x24\x1a\x81\x4f\x59\x9a\x66\xe7\xa6\xe6\xc7\xcf\x0e\x57\x51\x3a\xe9\x5b\x82\xf8\x9d\x5d\xfa\x12\x4c\xab\xef\xc6\x8c\x39\x85\xc7\x37\x09\x5d\x8f\xe8\x5c\x7e\x9f\x86\x23\xf1\x7e\x8b\xc6\xf2\x27\x5c\xc7\xb8\x1e\x00\x27\xcc\xdb\x7e\x49\xf3\x21\x5a\x74\x40\x3e\xfa\xbc\xe6\x18\xeb\xe2\xce\xfb\xaf\xaa\xfa\xa1\x64\xf8\xf1\x94\xe1\x55\xfb\xd1\x43\x32\xdc\xf4\x8c\xfa\xec\x46\xec\x6b\x4e\xbd\xfc\x29\xc3\x2b\x77\x59\x82\x50\xfd\x63\x1f\x72\x3f\xe1\x19\xcb\x25\xe8\x39\xa6\xf8\x65\x7c\x31\xb2\x93\x27\x5c\xb5\xb4\xc2\xd8\xe1\xed\xca\xf8\xb5\x72\x40\x37\x7a\xb1\x41\x2b\xa6\xef\x55\x3a\x97\xe9\x32\x33\x74\x80\x84\xeb\x1f\xbc\xda\xa6\x47\x7d\xf1\xeb\x07\x3e\x9f\xa8\x3b\x8c\x57\xf9\xa5\x7f\x74\xa2\x8a\xe1\xef\xc3\xbc\x36\x74\x34\x83\x58\xc6\xc1\x28\x72\xcb\x3e\x80\xc0\xc3\x0c\xf2\xe2\x71\xe8\x48\x79\x6f\xcf\x7c\xa7\x0f\xd0\xfd\x4f\x00\x00\x00\xff\xff\x0e\xe1\x8a\x2b\x3d\x3e\x00\x00")

func configYamlBytes() ([]byte, error) {
	return bindataRead(
		_configYaml,
		"config.yaml",
	)
}

func configYaml() (*asset, error) {
	bytes, err := configYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.yaml", size: 15933, mode: os.FileMode(420), modTime: time.Unix(1494893941, 0)}
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
	"config.yaml": configYaml,
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
	"config.yaml": &bintree{configYaml, map[string]*bintree{}},
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

