package pkg

import (
	"net/http"
	"path/filepath"
)

// NeuteredFileSystem is a wrapper around http.FileSystem that prevents directory listing.
// ref: https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings
type NeuteredFileSystem struct {
	FS http.FileSystem
}

// Open opens the named file. If the file is a directory, it attempts to open the index.html file within the directory.
// If the index.html file does not exist, it returns an error.
func (nfs NeuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.FS.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.FS.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
