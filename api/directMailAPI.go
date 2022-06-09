package api

import (
	"PushSystem/config"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dm20151123 "github.com/alibabacloud-go/dm-20151123/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"go.uber.org/zap"
)

func createClient(accessKeyId string, accessKeySecret string) (_result *dm20151123.Client, _err error) {
	c := &openapi.Config{
		AccessKeyId:     &accessKeyId,
		AccessKeySecret: &accessKeySecret,
	}
	c.Endpoint = tea.String("dm.aliyuncs.com")
	_result = &dm20151123.Client{}
	_result, _err = dm20151123.NewClient(c)
	return _result, _err
}

func SendMail(toAddress, title, textBody string) (_err error) {
	keyID := config.Conf.Aliyun.KeyID
	keySecret := config.Conf.Aliyun.Secret
	client, _err := createClient(keyID, keySecret)
	if _err != nil {
		return _err
	}

	singleSendMailRequest := &dm20151123.SingleSendMailRequest{
		AccountName:    tea.String("admin@imouto.site"),
		AddressType:    tea.Int32(1),
		ToAddress:      tea.String(toAddress),
		Subject:        tea.String(title),
		FromAlias:      tea.String("imouto"),
		HtmlBody:       tea.String(textBody),
		ReplyToAddress: tea.Bool(true),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := client.SingleSendMailWithOptions(singleSendMailRequest, runtime)
		if _err != nil {
			return _err
		}
		zap.L().Debug(fmt.Sprint(resp))
		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		util.AssertAsString(error.Message)
	}
	return _err
}
