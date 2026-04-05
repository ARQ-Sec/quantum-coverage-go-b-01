package legacy
import ("crypto/des"; "crypto/ecdsa"; "crypto/ecdh"; "crypto/md5"; "crypto/rand"; "crypto/rc4"; "crypto/rsa"; "crypto/sha1"; "crypto/tls"; mrand "math/rand"; "github.com/golang-jwt/jwt/v5"; "golang.org/x/crypto/pbkdf2")
func ExecuteCoverage() any {
    // rule_key: quantum.arq-q-0570-go
    // evidence_anchor: crypto/tls.Config
    // regex_sample: &tls.Config{MinVersion: tls.VersionTLS11, MaxVersion: tls.VersionTLS12}

tls.Config{InsecureSkipVerify: true}
    // keywords: crypto/tls | tls.Config | MinVersion | TLS1.0 | TLS1.1 | SSLv3
    &tls.Config{MinVersion: tls.VersionTLS11, MaxVersion: tls.VersionTLS12}

    tls.Config{InsecureSkipVerify: true}
    return nil
}
