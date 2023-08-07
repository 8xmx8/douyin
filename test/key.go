package test

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

var byteAdata = make([]byte, 0, 50)

func init() {
	byteAdata = append(byteAdata, "video:favorite_count/"...)
}

func ByteA(vid int64) string {
	t := make([]byte, 0, 50)
	copy(t, byteAdata)
	t = append(t, strconv.FormatInt(vid, 10)...)
	return string(t)
}

func ByteB(vid int64) string {
	t := make([]byte, 0, 50)
	copy(t, byteAdata)
	t = append(t, strconv.FormatInt(vid, 16)...)
	return string(t)
}

func ByteC(vid int64) string {
	t := make([]byte, 0, 50)
	copy(t, byteAdata)
	t = append(t, strconv.FormatInt(vid, 36)...)
	return string(t)
}

func Builder(vid int64) string {
	var builder strings.Builder

	builder.Grow(50)
	builder.WriteString("video:favorite_count/")
	builder.WriteString(strconv.FormatInt(vid, 10))
	return builder.String()
}

func Buffer(vid int64) string {
	var buffer bytes.Buffer

	buffer.Grow(50)
	buffer.WriteString("video:favorite_count/")
	buffer.WriteString(strconv.FormatInt(vid, 10))
	return buffer.String()
}

func Printf(vid int64) string {
	return fmt.Sprintf("video:favorite_count/%d", vid)
}
