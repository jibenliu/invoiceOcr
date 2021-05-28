package tools

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io"
	"os"
)

type RSAHelper struct {
	GenerateBits  int
	PrivateWriter *io.Writer
	PublicWriter  *io.Writer
	PrivateReader *io.Reader
	PublicReader  *io.Reader
}

// GenerateKey 生成密钥
func (r *RSAHelper) GenerateKey() error {
	if r.GenerateBits == 0 {
		r.GenerateBits = 1000
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, r.GenerateBits)
	if err != nil {
		return err
	}
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	publicKey := privateKey.PublicKey
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}

	if r.PrivateWriter == nil { //写入文件
		privateFile, err := os.Create("private.pem")
		if err != nil {
			return err
		}
		defer func(privateFile *os.File) {
			_ = privateFile.Close()
		}(privateFile)
	}
	err = pem.Encode(*r.PrivateWriter, &privateBlock)
	if err != nil {
		return err
	}
	if r.PublicWriter == nil {
		publicFile, err := os.Create("public.pem")
		if err != nil {
			panic(err)
		}
		defer func(publicFile *os.File) {
			_ = publicFile.Close()
		}(publicFile)
	}
	err = pem.Encode(*r.PublicWriter, &publicBlock)
	if err != nil {
		return err
	}
	return nil
}

// RsaEncrypt RSA加密
func (r *RSAHelper) RsaEncrypt(plainText []byte) []byte {
	if r.PublicReader == nil {
		fi, err := os.Open("public.pem")
		if err != nil {
			return []byte{}
		}
		defer func(fi *os.File) {
			_ = fi.Close()
		}(fi)
		r.PublicReader = interface{}(fi).(*io.Reader)
	}
	buf, err := io.ReadAll(*r.PublicReader)
	if err != nil {
		return []byte{}
	}
	block, _ := pem.Decode(buf)
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		panic(err)
	}
	return cipherText
}

// RsaDecrypt RSA解密
func (r *RSAHelper) RsaDecrypt(cipherText []byte) []byte {
	if r.PrivateReader == nil {
		fi, err := os.Open("private.pem")
		if err != nil {
			return []byte{}
		}
		defer func(fi *os.File) {
			_ = fi.Close()
		}(fi)
		r.PrivateReader = interface{}(fi).(*io.Reader)
	}
	buf, err := io.ReadAll(*r.PrivateReader)
	if err != nil {
		return []byte{}
	}
	block, _ := pem.Decode(buf)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	return plainText
}
