package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"jzlservice/weixinsender"
)

type WeixinSenderClient struct {
}

func (this *WeixinSenderClient) GetAccessToken(appid string, appsecret string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("获取公众号[%v]的AccessToken开始.", appid)

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.GetAccessToken(appid, appsecret)
	if err != nil {
		LOG_ERROR("获取公众号[%v]的AccessToken失败，失败原因：%v", appid, err)
		return "", err
	}

	LOG_INFO("获取公众号[%v]的AccessToken成功", appid)

	return

}

func (this *WeixinSenderClient) AddKFAccount(access_token string, kfaccount string, nickname string, password string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("添加微信客服[%v]开始.", kfaccount)

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.AddKFAccount(access_token, kfaccount, nickname, password)
	if err != nil {
		LOG_ERROR("添加微信客服[%v]失败，失败原因：%v", kfaccount, err)
		return "", err
	}

	LOG_INFO("添加微信客服[%v]成功", kfaccount)

	return
}

func (this *WeixinSenderClient) UpdateKFAccount(access_token string, kfaccount string, nickname string, password string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("更新微信客服[%v]开始.", kfaccount)

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.UpdateKFAccount(access_token, kfaccount, nickname, password)
	if err != nil {
		LOG_ERROR("更新微信客服[%v]失败，失败原因：%v", kfaccount, err)
		return "", err
	}

	LOG_INFO("更新微信客服[%v]成功", kfaccount)

	return
}

func (this *WeixinSenderClient) DeleteKFAccount(access_token string, kfaccount string, nickname string, password string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("删除微信客服[%v]开始.", kfaccount)

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.DeleteKFAccount(access_token, kfaccount, nickname, password)
	if err != nil {
		LOG_ERROR("删除微信客服[%v]失败，失败原因：%v", kfaccount, err)
		return "", err
	}

	LOG_INFO("删除微信客服[%v]成功", kfaccount)

	return
}

func (this *WeixinSenderClient) SetKFHeadImg(access_token string, kfaccount string, media string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("设置微信客服[%v]的头像开始.", kfaccount)

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.SetKFHeadImg(access_token, kfaccount, media)
	if err != nil {
		LOG_ERROR("设置微信客服[%v]的头像失败，失败原因：%v", kfaccount, err)
		return "", err
	}

	LOG_INFO("设置微信客服[%v]的头像成功", kfaccount)

	return
}

func (this *WeixinSenderClient) GetKFAccountList(access_token string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("获取微信客服列表开始.")

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.GetKFAccountList(access_token)
	if err != nil {
		LOG_ERROR("获取微信客服列表失败，失败原因：%v", err)
		return "", err
	}

	LOG_INFO("获取微信客服列表成功")

	return
}

func (this *WeixinSenderClient) SendMessage(access_token string, touser string, type_a1 string, data string, kfaccount string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("发送微信消息[%v]给[%v]开始.", data, touser)

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.SendMessage(access_token, touser, type_a1, data, kfaccount)
	if err != nil {
		LOG_ERROR("发送微信消息[%v]给[%v]失败，失败原因：%v", data, touser, err)
		return "", err
	}

	LOG_INFO("发送微信消息[%v]给[%v]成功", data, touser)

	return
}

func (this *WeixinSenderClient) UploadTempMedia(access_token string, type_a1 string, media string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("上传临时媒体资源开始.")

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.UploadTempMedia(access_token, type_a1, media)
	if err != nil {
		LOG_ERROR("上传临时媒体资源失败，失败原因：%v", err)
		return "", err
	}

	LOG_INFO("上传临时媒体资源成功")

	return
}

func (this *WeixinSenderClient) DownloadTempMedia(access_token string, media_id string) (r []byte, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("下载临时媒体资源[%v]开始.", media_id)

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return nil, fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return nil, fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return nil, err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return nil, err
	}

	r, err = client.DownloadTempMedia(access_token, media_id)
	if err != nil {
		LOG_ERROR("下载临时媒体资源[%v]失败，失败原因：%v", media_id, err)
		return nil, err
	}

	LOG_INFO("下载临时媒体资源[%v]成功", media_id)

	return
}

func (this *WeixinSenderClient) UploadPermanentMedia(access_token string, type_a1 string, media string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("上传永久媒体资源开始.")

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.UploadPermanentMedia(access_token, type_a1, media)
	if err != nil {
		LOG_ERROR("上传永久媒体资源失败，失败原因：%v", err)
		return "", err
	}

	LOG_INFO("上传永久媒体资源成功")

	return
}

