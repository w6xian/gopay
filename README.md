<div align=center><img width="240" height="240" alt="Logo was Loading Faild!" src="logo.png"/></div>

# GoPay

### 微信、支付宝、QQ、通联支付、拉卡拉、PayPal、扫呗、Apple支付的 Golang 版本SDK

[![Github](https://img.shields.io/github/followers/iGoogle-ink?label=Follow&style=social)](https://github.com/iGoogle-ink)
[![Github](https://img.shields.io/github/forks/go-pay/gopay?label=Fork&style=social)](https://github.com/w6xian/gopay/fork)

[![Golang](https://img.shields.io/badge/golang-1.23+-brightgreen.svg)](https://golang.google.cn)
[![GoDoc](https://img.shields.io/badge/doc-pkg.go.dev-informational.svg)](https://pkg.go.dev/github.com/w6xian/gopay)
[![Go](https://github.com/w6xian/gopay/actions/workflows/go.yml/badge.svg)](https://github.com/w6xian/gopay/actions/workflows/go.yml)
[![GitHub Release](https://img.shields.io/github/v/release/go-pay/gopay)](https://github.com/w6xian/gopay/releases)
[![License](https://img.shields.io/github/license/go-pay/gopay)](https://www.apache.org/licenses/LICENSE-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/w6xian/gopay)](https://goreportcard.com/report/github.com/w6xian/gopay)

---

# 一、安装

```bash
go get github.com/w6xian/gopay
```

#### 查看 GoPay 版本

  [版本更新记录](https://github.com/w6xian/gopay/blob/main/release_note.md)

```go
import (
    "github.com/w6xian/gopay"
    "github.com/go-pay/xlog"
)

func main() {
    xlog.Info("GoPay Version: ", gopay.Version)
}
```

---

<br>

# 二、文档目录

> ### 点击查看不同支付方式的使用文档。方便的话，请留下您认可的小星星，十分感谢！

* #### [支付宝支付（V3版）](https://github.com/w6xian/gopay/blob/main/doc/alipay_v3.md)
* #### [支付宝支付](https://github.com/w6xian/gopay/blob/main/doc/alipay.md)
* #### [微信支付（V3版）](https://github.com/w6xian/gopay/blob/main/doc/wechat_v3.md)
  * 微信商家转账产品升级，目前已支持新版商家转账接口
* #### [微信支付（V2版，不推荐）](https://github.com/w6xian/gopay/blob/main/doc/wechat_v2.md)
* #### [QQ支付](https://github.com/w6xian/gopay/blob/main/doc/qq.md)
* #### [通联支付](https://github.com/w6xian/gopay/blob/main/doc/allinpay.md)
* #### [拉卡拉支付](https://github.com/w6xian/gopay/blob/main/doc/lakala.md)
* #### [Paypal支付](https://github.com/w6xian/gopay/blob/main/doc/paypal.md)
* #### [Apple支付校验](https://github.com/w6xian/gopay/blob/main/doc/apple.md)
* #### [扫呗支付](https://github.com/w6xian/gopay/blob/main/doc/saobei.md)

---

<br>

# 三、其他说明

* 如需自定义Log输出，New Client 后，调用 `client.SetLogger()` 方法设置自定义Logger，自定义Logger实现 `xlog.XLogger` interface即可。

* 各支付方式接入，请仔细查看 `xxx_test.go` 使用方式
    * `gopay/wechat/v3/client_test.go`
    * `gopay/alipay/v3/client_test.go`
    * `gopay/alipay/client_test.go`
    * `gopay/qq/client_test.go`
    * `gopay/allinpay/client_test.go`
    * `gopay/lakala/client_test.go`
    * `gopay/paypal/client_test.go`
    * `gopay/apple/verify_test.go`
    * 或 examples
* 接入gopay示例项目(可参考接入使用方式)：[gopay-platform](https://github.com/w6xian/gopay-platform)
* 有问题请加微信群 或 关注抖音账号，加入首页粉丝群拉微信群。在此，非常感谢提出宝贵意见和反馈问题的同志们！
* 开发过程中，请尽量使用正式环境，1分钱测试法！
* 有偿承接技术咨询、开发，如需要加微信联系。

<br>

---

## 赞赏多少是您的心意，感谢支持！

微信赞赏码： <img width="200" height="200" src=".github/zanshang.png"/>
&nbsp;&nbsp;&nbsp;&nbsp;
支付宝赞助码： <img width="200" height="200" src=".github/zanshang_zfb.png"/>

---

## 问题沟通（非必要不用加微信）
> 微信: **[85411418](.github/wechat_jerry.png)**（加个人微信备注 gopay 并关注抖音账号，谢谢）。

[//]: # (关注抖音：)

[//]: # (<img width="240" height="240" src=".github/douyin_jerry.png"/>)

[//]: # (&nbsp;&nbsp;&nbsp;&nbsp;QQ群：)

[//]: # (<img width="240" height="240" src=".github/wx_gopay.png"/>)

---

<br>

## 鸣谢

> [GoLand](https://www.jetbrains.com/go/?from=gopay) A Go IDE with extended support for JavaScript, TypeScript, and Databases。

特别感谢 [JetBrains](https://www.jetbrains.com/?from=gopay) 为开源项目提供免费的 [GoLand](https://www.jetbrains.com/go/?from=gopay) 等 IDE 的授权  
<br>
[<img src=".github/jetbrains.png" width="300"/>](https://www.jetbrains.com/?from=gopay)

> Copyright © 2025 JetBrains s.r.o. JetBrains and the JetBrains logo are trademarks of JetBrains s.r.o.