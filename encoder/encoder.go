package encoder

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/wen-long/bdpass"
)

const (
	_stdFormat = "%s#%s#%d#%s"
	_pdlFormat = "%s|%d|%s|%s"
	_pcsFormat = "BaiduPCS-Go rapidupload -length=%d -md5=%s -slicemd5=%s -crc32=%d %q"
)

type (
	// 梦姬标准格式
	STD struct{}

	// PanDownload 格式
	PDL struct{}

	// BaiduPCS-Go 格式
	PCS struct{}

	Encoder interface {
		Encode(meta *bdpass.RapidUploadMeta) string
	}
)

func (*STD) Encode(meta *bdpass.RapidUploadMeta) string {
	return fmt.Sprintf(_stdFormat,
		strings.ToUpper(meta.ContentMD5),
		strings.ToUpper(meta.SliceMD5),
		meta.ContentLength,
		meta.Filename,
	)
}

func (*PDL) Encode(meta *bdpass.RapidUploadMeta) string {
	s := fmt.Sprintf(_pdlFormat,
		meta.Filename,
		meta.ContentLength,
		meta.ContentMD5,
		meta.SliceMD5,
	)
	return "bdpan://" + base64.URLEncoding.EncodeToString([]byte(s))
}

func (*PCS) Encode(meta *bdpass.RapidUploadMeta) string {
	return fmt.Sprintf(_pcsFormat,
		meta.ContentLength,
		meta.ContentMD5,
		meta.SliceMD5,
		meta.ContentCRC32,
		meta.Filename,
	)
}
