package legacy
import ("crypto/des"; "crypto/ecdsa"; "crypto/ecdh"; "crypto/md5"; "crypto/rand"; "crypto/rc4"; "crypto/rsa"; "crypto/sha1"; "crypto/tls"; mrand "math/rand"; "github.com/golang-jwt/jwt/v5"; "golang.org/x/crypto/pbkdf2")
func ExecuteCoverage() any {
    // rule_key: quantum.arq-q-0574-go
    // evidence_anchor: crypto/des
    // regex_sample: des.NewTripleDESCipher([]byte("123456789012345678901234"))

rc4.NewCipher([]byte("legacy-key"))
    // keywords: crypto/des | des.NewCipher | des.NewTripleDESCipher | crypto/rc4 | rc4.NewCipher
    des.NewTripleDESCipher([]byte("123456789012345678901234"))

    rc4.NewCipher([]byte("legacy-key"))
    return nil
}
