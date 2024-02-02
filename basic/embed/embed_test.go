package embed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed logo.png
var image []byte

// file binary seperti image, video, music
func TestByte(t *testing.T) {
	err := os.WriteFile("logo_baru.png", image, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed sample/a.txt
//go:embed sample/b.txt
var files embed.FS

// embed multiple files
func TestMultiFile(t *testing.T) {
	fileA, err := files.ReadFile("sample/a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(fileA)

	fileB, err := files.ReadFile("sample/b.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fileB))
}

//go:embed sample/*.txt
var path embed.FS

func TestPathMatch(t *testing.T) {
	dir, err := path.ReadDir("sample")
	if err != nil {
		panic(err)
	}
	for _, file := range dir {
		file.Type()
		if !file.IsDir() {
			fmt.Println(file.Name())
			b, _ := path.ReadFile("sample/" + file.Name())
			fmt.Println(string(b))
		}
	}
}
