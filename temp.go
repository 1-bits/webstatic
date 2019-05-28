package main

/*

import (
	"bytes"
	"io/ioutil"
	"math"
	"os"
	"strings"
	"time"

	"./source"
	"github.com/russross/blackfriday"
)

func getLayoutStart(title string) string {
	return `<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<link href="https://fonts.googleapis.com/css?family=IBM+Plex+Sans:300,400,400i,500" rel="stylesheet">
			<link href="https://fonts.googleapis.com/css?family=IBM+Plex+Mono:400" rel="stylesheet">
			<title>` + title + `</title>
			<style>
				body {
					font-family: 'IBM Plex Sans', sans-serif;
					line-height: 1.875;
					font-weight: 400;
					font-size: 16px;
				}

				body.light {
					background-color: #fdffff;
					color: #0f3763;
				}

				body.dark {
					background-color: #141c2b;
					color: #d9dce4;
				}

				h1, h2, h3 {
					font-weight: 300;
				}

				h1 {
					font-size: 26.79296875px;
					margin-top: 24px;
				}

				h2 {
					font-size: 22.5625px;
					margin-top: 40px;
				}

				h3 {
					font-size: 19px;
					margin-top: 40px;
				}

				body.light h1,
				body.light h2,
				body.light h3 {
					color: #003880;
				}

				body.dark h1,
				body.dark h2,
				body.dark h3 {
					color: #c8ddff;
				}

				@media (max-width: 1023.98px) {
					.container {
						margin: 28px auto 48px auto;
					}
				}

				@media (min-width: 1024px) {
					.container {
						margin: 48px auto;
					}
				}

				.container {
					max-width: 896px;
					padding: 0 16px;
				}

				nav ul {
					list-style-type: none;
					padding: 0;
				}

				nav li {
					margin-bottom: 8px;
				}

				nav li .date {
					display: inline-block;
					width: 104px;
				}

				.all-posts {
					font-size: 13.47368421052632px;
					margin: 16px 0 32px 0;
				}

				a {
					text-decoration: none;
				}

				a:hover {
					text-decoration: underline;
				}

				body.light a {
					color: #006aee;
				}

				body.dark a {
					color: #7da7ef;
				}

				.toggle-theme {
					font-size: 0.889rem;
					margin-top: 40px;
					padding-top: 24px;
				}

				body.light .toggle-theme {
					border-top: 1px solid #dbe0ed;
				}

				body.dark .toggle-theme {
					border-top: 1px solid #2b3241;
				}

				pre {
					overflow: auto;
					padding: 0.25rem 0.75rem;
					margin-bottom: 32px;
				}

				body.light pre {
					background-color: #f4f7ff;
				}

				body.dark pre {
					background-color: #1a2231;
				}

				code {
					font-size: 0.875em;
					font-family: 'IBM Plex Mono', monospace;
				}

				table {
					border-collapse: collapse;
					width: 100%;
					margin-bottom: 32px;
				}

				@media (max-width: 1023.98px) {
					table {
						font-size: 14px;
					}
				}

				@media (min-width: 1024px) {
					table {
						font-size: 15px;
					}
				}

				body.light tr {
					border-bottom: 0.5px solid #bdc5d8;
				}

				body.dark tr {
					border-bottom: 0.5px solid #424957;
				}

				th {
					text-align: left;
					font-weight: 500;
					padding: 12px;
					white-space: nowrap;
				}

				td {
					padding: 12px;
					white-space: nowrap;
				}
			</style>
			<script>
				document.addEventListener('DOMContentLoaded', function(event) {
					if (localStorage.getItem('theme') === 'dark') {
						setDarkTheme();
					} else {
						setLightTheme();
					}
				});

				function toggleTheme(event) {
					event.preventDefault();

					if (document.body.className === 'dark') {
						setLightTheme();
					} else {
						setDarkTheme();
					}
				}

				function setLightTheme() {
						document.body.className = 'light';
						document.getElementsByClassName('toggle-theme')[0].children[0].innerHTML = 'Dark';
						localStorage.setItem('theme', 'light');
				}

				function setDarkTheme() {
						document.body.className = 'dark';
						document.getElementsByClassName('toggle-theme')[0].children[0].innerHTML = 'Light';
						localStorage.setItem('theme', 'dark');
				}
			</script>
		</head>
		<body>
			<div class="container">`
}

func getLayoutEnd() string {
	return `
					<p class="toggle-theme">
						<a href="#" onclick="toggleTheme(event)">Dark</a>
					</p>
			</div>
		</body>
	</html>`
}

func getSiteTitle() string {
	return strings.Split(string(source.getFile("_sections/header.md")), "\n")[0][2:]
}

func getPostMeta(fi os.FileInfo) (string, string, string) {
	id := fi.Name()[:len(fi.Name())-3]
	date := fi.Name()[0:10]
	title := strings.Split(string(filesystem.getFile("_posts/"+fi.Name())), "\n")[0][2:]

	return id, date, title
}

func getPageMeta(fi os.FileInfo) (string, string) {
	id := fi.Name()[:len(fi.Name())-3]
	title := strings.Split(string(filesystem.getFile("_pages/"+fi.Name())), "\n")[0][2:]

	return id, title
}

func writeIndex() {
	var b bytes.Buffer
	b.WriteString(getLayoutStart(getSiteTitle()))
	b.Write(blackfriday.MarkdownCommon(getFile("_sections/header.md")))
	writePostsSection(&b)
	writePagesSection(&b)
	b.WriteString(getLayoutEnd())
	writeFile("index", b)
}

func writePostsSection(b *bytes.Buffer) {
	b.WriteString("<h2>Posts</h2><nav class=\"posts\"><ul>")

	posts := getDir("_posts")
	limit := int(math.Max(float64(len(posts))-5, 0))

	for i := len(posts) - 1; i >= limit; i-- {
		fileName, date, title := getPostMeta(posts[i])

		b.WriteString("<li><span class=\"date\">" + date +
			"</span><a href=\"posts/" +
			fileName + ".html\">" +
			title + "</a></li>\n")
	}

	b.WriteString("</ul></nav><p class=\"all-posts\"><a href=\"all-posts.html\">All posts</a></p>")
}

func writePagesSection(b *bytes.Buffer) {
	b.WriteString("<h2>Pages</h2><nav class=\"pages\"><ul>")

	pages := getDir("_pages")

	for i := 0; i < len(pages); i++ {
		id, title := getPageMeta(pages[i])

		b.WriteString("<li><a href=\"pages/" +
			id + ".html\">" +
			title + "</a></li>\n")
	}

	b.WriteString("</ul></nav>")
}

func writePosts() {
	posts := getDir("_posts")

	for i := 0; i < len(posts); i++ {
		id, date, title := getPostMeta(posts[i])

		var b bytes.Buffer
		b.WriteString(getLayoutStart(title + " – " + getSiteTitle()))
		b.WriteString("<p><a href=\"../index.html\">←</a></p>")
		b.WriteString("<p class=\"date\">" + date + "</p>")
		b.Write(blackfriday.MarkdownCommon(getFile("_posts/" + posts[i].Name())))
		b.WriteString("<p><a href=\"../index.html\">←</a></p>")
		b.WriteString(getLayoutEnd())

		writeFile("posts/"+id, b)
	}
}

func writePostsPage() {
	posts := getDir("_posts")
	var b bytes.Buffer

	b.WriteString(getLayoutStart("All posts – " + getSiteTitle()))
	b.WriteString("<p><a href=\"index.html\">←</a></p>")
	b.WriteString("<h1>All posts</h1>")
	b.WriteString("<nav class=\"posts\"><ul>")

	for i := len(posts) - 1; i >= 0; i-- {
		id, date, title := getPostMeta(posts[i])

		b.WriteString("<li><span class=\"date\">" + date +
			"</span><a href=\"posts/" +
			id + ".html\">" +
			title + "</a></li>\n")
	}

	b.WriteString("</ul></nav><p><a href=\"index.html\">←</a></p>")
	b.WriteString(getLayoutEnd())
	writeFile("all-posts", b)
}

func writePages() {
	pages := getDir("_pages")

	for i := 0; i < len(pages); i++ {
		fileName, title := getPageMeta(pages[i])

		var b bytes.Buffer
		b.WriteString(getLayoutStart(title + " – " + getSiteTitle()))
		b.WriteString("<p><a href=\"../index.html\">←</a></p>")
		b.Write(blackfriday.MarkdownCommon(getFile("_pages/" + pages[i].Name())))
		b.WriteString("<p><a href=\"../index.html\">←</a></p>")
		b.WriteString(getLayoutEnd())

		writeFile("pages/"+fileName, b)
	}
}

func createFilesAndDirs() {
	_= os.RemoveAll("server")
	os.MkdirAll("server", 0755)
	os.MkdirAll("_sections", 0755)
	os.MkdirAll("_posts", 0755)
	os.MkdirAll("_pages", 0755)
	os.MkdirAll("/server/posts", 0755)
	os.MkdirAll("/server/pages", 0755)
}

func main() {
	createFilesAndDirs()
	//writeIndex()
	//writePosts()
	//writePostsPage()
	//writePages()
}
*/
