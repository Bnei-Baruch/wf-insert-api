package main

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

type props struct {
	Filename string `json:"filename"`
	Sha1     string `json:"sha1"`
	Size     int64  `json:"size"`
	Mimetype string `json:"type"`
	Url      string `json:"url"`
}

func (s *props) fileProps(filepath string) error {

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}

	fi, err := f.Stat()
	if err != nil {
		return err
	}

	s.Size = fi.Size()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return err
	}

	s.Sha1 = hex.EncodeToString(h.Sum(nil))

	err = os.Rename(filepath, *storage + s.Sha1)
	if err != nil {
		return err
	}

	s.Url = *url + s.Sha1

	return nil
}
