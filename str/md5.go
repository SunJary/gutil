package str

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
)

func MD5(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

func MD5File(header *multipart.FileHeader) (string, error) {
	fileIO, openError := header.Open()
	if openError != nil {
		return "", openError
	}

	defer fileIO.Close()

	r := bufio.NewReader(fileIO)

	h := md5.New()

	_, err := io.Copy(h, r)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
