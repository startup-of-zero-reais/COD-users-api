package x_api_key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

type (
	// Encryptor é a estrutura de encriptação das chaves de api
	Encryptor struct {
		PrivateKey *rsa.PrivateKey
		PublicKey  *rsa.PublicKey
	}
)

// NewEncryptor é o construtor de Encryptor
func NewEncryptor() *Encryptor {
	encryptor := &Encryptor{}

	privateKey, err := encryptor.GetRSA()
	if err != nil {
		log.Printf("Key pairs não encontrada: %s\nGerando novo par de chaves...", err.Error())
		err = encryptor.GenerateKeyPairs()
		if err != nil {
			log.Printf("Erro ao gerar key pairs: %s", err.Error())
			return encryptor
		}

		return encryptor
	}

	block, _ := pem.Decode(privateKey)
	encryptor.PrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Printf("Erro ao recuperar private key: %s", err.Error())
		return encryptor
	}

	publicKey, err := encryptor.GetRSAPub()
	if err != nil {
		log.Printf("Erro ao recuperar id_rsa.pub: %s", err.Error())
		return encryptor
	}

	block, _ = pem.Decode(publicKey)
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Printf("Erro ao recuperar public key: %s", err.Error())
		return encryptor
	}
	encryptor.PublicKey = pubKey.(*rsa.PublicKey)

	return encryptor
}

const (
	_ = iota
	// BitsBase é o tamanho de 512 bits de criptografia
	BitsBase = 512 * iota
	// BitsLv2 é o tamanho de 1024 bits de criptografia
	BitsLv2 = BitsBase * iota
	_       = iota
	// BitsLv3 é o tamanho de 4096 bits de criptografia
	BitsLv3 = BitsLv2 * iota
)

// GenerateKeyPairs gera os arquivos das chaves pública e privada RSA
func (e *Encryptor) GenerateKeyPairs() error {
	privateKey, _ := rsa.GenerateKey(rand.Reader, BitsLv3)
	publicKey := &privateKey.PublicKey

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	privateIDRsa, err := os.Create("id_rsa")
	if err != nil {
		return err
	}

	err = pem.Encode(privateIDRsa, privateKeyBlock)
	if err != nil {
		return err
	}
	e.PrivateKey = privateKey

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	publicIDRsa, err := os.Create("id_rsa.pub")
	if err != nil {
		return err
	}

	err = pem.Encode(publicIDRsa, publicKeyBlock)
	if err != nil {
		return err
	}
	e.PublicKey = publicKey

	return nil
}

// GetFileContent lê os 'bytes' do arquivo definido no parâmetro
func (e *Encryptor) GetFileContent(file string) ([]byte, error) {
	_, err := os.ReadFile(getKeysLocation(file))
	if err != nil {
		return nil, err
	}

	f, err := os.Open(getKeysLocation(file))
	if err != nil {
		return nil, err
	}

	b1 := make([]byte, 4096)
	n1, err := f.Read(b1)
	if err != nil {
		return nil, err
	}

	return b1[:n1], nil
}

// GetRSA pega o conteúdo do arquivo de chave id_rsa
func (e *Encryptor) GetRSA() ([]byte, error) {
	return e.GetFileContent("id_rsa")
}

// GetRSAPub pega o conteúdo do arquivo de chave pública id_rsa.pub
func (e *Encryptor) GetRSAPub() ([]byte, error) {
	return e.GetFileContent("id_rsa.pub")
}

// O getKeysLocation é para resolver o caminho dos arquivos baseado no 'Environment'
func getKeysLocation(file string) string {
	if e := os.Getenv("APPLICATION_ENV"); e != "" {
		if e == "development" {
			return fmt.Sprintf("/go/src/%s", file)
		}
	}

	return fmt.Sprintf("/go/build/%s", file)
}
