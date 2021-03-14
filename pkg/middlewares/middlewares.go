package middlewares

import (
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	_ "strings"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// TODO checkAuth


func init() {
	flag.StringVar(
		&SecretKey, 
		"sk", 
		"my_secret_key", 
		"secret key for allows methods",
	)
}

var ( 
	expectKey []byte
	SecretKey string
	ifInit bool

	Logger *log.Logger
)

func initExpectKey() {
	if !ifInit {
		sha := sha512.New()
		sha.Write([]byte(SecretKey))
		expectKey = sha.Sum(nil)
		ifInit = true
	}
}

func initLogger() {
	if Logger == nil {
		Logger = logrus.StandardLogger()
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
	initLogger()

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

		key, err := hex.DecodeString(token)
		if err != nil {
			Logger.WithFields(log.Fields{
				"package": "middlewares",
				"msg": err.Error(),
			}).Info()
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !bytes.Equal(expectKey, key) {
			Logger.WithFields(log.Fields{
				"package": "middlewares",
				"msg": "don't equal",
				"expect": string(expectKey),
				"sekretKey": SecretKey,
				"key": key,
			}).Info()

			w.WriteHeader(http.StatusForbidden)
			return
		}
		
		next.ServeHTTP(w,r)
	})
}

func CheckBodyIfEmpty(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := ioutil.ReadAll(r.Body); err == io.EOF {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Body is null")
			return
		} else if err != nil {
			log.Errorf("Err in middlewares %v", err)
			return
		}
		next.ServeHTTP(w, r)
	})
}