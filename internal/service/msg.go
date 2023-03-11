package service

import (
	"strconv"
	"strings"
)

type Msg struct {
	Code    int
	Method  string
	URL     string
	Address string
}

func (m *Msg) Read(p []byte) (n int, err error) {
	buf := make([]byte, 0, len(p))
	for _, v := range p {
		if v != 0 {
			buf = append(buf, v)
		}
	}
	content := strings.Split(string(buf), "\n")

	m.Code, err = strconv.Atoi(content[0])
	if err != nil {
		return len(buf), err
	}
	m.Method = content[1]
	m.URL = content[2]
	m.Address = content[3]

	return len(buf), nil
}
