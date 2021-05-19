package crypt

import (
	"encoding/base64"
	"app/system/openssl"
)

var (
	passphrase = "DEL!@12sha"
	o = openssl.New()
)
// Encrypt Methods Defination
func Encrypt(plaintext string) string {
	enc, _ := o.EncryptBytes(passphrase, []byte(plaintext), openssl.PBKDF2SHA256)
	return base64.URLEncoding.EncodeToString(enc)
}
// Decrypt Methods Defination
func Decrypt(opensslEncrypted string) string {
	decoded, _ := base64.URLEncoding.DecodeString(opensslEncrypted)
	dec, _ := o.DecryptBytes(passphrase, []byte(decoded), openssl.PBKDF2SHA256)
	return string(dec)
}