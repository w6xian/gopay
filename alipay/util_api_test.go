package alipay

import (
	"testing"

	"github.com/go-pay/xlog"
	"github.com/w6xian/gopay"
	"github.com/w6xian/gopay/alipay/cert"
)

func TestClient_OpenAuthTokenApp(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("grant_type", "authorization_code").
		Set("code", "866185490c4e40efa9f71efea6766X02")

	// 发起请求
	aliRsp, err := client.OpenAuthTokenApp(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_OpenAuthTokenAppQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set(AppAuthToken, "202212BB9e1cd0c2e0ab489393aa2570ec4faX87")

	// 发起请求
	aliRsp, err := client.OpenAuthTokenAppQuery(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_UserInfoAuth(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	// 接口权限值，目前只支持auth_user和auth_base两个值。具体说明看文档介绍
	bm.Set("scopes", []string{"auth_user"}).
		Set("state", "init")

	// 发起请求
	html, err := client.UserInfoAuth(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debugf("html: %s", string(html))
}

func TestClient_UserInfoShare(t *testing.T) {
	// 发起请求
	aliRsp, err := client.UserInfoShare(ctx, "authbseBb6dd42c0d93a47dfa1b23aa725778X18")
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)

	// 同步返回验签
	ok, err := VerifySyncSignWithCert(cert.AlipayPublicContentRSA2, aliRsp.SignData, aliRsp.Sign)
	if err != nil {
		xlog.Error(err)
	}
	xlog.Debug("ok:", ok)
}

func TestClient_PublicCertDownload(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("alipay_cert_sn", "52c63ed47b57c049b4bc9bea9da02c2a")

	// 发起请求
	aliRsp, err := client.PublicCertDownload(ctx, bm)
	if err != nil {
		//xlog.Errorf("client.UserInfoShare(),error:%+v", err)
		return
	}
	xlog.Debugf("aliRsp.Response.AlipayCertContent:\n %s", aliRsp.Response.AlipayCertContent)
}
