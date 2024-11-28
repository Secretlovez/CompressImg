package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
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
			thumbFilepath := filepath.Join(currentDir, cleanFileName+"_thumb"+ext)

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
