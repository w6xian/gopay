package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/util/js"
	"github.com/w6xian/gopay"
)

// 合单下单-APP
// Code = 0 is success
func (c *ClientV3) V3CombineTransactionApp(ctx context.Context, bm gopay.BodyMap) (wxRsp *PrepayRsp, err error) {
	if bm.GetString("combine_mchid") == gopay.NULL {
		bm.Set("combine_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3CombinePayApp, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3CombinePayApp, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PrepayRsp{Code: Success, SignInfo: si, Response: new(Prepay)}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 合单下单-JSAPI/小程序
// Code = 0 is success
func (c *ClientV3) V3CombineTransactionJsapi(ctx context.Context, bm gopay.BodyMap) (wxRsp *PrepayRsp, err error) {
	if bm.GetString("combine_mchid") == gopay.NULL {
		bm.Set("combine_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3CombinePayJsapi, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3CombinePayJsapi, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PrepayRsp{Code: Success, SignInfo: si, Response: new(Prepay)}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 合单下单-NATIVE
// Code = 0 is success
func (c *ClientV3) V3CombineTransactionNative(ctx context.Context, bm gopay.BodyMap) (wxRsp *NativeRsp, err error) {
	if bm.GetString("combine_mchid") == gopay.NULL {
		bm.Set("combine_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3CombineNative, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3CombineNative, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &NativeRsp{Code: Success, SignInfo: si, Response: new(Native)}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 合单下单-H5
// Code = 0 is success
func (c *ClientV3) V3CombineTransactionH5(ctx context.Context, bm gopay.BodyMap) (wxRsp *H5Rsp, err error) {
	if bm.GetString("combine_mchid") == gopay.NULL {
		bm.Set("combine_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3CombinePayH5, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3CombinePayH5, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &H5Rsp{Code: Success, SignInfo: si, Response: new(H5Url)}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 合单QQ小程序下单-H5
// Code = 0 is success
func (c *ClientV3) V3CombineQQTransactionH5(ctx context.Context, qqAppid, accessToken, realNotifyUrl string, bm gopay.BodyMap) (wxRsp *H5Rsp, err error) {
	if bm.GetString("combine_mchid") == gopay.NULL {
		bm.Set("combine_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3CombinePayH5, bm)
	if err != nil {
		return nil, err
	}
	path := "/wxpay/v3/combine-transactions/h5?appid=" + qqAppid + "&access_token=" + accessToken + "&real_notify_url=" + realNotifyUrl
	res, si, bs, err := c.doProdPostWithHost(ctx, bm, "https://api.q.qq.com", path, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &H5Rsp{Code: Success, SignInfo: si, Response: new(H5Url)}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 合单查询订单
// Code = 0 is success
func (c *ClientV3) V3CombineQueryOrder(ctx context.Context, traderNo string) (wxRsp *CombineQueryOrderRsp, err error) {
	uri := fmt.Sprintf(v3CombineQuery, traderNo)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &CombineQueryOrderRsp{Code: Success, SignInfo: si, Response: new(CombineQueryOrder)}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 合单关闭订单
// Code = 0 is success
func (c *ClientV3) V3CombineCloseOrder(ctx context.Context, tradeNo string, bm gopay.BodyMap) (wxRsp *EmptyRsp, err error) {
	url := fmt.Sprintf(v3CombineClose, tradeNo)
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
