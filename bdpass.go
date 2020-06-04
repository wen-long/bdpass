package bdpass

import (
	"crypto/md5"
	"encoding/hex"
	"hash/crc32"
	"io"
	"os"
	"path/filepath"
)

const (
	_size = 256 * 1024
)

type RapidUploadMeta struct {
	Filename      string
	ContentLength int64
	ContentMD5    string
	SliceMD5      string
	ContentCRC32  uint32
}

func Stat(filename string) (*RapidUploadMeta, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	meta := &RapidUploadMeta{
		Filename: filepath.Base(filename),
	}

	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	meta.ContentLength = fi.Size()

	data := make([]byte, _size)
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}
	sliceMD5 := md5.Sum(data)
	meta.SliceMD5 = hex.EncodeToString(sliceMD5[:])

	hash := md5.New()
	_, err = hash.Write(data)
	if err != nil {
		return nil, err
	}
	hash32 := crc32.NewIEEE()
	_, err = hash32.Write(data)
	if err != nil {
		return nil, err
	}
	dst := io.MultiWriter(hash, hash32)
	_, err = io.Copy(dst, file)
	if err != nil {
		return nil, err
	}
	meta.ContentMD5 = hex.EncodeToString(hash.Sum(nil))
	meta.ContentCRC32 = hash32.Sum32()

	return meta, nil
}
