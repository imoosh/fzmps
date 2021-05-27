package utils

import (
	"centnet-fzmps/common/log"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func AesCBCDecrypt(encryptData, key, iv []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error(err)
		return nil, errors.New("aes.NewCipher error: " + err.Error())
	}

	blockSize := block.BlockSize()
	if len(encryptData) < blockSize {
		log.Error("cipher text is too short")
		return nil, errors.New("cipher text is too short")
	}

	if len(encryptData)%blockSize != 0 {
		log.Error("cipher text is not a multiple of the block size")
		return nil, errors.New("cipher text is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)

	decryptedData := make([]byte, len(encryptData))
	mode.CryptBlocks(decryptedData, encryptData)
	decryptedData = PKCS7UnPadding(decryptedData)
	return decryptedData, nil
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
