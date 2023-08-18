package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// in the case of GET, the parameter queryParameters is transferred to the URL as query parameters
// in the case of POST, the parameter body, an io.Reader, is used
func MakeHTTPRequest[T any](fullUrl string, httpMethod string, headers map[string]string, queryParameters url.Values, body io.Reader, responseType T) (T, http.Header, error) {
	client := http.Client{}
	u, err := url.Parse(fullUrl)
	if err != nil {
		return responseType, nil, err
	}

	// if it's a GET, we need to append the query parameters.
	if httpMethod == "GET" {
		q := u.Query()

		for k, v := range queryParameters {
			// this depends on the type of api, you may need to do it for each of v
			q.Set(k, strings.Join(v, ","))
		}
		// set the query to the encoded parameters
		u.RawQuery = q.Encode()
	}

	// regardless of GET or POST, we can safely add the body
	req, err := http.NewRequest(httpMethod, u.String(), body)
	if err != nil {
		return responseType, nil, err
	}

	// for each header passed, add the header value to the request
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	// Set the token in the header
	//req.Header.Add("X-Auth-Token", c.Token)
	// finally, do the request
	res, err := client.Do(req)
	if err != nil {
		return responseType, nil, err
	}

	//appLogger.Infof("response object from request call: %+v", res)
	if res == nil {
		return responseType, nil, fmt.Errorf("error: calling %s returned empty response", u.String())
	}

	responseData, err := io.ReadAll(res.Body)

	if err != nil {
		return responseType, nil, err
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusCreated, http.StatusOK:
		var responseObject T
		err = json.Unmarshal(responseData, &responseObject)

		if err != nil {

			return responseType, res.Header, err
		}

		return responseObject, res.Header, err
	default:
		return responseType, res.Header, fmt.Errorf("error calling %s:\nstatus: %s\nresponseData: %s", u.String(), res.Status, responseData)
	}

}

