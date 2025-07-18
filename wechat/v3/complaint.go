package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-pay/util/js"
	"github.com/w6xian/gopay"
)

// 创建投诉通知回调地址API
// Code = 0 is success
func (c *ClientV3) V3ComplaintNotifyUrlCreate(ctx context.Context, url string) (wxRsp *ComplaintNotifyUrlRsp, err error) {
	bm := make(gopay.BodyMap)
	bm.Set("url", url)
	authorization, err := c.authorization(MethodPost, v3ComplaintNotifyUrlCreate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ComplaintNotifyUrlCreate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ComplaintNotifyUrlRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	wxRsp.Response = new(ComplaintNotifyUrl)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询投诉通知回调地址API
// Code = 0 is success
func (c *ClientV3) V3ComplaintNotifyUrlQuery(ctx context.Context) (wxRsp *ComplaintNotifyUrlRsp, err error) {
	authorization, err := c.authorization(MethodGet, v3ComplaintNotifyUrlQuery, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, v3ComplaintNotifyUrlQuery, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ComplaintNotifyUrlRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	wxRsp.Response = new(ComplaintNotifyUrl)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 更新投诉通知回调地址API
// Code = 0 is success
func (c *ClientV3) V3ComplaintNotifyUrlUpdate(ctx context.Context, url string) (wxRsp *ComplaintNotifyUrlRsp, err error) {
	bm := make(gopay.BodyMap)
	bm.Set("url", url)
	authorization, err := c.authorization(MethodPut, v3ComplaintNotifyUrlUpdate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPut(ctx, bm, v3ComplaintNotifyUrlUpdate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ComplaintNotifyUrlRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	wxRsp.Response = new(ComplaintNotifyUrl)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 删除投诉通知回调地址API
// Code = 0 is success
func (c *ClientV3) V3ComplaintNotifyUrlDelete(ctx context.Context) (wxRsp *EmptyRsp, err error) {
	authorization, err := c.authorization(MethodDelete, v3ComplaintNotifyUrlDelete, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdDelete(ctx, nil, v3ComplaintNotifyUrlDelete, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 商户上传反馈图片
// 注意：图片不能超过2MB
// Code = 0 is success
func (c *ClientV3) V3ComplaintUploadImage(ctx context.Context, fileName, fileSha256 string, img *gopay.File) (wxRsp *MediaUploadRsp, err error) {
	bmFile := make(gopay.BodyMap)
	bmFile.Set("filename", fileName).Set("sha256", fileSha256)
	authorization, err := c.authorization(MethodPost, v3ComplaintUploadImage, bmFile)
	if err != nil {
		return nil, err
	}

	bm := make(gopay.BodyMap)
	bm.SetBodyMap("meta", func(bm gopay.BodyMap) {
		bm.Set("filename", fileName).Set("sha256", fileSha256)
	}).SetFormFile("file", img)
	res, si, bs, err := c.doProdPostFile(ctx, bm, v3ComplaintUploadImage, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &MediaUploadRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	wxRsp.Response = new(MediaUpload)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 商户反馈图片请求
// Code = 0 is success
func (c *ClientV3) V3ComplaintImage(ctx context.Context, mediaId string) (wxRsp *ComplaintImageRsp, err error) {
	uri := fmt.Sprintf(v3ComplaintImage, mediaId)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	contentType := res.Header.Get("Content-Type")
	wxRsp = &ComplaintImageRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if strings.HasPrefix(contentType, "image/") {
		wxRsp.ImageData = bs
		return wxRsp, nil
	}
	wxRsp.Response = new(ComplaintImage)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询投诉单列表API
// Code = 0 is success
func (c *ClientV3) V3ComplaintList(ctx context.Context, bm gopay.BodyMap) (wxRsp *ComplaintListRsp, err error) {
	uri := v3ComplaintList + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ComplaintListRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	wxRsp.Response = new(ComplaintList)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询投诉协商历史API
// Code = 0 is success
func (c *ClientV3) V3ComplaintNegotiationHistory(ctx context.Context, complaintId string, bm gopay.BodyMap) (wxRsp *ComplaintNegotiationHistoryRsp, err error) {
	uri := fmt.Sprintf(v3ComplaintNegotiationHistory, complaintId) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ComplaintNegotiationHistoryRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	wxRsp.Response = new(ComplaintNegotiationHistory)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询投诉单详情API
// Code = 0 is success
func (c *ClientV3) V3ComplaintDetail(ctx context.Context, complaintId string) (wxRsp *ComplaintDetailRsp, err error) {
	url := fmt.Sprintf(v3ComplaintDetail, complaintId)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ComplaintDetailRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	wxRsp.Response = new(ComplaintDetail)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 回复用户API
// Code = 0 is success
func (c *ClientV3) V3ComplaintResponse(ctx context.Context, complaintId string, bm gopay.BodyMap) (wxRsp *EmptyRsp, err error) {
	url := fmt.Sprintf(v3ComplaintResponse, complaintId)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 反馈处理完成API
// Code = 0 is success
func (c *ClientV3) V3ComplaintComplete(ctx context.Context, complaintId string, bm gopay.BodyMap) (wxRsp *EmptyRsp, err error) {
	url := fmt.Sprintf(v3ComplaintComplete, complaintId)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 更新退款审批结果
// Code = 0 is success
func (c *ClientV3) V3ComplaintUpdateRefundProgress(ctx context.Context, complaintId string, bm gopay.BodyMap) (wxRsp *EmptyRsp, err error) {
	url := fmt.Sprintf(v3ComplaintUpdateRefundProgress, complaintId)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}
