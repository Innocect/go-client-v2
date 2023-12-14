package httpclientutils

import (
	"clientv1/data"
	"fmt"
	"io/ioutil"
	"os"

	commonpb "golang.frontdoorhome.com/software/protos/go/common"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestGetApiCall() {
	for _, api := range data.APIInfo {
		if api.Ignore {
			continue
		}

		errResp := &commonpb.ErrorResponse{}
		APIURL := fmt.Sprintf("%s%s", data.GetEnvURL(api.Env), api.Path)
		token := os.Getenv("DEV_TOKEN")

		res := ApiCall(APIURL, api.Request, api.Response, errResp, token, api.MethodType)

		protoJson := protojson.MarshalOptions{
			EmitUnpopulated: false,
			Multiline:       true,
		}

		js := protoJson.Format(res)
		err := ioutil.WriteFile(api.Name+".json", []byte(js), 0644)
		if err != nil {
			panic(err)
		}

	}
}
