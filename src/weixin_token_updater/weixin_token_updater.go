package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ewangplay/jzlconfig"
	"github.com/outmana/log4jzl"
	"os"
	"sync"
	"time"
)

const (
	GET_WEIXIN_ACCOUNT_ALL_SQL string = "SELECT cid,account_id,appkey,appsecret,access_token,access_token_updated_at,access_token_expires_in FROM jzl_DB.jzl_weixin_account WHERE is_delete=0"
	UPDATE_ACCESS_TOKEN_SQL    string = "UPDATE jzl_weixin_account SET access_token=?,access_token_updated_at=?,access_token_expires_in=? WHERE cid = ? AND account_id = ? AND is_delete=0"
)

var g_logger *log4jzl.Log4jzl
var g_config jzlconfig.JZLConfig
var g_mysqladaptor *MysqlDBAdaptor
var g_snsClient *SNSClient
var g_waitgroup *sync.WaitGroup

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [--config path_to_config_file]")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	var err error

	//parse command line
	var configFile string
	flag.Usage = Usage
	flag.StringVar(&configFile, "config", "weixin_token_updater.conf", "specified config filename")
	flag.Parse()

	fmt.Println("config file: ", configFile)

	//read config file
	err = g_config.Read(configFile)
	if err != nil {
		fmt.Println("Read config file fail: %v", err)
		return
	}
	fmt.Println(g_config)

	//init logger
	g_logger, err = log4jzl.New("weixin_token_updater")
	if err != nil {
		fmt.Println("Open log file fail: %v", err)
		return
	}

	//init log level object
	g_logLevel, err = NewLogLevel()
	if err != nil {
		LOG_ERROR("创建NewLogLevel对象失败，失败原因: %v", err)
		return
	}

	//init mysql db adaptor
	g_mysqladaptor, err = NewMysqlDBAdaptor()
	if err != nil {
		LOG_ERROR("create MysqlDBAdaptor object fail.", err)
		os.Exit(1)
	}
	defer g_mysqladaptor.Release()

	//init SNS client
	g_snsClient, err = NewSNSClient()
	if err != nil {
		LOG_ERROR("创建SNSClient对象失败，失败原因: %v", err)
		os.Exit(1)
	}

	g_waitgroup = &sync.WaitGroup{}
	ticker := time.NewTicker(1 * time.Minute)
	for {
		select {
		case <-ticker.C:

			LOG_DEBUG("Start to update all weixin account token...")

			//从DB中获取cid,account_id,access_token
			rows, err := g_mysqladaptor.Query(GET_WEIXIN_ACCOUNT_ALL_SQL)
			if err != nil {
				LOG_ERROR("query all weixin account error: %v", err)
				continue
			}

			for rows.Next() {
				var cid, account_id, access_token_updated_at, access_token_expires_in int64
				var appkey, appsecret, access_token string

				err = rows.Scan(&cid, &account_id, &appkey, &appsecret, &access_token, &access_token_updated_at, &access_token_expires_in)
				if err != nil {
					LOG_ERROR("scan weixin account info error: %v", err)
					continue
				}

				if cid == 0 || account_id == 0 || appkey == "" || appsecret == "" {
					continue
				}

				LOG_DEBUG("[Cid: %v, ACCOUNT_ID: %v] Start to update weixin token", cid, account_id)

				g_waitgroup.Add(1)
				go UpdateWeixinAccontToken(cid, account_id, appkey, appsecret, access_token, access_token_updated_at, access_token_expires_in)
			}

			//同步协程结束
			g_waitgroup.Wait()

			rows.Close()

			LOG_DEBUG("End to update all weixin account token")
		}
	}
}

type AccessTokenInfo struct {
	Access_token string
	Expires_in   int64
}

func UpdateWeixinAccontToken(cid, account_id int64, appkey, appsecret, access_token string, token_updated_at, token_expires_in int64) error {
	defer g_waitgroup.Done()

	now := time.Now().Unix()
	if access_token == "" || (now-token_updated_at >= token_expires_in-60) {
		result, err := g_snsClient.GetWeixinSender().GetAccessToken(appkey, appsecret)
		if err == nil {
			var info AccessTokenInfo
			err = json.Unmarshal([]byte(result), &info)
			if err == nil {
				LOG_DEBUG("[CID:%v, ACCOUNT_ID:%v]刷新到新的AccessToken信息：%v", cid, account_id, info)

				//更新数据库中的AccessToken
				err := g_mysqladaptor.ExecFormat(UPDATE_ACCESS_TOKEN_SQL, info.Access_token, now, info.Expires_in, cid, account_id)
				if err != nil {
					LOG_ERROR("更新[%v_%v]数据库中的AccessToken失败，失败原因：%v", cid, account_id, err)
				}
			}
		}
	}
	return nil
}
