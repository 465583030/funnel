// Code generated by go-bindata.
// sources:
// config/default-config.yaml
// config/gridengine-template.txt
// config/htcondor-template.txt
// config/pbs-template.txt
// config/slurm-template.txt
// DO NOT EDIT!

package config

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

var _configDefaultConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x59\x5f\x6f\xe3\xb8\x11\x7f\xf7\xa7\x98\x26\x5b\xf4\x0e\xb0\x14\xa7\x87\x2b\x7a\x06\xf2\xe0\x38\xd9\x6c\x70\xc9\xae\x1b\xfb\x9a\xb6\x2f\x01\x25\x8d\x25\x9e\x29\x52\x4b\x52\xf6\x69\xd3\x7c\xf7\x62\x48\x4a\xb2\xb3\xce\x26\xdb\x0b\x0e\xfd\xb3\xc0\x1d\x62\x71\xe6\x37\x43\xce\x3f\xce\x70\x8e\x7a\x8d\x7a\x3c\x00\x38\x84\x77\xca\x58\xc9\x4a\x04\xb5\x04\x5b\x20\xbc\xad\xa5\x44\x01\xc6\x91\xc4\x70\xcd\xb8\x14\xcd\x10\x6c\xc1\x0d\x70\x03\xb5\xc1\x0c\x92\x06\x58\x6d\x55\x64\x52\x26\x50\x1b\x87\x63\x15\xa4\x4a\x2e\x79\x5e\x6b\x84\x8d\xd2\x2b\xd4\x26\x1e\x80\xc3\x7f\xcf\x4a\x1c\x83\x50\x29\x13\x85\x32\x76\xe0\x18\x66\x4a\x5b\x0f\xb7\x54\x1a\xde\x2d\x16\x33\x48\x55\x59\xd6\x92\xa7\xcc\x72\x25\x81\xc9\xcc\x69\xb4\xc1\x04\x32\x66\x8a\x44\x31\x9d\x39\xc8\xc5\x62\x46\xdc\x63\xf8\xf3\x68\x34\xda\x87\x76\x33\x9b\xee\x82\x11\xdb\xcd\x6c\xea\xb9\x7e\x18\xfd\x10\xb8\x6e\xf0\x63\xcd\x35\x42\xc2\x0c\x4f\x69\x4f\x05\x4a\xdb\xca\x27\x20\x92\xef\x8f\x02\x26\xb3\x4b\xda\x3e\x97\x39\x30\xa8\x98\x31\x1b\xe5\xd5\x39\x84\xcb\xa5\x13\x3d\x84\x92\xad\x10\x0c\x9d\x80\x55\x50\x69\x55\xa1\x16\x0d\x68\x34\x56\xf3\xd4\x02\x4b\x53\x34\x86\xd6\x08\xd7\x1f\x17\x2c\xb9\x40\x87\xf2\x0d\xc6\x79\x0c\x69\x51\xaa\x0c\xfe\x34\x1a\xc1\xd2\x59\x22\xf6\x64\x71\x53\x8a\x6f\xfd\x4e\x83\xe8\x31\xb0\x24\x3d\xfe\xe3\x77\x7e\x27\x97\x32\x15\x75\x86\xc0\xe0\x60\xca\xd2\x02\xa3\xa9\x92\x56\x2b\x31\x06\xa9\x22\x63\x95\xc6\x03\x7f\xc6\x05\xb2\x0c\x35\x70\x09\x17\x68\x8f\xae\xb8\xb1\xa4\x5f\xa5\xa4\xc1\xce\x90\x95\xc6\x35\x4a\x0b\x29\x4b\x0b\xda\x6f\xd2\x00\x97\x16\x75\x89\x19\x67\xba\x71\x27\xc2\x53\x74\xf6\x3d\xe3\x86\x25\x02\x09\xdb\x09\x1e\x83\xd5\x35\x7a\xa5\xae\x78\xc9\xad\x3f\x43\xfe\xc9\x7b\x18\x33\x2b\xc0\x5f\x30\xad\xad\xd2\x20\x54\x6e\xe0\x1b\x63\x33\x55\xdb\x23\xd4\xfa\xdb\x21\xe9\x95\x34\xd6\x43\x5f\xb3\x5f\xce\x03\xe9\x95\xca\xe7\xfc\x13\x8e\xe1\x78\x34\x1a\x8d\xe0\x10\x8e\x47\xf0\xe3\xa9\x97\xb2\x28\x10\xb6\x3d\x98\xa5\x96\xaf\x3b\xb3\x65\xcc\xb2\x84\x19\xb2\x71\xba\x42\x99\x39\x96\xc9\x9a\x71\x41\x6a\xb7\x5f\xcd\x18\x12\x25\x6c\x96\x0c\x21\x6b\x24\x2b\x55\x96\xd0\xde\x02\x6f\xbb\xb8\xf5\xc9\x50\xf4\x00\x9c\x2a\x61\xcf\x4e\xfd\xdf\xde\x38\xb6\x68\xed\xdb\x49\x0e\x16\xa6\x7f\xb4\x3e\x86\xf8\xc8\xdb\x36\xa2\x38\x89\x32\xae\xc3\xef\x38\x4b\x06\x03\x0f\x74\xe6\xb4\x68\xa1\x0f\xdd\x7f\xa7\xcc\xa0\xdb\xa8\x55\xe4\x6f\xce\x41\x5b\x6d\xc1\xd2\x76\x4c\x47\xbd\xa0\x9f\x2d\xc3\x18\x0e\xbc\x80\x83\x2d\xb4\xc9\xed\x1c\x26\xde\x21\x57\xd8\xc0\xe5\x59\xb7\xf6\x23\x36\x63\x38\x78\x4c\x3b\xc7\x54\xa3\x6d\x59\x7e\xc4\xa6\x5b\xf7\x2b\x7b\x58\x34\xe6\x5c\xc9\xee\xe3\x8d\xfb\xe9\xe8\x06\x00\x57\x2a\xcf\x7d\x12\x72\xae\xa2\xf2\x9c\x5c\x4d\xe0\x1a\x85\x19\x43\x86\x49\x9d\x93\x3b\x2c\xd5\x10\x50\x6b\xa5\x1d\xe1\x15\x2d\x8f\xdd\xe7\xc0\x78\xab\xb9\x45\xef\x49\xee\xdc\xb9\x81\x8a\xd9\x22\xa6\xa0\xc4\xb2\xb2\xcd\xd0\x2f\x32\x4a\x4b\x9a\x5b\x8b\x92\x08\x8d\xcd\x50\xeb\xd8\x81\x7c\xa8\x6d\x55\xdb\xb7\x5c\xa0\xd7\xcd\xbb\x94\x49\x0b\xcc\x6a\x41\xb1\x62\xfa\xcc\x42\x86\xbd\x98\x9e\x0f\xe1\x43\x85\xd2\x58\x96\xae\x86\x2e\x51\x5d\x33\x59\x33\x41\x49\xa7\xaa\x6d\xef\x56\x31\x0c\xe6\x2d\x4e\x9b\x6f\x37\xa0\x96\x41\x0b\x5d\x4b\x60\xdb\x92\x2c\xea\x2e\x5f\x51\x50\x83\x64\x52\x19\x4c\x15\x61\x0d\x00\x5a\xb0\x1b\x66\xdb\x60\x18\xb5\x11\x01\x9e\xae\x93\x52\x32\xd9\xb8\x80\x73\x27\xd3\x0a\xa1\x08\x53\x12\x77\x45\xb5\xb0\xd3\xa2\x96\x2b\xc2\xed\x40\x84\x92\x39\xb1\x6f\x18\xb7\x90\xa0\xdd\x20\x4a\xa8\xab\x8c\x59\x34\x90\xe0\x52\x69\x84\x92\xe9\x95\x4f\x8b\x52\x65\x08\x19\xb2\xec\x29\xfd\xdf\xab\x0c\x67\x5c\xe6\x0b\x5e\xa2\xaa\xed\x98\x12\xdd\xce\x1e\x4a\x2e\x6b\x8b\xfb\xc5\xd3\xf9\x07\x19\xce\x82\x4c\xdb\xe1\x63\x1d\xc8\x3e\x2f\xd2\xe2\x52\x72\xdb\x69\xf1\xdd\x68\x47\x8d\xef\x83\x1a\x26\xd0\xb6\x4e\xda\xbb\x54\x50\xe3\xf2\x0c\x36\x5c\x08\x48\xd0\x15\xc4\x92\x51\xf1\x10\xa2\x81\x1c\x25\x1d\x2f\x66\xde\xc3\x2e\xcf\xb6\xa2\x83\x8a\x13\xcb\x32\x4d\x61\xb4\xaf\xe8\x3a\x32\x5f\xa2\x27\x9e\x6c\xab\x76\x8e\xdb\xe2\xb5\x5d\x0a\xdc\xc9\xf4\x35\x2c\x86\x0d\xa7\x34\xb4\x5b\xbb\xe2\xc0\xe4\x91\xf7\x55\x11\x5a\xa5\x28\x30\x90\x6a\x24\xe5\x21\xab\x35\x1d\x6a\xa5\x15\x05\x3d\xfd\xd9\x6e\xb7\x0d\x25\x2e\x7d\xcc\x65\x5c\x63\x6a\x95\x6e\xbc\x98\x5b\xa5\x57\x67\x5c\xef\x49\x75\xfd\x51\x76\xa6\x2a\x18\xb9\x12\x81\x65\xc2\x67\x35\x32\x3b\x52\xb0\x31\xe9\xc8\xac\x37\xd4\x10\xb8\xf5\x1a\x98\xa2\xb6\x90\xa9\x8d\x6c\x77\x15\x1d\x43\x89\x4c\x1a\x22\xd7\x48\x21\x2b\x55\xcb\x16\xc3\xa8\x5d\xf4\x1f\x80\x97\xae\x94\x59\x14\x0d\xb0\xa5\x45\x1f\xd7\x4b\xae\x8d\x75\x11\xe3\x51\x3b\xf7\x88\x8e\xdb\xe3\x99\x38\x7f\xf0\x3a\xec\x5a\xdc\xea\x86\xbc\x32\x43\x8b\xa9\x85\x4d\xc1\x5c\x5d\x55\xb5\x4e\xd1\x27\x1f\xd6\xd5\x1c\xab\x80\xdb\x18\xda\x4c\x8f\x4b\x2e\xe9\x68\x6f\x3a\x72\xee\x77\xed\x44\xb5\x97\x2a\x7f\x27\x51\x6b\xd4\x9a\x67\x68\xfc\xa9\x27\x58\xb0\x35\x57\x21\x8f\x75\x00\x7d\x3d\x9a\xce\x7e\x32\xbd\xe4\xb8\xff\x5e\xd5\x66\x0c\xc1\x91\x9c\x53\x4e\xae\x7b\x3a\x57\x86\x2f\x4e\x7b\xf2\x1b\x56\x5e\x24\x63\x18\xc5\x5b\x1c\x67\xdc\xac\xc0\x54\x2c\xc5\x2f\x30\x12\xd1\x67\x9c\x6f\x9d\x85\x37\x91\xcb\xf5\x60\x6b\xda\x7d\xcf\xb2\x93\x20\x4d\x23\xd3\xde\x9b\x77\xef\xa6\x1d\xc7\xe7\x01\x4e\xff\x7e\x72\x49\xca\x27\xca\xef\x77\xa3\x3b\x50\xf6\xdb\x9b\x4d\x3b\xd7\x20\xf7\xf3\xf9\xed\xc8\x09\x27\xeb\xbe\x4c\x54\xe7\x2e\xfb\xd3\x72\xa0\xdd\x2e\x7b\x5f\x55\xfa\xf6\x95\xbf\x57\x2b\x81\xfb\xca\xe0\xe0\xc9\xab\xd5\xa3\x2a\x37\xd8\x7f\xa1\x72\x39\x6b\x08\x85\xa5\xed\x2b\x3d\x04\x23\x6a\x5d\x0e\xa1\x4a\xcc\x10\x72\xf2\x62\x99\x73\x89\x74\x6b\xa6\xda\x39\x84\x3c\xc5\x21\xa8\xb6\xae\x0e\x4e\x3d\x52\x00\x22\x75\xba\x52\xda\x4a\x09\xd1\xd1\x92\x3a\xc7\x7f\xb7\x98\x3a\x79\xfe\x88\x17\x58\x56\xc2\xf9\xc0\x3f\xc3\x46\x6b\xc9\xd7\xa8\x0d\xc2\x09\xac\x99\xe4\x42\xb0\xb0\x90\xa3\x45\xb9\x86\x13\x58\xd0\x05\xd6\x7f\xf3\x57\x55\xb7\xb1\x13\xb8\xbf\x8f\xcf\xbb\xdf\x0f\x0f\x81\x84\xe9\xbc\x2e\x51\x5a\x03\x27\xa1\xef\x71\x65\x3d\x8a\xc2\xfd\xfe\xfe\x3e\x9e\xba\xbf\x1e\x1e\x20\x8a\x28\xbb\x44\x3c\xa3\xaf\x0b\x66\x56\x97\x59\x87\x23\x54\xee\x65\x84\xdc\xf9\xf0\x70\xe4\x0f\x2e\x72\x97\xf1\x48\xa8\xbc\x55\x8a\x7c\xe2\x31\x6d\x48\xb3\xde\xa8\x81\x50\x39\x8b\x3e\x4d\xa9\x6a\x1b\x28\x4d\xa1\x6a\x91\xdd\x59\xcd\xa4\x59\xa2\xbe\x5b\xba\x32\x70\x02\x7f\x3f\x9f\x07\x8a\x4d\x81\xf2\xce\xaa\x9e\xa4\x03\xff\xf0\xfe\xee\xfc\x6f\x97\x8b\xbb\x0f\x37\x77\xe7\x7f\xbd\x9c\x2e\x02\xc3\xfd\x3d\x5f\x82\x44\x88\x29\xd7\xc0\x08\xa2\x6e\xa7\xf7\xf7\x95\xe6\xd2\x2e\xe1\x40\xe3\xc7\x1a\x8d\xbd\x4b\x89\xe4\x04\x7e\x9f\x1d\x78\xf2\x2d\xd2\x08\x50\x66\x5b\xbf\x03\xa8\xcb\x48\x94\x56\xbe\x88\x5b\x62\xa9\x74\x43\xc8\xf1\x68\x09\x17\xa7\x07\x81\xf1\x79\x7c\x9f\xb8\x9e\x15\x90\x51\x12\xdc\x86\xf7\x7c\x7b\xf0\xc3\x87\x8f\x35\xfa\xee\x68\x76\x3a\x7f\xca\x43\x0f\x7f\x97\x70\x79\x94\x30\x53\xb4\x1f\x66\xa7\x73\x88\xde\x93\x19\xdd\xd5\xa5\xd7\xd7\xaf\xa8\xe7\x0d\xec\x09\xf1\x79\x9f\x79\x89\xd9\x3c\x98\x70\xd5\xdb\x9c\x1c\x8f\xab\x4a\x9e\xbc\x9a\xed\x5a\xf0\x12\xcb\x13\x3a\xd7\x3c\x79\x35\xab\xb5\xd0\xe4\xdd\x3d\xf6\x73\x26\x7b\x14\xf6\xff\x66\x90\x0f\x00\x2e\x34\xcf\xce\x5d\xce\x7b\xb9\xe5\xdf\x3c\x61\xf7\x37\x2f\xb3\xfa\x9b\x17\xd9\x9c\xc8\x3a\x6b\x7e\x8d\x1f\xbc\x81\xa8\x42\x28\x2b\xfe\x7a\xa1\xeb\x75\x29\xee\xd6\xad\xfd\x2f\x5e\xcf\xfc\x01\x7a\x69\xf8\x27\xec\xb0\x7f\x1b\xf3\x03\x5d\xf7\xe6\x57\x3f\xdd\x5c\x3f\x6d\xfb\xa3\xc7\xc6\x9f\x9f\x4e\x16\xd3\x77\x10\x45\x3f\xab\x24\x72\x75\x78\x8f\x27\x74\x44\xd2\xf7\x7a\xc7\x9f\x2d\xf8\x7a\xf1\xbc\x17\x74\x0c\x21\xb5\x3f\xeb\x5e\x2f\xf2\x91\x0e\x95\x92\x7c\x54\xa1\x76\xe7\xf3\x8a\x0e\xd3\x09\x28\xb1\x74\x99\xf8\x15\xf3\x7c\x0f\x6e\xcb\xaa\x07\xff\xed\x52\xc6\xf4\x7c\xdc\x35\x7b\x7e\xb4\xc4\xd2\x54\xd5\xd2\x52\x8f\x96\xa1\xb4\x9c\x09\xe3\x12\x5a\xdc\x77\x56\x95\x32\x86\xbb\xeb\xb8\xbb\x0c\xee\x6f\x51\x33\x6e\x52\x6a\x27\xda\x1e\x75\xe2\x71\xbb\xfb\x9f\x47\xbb\x50\x2a\x17\x08\x53\xa1\xea\x8c\xfa\xc0\x9f\xa9\xbb\xb9\x3c\xfb\xb5\xc2\x66\x1e\xe9\x29\x41\x9f\x94\xfc\xd5\xfb\xf9\x87\x92\xfd\x46\x6e\x91\xe7\x85\xdd\x6a\x8c\x66\x1a\x97\xa8\x7d\xce\xa3\x3e\xd3\xfa\x79\x02\xd4\x15\x7c\xac\x79\xba\x12\x4d\x7f\xe9\x7f\xdf\x13\xb9\x46\x4e\x68\x64\x59\x03\x4a\x0a\x2e\xa9\x7b\x5d\x23\x70\xba\xb3\xca\x00\x52\x57\xae\x9b\x68\x01\xbc\xa8\xbf\x10\xea\xdc\x2f\x8f\xe1\x38\x1e\x85\xed\x6d\x4f\x39\x52\x96\xfa\xf1\x12\x35\xec\xd4\x3e\xd6\xc2\x1a\xf8\xa6\x74\x63\x58\x04\xc1\x8d\x1d\x82\x0d\x99\xc3\x0c\x01\x6d\xfa\x6d\x80\x09\x63\x10\x8d\x4b\x8d\xa6\xe8\x1a\x2a\x37\x92\x5d\x2c\xae\x9e\x1c\xb4\x0c\x6e\x9d\x6b\xfa\xd9\xd4\xab\x34\xff\x5f\x68\xfd\x3b\x21\x60\xac\xd2\x2c\x47\x30\x8d\xb1\x58\xba\x46\xea\x30\x0c\x73\xdb\x77\x86\xda\x59\xd5\xa0\x7d\x64\xe7\xa4\xf1\x93\x8d\xb6\x17\x18\x42\x52\x5b\x68\x54\x0d\x25\xd9\x18\x24\x62\xe6\xd4\x72\x78\x7c\x49\x4b\x7f\xd0\xe8\x4d\xe3\xc7\x54\x21\x1c\x7d\xdf\xe1\x2d\x3d\xf7\x0a\xf5\xe3\xc8\x94\xf9\x7b\x42\x50\xd1\x1f\xa7\xfb\xdc\x3b\xd1\x6d\xc1\x2d\x92\x55\xa8\x3f\x72\x5d\x4a\x7f\x14\xae\x05\x33\xb0\x29\x78\x5a\xb4\x8d\x2b\x37\xc0\x84\x50\x1b\x52\x50\x85\xd7\x80\xd6\x4b\x26\x7e\xe1\x8c\x6b\xd3\x37\x87\x11\xc4\x47\xed\xe8\x61\xfe\x5d\x3f\x0b\x8e\xe0\x5c\x66\x95\xe2\xd2\xf6\xdf\xfc\xdc\x76\xeb\x67\x18\xcb\xb6\xfc\x17\xf3\xed\xc5\x97\xe7\x13\xff\xff\x0f\x95\xe5\x4a\x32\x11\xef\xc6\xe3\x36\xd3\x73\x61\xb9\x83\xb7\xd4\xaa\x74\x76\x44\xb9\xe6\x5a\x49\xea\xa0\xe2\xad\xfd\x6d\x67\xa3\x1d\xc6\xc9\x5e\xf4\x5d\xed\xbf\x88\x0d\xf0\x56\xab\xf2\x5c\xae\xb7\x5f\x2a\x9e\x1a\x49\x3c\x1e\x47\x50\xcf\xe9\xca\x17\x75\x6b\xed\x18\x34\xcc\xfb\x3f\x9b\x4e\xec\x1b\x17\xbc\x60\x2a\xe1\x38\xaf\xd9\x2f\xbc\xac\xcb\x5e\xd8\xf6\x33\x89\x7b\x4e\xe9\xdf\x4a\x48\x81\xa4\x5e\x52\x3e\x7b\x34\xa1\x25\x89\xa7\x6e\xe5\x6b\x5f\x4f\x9c\x5c\xed\xdf\x8a\x42\xb3\x1d\x3f\xf9\x74\xa2\xab\x74\xe7\xdd\x84\x2a\xd8\x8d\xe3\x75\x6b\x3b\x5f\x4c\x6b\x4e\xfa\x19\xc6\xd3\xce\x60\x3e\x59\xb8\x09\x17\xe4\x37\xb3\xa9\x1f\x69\xcd\xa6\xe3\x9d\x31\xcd\xb3\x83\xd3\x2f\x8c\x4e\xdd\x6b\x61\x57\x00\xbe\x7e\x76\xfa\xfc\xf4\x74\x6b\x4f\x3e\x11\x88\xc6\x6f\xae\x7d\xce\x89\xa1\x1b\xfb\xfd\xcf\x3f\xf0\x3c\xe9\x5b\xfe\x71\x91\x8a\xc8\x8b\x9d\x4b\xa8\x7c\xc7\xc3\x26\x0e\xe9\x9c\x80\xdc\xf4\x4b\x9b\xb1\x3f\xd8\x28\x38\x1c\xfd\xe5\x27\x26\x3b\x44\x41\x55\x3f\x31\x73\x7a\x98\xff\x52\x8f\xdb\x56\xf1\xb7\x99\x5d\x1e\xc2\x42\x33\x69\xe8\xee\xd1\x1e\x1d\x93\x99\x2b\xe5\x48\xea\x73\x0d\xa9\x92\x96\x6c\xcb\x65\xef\xf1\xff\x37\xfe\xfe\x1f\xfb\xa0\xf9\xaf\x00\x00\x00\xff\xff\xf9\x68\xa8\x2d\xef\x21\x00\x00")

func configDefaultConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_configDefaultConfigYaml,
		"config/default-config.yaml",
	)
}

func configDefaultConfigYaml() (*asset, error) {
	bytes, err := configDefaultConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/default-config.yaml", size: 8687, mode: os.FileMode(420), modTime: time.Unix(1506638476, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGridengineTemplateTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xcd\x4a\x03\x31\x14\x85\xf7\xf3\x14\xd7\xa9\x5d\x26\x33\x2f\xe0\xaa\x95\xe2\xc6\x85\x08\x2e\x25\xd3\xdc\xd8\x4b\x66\x92\x21\x3f\x2a\x86\xfb\xee\x32\x69\x11\x0a\x63\x77\x87\xc3\x77\xbe\xc5\xd9\xdc\x75\x03\xb9\x6e\x50\xf1\xd4\x6c\xee\x41\x3c\x43\x29\xf2\x55\x45\xfb\xa4\x99\x6b\xe3\x97\xe6\xcd\x07\xbb\xa7\xc0\xdc\x99\xec\x1c\x8e\x22\x26\xed\x73\xaa\x00\xfe\x07\x60\x08\x4d\x29\x64\xc0\x21\xc8\xdd\x9c\x23\xf4\x20\x98\x9b\x52\xe6\x40\x2e\x19\x68\x97\xf9\x8c\x30\xcd\x04\x5b\xdd\x9e\xa1\x0a\x08\x40\xa7\x6b\xba\xcc\x5f\xd4\x74\x18\xa0\x97\x6b\x86\x11\x4e\xef\x9f\x13\x4e\x0f\x5b\xd9\x9b\x43\x7b\x81\xd7\x3d\x7b\x8a\xf6\xa6\xc8\x44\xfa\xc1\x3f\xd3\x19\xbf\x52\x35\xa5\xc8\xc7\x6f\x3c\xe6\xa4\x86\x11\x99\xe1\xcb\x07\x8b\x01\x42\x76\x20\xc4\xd1\x3b\x43\x1f\xcb\x23\xbb\x9a\x98\x41\x88\xa4\xa2\x15\xa4\xaf\xae\xfd\x0d\x00\x00\xff\xff\x24\x76\x8a\x2a\x79\x01\x00\x00")

