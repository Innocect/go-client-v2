package httpclientutils

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"google.golang.org/protobuf/proto"
)

var URL, URLNew, URLNewV1, URLNewV2 string

var client http.Client

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = http.Client{Timeout: time.Duration(20) * time.Second, Transport: tr}
}

// func APICall(apiName string, path string, request proto.Message, response proto.Message, scenario string, env Environment) *csvutils.APIComparisonRow {
// 	errResp := &commonpb.ErrorResponse{}
// 	APIURL := fmt.Sprintf("%s%s", getEnvURL(env), path)
// 	token := os.Getenv("token")

// 	resp, body := apiCall(APIURL, request, response, errResp, token)
// 	if resp != nil && resp.StatusCode != 404 {
// 		body = "-"
// 	}
// 	if resp != nil {
// 		apiComparisonRow := &csvutils.APIComparisonRow{Name: apiName, Path: path, Version: getEnvName(env), Case: scenario, Success: false, StatusCode: resp.StatusCode, ProtoResponse: proto.Clone(response), ErrResponse: errResp, StringResponse: body}
// 		return apiComparisonRow
// 	}
// 	return nil
// }

func ApiCall(URL string, request proto.Message, response proto.Message, errResponse proto.Message, token, method string) proto.Message {
	urlWithProtoBody := fmt.Sprintf("%s?proto_body=%s", URL, getEncodedProtoBody(request))
	fmt.Println(urlWithProtoBody)
	req, err := http.NewRequest(method, urlWithProtoBody, nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return nil
	}
	req.Header.Add("content-type", `application/x-protobuf`)
	if len(token) > 0 {
		req.Header.Add("Authorization", "Bearer "+token)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error Response %s", err)
		return nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil
	}
	proto.Unmarshal(body, response)
	proto.Unmarshal(body, errResponse)
	return response
}

func getEncodedProtoBody(request proto.Message) string {
	marshalledRequest, _ := proto.Marshal(request)
	encoded := base64.RawURLEncoding.EncodeToString(marshalledRequest)
	return encoded
}
