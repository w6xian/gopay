package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/w6xian/gopay"
)

// alipay.trade.royalty.relation.bind(分账关系绑定)
// 文档地址：https://opendocs.alipay.com/open/c21931d6_alipay.trade.royalty.relation.bind
func (a *Client) TradeRelationBind(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeRelationBindResponse, err error) {
	err = bm.CheckEmptyError("receiver_list", "out_request_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.royalty.relation.bind"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeRelationBindResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.trade.royalty.relation.unbind(分账关系解绑)
// 文档地址：https://opendocs.alipay.com/open/3613f4e1_alipay.trade.royalty.relation.unbind
func (a *Client) TradeRelationUnbind(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeRelationUnbindResponse, err error) {
	err = bm.CheckEmptyError("receiver_list", "out_request_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.royalty.relation.unbind"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeRelationUnbindResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.trade.royalty.relation.batchquery(分账关系查询)
// 文档地址：https://opendocs.alipay.com/open/1860be54_alipay.trade.royalty.relation.batchquery
func (a *Client) TradeRelationBatchQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeRelationBatchQueryResponse, err error) {
	err = bm.CheckEmptyError("out_request_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.royalty.relation.batchquery"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeRelationBatchQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.trade.settle.confirm(统一收单确认结算接口)
// 文档地址：https://opendocs.alipay.com/open/028xqy
func (a *Client) TradeSettleConfirm(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeSettleConfirmResponse, err error) {
	err = bm.CheckEmptyError("out_request_no", "trade_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.settle.confirm"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeSettleConfirmResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.trade.order.settle(统一收单交易结算接口)
// 文档地址：https://opendocs.alipay.com/open/c3b24498_alipay.trade.order.settle
func (a *Client) TradeOrderSettle(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeOrderSettleResponse, err error) {
	err = bm.CheckEmptyError("out_request_no", "trade_no", "royalty_parameters")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.order.settle"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeOrderSettleResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.trade.order.settle.query(交易分账查询接口)
// 文档地址：https://opendocs.alipay.com/open/02pj6l
func (a *Client) TradeOrderSettleQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeOrderSettleQueryResponse, err error) {
	err = bm.CheckEmptyError("settle_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.order.settle.query"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeOrderSettleQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.trade.order.onsettle.query(分账剩余金额查询)
// 文档地址：https://opendocs.alipay.com/open/d87dc009_alipay.trade.order.onsettle.query
func (a *Client) TradeOrderOnSettleQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeOrderOnSettleQueryResponse, err error) {
	err = bm.CheckEmptyError("trade_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.order.onsettle.query"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeOrderOnSettleQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.trade.royalty.rate.query(分账比例查询)
// 文档地址：https://opendocs.alipay.com/open/6f314ee9_alipay.trade.royalty.rate.query
func (a *Client) TradeRoyaltyRateQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeRoyaltyRateQueryResponse, err error) {
	err = bm.CheckEmptyError("out_request_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.royalty.rate.query"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeRoyaltyRateQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.security.risk.customerrisk.send(商户数据同步)
// 文档地址：https://opendocs.alipay.com/open/02qth4
func (a *Client) SecurityCustomerRiskSend(ctx context.Context, bm gopay.BodyMap) (aliRsp *SecurityCustomerRiskSendRsp, err error) {
	err = bm.CheckEmptyError("process_code", "trade_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.security.risk.customerrisk.send"); err != nil {
		return nil, err
	}
	aliRsp = new(SecurityCustomerRiskSendRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.pay.app.marketing.consult(商户前置内容咨询接口)
// 文档地址：https://opendocs.alipay.com/pre-open/296d225f_alipay.pay.app.marketing.consult
func (a *Client) PayAppMarketingConsult(ctx context.Context, bm gopay.BodyMap) (aliRsp *PayAppMarketingConsultRsp, err error) {
	err = bm.CheckEmptyError("biz_scene", "total_amount", "product_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.pay.app.marketing.consult"); err != nil {
		return nil, err
	}
	aliRsp = new(PayAppMarketingConsultRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}
