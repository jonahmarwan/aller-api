package multi

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/armon/go-radix"
)

type Polyglot struct {
	PathTree *radix.Tree
}

func (p *Polyglot) ReadServices() {
	_, err := os.Stat("../services/")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("no exist")
			return
		}
		fmt.Println(err)
	}
	err = filepath.WalkDir("../services", func(path string, d fs.DirEntry, err error) error {
		p.PathTree.Insert(path, d.IsDir())
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func (p *Polyglot) InitPolyglot() {
	p.PathTree = radix.New()
}
