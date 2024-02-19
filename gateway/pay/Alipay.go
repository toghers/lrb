package pay

import (
	"fmt"
	"net/url"

	"github.com/smartwalle/alipay/v3"

	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/svc"
)

var client *alipay.Client

func AlipayGoods(c *svc.ServiceContext, subject string, outNo string, prIce string) string {
	client, _ = alipay.New(c.Config.Alipays.AppId, c.Config.Alipays.PrivateKey, false)

	client.LoadAliPayPublicKey(c.Config.Alipays.AliPublicKey)

	var p = alipay.TradePagePay{}
	p.NotifyURL = c.Config.Alipays.NotifyUrl
	p.ReturnURL = "http://xxx"
	p.Subject = subject
	p.OutTradeNo = outNo
	p.TotalAmount = prIce
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	var url, err = client.TradePagePay(p)
	if err != nil {
		fmt.Println(err)
	}

	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	var payURL = url.String()
	fmt.Println(payURL)
	return payURL
}

func NotifyUpdateStatus(formData url.Values) (string, int, error) {
	notification, err := client.DecodeNotification(formData)
	if err != nil {
		return "", 0, err
	}
	switch alipay.TradeStatusFinished {
	case alipay.TradeStatusClosed:
		return notification.OutTradeNo, 3, nil
	case alipay.TradeStatusSuccess:
		return notification.OutTradeNo, 2, nil
	case alipay.TradeStatusFinished:
		return notification.OutTradeNo, 2, nil
	case alipay.TradeStatusWaitBuyerPay:
		return notification.OutTradeNo, 1, nil
	default:
		return "", -2, err
	}

}
