package crowdin

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type postOptions struct {
	urlStr      string
	params      map[string]string
	paramsArray map[string][]string
	files       map[string]string
}

type getOptions struct {
	urlStr string
	params map[string]string
}

// params - extra params
// fileNames - key = dir
func (crowdin *Crowdin) post(options *postOptions) ([]byte, error) {

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	if options.params != nil {
		for k, v := range options.params {
			fw, err := writer.CreateFormField(k)
			if err != nil {
				return nil, err
			}
			if _, err = fw.Write([]byte(v)); err != nil {
				return nil, err
			}
		}
	}

	if options.paramsArray != nil {
		for k, arr := range options.paramsArray {
			for _, v := range arr {
				fw, err := writer.CreateFormField(k)
				if err != nil {
					return nil, err
				}
				if _, err = fw.Write([]byte(v)); err != nil {
					return nil, err
				}
			}
		}
	}

	if options.files != nil {
		for key, filePath := range options.files {
			file, err := os.Open(filePath)
			if err != nil {
				return nil, err
			}

			defer file.Close()

			fw, err := writer.CreateFormFile(key, filePath)
			if err != nil {
				return nil, err
			}
			if _, err = io.Copy(fw, file); err != nil {
				return nil, err
			}

		}
	}

	writer.Close()

	req, err := http.NewRequest("POST", options.urlStr, &buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	response, err := crowdin.config.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return bodyResponse, APIError{What: fmt.Sprintf("Status code: %v", response.StatusCode)}
	}

	return bodyResponse, nil
}

func (crowdin *Crowdin) get(options *getOptions) ([]byte, error) {

	response, err := crowdin.getResponse(options)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return bodyResponse, APIError{What: fmt.Sprintf("Status code: %v", response.StatusCode)}
	}

	return bodyResponse, nil
}

func (crowdin *Crowdin) getResponse(options *getOptions) (*http.Response, error) {

	if options != nil && options.params != nil {
		for k, v := range options.params {
			options.urlStr += "&" + k + "=" + v
		}
	}

	req, err := http.NewRequest("GET", options.urlStr, nil)
	if err != nil {
		return nil, err
	}

	response, err := crowdin.config.client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (crowdin *Crowdin) log(a interface{}) {
	if crowdin.debug {
		log.Println(a)
		if crowdin.logWriter != nil {
			timestamp := time.Now().Format(time.RFC3339)
			msg := fmt.Sprintf("%v: %v", timestamp, a)
			fmt.Fprintln(crowdin.logWriter, msg)
		}
	}
}

// APIError holds data of errors returned from the API.
type APIError struct {
	What string
}

func (e APIError) Error() string {
	return fmt.Sprintf("%v", e.What)
}
