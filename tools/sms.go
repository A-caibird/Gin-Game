package tools

import (
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
)
import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
)

func SendSMS(phone string, code string) (*dysmsapi.SendSmsResponse, error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(Conf.Aliyun.AccessKeyId),
		AccessKeySecret: tea.String(Conf.Aliyun.AccessKeySecret),
	}
	config.Endpoint = tea.String(Conf.Aliyun.SMS.Domain)
	client, _ := dysmsapi.NewClient(config)
	request := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		SignName:      tea.String(Conf.Aliyun.SMS.SignName),
		TemplateCode:  tea.String(Conf.Aliyun.SMS.TemplateCode),
		TemplateParam: tea.String("{\"code\":\"" + code + "\"}"),
	}
	response, err := client.SendSms(request)
	fmt.Println(response)
	if err != nil {
		fmt.Println(err)
		return response, err
	}
	return response, nil
}
