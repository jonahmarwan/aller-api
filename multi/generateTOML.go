package multi

import (
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

type Method struct {
	Use  string `toml:"use"`
	Path string `toml:"path"`
}

type Build struct {
	Build   string `toml:"build"`
	Version string `toml:"version"`
}

type Service struct {
	Name   string `toml:"name"`
	path   string
	Method Method `toml:"method"`
	Build  Build  `toml:"build"`
}
type Config struct {
	Service Service `toml:"service"`
}

//Fix syntax in future

func GenerateTOMLfromMeta(Meta []string, path string) error {
	broken := []string{}
	var cfg Config
	cfg.Service.path = path
	pathElem := strings.Split(path, "\\")
	cfg.Service.Name = pathElem[len(pathElem)-2]
	for _, line := range Meta {
		broken = strings.Split(line, " ")

		switch broken[0] {
		case "use":
			method := broken[1]
			path := broken[2]
			cfg.Service.Method.Use = method
			cfg.Service.Method.Path = path
		case "build":
			lang := broken[1]
			vers := broken[2]
			cfg.Service.Build.Build = lang
			cfg.Service.Build.Version = vers

		default:
			continue
		}
	}
	file, err := os.Create(strings.Join(pathElem[:len(pathElem)-1], "\\") + "\\" + cfg.Service.Name + ".toml")
	if err != nil {
		return err
	}
	defer file.Close()

	err = toml.NewEncoder(file).Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}