func (this *WeixinSenderClient) DownloadPermanentMedia(access_token string, media_id string) (r []byte, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("下载永久媒体资源[%v]开始.", media_id)

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return nil, fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return nil, fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return nil, err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return nil, err
	}

	r, err = client.DownloadPermanentMedia(access_token, media_id)
	if err != nil {
		LOG_ERROR("下载永久媒体资源[%v]失败，失败原因：%v", media_id, err)
		return nil, err
	}

	LOG_INFO("下载永久媒体资源[%v]成功", media_id)

	return
}

func (this *WeixinSenderClient) DeletePermanentMedia(access_token string, media_id string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("删除永久媒体资源开始.")

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.DeletePermanentMedia(access_token, media_id)
	if err != nil {
		LOG_ERROR("删除永久媒体资源失败，失败原因：%v", err)
		return "", err
	}

	LOG_INFO("删除永久媒体资源成功")

	return
}

func (this *WeixinSenderClient) UploadNews(access_token string, news []byte) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("上传图文信息[%v]开始.", string(news))

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.UploadNews(access_token, news)
	if err != nil {
		LOG_ERROR("上传图文信息[%v]失败，失败原因：%v", string(news), err)
		return "", err
	}

	LOG_INFO("上传图文信息[%v]成功", string(news))

	return
}

func (this *WeixinSenderClient) SendNews(access_token string, is_to_all bool, group_id string, msg_type string, content string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("群发图文信息[%v]开始.", content)

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		outputStr = fmt.Sprintf("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		LOG_ERROR(outputStr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.SendNews(access_token, is_to_all, group_id, msg_type, content)
	if err != nil {
		LOG_ERROR("群发图文信息[%v]失败，失败原因：%v", content, err)
		return "", err
	}

	LOG_INFO("群发图文信息[%v]成功", content)

	return
}

func (this *WeixinSenderClient) DeleteNews(access_token string, msg_id string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("删除图文信息[%v]开始.", msg_id)

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.DeleteNews(access_token, msg_id)
	if err != nil {
		LOG_ERROR("删除图文信息[%v]失败，失败原因：%v", msg_id, err)
		return "", err
	}

	LOG_INFO("删除图文信息[%v]成功", msg_id)

	return
}

func (this *WeixinSenderClient) PreviewNewsByOpenId(access_token string, touser string, msg_type string, content string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("预览发送图文消息[%v]给[%v]开始.", content, touser)

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.PreviewNewsByOpenId(access_token, touser, msg_type, content)
	if err != nil {
		LOG_ERROR("预览发送图文消息[%v]给[%v]失败，失败原因：%v", content, touser, err)
		return "", err
	}

	LOG_INFO("预览发送图文消息[%v]给[%v]成功", content, touser)

	return
}

func (this *WeixinSenderClient) GetNewsStatus(access_token string, msg_id string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("获取图文消息[%v]的状态开始.", msg_id)

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.GetNewsStatus(access_token, msg_id)
	if err != nil {
		LOG_ERROR("获取图文消息[%v]的状态失败，失败原因：%v", msg_id, err)
		return "", err
	}

	LOG_INFO("获取图文消息[%v]的状态成功", msg_id)

	return
}

func (this *WeixinSenderClient) CreateUserGroup(access_token string, group_name string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("创建微信用户组[%v]开始.", group_name)

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.CreateUserGroup(access_token, group_name)
	if err != nil {
		LOG_ERROR("创建微信用户组[%v]失败，失败原因：%v", group_name, err)
		return "", err
	}

	LOG_INFO("创建微信用户组[%v]成功", group_name)

	return
}

func (this *WeixinSenderClient) UpdateUserGroup(access_token string, group_id string, new_group_name string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("微信用户组[%v]更名为[%v]开始.", group_id, new_group_name)

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.UpdateUserGroup(access_token, group_id, new_group_name)
	if err != nil {
		LOG_ERROR("微信用户组[%v]更名为[%v]失败，失败原因：%v", group_id, new_group_name, err)
		return "", err
	}

	LOG_INFO("微信用户组[%v]更名为[%v]成功", group_id, new_group_name)

	return
}

func (this *WeixinSenderClient) GetUserGroupList(access_token string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("获取微信用户组列表开始.")

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.GetUserGroupList(access_token)
	if err != nil {
		LOG_ERROR("获取微信用户组列表失败，失败原因：%v", err)
		return "", err
	}

	LOG_INFO("获取微信用户组列表成功")

	return
}

func (this *WeixinSenderClient) GetUserGroupByOpenID(access_token string, openid string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("获取微信用户[%v]的组开始.", openid)

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.GetUserGroupByOpenID(access_token, openid)
	if err != nil {
		LOG_ERROR("获取微信用户[%v]的组失败，失败原因：%v", openid, err)
		return "", err
	}

	LOG_INFO("获取微信用户[%v]的组成功", openid)

	return
}

func (this *WeixinSenderClient) MoveUserToGroup(access_token string, openid_list []string, to_groupid string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("移动用户%v到组[%v]开始.", openid_list, to_groupid)

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.MoveUserToGroup(access_token, openid_list, to_groupid)
	if err != nil {
		LOG_ERROR("移动用户%v到组[%v]失败，失败原因：%v", openid_list, to_groupid, err)
		return "", err
	}

	LOG_INFO("移动用户%v到组[%v]成功", openid_list, to_groupid)

	return
}

func (this *WeixinSenderClient) RemarkUser(access_token string, openid string, remark string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("备注用户[%v]为[%v]开始.", openid, remark)

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.RemarkUser(access_token, openid, remark)
	if err != nil {
		LOG_ERROR("备注用户[%v]为[%v]失败，失败原因：%v", openid, remark, err)
		return "", err
	}

	LOG_INFO("备注用户[%v]为[%v]成功", openid, remark)

	return
}

func (this *WeixinSenderClient) GetUserInfo(access_token string, openid string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("获取用户[%v]的详细信息开始.", openid)

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.GetUserInfo(access_token, openid)
	if err != nil {
		LOG_ERROR("获取用户[%v]的详细信息失败，失败原因：%v", openid, err)
		return "", err
	}

	LOG_INFO("获取用户[%v]的详细信息成功", openid)

	return
}

func (this *WeixinSenderClient) GetUserList(access_token string, next_openid string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("获取微信用户列表开始.")

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.GetUserList(access_token, next_openid)
	if err != nil {
		LOG_ERROR("获取微信用户列表失败，失败原因：%v", err)
		return "", err
	}

	LOG_INFO("获取微信用户列表成功")

	return
}

func (this *WeixinSenderClient) CreateMenu(access_token string, menu_data []byte) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("创建微信菜单开始.")

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.CreateMenu(access_token, menu_data)
	if err != nil {
		LOG_ERROR("创建微信菜单失败，失败原因：%v", err)
		return "", err
	}

	LOG_INFO("创建微信菜单成功")

	return

}

