package httpclient

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type xWwwFormUrlencodedHelperOptions struct {
	Data   url.Values
	Header map[string]string
}

func NewXWwwFormUrlencodedHelperOptions() *xWwwFormUrlencodedHelperOptions {
	// 옵셔널 값 기본 값 설정
	return &xWwwFormUrlencodedHelperOptions{
		Data: url.Values{},
	}
}

// [http Content-Type application/x-www-form-urlencoded 를 쉽게 요청할 수 있는 헬퍼 함수이며 해당 글을 참고하여 제작되었습니다.](https://golang.cafe/blog/how-to-make-http-url-form-encoded-request-golang.html)
// [자세한 설명](https://golangbyexample.com/http-client-urlencoded-body-go)
func (o *xWwwFormUrlencodedHelperOptions) Req(method Method, endpoint string) (body []byte, err error) {
	// URL-encoded payload
	req, err := http.NewRequest(method.String(), endpoint, strings.NewReader(o.Data.Encode()))
	if err != nil {
		return nil, errors.Wrap(err, "http request 생성 실패")
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	req.Header.Add("Content-Length", strconv.Itoa(len(o.Data.Encode())))
	for i, v := range o.Header {
		req.Header.Add(i, v)
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "http request 실패")
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "http response body read 실패")
	}

	return body, nil
}
