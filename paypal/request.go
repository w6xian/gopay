package paypal

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/w6xian/gopay"
	"github.com/w6xian/gopay/pkg/xhttp"
)

type HeaderKeyType string

func (c *Client) doPayPalGet(ctx context.Context, uri string) (res *http.Response, bs []byte, err error) {
	var url = c.baseUrlProd + uri
	if !c.IsProd {
		url = c.baseUrlSandbox + uri
	}
	req := c.hc.Req() // default json
	c.setPaypalHeader(ctx, req)
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("PayPal_Url: %s", url)
		c.logger.Debugf("PayPal_Req_Headers: %#v", req.Header)
	}
	res, bs, err = req.Get(url).EndBytes(ctx)
	if err != nil {
		return nil, nil, err
	}
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("PayPal_Response: %d > %s", res.StatusCode, string(bs))
		c.logger.Debugf("PayPal_Rsp_Headers: %#v", res.Header)
	}
	return res, bs, nil
}

func (c *Client) DoPayPalPost(ctx context.Context, bm gopay.BodyMap, path string) (res *http.Response, bs []byte, err error) {
	return c.doPayPalPost(ctx, bm, path)
}

func (c *Client) doPayPalPost(ctx context.Context, bm gopay.BodyMap, path string) (res *http.Response, bs []byte, err error) {
	var url = c.baseUrlProd + path
	if !c.IsProd {
		url = c.baseUrlSandbox + path
	}
	req := c.hc.Req() // default json
	c.setPaypalHeader(ctx, req)
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("PayPal_Url: %s", url)
		c.logger.Debugf("PayPal_Req_Body: %s", bm.JsonBody())
		c.logger.Debugf("PayPal_Req_Headers: %#v", req.Header)
	}
	res, bs, err = req.Post(url).SendBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil, nil, err
	}
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("PayPal_Response: %d > %s", res.StatusCode, string(bs))
		c.logger.Debugf("PayPal_Rsp_Headers: %#v", res.Header)
	}
	return res, bs, nil
}

func (c *Client) doPayPalPut(ctx context.Context, bm gopay.BodyMap, path string) (res *http.Response, bs []byte, err error) {
	var url = c.baseUrlProd + path
	if !c.IsProd {
		url = c.baseUrlSandbox + path
	}
	req := c.hc.Req() // default json
	c.setPaypalHeader(ctx, req)
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("PayPal_Url: %s", url)
		c.logger.Debugf("PayPal_Req_Body: %s", bm.JsonBody())
		c.logger.Debugf("PayPal_Req_Headers: %#v", req.Header)
	}
	res, bs, err = req.Put(url).SendBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil, nil, err
	}
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("PayPal_Response: %d > %s", res.StatusCode, string(bs))
		c.logger.Debugf("PayPal_Rsp_Headers: %#v", res.Header)
	}
	return res, bs, nil
}

func (c *Client) doPayPalPatch(ctx context.Context, patchs []*Patch, path string) (res *http.Response, bs []byte, err error) {
	var url = c.baseUrlProd + path
	if !c.IsProd {
		url = c.baseUrlSandbox + path
	}
	req := c.hc.Req() // default json
	c.setPaypalHeader(ctx, req)
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("PayPal_Url: %s", url)
		body, _ := json.Marshal(patchs)
		c.logger.Debugf("PayPal_Req_Body: %s", string(body))
		c.logger.Debugf("PayPal_Req_Headers: %#v", req.Header)
	}
	res, bs, err = req.Patch(url).SendStruct(patchs).EndBytes(ctx)
	if err != nil {
		return nil, nil, err
	}
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("PayPal_Response: %d > %s", res.StatusCode, string(bs))
		c.logger.Debugf("PayPal_Headers: %#v", res.Header)
	}
	return res, bs, nil
}

func (c *Client) doPayPalDelete(ctx context.Context, path string) (res *http.Response, bs []byte, err error) {
	var url = c.baseUrlProd + path
	if !c.IsProd {
		url = c.baseUrlSandbox + path
	}
	req := c.hc.Req() // default json
	c.setPaypalHeader(ctx, req)
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("PayPal_Url: %s", url)
		c.logger.Debugf("PayPal_Req_Headers: %#v", req.Header)
	}
	res, bs, err = req.Delete(url).EndBytes(ctx)
	if err != nil {
		return nil, nil, err
	}
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("PayPal_Response: %d > %s", res.StatusCode, string(bs))
		c.logger.Debugf("PayPal_Rsp_Headers: %#v", res.Header)
	}
	return res, bs, nil
}

// setPaypalHeader 给paypal设定header  可以增加paypal的一些指定的header 示例： 'Prefer': 'return=representation',
func (c *Client) setPaypalHeader(ctx context.Context, req *xhttp.Request) {
	req.Header.Set(HeaderAuthorization, AuthorizationPrefixBearer+c.AccessToken)
	req.Header.Set("Accept", "*/*")

	// 尝试从 context 中获取header 如果数据为空则不设置
	if len(c.headerKeyMap) > 0 {
		for headerKey, defaultVal := range c.headerKeyMap {
			req.Header.Set(headerKey, defaultVal)
			if value := ctx.Value(headerKey); value != nil {
				// 存在自定义，则覆盖默认值
				if v, ok := value.(string); ok && v != "" {
					req.Header.Set(headerKey, v)
				}
			}
		}
	}
}
