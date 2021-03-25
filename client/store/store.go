package store

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"log"
	"net"
	"net/http"
	"time"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

// Client interface for basic client structure
type Client struct {
	Logger     *logrus.Logger
	BaseURL string
	HTTPClient *http.Client
}

type Response struct {
	Name  string
	Hash string
	Size string
}
// NewClient creates a new client
func NewClient() *Client {
	return &Client{
		BaseURL: viper.GetString("server-url"),
		Logger: logrus.New(),
		HTTPClient: &http.Client {
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout: 5 * time.Second,
				}).Dial,
			},
			Timeout: time.Second * 10,
		},
	}
}


// multipartBody adds files to a multipart writer
func multipartBody(files []string) (*bytes.Buffer, string, error) {
	bodyBuffer := new(bytes.Buffer)
	bodyWriter := multipart.NewWriter(bodyBuffer)
	for _, fn := range files {
		file, err := os.Open(fn)
		if err != nil {
			return nil, "", err
		}
		defer file.Close()
		fi, err := file.Stat()
		if err != nil {
			return nil, "", err
		}
		part, err := bodyWriter.CreateFormFile(fn, fi.Name())
		if err != nil {
			return nil, "", err
		}
		_, err = io.Copy(part, file)
		if err != nil {
			return nil, "", err
		}
	}
	err := bodyWriter.Close()
	if err != nil {
		return nil, "", err
	}
	return bodyBuffer, bodyWriter.FormDataContentType(), nil
}

func (c *Client) AddAndSave(files []string)  error {
	hash, err := c.Add(files)
	if err != nil {
		c.Logger.Fatalf("Could not add files %v", err)
		return err
	}
	return c.SaveMultiHash(hash)
}
// Add adds files to the store
func (c *Client) Add(files []string) (string, error) {
	bodyBuffer, contentType, err := multipartBody(files) 
	if err != nil {
		c.Logger.Fatalf("Could not build request %v", err)
		return "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/add", c.BaseURL), bodyBuffer)
	c.Logger.Debugf("req %v", req)
	if err != nil {
		c.Logger.Fatalf("Could not build request %v", err)
		return "", err
	}

	req.Header.Add("Content-Type", contentType)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		c.Logger.Fatalf("Could not get response %v", err)
		return "", err
	}
	
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated) || (err != nil) {
		c.Logger.Errorf("HTTPStatusCode: '%d'; ResponseMessage: '%s'; ErrorMessage: '%v'", resp.StatusCode, string(b), err)
		return "", fmt.Errorf("HTTPStatusCode: '%d'; ResponseMessage: '%s'; ErrorMessage: '%v'", resp.StatusCode, string(b), err)
	}
	c.Logger.Infof("HTTPStatusCode: '%d'; ResponseMessage: '%s'", resp.StatusCode, string(b))

	response := Response{}
	jsonErr := json.Unmarshal(b, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return response.Hash, nil
}

func (c *Client) SaveMultiHash(multihash string)  error {
	// connect to an ethereum node  hosted by infura
	client, err := ethclient.Dial("http://localhost:7545")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
		return err
	}

	// Create a new instance of the ipfs contract bound to a specific deployed contract
  	contract, err := NewIPFSStorage(common.HexToAddress("0xdFa8d5DF8Cf855d7A1135Bb01E14d00946B67443"), client)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of contract:%v\n")
		return err
	}

	privateKey, err := crypto.HexToECDSA("868d0895b14def4d83e7a55be9c9e6e0d076fbd9719861faa2663891c481e242")
	if err != nil {
		log.Fatal(err)
		return err
	}

	auth := bind.NewKeyedTransactor(privateKey)

	fmt.Println(contract.SetEntry(&bind.TransactOpts{
		From:auth.From,
		Signer:auth.Signer,
		Value: nil,
	}, multihash))
	return nil
}