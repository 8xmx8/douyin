package server

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func videoGet(t *testing.T, data *Data) map[string]any {
	return newReq(t, "GET", "/douyin/feed?repeat=true", data)
}

func videoAction(t *testing.T, data *Data) map[string]any {
	return newReq(t, "POST", "/douyin/publish/action/", data)
}

func videoActionurl(t *testing.T, data *Data) map[string]any {
	return newReq(t, "POST", "/douyin/publish/actionUrl/", data)
}

func videoList(t *testing.T, data *Data) map[string]any {
	return newReq(t, "GET", "/douyin/publish/list/", data)
}

func videoFollowList(t *testing.T, data *Data) map[string]any {
	return newReq(t, "GET", "/douyin/publish/follow/", data)
}

func userRegister(t *testing.T, data *Data) map[string]any {
	return newReq(t, "POST", "/douyin/user/register/", data)
}

func userLogin(t *testing.T, data *Data) map[string]any {
	return newReq(t, "POST", "/douyin/user/login/", data)
}

func userInfo(t *testing.T, data *Data) map[string]any {
	return newReq(t, "GET", "/douyin/user/", data)
}

func favoriteAction(t *testing.T, data *Data) map[string]any {
	return newReq(t, "POST", "/douyin/favorite/action/", data)
}

func favoriteList(t *testing.T, data *Data) map[string]any {
	return newReq(t, "GET", "/douyin/favorite/list/", data)
}

func commentAction(t *testing.T, data *Data) map[string]any {
	return newReq(t, "POST", "/douyin/comment/action/", data)
}

func commentList(t *testing.T, data *Data) map[string]any {
	return newReq(t, "GET", "/douyin/comment/list/", data)
}

func relationAction(t *testing.T, data *Data) map[string]any {
	return newReq(t, "POST", "/douyin/relation/action/", data)
}

func relationFollowList(t *testing.T, data *Data) map[string]any {
	return newReq(t, "GET", "/douyin/relation/follow/list/", data)
}

func relationFollowerList(t *testing.T, data *Data) map[string]any {
	return newReq(t, "GET", "/douyin/relation/follower/list/", data)
}

func relationFriendList(t *testing.T, data *Data) map[string]any {
	return newReq(t, "GET", "/douyin/relation/friend/list/", data)
}

func messageChat(t *testing.T, data *Data) map[string]any {
	return newReq(t, "GET", "/douyin/message/chat/", data)
}

func messageAction(t *testing.T, data *Data) map[string]any {
	return newReq(t, "POST", "/douyin/message/action/", data)
}

type Data struct {
	Json  H
	Form  *F
	Query S
}
type (
	H map[string]any
	S map[string]string
	F struct {
		r io.Reader
		w *multipart.Writer
	}
)

func newReq(t *testing.T, mod, url string, data *Data) map[string]any {
	w := httptest.NewRecorder()
	var req *http.Request
	if data != nil && data.Json != nil {
		jsonData, _ := json.Marshal(data.Json)
		req, _ = http.NewRequest(mod, url, bytes.NewBuffer(jsonData))
	} else if data != nil && data.Form != nil {
		req, _ = http.NewRequest(mod, url, data.Form.r)
		req.Header.Set("Content-Type", data.Form.w.FormDataContentType())
	} else {
		req, _ = http.NewRequest(mod, url, nil)
	}
	if data != nil && data.Query != nil {
		q := req.URL.Query()
		for k, v := range data.Query {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	code, msg, _data := unmarshal(w.Body.Bytes())
	assert.Equal(t, float64(0), code, msg, _data)
	return _data
}

func unmarshal(body []byte) (float64, string, map[string]any) {
	m := make(map[string]any)
	if err := json.Unmarshal(body, &m); err != nil {
		return 1, "json 解析失败", m
	}
	if v, ok := m["status_code"]; ok {
		if v != "0" && v != float64(0) && v != 0 {
			return 1, "错误的status_code", m
		}
	} else {
		return 1, "无 status_code", m
	}
	return m["status_code"].(float64), "", m
}
