package service

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"shopping/conf"
	"strconv"
)

func uploadAvatarToLocalStatic(file multipart.File, userId uint, username string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId)) // 路径拼接
	basePath := "." + conf.AvatarPath + "user" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}

	avatarPath := basePath + username + ".jpg" // todo 把file后缀提取出来
	// 转为byte类型
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return "", err
	}

	return "user" + bId + "/" + username + ".jpg", nil
}

func uploadProductToLocalStatic(file multipart.File, userId uint, productName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId)) // 路径拼接
	basePath := "." + conf.ProductPath + "boss" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}

	productPath := basePath + productName + ".jpg" // todo 把file后缀提取出来
	// 转为byte类型
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(productPath, content, 0666)
	if err != nil {
		return "", err
	}

	return "boss" + bId + "/" + productName + ".jpg", nil
}

// 判断文件夹是否存在
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// CreateDir 创建文件夹
func CreateDir(dirName string) bool {
	// 第一位（7）：文件所有者的权限。
	// 第二位（5）：文件所属组的权限。
	// 第三位（5）：其他用户的权限。
	// 4 表示读取权限（r）。
	// 2 表示写入权限（w）。
	// 1 表示执行权限（x）。
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		return false
	}
	return true
}
