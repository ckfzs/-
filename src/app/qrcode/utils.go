package qrcode

import (
    QRCEncoder "github.com/skip2/go-qrcode"
    QRCDecoder "github.com/tuotoo/qrcode"
    "io"
    "fmt"
    "errors"
)

func Decode(fi io.Reader) (content string, err error) {
    defer func() {
        if r := recover(); r != nil {
            content = ""
            err = fmt.Errorf("%v", r)
        }
    }()
    qrmatrix, _err := QRCDecoder.Decode(fi)
    if _err != nil {
        return "", _err
    } else {
        return qrmatrix.Content, nil
    }
}

func Encode(content string) (bytes []byte, err error) {
    if len(content) == 0 {
        return nil, errors.New("content can't be empty")
    }
    defer func() {
        if r := recover(); r != nil {
            bytes = nil
            err = fmt.Errorf("%v", r)
        }
    }()
    png, _err := QRCEncoder.Encode(content, QRCEncoder.Medium, 256)
    if _err != nil {
        return nil, _err
    } else {
        return png, nil
    }
}

