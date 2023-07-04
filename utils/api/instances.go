package api

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/vicanso/go-axios"
	"golang.org/x/oauth2"
	"khl-meobot/model"
	"log"
	"net/http"
	"net/url"
	"time"
)

const KookBaseUrl = "https://www.kookapp.cn/api/v3"

const KookToken = "Bot xxxxx"

const GameApiToken = "xxxx"

const SPOTIFY_CLIENT_ID = "xxxx"

const SPOTIFY_CLIENT_SECRET = "xxx"

const SpotifyTokenBaseUrl = "https://accounts.spotify.com/api"

const SpotifyApiBaseUrl = "https://api.spotify.com/v1"

const SpotifyBasicToken = "Basic xxx"

const KookOAuthCallback = "https://xxxxx/spotifyapp/callback"

func KookInstance() *axios.Instance {
	headers := make(http.Header)
	headers.Add("Authorization", KookToken)
	headers.Add("Content-Type", "application/json")
	return axios.NewInstance(&axios.InstanceConfig{
		BaseURL: KookBaseUrl,
		Headers: headers,
	})
}

func GameInstance() *axios.Instance {
	return axios.NewInstance(&axios.InstanceConfig{
		BaseURL: "https://api.mozambiquehe.re",
	})
}

func SpotifyTokenInstance() *axios.Instance {
	headers := make(http.Header)
	headers.Add("Authorization", SpotifyBasicToken)
	headers.Add("Content-Type", "application/x-www-form-urlencoded")
	return axios.NewInstance(&axios.InstanceConfig{
		BaseURL: SpotifyTokenBaseUrl,
		Headers: headers,
	})
}

func RenewToken() (*oauth2.Token, error) {
	rt, _ := model.Rdb().Get(context.Background(), "refresh_token").Result()
	if rt == "" {
		return nil, errors.New("请登录")
	}
	vals := make(url.Values)
	vals.Set("grant_type", "refresh_token")
	vals.Set("refresh_token", rt)
	resp, err := SpotifyTokenInstance().Post("/token", vals)
	var tok oauth2.Token
	json.Unmarshal(resp.Data, &tok)
	if err != nil {
		return nil, err
	}
	pipe := model.Rdb().TxPipeline()
	pipe.Set(context.Background(), "access_token", tok.AccessToken, 3480*time.Second)
	_, err = pipe.Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return &tok, nil
}

type SpotifyResponseError struct {
	Error SpotifyError `json:"error,omitempty"`
}

type SpotifyError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func SpotifyInstance() *axios.Instance {
	headers := make(http.Header)
	at, _ := model.Rdb().Get(context.Background(), "access_token").Result() // 从缓存读取access_token, 缓存中的access_token仅3600秒有效期
	headers.Add("Accept", "application/json")
	headers.Add("Authorization", "Bearer "+at) // 请求头携带access_token
	headers.Add("Content-Type", "application/json")
	return axios.NewInstance(&axios.InstanceConfig{
		BaseURL: SpotifyApiBaseUrl,
		Headers: headers,
		RequestInterceptors: []axios.RequestInterceptor{
			func(config *axios.Config) (err error) {
				if at == "" {
					to, err := RenewToken() // 每次请求前读取access_token, 如果过期，使用refresh_token重新获取access_token
					if err != nil {
						return err
					}
					config.Request.Header.Del("Authorization")
					config.Request.Header.Add("Authorization", "Bearer "+to.AccessToken) // 使用新获取的access_token再次请求
				}
				log.Printf("REQUSET: URL: %v %v", config.URL, config.Query)
				return
			},
		},
		ResponseInterceptors: []axios.ResponseInterceptor{
			func(resp *axios.Response) (err error) {
				var errorResp SpotifyResponseError
				json.Unmarshal(resp.Data, &errorResp)
				if errorResp.Error.Status != 0 {
					return errors.New(errorResp.Error.Message)
				}
				return
			},
		},
	})
}
