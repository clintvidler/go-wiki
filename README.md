# go-wiki

View, edit, and save text files via http to the file system.

(WIP) https://golang.org/doc/articles/wiki/

## Instructions

`go run main.go`

In a web browser, visit localhost:8080

/view/x will redirect to create a page with title 'x'
/edit/x will edit page with title 'x' (Note: editing a non-existent url will create a new page)
