package legacy
import ("crypto/des"; "crypto/ecdsa"; "crypto/ecdh"; "crypto/md5"; "crypto/rand"; "crypto/rc4"; "crypto/rsa"; "crypto/sha1"; "crypto/tls"; mrand "math/rand"; "github.com/golang-jwt/jwt/v5"; "golang.org/x/crypto/pbkdf2")
func ExecuteCoverage() any {
    // rule_key: quantum.arq-q-0571-go
    // evidence_anchor: tls.Config{InsecureSkipVerify:true}
    // regex_sample: tls.Config{InsecureSkipVerify: true}

des.NewTripleDESCipher([]byte("123456789012345678901234"))
    // keywords: InsecureSkipVerify | tls.Config | crypto/tls | crypto/des | des.NewCipher | des.NewTripleDESCipher
    tls.Config{InsecureSkipVerify: true}

    des.NewTripleDESCipher([]byte("123456789012345678901234"))
    return nil
}
