package multi

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

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


func NewPolyglot() Polyglot {
	var p Polyglot
	p.PathTree = radix.New()
	return p
}