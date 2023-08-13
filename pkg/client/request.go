package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func MyRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Test", "true")
	return req, nil
}

// Implement functions for creating and sending specific API requests (e.g., GET and POST).
// This file can also include utility functions for setting headers, authentication, etc.
func Get[T any](ctx context.Context, url string) (T, error) {
	var m T
	r, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return m, err
	}
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return m, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return m, err
	}
	return parseJSON[T](body)
}

/* func PPost(url string, data any) (something, error){
	b, err := toJSON(data)
	if err != nil {
		return b, err
	}
	byteReader := bytes.NewReader(b)
	r, err := http.NewRequest("POST", url, byteReader)
	if err != nil {
		return r, err
	}
	// TODO add setting header function for tokens
	r.Header.Add("Content-Type", "application/json")

	// Send the request
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return res, err
	}
	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()


} */
func Post[T any](ctx context.Context, url string, data any) (T, error) {
	var m T
	b, err := toJSON(data)
	if err != nil {
		return m, err
	}
	byteReader := bytes.NewReader(b)
	r, err := http.NewRequestWithContext(ctx, "POST", url, byteReader)
	if err != nil {
		return m, err
	}
	// Important to set
	r.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return m, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return m, err
	}
	return parseJSON[T](body)
}

func toJSON(T any) ([]byte, error) {
	return json.Marshal(T)
}

func parseJSON[T any](s []byte) (T, error) {
	var r T
	if err := json.Unmarshal(s, &r); err != nil {
		return r, err
	}
	return r, nil
}

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

	// optional: log the request for easier stack tracing
	//appLogger.Infof("Logging request: %s %s\n", httpMethod, req.URL.String())

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

	// Not a good idea soince we can get different repsones according to the request performed
	// Idea extract this functionality if needed and call it after beeing returned
	/* if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusOK {
		appLogger.Info(res.StatusCode)
		return responseType, res.Header, fmt.Errorf("error calling %s:\nstatus: %s\nresponseData: %s", u.String(), res.Status, responseData)
	}  */
	switch res.StatusCode {
	case http.StatusCreated, http.StatusOK:
		var responseObject T
		err = json.Unmarshal(responseData, &responseObject)
		// Now rsp should have the response data, print it:
		//fmt.Printf("Payload converted, JSON: %s\n", responseData)
		if err != nil {

			return responseType, res.Header, err
		}

		return responseObject, res.Header, err
	default:
		return responseType, res.Header, fmt.Errorf("error calling %s:\nstatus: %s\nresponseData: %s", u.String(), res.Status, responseData)
	}

}
