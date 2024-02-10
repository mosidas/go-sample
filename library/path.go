package library

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
)

func Path() {
	fmt.Println("=== path ===")
	pt := "../library/path.go"
	fmt.Println(path.Base("../../go-sample/library/path.go"))
	fmt.Println(path.Clean(pt))
	fmt.Println(path.Dir(pt))
	fmt.Println(path.Ext(pt))
	fmt.Println(path.IsAbs(pt))
	fmt.Println(path.Join("go-sample", "library", "path.go"))
	fmt.Println(path.Split(pt))
	fmt.Println(path.Match("*.go", pt))
}

func Filepath() {
	fmt.Println("=== filepath ===")
	pt := "../library/path.go"
	fmt.Println(pt)
	fmt.Println("Base:", filepath.Base("../../go-sample/library/path.go"))
	fmt.Println("Clean:", filepath.Clean(pt))
	fmt.Println("Dir:", filepath.Dir(pt))
	fmt.Println("Ext:", filepath.Ext(pt))
	fmt.Println("IsAbs:", pt, filepath.IsAbs(pt))
	fmt.Println("Join:", filepath.Join("go-sample", "library", "path.go"))
	dir, fl := filepath.Split(pt)
	fmt.Println("Split:", dir, fl)
	match, _ := filepath.Match(".go", pt)
	fmt.Println("Match:", ".go", match)
	abs, _ := filepath.Abs(pt)
	fmt.Println("Abs:", abs)
	eval, _ := filepath.EvalSymlinks(pt)
	fmt.Println("EvalSymlinks:", eval)
	fmt.Println("FromSlash:", filepath.FromSlash(pt))
	glob, _ := filepath.Glob("*.go")
	fmt.Println("Glog:", "*.go", glob)
	rel, _ := filepath.Rel("../", pt)
	fmt.Println("Rel:", rel)
	for _, p := range filepath.SplitList("a;b;c") {
		fmt.Println("SplitList:", p)
	}
	fmt.Println("ToSlash:", filepath.ToSlash(filepath.FromSlash(pt)))
	winAbsPath := "C:/go-sample/library/path.go"
	fmt.Println("VolumeName:", winAbsPath, filepath.VolumeName(winAbsPath))
	unixAbsPath := "/go-sample/library/path.go"
	fmt.Println("VolumeName:", unixAbsPath, filepath.VolumeName(unixAbsPath))

	wd, _ := filepath.Abs(".")
	filepath.WalkDir(wd, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// if *.go file is found, print the path
		if filepath.Ext(path) == ".go" {
			fmt.Println("WalkDir:", path)
		}
		return nil
	})
}
