package utils

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
	"fmt"
)

var Encrypt *Encryption

// AES 对称加密
type Encryption struct {
	key string
}

func init() {
	Encrypt = NewEncryption()
}

func NewEncryption() *Encryption {
	return &Encryption{}
}

// PadPwd 填充密码长度
func PadPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize
	// byte(padNum) 将需要填充的字节数转换为一个字节值。
	// 生成一个长度为 padNum 的字节切片，每个字节的值都是 padNum。
	// 第二个输入参数是切片长度
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}

// AesEncoding 加密
func (k *Encryption) AesEncoding(src string) string {
	srcByte := []byte(src)
	block, err := aes.NewCipher([]byte(k.key)) // 创建一个 AES 加密块
	// 如果密钥长度不符合要求，aes.NewCipher 会返回错误，程序会 panic 并终止运行。
	if err != nil {
		fmt.Println("获取AES块数失败: ", err.Error())
		return src
	}

	NewSrcByte := PadPwd(srcByte, block.BlockSize()) // 填充后的数据
	dst := make([]byte, len(NewSrcByte))             // 创建与填充后的数据相同的新字节数组
	block.Encrypt(dst, NewSrcByte)                   // 进行加密，并将结果存储到 dst 中。
	// base64 编码 目的是将二进制数据转换为可读的字符串格式，便于存储或传输。
	pwd := base64.StdEncoding.EncodeToString(dst)
	return pwd
}

// UnPadPwd 去掉填充部分
func UnPadPwd(dst []byte) ([]byte, error) {
	if len(dst) <= 0 {
		return dst, errors.New("长度有误")
	}

	// 最后一个的值肯定是，要填充的数字值
	unPadNum := int(dst[len(dst)-1])
	strErr := "error"
	op := []byte(strErr)
	if len(dst) < unPadNum {
		return op, nil
	}

	str := dst[:len(dst)-unPadNum] // 切片
	return str, nil
}

// AesDecoding 解密
func (k *Encryption) AesDecoding(pwd string) string {
	pwdByte := []byte(pwd)
	pwdByte, err := base64.StdEncoding.DecodeString(string(pwdByte)) // base64解密
	if err != nil {
		return pwd
	}

	block, err := aes.NewCipher([]byte(k.key)) // 创建 AES 加密块
	if err != nil {
		return pwd
	}

	dst := make([]byte, len(pwdByte))
	block.Decrypt(dst, pwdByte) // 解密
	dst, err = UnPadPwd(dst)    // 去掉填充后的字节数组
	if err != nil {
		return "0"
	}

	return string(dst)
}

func (k *Encryption) SetKey(key string) {
	k.key = key
}
