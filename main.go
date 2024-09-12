package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"


	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
)

type Post struct {
	Id string
	Title string
	Description string

}

type Blog struct {
	Posts []Post

}	

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))

	about_me := elem.Div(attrs.Props{
		attrs.ID: ("about_me"),
		attrs.Class: ("container"),
	},
		elem.Raw(markdownToHTML("./posts/about_me.md")),
	)
	
	tmpl.Execute(w, template.HTML(about_me.Render()))

}

func blog(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/blog.html"))
	blogs := readPostsDit("./posts")
	
	blog := make(map[string][]Post)
	
	for _, post := range blogs {
		blog["Blog"] = append(blog["Blog"], post)
	}


	tmpl.Execute(w, blog)

}

func blogpost(w http.ResponseWriter, r *http.Request) {


	tmpl := template.Must(template.ParseFiles("./templates/blogpost.html"))

	blogpost := elem.Div(attrs.Props{
		attrs.ID: ("blogpost"),
		attrs.Class: ("container"),
	},
		elem.Raw(markdownToHTML("./" + r.URL.Path + ".md")),
	)
	

	tmpl.Execute(w, template.HTML(blogpost.Render()))

}

func markdownToHTML(filename string) string {  
    var buf bytes.Buffer  
    md := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("dracula"),
			),
			extension.Table,
		),

	)  
    if err := md.Convert(readMarkdownFile(filename), &buf); err != nil {  
       log.Fatal(err)  
    }    
	return buf.String()  
}

func readMarkdownFile(filename string) []byte {  
	markdown, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer markdown.Close()

	content, err := ioutil.ReadAll(markdown)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

func readPostsDit(dir string) []Post {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var posts []Post
	
	for _, file := range files {
		posts = append(posts, Post{Id: strings.Split(file.Name(), ".")[0], Title: strings.Split(file.Name(), ".")[0] , Description: "Description"})
	}

	return posts

}




func main (){
	fmt.Println("Server started at http://localhost:8000")
	http.HandleFunc("/", index)
	http.HandleFunc("/posts", blog)
	http.HandleFunc("/posts/{id}", blogpost)
	readPostsDit("./posts")
	markdownToHTML("./posts/kubernetes.md")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
