package middlewares

import (
	"bytes"
	"crypto/sha512"
	"flag"
	"net/http"
	"regexp"
	_ "strings"

	log "github.com/sirupsen/logrus"
)

// TODO checkAuth


func init() {
	flag.StringVar(
		&secretKey, 
		"sk", 
		"my_secret_key", 
		"secret key for allows methods",
	)
}

var ( 
	expectKey []byte
	secretKey string
	ifInit bool

	Logger *log.Logger
)

func initExpectKey() {
	if !ifInit {
		sha := sha512.New()
		sha.Write([]byte(secretKey))
		expectKey = sha.Sum(nil)
		ifInit = true
	}
}

const (
	authHeaderPattern = `^Token ([^\s]{1,})`
)

// ContentTypeJSONMiddleware set header content-type to application json
func ContentTypeJSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w,r)
	})
}

func CheckTokenIfFromService(next http.Handler) http.Handler {
	initExpectKey()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Example Authorization: Token dasdawfegjkdsa
		authHeader := r.Header.Get("Authorization")
		re := regexp.MustCompile(authHeaderPattern)
		s := re.FindStringSubmatch(authHeader)

		if len(s) == 0 {
			Logger.WithFields(log.Fields{
				"package": "middlewares",
				"msg": "don't find token",
			}).Info()

			w.WriteHeader(http.StatusForbidden)
			return
		}

		if len(s) != 2 {
			Logger.WithFields(log.Fields{
				"package": "middlewares",
				"msg": "don't have matches",
			}).Info()
			
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		token := s[1]

		if !bytes.Equal(expectKey, []byte(token)) {
			Logger.WithFields(log.Fields{
				"package": "middlewares",
				"msg": "don't equal",
			}).Info()

			w.WriteHeader(http.StatusForbidden)
			return
		}
		
		next.ServeHTTP(w,r)
	})
}