func configGridengineTemplateTxtBytes() ([]byte, error) {
	return bindataRead(
		_configGridengineTemplateTxt,
		"config/gridengine-template.txt",
	)
}

func configGridengineTemplateTxt() (*asset, error) {
	bytes, err := configGridengineTemplateTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gridengine-template.txt", size: 377, mode: os.FileMode(420), modTime: time.Unix(1505922697, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configHtcondorTemplateTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xce\xcd\x6e\xe2\x30\x10\xc0\xf1\xbb\x9f\x62\x84\xb4\x47\x67\x79\x81\x5c\x16\x22\xc4\x65\x91\x68\xd4\x8f\x53\x64\xe2\x49\xb0\xe2\x8c\x61\x6c\x87\x56\x91\xdf\xbd\x0a\x20\x2a\xaa\xd2\xdb\x24\xf3\x9f\x9f\x1c\xc9\x0c\xc8\x1e\x21\x87\x41\x91\xb1\x56\x89\x16\x03\xd2\x00\x39\x94\x1c\x51\xe0\x3b\xd6\x31\xa8\x9d\x9d\x92\x71\xcc\x8a\xdb\x77\x4a\x42\x71\x1b\x7b\xa4\xe0\x21\x87\x93\xe3\x0e\x19\x38\x12\x48\x59\x3b\x6a\x4c\x3b\xf5\x8b\xf3\x94\x12\x48\x19\x94\xef\xa4\xd1\xd3\xdf\x52\xf9\x6e\xad\x53\x12\xd6\xb5\x17\xf7\xc5\x71\xb7\x34\x9c\xd2\xdf\xda\x91\x76\x2c\x71\x40\x0a\xd2\xba\x56\x20\xb3\xe3\xef\x55\x13\x89\xd0\x4a\x1f\x34\x32\x0b\x17\xc3\x21\x86\xc7\x8d\x8b\x41\xf8\xbd\x8b\x56\x57\x81\x15\xf9\x06\xb9\x6a\x8c\xc5\xe9\xe1\x6f\xc5\x93\x38\xed\x91\xaa\xe0\xbe\x96\x37\x70\xf3\xbf\x2a\x5e\xd7\x65\xb5\xd9\x56\xc5\xf3\x7a\x51\x8a\x71\x34\x0d\x10\x42\xb6\x38\x44\x0f\x73\x90\x29\x89\x71\x3c\xb0\xa1\xd0\xc0\x8c\xf1\x18\xd1\x87\xaa\x9e\x96\x39\xfc\xd1\xb3\x4b\x78\x8e\x24\x20\xe9\xf3\x74\x25\xb6\xaa\x5f\xed\x60\x9e\x3d\x52\x7a\xec\x1d\x7f\x4c\x4e\x36\x6f\x60\xf5\x6f\x76\x3d\xf9\x59\x5b\x1a\xdf\xfd\xca\x69\xe3\xbb\x3b\xec\x72\x71\xa7\x89\x63\xc4\x88\xe2\x33\x00\x00\xff\xff\x5f\xc0\xd7\xf8\x18\x02\x00\x00")

