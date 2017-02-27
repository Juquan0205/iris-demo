package util

import (
	"errors"
	"github.com/parnurzeal/gorequest"
)

func HttpGet(url string, params interface{}) (string, error) {
	Logger.Warningln("get请求链接地址:", url, params)
	request := gorequest.New()
	resp, body, errs := request.Get(url).Query(params).End()
	for _, err := range errs {
		if err != nil {
			Logger.Warningln("get请求链接地址:", resp.Request.URL, params, "请求出错: err=", errs)
			return "", errors.New("请求第三方网络发生错误")
		}
	}
	Logger.Warningln("请求链接地址:", resp.Request.URL, params, "返回的结果是: ", body)
	return body, nil
}

func HttpPostForm(url string, params interface{}) (string, error) {
	Logger.Warningln("postForm请求链接地址:", url, params)
	request := gorequest.New()
	resp, body, errs := request.Post(url).Type("multipart").Send(params).End()
	for _, err := range errs {
		if err != nil {
			Logger.Warningln("postForm请求链接地址:", resp.Request.URL, params, "请求出错: err=", errs)
			return "", errors.New("请求第三方网络发生错误")
		}
	}
	Logger.Warningln("请求链接地址:", resp.Request.URL, params, "返回的结果是: ", body)
	return body, nil
}

func HttpPostJson(url string, params interface{}) (string, error) {
	Logger.Warningln("postJson请求链接地址:", url, params)
	request := gorequest.New()
	resp, body, errs := request.Post(url).Send(params).End()
	for _, err := range errs {
		if err != nil {
			Logger.Warningln("postJson请求链接地址:", resp.Request.URL, params, "请求出错: err=", errs)
			return "", errors.New("请求第三方网络发生错误")
		}
	}
	Logger.Warningln("请求链接地址:", resp.Request.URL, params, "返回的结果是: ", body)
	return body, nil
}
