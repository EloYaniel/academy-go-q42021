package apiclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var instance *HttpApiClient

type HttpApiClient struct {
	client *http.Client
}

func GetHttpApiClientInstance() *HttpApiClient {
	once := sync.Once{}
	once.Do(func() {
		client := &http.Client{
			Timeout: time.Second * 10,
		}
		instance = &HttpApiClient{client: client}
	})

	return instance
}

func (api *HttpApiClient) Get(url string, params map[string]interface{}, response interface{}) error {
	resp, err := api.client.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return errors.New(fmt.Sprint("error reading body response:", err.Error()))
	}

	err = json.Unmarshal(buf, &response)
	if err != nil {
		return errors.New(fmt.Sprint("error parsing body response:", err.Error()))
	}

	return nil
}
