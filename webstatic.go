package main

import (
	"log"
	"os"

	"github.com/otiai10/copy"
	"./source/filesystem"
)


func writeIndex() {
	var b bytes.Buffer
	Filesystem.
	b.WriteString(getLayoutStart(getSiteTitle()))
	b.Write(blackfriday.MarkdownCommon(getFile("_sections/header.md")))
	writePostsSection(&b)
	writePagesSection(&b)
	b.WriteString(getLayoutEnd())
	writeFile("index", b)
}

func createFilesAndDirs() {
	_ = os.RemoveAll("server")
	os.MkdirAll("server", 0755)
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	pwd += "/server/"
	_ = copy.Copy("./template/static", pwd)
	os.MkdirAll("_sections", 0755)
	os.MkdirAll("_sections/_posts", 0755)
	os.MkdirAll("_sections/_pages", 0755)
	os.MkdirAll("/server/posts", 0755)
	os.MkdirAll("/server/pages", 0755)
}

func main() {
	createFilesAndDirs()
}
