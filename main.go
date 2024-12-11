package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 定义一个命令行参数，用于接收目录路径
	dirPtr := flag.String("dir", "", "The directory to scan for images")
	flag.Parse()

	var currentDir string
	var err error

	// 检查是否提供了目录路径
	if *dirPtr == "" {
		// 如果没有提供目录路径，使用当前工作目录
		currentDir, err = os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			return
		}
	} else {
		// 如果提供了目录路径，处理相对路径和绝对路径
		currentDir = filepath.Clean(*dirPtr)
	}

	fmt.Printf("Current directory: %s\n", currentDir)

	// 遍历当前目录及其子目录中的所有文件
	err = filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 跳过 compress_file 目录
		if info.IsDir() && info.Name() == "compress_file" {
			return filepath.SkipDir
		}

		if !info.IsDir() {
			// 提取文件名（不包含路径）
			fileName := info.Name()
			if fileName == "compressImg" || fileName == "compressImg.exe" {
				return nil
			}
			// 去掉文件后缀
			baseName := filepath.Base(fileName)
			ext := filepath.Ext(baseName)
			cleanFileName := baseName[:len(baseName)-len(ext)]

			if strings.HasSuffix(cleanFileName, "_thumb") {
				return nil
			}
			thumbDir := filepath.Join(currentDir, "compress_file")
			thumbFilepath := filepath.Join(thumbDir, cleanFileName+"_thumb"+ext)

			// 检查并创建 compress_file 文件夹
			if err := os.MkdirAll(thumbDir, os.ModePerm); err != nil {
				fmt.Printf("Failed to create directory %s: %v\n", thumbDir, err)
				return err
			}

			// 调用 compress 函数压缩文件
			if err := CompressImage(path, thumbFilepath, 50); err != nil {
				fmt.Printf("Failed to compress %s: %v\n", path, err)
			} else {
				fmt.Printf("Successfully compressed %s\n", path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path:", err)
	}
}