func (this *WeixinSenderClient) DeleteMenu(access_token string) (r string, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("删除微信菜单开始.")

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return "", fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return "", err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return "", err
	}

	r, err = client.DeleteMenu(access_token)
	if err != nil {
		LOG_ERROR("删除微信菜单失败，失败原因：%v", err)
		return "", err
	}

	LOG_INFO("删除微信菜单成功")

	return

}

func (this *WeixinSenderClient) GetMenu(access_token string) (r []byte, err error) {
	var outputStr string
	var networkAddr string
	var addr, port string
	var addrIsSet, portIsSet bool

	LOG_INFO("获取微信菜单开始.")

	addr, addrIsSet = g_config.Get("weixin_sender.addr")
	port, portIsSet = g_config.Get("weixin_sender.port")

	if addrIsSet && portIsSet {
		if addr != "" && port != "" {
			networkAddr = fmt.Sprintf("%s:%s", addr, port)
		} else {
			outputStr = "WeixinSender服务的网络地址设置错误"
			LOG_ERROR(outputStr)
			return nil, fmt.Errorf(outputStr)
		}
	} else {
		outputStr = "WeixinSender服务的网络地址没有设置"
		LOG_ERROR(outputStr)
		return nil, fmt.Errorf(outputStr)
	}

	trans, err := thrift.NewTSocket(networkAddr)
	if err != nil {
		LOG_ERROR("创建到WeixinSender服务[%v]的连接失败", networkAddr)
		return nil, err
	}
	defer trans.Close()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	client := weixinsender.NewWeixinSenderClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		LOG_ERROR("打开到WeixinSender服务的连接失败, WeixinSender服务可能没有就绪，请检查服务状态")
		return nil, err
	}

	r, err = client.GetMenu(access_token)
	if err != nil {
		LOG_ERROR("获取微信菜单失败，失败原因：%v", err)
		return nil, err
	}

	LOG_INFO("获取微信菜单成功")

	return
}

type SNSClient struct {
	weixinSenderClient *WeixinSenderClient
}

func NewSNSClient() (*SNSClient, error) {
	snsClient := &SNSClient{}
	snsClient.weixinSenderClient = &WeixinSenderClient{}
	return snsClient, nil
}

func (this *SNSClient) GetWeixinSender() *WeixinSenderClient {
	return this.weixinSenderClient
}
