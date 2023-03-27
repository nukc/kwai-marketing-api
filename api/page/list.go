package lp

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/page"
)

// PageList 获取魔力建站落地页组信息列表
func PageList(clt *core.SDKClient, accessToken string, req *page.ListRequest) (*page.ListResponse, error) {
	var resp page.ListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
