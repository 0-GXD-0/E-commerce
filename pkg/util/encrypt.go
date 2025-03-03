package util

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"log"
)

var Encrypt *Encryption

// AES对称加密
type Encryption struct {
	key string
}

func init() {
	Encrypt = NewEncryption()
}

func NewEncryption() *Encryption {
	return &Encryption{}
}

// PadPwd填充密码长度
func PadPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}

// 加密
func (k *Encryption) AesEncoding(src string) string {
	//将输入字符串转换为字节数组
	srcByte := []byte(src)
	//创建 AES 密码块
	block, err := aes.NewCipher([]byte(k.key))
	if err != nil {
		return ""
	}
	//密码填充，由于字节长度不够，需要填充
	NewSrcByte := PadPwd(srcByte, block.BlockSize())
	//加密
	dst := make([]byte, len(NewSrcByte))
	block.Encrypt(dst, NewSrcByte)
	//base64编码
	pwd := base64.StdEncoding.EncodeToString(dst)
	return pwd
}

// UnPadPwd去除填充
func UnPadPwd(dst []byte) ([]byte, error) {
	if len(dst) <= 0 {
		return []byte{}, nil
	}
	//去掉的长度
	unpadNum := int(dst[len(dst)-1])
	strErr := "error"
	op := []byte(strErr)
	if len(dst) < unpadNum {
		return op, nil
	}
	str := dst[:(len(dst) - unpadNum)]
	return str, nil
}

// AesDecoding解密
func (k *Encryption) AesDecoding(pwd string) string {
	pwdByte := []byte(pwd)
	//解码Base64编码的字符串
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		log.Printf("解码错误", err.Error())
		return ""
	}
	//创建AES密码块
	block, errBlock := aes.NewCipher([]byte(k.key))
	if errBlock != nil {
		log.Printf("创建AES密码块出错", errBlock.Error())
		return ""
	}
	//解密
	dst := make([]byte, len(pwdByte))
	block.Decrypt(dst, pwdByte)
	dst, err = UnPadPwd(dst)
	if err != nil {
		log.Printf("解密出错", err.Error())
		return ""
	}
	log.Printf("解密后money为", string(dst))
	return string(dst)
}

func (k *Encryption) SetKey(key string) {
	k.key = key
}