func configHtcondorTemplateTxtBytes() ([]byte, error) {
	return bindataRead(
		_configHtcondorTemplateTxt,
		"config/htcondor-template.txt",
	)
}

func configHtcondorTemplateTxt() (*asset, error) {
	bytes, err := configHtcondorTemplateTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/htcondor-template.txt", size: 536, mode: os.FileMode(420), modTime: time.Unix(1505922697, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configPbsTemplateTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xd0\xcf\x4a\x03\x31\x10\xc7\xf1\xfb\x3e\xc5\xd8\xd2\x63\xb2\xeb\x55\xd8\x8b\xad\x88\x17\x11\x15\x3c\x27\xcd\xa4\x86\xec\x4e\x96\xfc\x41\x21\xcc\xbb\x4b\x77\x17\xa4\x60\xbd\x0d\xc3\x97\xcf\xe1\xb7\xbd\x69\xb5\xa3\x56\xab\xf4\xd9\x6c\x5f\xee\xdf\x40\x3c\x43\xad\xf2\x5d\x25\xff\x64\x98\xd7\x5f\x38\xff\x3e\x42\xf4\x07\x17\x99\x5b\x5b\x88\x70\x10\x29\x9b\x50\xf2\x9a\xe0\xb5\x04\x63\x6c\x6a\x75\x16\x08\x41\xee\xa7\x92\xa0\x03\xc1\xdc\xd4\x3a\x45\x47\xd9\xc2\x66\x01\x06\xa0\x60\x30\xf5\xb7\x77\xd3\x44\xfd\xce\x6c\x96\x7a\x2e\x05\x20\x99\xf9\x5a\x9d\x57\x35\x3e\x6a\xe8\xe4\x35\x6a\xc4\xb1\xdf\xc9\xce\x9e\xf4\x66\x8d\xff\x76\x0e\x2e\xf9\x7f\x21\xeb\x06\xfc\x95\x96\xfc\x82\x6a\x6a\x95\x0f\xdf\x78\x2c\x59\xe9\x01\x99\xe1\x2b\x44\x8f\x11\x62\x21\x10\xe2\x18\xc8\xba\xd3\x79\x9a\xfd\x7c\x31\x83\x10\x59\x25\x2f\x9c\xb9\xd8\xf9\x27\x00\x00\xff\xff\x80\x45\x84\x4d\x88\x01\x00\x00")

