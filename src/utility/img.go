package utility

import (
  "fmt"
	"os"
	"encoding/base64"
	"strings"
)

func UploadImage(raw, path, name string) (string, error) {
	i := strings.Index(raw, "base64,") + 7
	size := len(raw);
	val := raw[i:size - 8]

  full := fmt.Sprintf("%s/%s", path, name)
	data, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(val)
	if err == nil {
		size = len(data)
		// open output file
    fo, err := os.Create(full)
    if err == nil {
			// write a chunk
      _, err = fo.Write(data[:size])
 			// close fo on exit
      if err == nil {
        err = fo.Close();
      }
		}
	}

  return full, err
}

func CreatePath(path string) {
  os.MkdirAll(path, os.ModePerm)
}
