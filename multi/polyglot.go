package multi

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/armon/go-radix"
)

type Polyglot struct {
	PathTree *radix.Tree
}

func (p *Polyglot) ReadServices() error {
	path := os.Getenv("API_SERVICES")
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("Directory does not exist.")
		}
		return errors.New(err.Error())
	}
	err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		p.PathTree.Insert(path, d.IsDir())
		return nil
	})
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}


func NewPolyglot() (Polyglot, error) {
	var p Polyglot
	p.PathTree = radix.New()
	return p, nil
}

func (p *Polyglot) ExecServices() error {
	var e error = nil
	p.PathTree.Walk(func(s string, v interface{}) bool {
		if v == false {
			pathElem := strings.Split(s, "\\")
			if strings.Contains(pathElem[len(pathElem)-1], "_") {
				meta, err := ReadMeta(s)
				if err != nil {
					e = err
					return true
				}
				err = GenerateTOMLfromMeta(meta, s)
				if err != nil {
					e = err
					return true
				}

			}
		}
		return false
	})
	return e
}