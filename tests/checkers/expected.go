package checkers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

const (
	KEY_STATUS = "expected-status"
	KEY_BODY   = "expected-body"
	KEY_FILTER = "expected-filter"
)

type Filter func(interface{}) interface{}

type Validator struct {
	report *Report
	code   int
	body   interface{}
	filter Filter
}

func Expected(statusCode int, body interface{}, prepare Filter) context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, KEY_STATUS, statusCode)
	if body != nil {
		ctx = context.WithValue(ctx, KEY_BODY, body)
	}
	if prepare != nil {
		ctx = context.WithValue(ctx, KEY_FILTER, prepare)
	}
	return ctx

}

func NewValidator(ctx context.Context, report *Report) *Validator {
	v := Validator{report: report}
	if ctx != nil {
		if ctx.Value(KEY_STATUS) != nil {
			v.code = ctx.Value(KEY_STATUS).(int)
			if ctx.Value(KEY_BODY) != nil {
				v.body = ctx.Value(KEY_BODY)
			}
			if ctx.Value(KEY_FILTER) != nil {
				v.filter = ctx.Value(KEY_FILTER).(Filter)
			}
		}
	}
	if v.filter == nil {
		v.filter = func(data interface{}) interface{} {
			return data
		}
	}
	return &v
}

func (self *Validator) validate(req *http.Request, res *http.Response) bool {
	defer func() {
		if r := recover(); r != nil {
			self.report.AddError(r)
		}
	}()
	if self.code != 0 {
		body := []byte{}
		if res.Body != nil {
			ibody := res.Body
			defer ibody.Close()
			var err error
			body, err = ioutil.ReadAll(ibody)
			if err != nil {
				panic(err)
			}
		}

		if res.StatusCode != self.code {
			message := fmt.Sprintf("Unexpected status code: %d (expected %d)", res.StatusCode, self.code)
			self.report.RoundTrip(req, res, self.Example(), &message)
			self.report.result = REPORT_FAILED
		}
		delta := GetDelta(body, self.body, self.filter)
		if (res.StatusCode != self.code) || (delta != "") {
			log.Println("----------------")
			log.Println(string(body))
			expected_json, _ := json.MarshalIndent(self.body, "", "  ")
			log.Println(string(expected_json))
			log.Println("++++++++++++++++")

			log.Println("Unexpected status code:", res.StatusCode, "!=", self.code, string(body))
			panic("Ops...")
			self.report.RoundTrip(req, res, self.Example(), &delta)
			self.report.result = REPORT_FAILED
		} else {
			self.report.RoundTrip(req, res, nil, nil)
		}

		if res.Body != nil {
			res.Body = ioutil.NopCloser(bytes.NewReader(body))
		}
	}
	return true
}

func (self *Validator) Example() *http.Response {
	if self.body == "" {
		return nil
	}
	json_body := ToJson(self.body)
	return &http.Response{
		StatusCode: self.code,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(json_body))),
	}
}

func (self *Validator) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Println(*req)
	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		self.report.AddError(err)
		return nil, err
	}
	if self.validate(req, res) {
		return res, nil
	}
	return nil, errors.New("Unexpected error")
}

func ToJson(obj interface{}) string {
	data, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func GetDiff(actual string, expected string) string {
	if actual == expected {
		return ""
	}
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(actual, expected, false)
	return dmp.DiffPrettyText(diffs)
}

func GetDelta(data []byte, expected interface{}, prepare Filter) string {
	if expected == nil {
		return ""
	}
	expected_json := ToJson(prepare(expected))
	var actual interface{} = reflect.New(reflect.TypeOf(expected).Elem()).Interface()
	if err := json.Unmarshal(data, actual); err != nil {
		return GetDiff(string(data), expected_json)
	}

	actual_json := ToJson(prepare(actual))
	if expected_json == actual_json {
		return ""
	}
	return GetDiff(actual_json, expected_json)
}
