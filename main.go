package main

import (
	"fmt"
	"go/ast"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hori-ryota/go-genconstructor/genconstructor"
)

func main() {
	if err := Main(os.Args); err != nil {
		log.Print(err)
		fmt.Printf(`
Useage: %s [targetDir]
`, os.Args[0])
	}
}

func Main(args []string) error {
	targetDir := "."
	if len(args) > 1 {
		targetDir = args[1]
	}

	if err := genconstructor.Run(
		targetDir,
		func(pkg *ast.Package) io.Writer {
			dstFileName := fmt.Sprintf("%s_constructor_gen.go", pkg.Name)
			dstFilePath := filepath.Join(filepath.FromSlash(targetDir), dstFileName)
			f, err := os.Create(dstFilePath)
			if err != nil {
				panic(err)
			}
			return f
		},
		genconstructor.WithFileFilter(
			func(finfo os.FileInfo) bool {
				return !strings.HasSuffix(finfo.Name(), "_test.go")
			},
		),
	); err != nil {
		return err
	}
	return nil
}
