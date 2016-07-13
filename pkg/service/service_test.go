// Copyright (c) 2016 Kazumasa Kohtaka. All rights reserved.
// This file is available under the MIT license.

package service

import (
	"io/ioutil"
	"net/http"
	"testing"
)

type testCase struct {
	url      string
	expected string
}

func TestRun(t *testing.T) {
	svc := NewService()
	go svc.Run()

	cases := []testCase{
		{
			url:      "http://localhost:8080/entrypoint",
			expected: "Hello, Web App.",
		},
	}
	for _, c := range cases {
		resp, err := http.Get(c.url)
		if err != nil {
			t.Errorf("Failed to request: %s, %v", c.url, err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		actual := string(body)
		if actual != c.expected {
			t.Errorf("Failed to test, actual: %v, expected: %v", actual, c.expected)
		}
	}
}
