# go-wiki

View, edit, and save text files via http to the file system.

(DONE) https://golang.org/doc/articles/wiki/

## Instructions

`go run main.go`

In a web browser, visit localhost:8080

- /view/x will redirect to create a page with title 'x'

- /edit/x will edit page with title 'x' (Note: editing a non-existent url will create a new page)

## Generate self-signed certificates

```
mkdir certs
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
    -keyout certs/localhost.key -out certs/localhost.crt \
    -subj "/C=AU/ST=Australia/L=Canberra/O=Clint Vidler/OU=Dev/CN=localhost/emailAddress=clint@vidler"
```
