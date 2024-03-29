package x_api_key

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/router"
	"github.com/startup-of-zero-reais/COD-users-api/domain/utilities"
	"strings"
)

type (
	// XApiKey é a estrutura de chaves de api
	XApiKey struct {
		encryptor *Encryptor
		pk        *rsa.PrivateKey
	}
)

// NewXApiKey é o construtor de XApiKey
func NewXApiKey() *XApiKey {
	encryptor := NewEncryptor()
	return &XApiKey{
		encryptor: encryptor,
		pk:        encryptor.PrivateKey,
	}
}

// KeyAuth retorna o middleware que valida o cabeçalho X-Api-Key
func (x *XApiKey) KeyAuth() router.MiddlewareHandler {
	return (router.MiddlewareHandler)(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup:  "header:X-Api-Key",
		AuthScheme: "X-Api-Key",
		Validator:  x.IsValidKey,
	}))
}

// CheckApplication é o middleware que valida o cabeçalho Application
func (x *XApiKey) CheckApplication() router.MiddlewareHandler {
	return (router.MiddlewareHandler)(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup:  "header:Application",
		AuthScheme: "Application",
		Validator:  x.ValidApplication,
	}))
}

// ValidApplication valida se o cabeçalho Application possui um valor aceito e se
// é um valor permitido
func (x *XApiKey) ValidApplication(key string, _ echo.Context) (bool, error) {
	validPlatforms := strings.Split(utilities.GetEnv("PERMIT_APPLICATIONS", ""), ",")
	if len(validPlatforms) == 0 {
		return false, errors.New("requisição nao autorizada")
	}
	b, err := json.Marshal(validPlatforms)
	if err != nil {
		return false, err
	}

	isPermitted := bytes.Contains(b, []byte(key))
	return isPermitted, nil
}

// IsValidKey valida se o cabeçalho de X-Api-Key é válido e autorizado
func (x *XApiKey) IsValidKey(key string, c echo.Context) (bool, error) {
	application := c.Request().Header.Get("Application")
	hashApp := sha1.New()
	hashApp.Write([]byte(application))
	hashSum := hashApp.Sum(nil)

	headerBuff := &bytes.Buffer{}
	headerBuff.Write([]byte(key))

	hashes := strings.Split(headerBuff.String(), ".")
	app := hashes[0]
	appBytes, _ := base64.StdEncoding.DecodeString(app)

	if app != base64.StdEncoding.EncodeToString(hashSum) {
		return false, errors.New("unauthorized")
	}

	signature := hashes[1]
	signBytes, _ := base64.StdEncoding.DecodeString(signature)

	err := rsa.VerifyPSS(x.encryptor.PublicKey, crypto.SHA1, appBytes, signBytes, nil)
	if err != nil {
		return false, err
	}

	return app == base64.StdEncoding.EncodeToString(hashSum), nil
}

// GenerateApiKey é responsável por criar uma chave de api baseado no cabeçalho application
func (x *XApiKey) GenerateApiKey(application string) (string, error) {
	secret := []byte(application)

	secretHash := sha1.New()
	_, err := secretHash.Write(secret)
	if err != nil {
		return "", err
	}
	hashSum := secretHash.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, x.pk, crypto.SHA1, hashSum, nil)
	signBuff := &bytes.Buffer{}
	signBuff.Write(signature)
	if err != nil {
		return "", err
	}

	hash64 := base64.StdEncoding.EncodeToString(hashSum)
	sign64 := base64.StdEncoding.EncodeToString(signBuff.Bytes())
	key := fmt.Sprintf("%s.%s", hash64, sign64)
	return key, nil
}
