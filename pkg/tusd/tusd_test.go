package main

import (
	"github.com/tus/tusd"
	"github.com/tus/tusd/filestore"
	"net/http"
	"testing"
)

func Test_Upload(t *testing.T) {
	store := filestore.FileStore{
		Path: "./uploads",
	}
	composer := tusd.NewStoreComposer()
	store.UseIn(composer)

	handler, err := tusd.NewHandler(tusd.Config{
		BasePath:      "/files/",
		StoreComposer: composer,
	})
	if err != nil {
		t.Errorf("Unable to create handler:  %s", err)
	}

	http.Handle("/files/", http.StripPrefix("/files/", handler))
	err = http.ListenAndServe(":4001", nil)
	if err != nil {
		t.Errorf("Unable to listen: %s", err)
	}
}
