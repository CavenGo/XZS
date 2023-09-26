package encryptutil

import (
	"encoding/base64"
	"github.com/forgoer/openssl"
	"xzs/config"
)

// RsaEncode 加密 公钥使用pkcs8填充
func RsaEncode(src string) (string, error) {
	encrypt, err := openssl.RSAEncrypt([]byte(src), []byte(config.GlobalConf.System.PwdKey.PublicKey))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encrypt), nil
}

// RsaDecode 解密 私钥使用pkcs1填充，这里要注意对方提供的是pkcs1还是pkcs8，如果是pkcs8需要转为pkcs1
// 在线转换网站http://www.metools.info/code/c87.html，也可以在命令行使用openssl命令行转换，特别要注意的就是格式
func RsaDecode(src string) (string, error) {
	base64Decode, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}
	decrypt, err := openssl.RSADecrypt(base64Decode, []byte(config.GlobalConf.System.PwdKey.PrivateKey))
	if err != nil {
		return "", err
	}
	return string(decrypt), nil
}
/*
这段代码实现了使用RSA算法对字符串进行加密和解密的功能，并且在加密和解密过程中使用了配置文件中指定的公钥和私钥。它引用了以下两个包：

"encoding/base64"：用于对加密后的二进制数据进行Base64编码。
"github.com/forgoer/openssl"：一个Go语言的OpenSSL库，提供了RSAEncrypt和RSADecrypt等RSA算法的实现。
该代码包含了两个函数：

RsaEncode
该函数用于将输入的字符串进行RSA加密，实现过程如下：

首先定义了一个字符串src，表示要加密的原始数据。

调用openssl.RSAEncrypt对src进行加密，使用config.GlobalConf.System.PwdKey.PublicKey作为加密所使用的公钥。

如果加密过程中出现了错误，则返回错误信息；否则，将加密后的结果使用base64.StdEncoding.EncodeToString进行Base64编码，并返回编码后的字符串和nil。

RsaDecode

该函数用于将输入的字符串进行RSA解密，实现过程如下：

首先定义了一个字符串src，表示要解密的数据。
调用base64.StdEncoding.DecodeString将src进行Base64解码，得到加密后的二进制数据base64Decode。
调用openssl.RSADecrypt对base64Decode进行解密，使用config.GlobalConf.System.PwdKey.PrivateKey作为解密所使用的私钥。
如果解密过程中出现了错误，则返回错误信息；否则，将解密后的结果转换为字符串，并返回该字符串和nil。
需要注意的是，这里要求私钥使用pkcs1填充，而公钥使用pkcs8填充。如果提供的是pkcs8格式的公钥，需要先进行格式转换。转换方式可以通过在线转换网站或使用OpenSSL命令行进行转换。另外，还要确保使用的公钥和私钥与对方提供的一致。

通过调用这两个函数，可以方便地实现RSA加密和解密功能，并且可以通过配置文件指定使用的公钥和私钥，提高了代码的灵活性和可维护性。
*/