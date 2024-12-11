package main

import (
	"fmt"
	"os"

	"github.com/disintegration/imaging"
)

// CompressImage 接收图片路径并进行压缩
// 返回错误信息
func CompressImage(originImagePath string, thumbImagePath string, quality int) error {
	fmt.Printf("start compressing image, originImagePath=%s, quality=%d", originImagePath, quality)

	originImage, err := imaging.Open(originImagePath, imaging.AutoOrientation(true))
	if err != nil {
		fmt.Printf("failed to open image: %v", err)
		return err
	}

	thumbImageFile, err := os.Create(thumbImagePath)
	if err != nil {
		fmt.Printf("failed to create thumb image: %v", err)
		return err
	}
	defer thumbImageFile.Close()

	compressedImage := imaging.Resize(originImage, 600, 0, imaging.Lanczos)
	err = imaging.Encode(thumbImageFile, compressedImage, imaging.JPEG)
	if err != nil {
		fmt.Printf("failed to save image: %v", err)
		os.Remove(thumbImagePath)
		return err
	}

	fmt.Printf("compressing image finished, originImagePath=%s, quality=%d", originImagePath, quality)
	return nil
}
