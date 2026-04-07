package legacy
import ("crypto/des"; "crypto/ecdsa"; "crypto/ecdh"; "crypto/md5"; "crypto/rand"; "crypto/rc4"; "crypto/rsa"; "crypto/sha1"; "crypto/tls"; mrand "math/rand"; "github.com/golang-jwt/jwt/v5"; "golang.org/x/crypto/pbkdf2")
func ExecuteCoverage() any {
    // rule_key: quantum.arq-q-0575-go
    // evidence_anchor: crypto/rc4
    // regex_sample: rc4.NewCipher([]byte("legacy-key"))

NewECBEncrypter(block, iv)
    // keywords: crypto/rc4 | rc4.NewCipher | NewECBEncrypter | NewECBDecrypter | ECB
    rc4.NewCipher([]byte("legacy-key"))

    NewECBEncrypter(block, iv)
    return nil
}
