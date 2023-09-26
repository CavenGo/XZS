package encryptutil

import (
	"encoding/base64"
	"github.com/forgoer/openssl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRsaEncode(t *testing.T) {
	src := "123456"
	// 公钥使用pks8
	pubKey := "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDb6bxpsivyFxIC8suHnbMIPLYX\nN6amWYY2rarcxnkcyvaECJ2PDLiouo0uc0IFH/Dez8HmN0aL9rZn2za09lnSky9o\nUyJKZlwY7vDmiYeVyuNHvcYLZWT5+0OLbQ22PoEiGrhfCMxzqYdb0FkfDfWU2M/k\nxCAwhn2Hdu04VJz71QIDAQAB\n-----END PUBLIC KEY-----"
	encrypt, err := openssl.RSAEncrypt([]byte(src), []byte(pubKey))
	assert.NoError(t, err)
	toString := base64.StdEncoding.EncodeToString(encrypt)
	t.Log(toString)
}

func TestRsaDecode(t *testing.T) {
	src := "vjzxPjmnsRkSxJfPpGX9ud3CIr7eA0otUnQhmLfDTxst2JCoAa4gBrHRn3G58WeQGFhTxlL5F6xhlQPu6j9RAYz1qUpPDCCpVe8zx2WDPAOkx/2cDB3t5MMP2M1UTXVBTmwz07I7dI6RFuqeIqT9s2D7I/vUYMp1tHypYA41P5M="
	privateKey := "-----BEGIN RSA PRIVATE KEY-----\nMIICXgIBAAKBgQDb6bxpsivyFxIC8suHnbMIPLYXN6amWYY2rarcxnkcyvaECJ2P\nDLiouo0uc0IFH/Dez8HmN0aL9rZn2za09lnSky9oUyJKZlwY7vDmiYeVyuNHvcYL\nZWT5+0OLbQ22PoEiGrhfCMxzqYdb0FkfDfWU2M/kxCAwhn2Hdu04VJz71QIDAQAB\nAoGBAND8woJLwTmStRo6NDOQKVi1oXJU/7lssIB78DlZIDW9qCH3sgwE0eP/TTYM\ncHxAS37jP2iRtShD8Dqod8fnqZkSL6glUUj2Aui120VwiR2xwJDnsYlpRdKSNRoV\nIgH5Dyc0fKoCo0Yq/pZcABL55ITsUXLyXJW4NZeUA04docSNAkEA7VaV7R52Xfxt\n8tUy4kAbwHaXGUFM2NU3pKMLuVVQoQ/qPDpBvm/yBh2ApCX7qqgh+ZsfXDMwY6n3\nwbLU4yVdnwJBAO00ZWtjvDH86qfpKcRDc6IVxMt91qMGIZOIOch6OxOAL+7zg4G5\nI6UHpy6Orl1n+2ycVfL1PdaQmHXpczHOSgsCQGQlsyHZRs0l5Scge1YpAwzVfbC0\ncz7TyaT4/8t2io1L7+T2GCPJjPCzpkKdnHJIe/2dTUBUgUiswdTEJzyp2bUCQQCK\nUpil1AYlvE/2VKB3g8IFjd4xsBMfA+9GghT4FFco2wKYvEY+uoDPtrPGEYwaig1y\n24O/Z0WFPtK5R8ZWD+7bAkEA6r4S4sdispPMjKVwZhi55t+nH1gRNb/f+HTdQFVW\nh0IznBGyVeoDDrsZuGhIOsxpFpeLek5NJVIcqqhUuZFKog==\n-----END RSA PRIVATE KEY-----"
	base64Decode, _ := base64.StdEncoding.DecodeString(src)
	decrypt, err := openssl.RSADecrypt(base64Decode, []byte(privateKey))
	assert.NoError(t, err)
	t.Logf("decrypt：%s", decrypt)
}
/*
这段代码实现了使用RSA算法对字符串进行加密和解密的功能。它引用了以下三个包：

"encoding/base64"：用于对加密后的二进制数据进行Base64编码。
"github.com/forgoer/openssl"：一个Go语言的OpenSSL库，提供了RSAEncrypt和RSADecrypt等RSA算法的实现。
"github.com/stretchr/testify/assert"：一个测试断言库，用于编写单元测试。
该代码包含了两个测试函数：

TestRsaEncode
该函数用于测试RSA加密功能，实现过程如下：

首先定义了一个字符串src，表示要加密的原始数据。

然后定义了一个公钥pubKey，使用PKCS8格式存储，表示用于加密的公钥。

调用openssl.RSAEncrypt对src进行加密，得到加密后的二进制数据encrypt。

调用base64.StdEncoding.EncodeToString将encrypt进行Base64编码，得到一个字符串toString。

最后使用testify/assert库的NoError方法断言加密过程中没有出现错误，并将加密后的结果输出到日志中。

TestRsaDecode

该函数用于测试RSA解密功能，实现过程如下：

首先定义了一个字符串src，表示要解密的数据，它是通过TestRsaEncode函数加密后得到的结果。
然后定义了一个私钥privateKey，表示用于解密的私钥。
调用base64.StdEncoding.DecodeString将src进行Base64解码，得到加密后的二进制数据base64Decode。
调用openssl.RSADecrypt对base64Decode进行解密，得到解密后的二进制数据decrypt。
最后使用testify/assert库的NoError方法断言解密过程中没有出现错误，并将解密后的结果输出到日志中。
通过编写这两个测试函数，可以对RSA加密和解密功能进行全面的测试，确保其正确性和稳定性。
*/