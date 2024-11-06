package jdb

import (
	"os"
	"time"
	"btree"
	"strings"
	"path/filepath"
)

// holds metadata and binary data
type FileData struct {
	Name string
	Path string
	Size int64
	Hash string
	Modified bool
	Mod_time time.Time
	Data []byte // store data as byte slice
}

type Directory struct {
	Name string
	Path string
	Size int64
	Files []*FileData
	Subdirs []*Directory
}

type Database struct {
	Name string
	Size int64
	Root Directory
}

func calcHash(data []byte) string {
	// simple md5 example, use sha256 for serious hashing
	return fmt.Sprintf("%x", md5.Sum(data))
}

func ReadFile(path string) (*FileData, error) {
	// read data into byte slice
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Error: reading file %s: %v", path, err)
	}
	// calculate file size
	size := int64(len(data))
	// calculate simple hash (md5)
	hash := calcHash(data)
	// handle metadata
	info, err := os.Stat(path)
	modtime := info.ModTime()
	filename := filepath.Base(path)
	name := strings.TrimSuffix(filename, filepath.Ext(filename))

	// return FileData object
	return &FileData{
		Name: name,
		Path: path,
		Size: size,
		Hash: Hash,
		Modified: true,
		Mod_time: modtime,
		Data: data,

	}, nil
}

// store file in memory
func storeInDir(dir *Directory, filedata *FileData) {
	dir.Files = append(dir.Files, *filedata)
}

func storeInBin(filename string, filedata *FileData) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Error: creating file %s: %v", filename, err)
	}
	defer file.Close()

	// write metadata (path, size, hash, data) to bin file
	_, err = file.Write([]byte(filedata.Path)) // write path as str
	if err != nil {
		return fmt.Errorf("Error: writing file path %s: %v", filename, err)
	}
	_, err = file.Write([]byte(fmt.Sprintf("%d", filedata.Size)))
	if err != nil {
		return fmt.Errorf("Error: writing file size %s: %v", filename, err)
	}
	_, err = file.Write([]byte(filedata.Hash))
	if err != nil {
		return fmt.Errorf("Error: writing file hash %s: %v", filename, err)
	}
	_, err = file.Write(filedata.Data)
	if err != nil {
		return fmt.Errorf("Error: writing file data %s: %v", filename, err)
	}
	
	return nil
}

func readBinFile(filename string) (*FileData, error) {
	file, err := os.Open(filname)
	if err != nil {
		return nil, fmt.Errorf("Error: opening bin file %s: %v", filename, err)
	}
	defer file.Close()
	var path string
	var size int64
	var hash string
	
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Error: reading bin file: %v", err)
	}

	return &FileData{
		Path: path,
		Size: size,
		Hash: hash,
		Data: data,
	}, nil
}