func configPbsTemplateTxtBytes() ([]byte, error) {
	return bindataRead(
		_configPbsTemplateTxt,
		"config/pbs-template.txt",
	)
}

func configPbsTemplateTxt() (*asset, error) {
	bytes, err := configPbsTemplateTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/pbs-template.txt", size: 392, mode: os.FileMode(420), modTime: time.Unix(1505922697, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configSlurmTemplateTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\xcf\x6a\xc3\x30\x0c\x87\xef\x79\x0a\x2d\xa5\x47\x27\xd9\x23\xac\xe9\xe8\x76\xdd\x0a\x3b\x3b\x8d\xb2\x79\x6e\x64\x23\xdb\x6c\x60\xf4\xee\x23\x69\xa1\x29\xac\xf4\xf6\xc3\xfe\xf4\xa1\x3f\xab\x87\xba\x33\x54\x77\x3a\x7c\x15\xab\xf7\xcd\xd3\xbe\x7d\x01\xa5\xbe\x5d\xa7\x48\x8f\x08\x39\x57\x7b\x1d\xec\x6b\x2f\xb2\xf8\xa6\xa8\x83\x0d\xf0\xb8\x78\x42\x66\xc7\x13\xfe\xe1\xd8\x6e\x0d\x8b\xd4\x43\x22\xc2\xa3\x0a\xb1\x47\xe6\x05\xea\x52\xf4\x29\xde\x62\x5d\x8a\x45\xce\x66\x00\x42\xa8\x5a\x9f\x02\x34\xa0\x44\x8a\x9c\x3d\x1b\x8a\x03\x94\x17\xd3\xc1\xa7\xa0\x3c\xb2\x9a\xfa\x81\x75\x5f\x9e\x2a\x66\x5a\x01\x52\x3f\xa7\xb3\xeb\x4d\x8f\xbb\x0e\x9a\xea\xb6\x6e\xc4\x11\xd6\x55\x33\xec\x36\xe5\x19\xff\xdf\xb4\x35\xc1\xde\x51\xc5\xd1\x5f\x54\x27\xfe\xca\x55\xe4\x5c\x3d\xff\xe2\x21\x45\xdd\x1d\x51\x04\x7e\x1c\x5b\x64\xe0\x44\xd3\x5c\x8e\x06\xf3\x39\x6d\xa8\x9d\x93\xc8\xa4\xd4\xc1\x2a\xd3\x5f\x9d\xe4\x2f\x00\x00\xff\xff\x36\x8f\xbd\x5a\xbe\x01\x00\x00")

func configSlurmTemplateTxtBytes() ([]byte, error) {
	return bindataRead(
		_configSlurmTemplateTxt,
		"config/slurm-template.txt",
	)
}

func configSlurmTemplateTxt() (*asset, error) {
	bytes, err := configSlurmTemplateTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/slurm-template.txt", size: 446, mode: os.FileMode(420), modTime: time.Unix(1505922697, 0)}
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
	"config/default-config.yaml":     configDefaultConfigYaml,
	"config/gridengine-template.txt": configGridengineTemplateTxt,
	"config/htcondor-template.txt":   configHtcondorTemplateTxt,
	"config/pbs-template.txt":        configPbsTemplateTxt,
	"config/slurm-template.txt":      configSlurmTemplateTxt,
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
	"config": {nil, map[string]*bintree{
		"default-config.yaml":     {configDefaultConfigYaml, map[string]*bintree{}},
		"gridengine-template.txt": {configGridengineTemplateTxt, map[string]*bintree{}},
		"htcondor-template.txt":   {configHtcondorTemplateTxt, map[string]*bintree{}},
		"pbs-template.txt":        {configPbsTemplateTxt, map[string]*bintree{}},
		"slurm-template.txt":      {configSlurmTemplateTxt, map[string]*bintree{}},
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
