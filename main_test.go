package main

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// creates a testing page and returns the page's title
func setup(t *testing.T) (string, func()) {
	p := &Page{
		Title: "testing-" + time.Now().Local().String(),
	}

	err := p.save()
	if err != nil {
		t.Fatal(err)
	}

	return p.Title, func() {
		os.Remove("pages/" + p.Title + ".txt")
	}
}

// load a page that doesn't exists
func TestViewNotFound(t *testing.T) {
	_, err := loadPage("/")
	assert.Error(t, err)
}

// load a page that does exists
func TestViewFound(t *testing.T) {
	title, cleanup := setup(t)
	defer cleanup()

	_, err := loadPage(title)
	assert.NoError(t, err)
}

// create a page, check that the page's file exists
func TestSave(t *testing.T) {
	title, cleanup := setup(t)
	defer cleanup()

	_, err := os.Stat("pages/" + title + ".txt")
	assert.NoError(t, err)
}

// create a page, edit the page's body, reload the page, check that the page's body has changed
func TestEdit(t *testing.T) {
	title, cleanup := setup(t)
	defer cleanup()

	p, err := loadPage(title)
	assert.NoError(t, err)

	oldBody := p.Body

	p.Body = []byte("edited-" + time.Now().Local().String())

	err = p.save()
	assert.NoError(t, err)

	pp, err := loadPage(title)
	assert.NoError(t, err)

	assert.NotEqual(t, pp.Body, oldBody)
}
