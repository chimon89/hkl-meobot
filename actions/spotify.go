package actions

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"khl-meobot/model"
	"khl-meobot/utils"
	"khl-meobot/utils/api"
	"log"
	"net/url"
	"time"
)

var (
	RedirectUrl = "https://kookbot.dailypics.cn/spotify/callback"
	AuthConfig  = spotifyauth.New(
		spotifyauth.WithClientID(api.SPOTIFY_CLIENT_ID),
		spotifyauth.WithClientSecret(api.SPOTIFY_CLIENT_SECRET),
		spotifyauth.WithRedirectURL(RedirectUrl),
		spotifyauth.WithScopes(
			spotifyauth.ScopeUserReadPrivate,
			spotifyauth.ScopeUserReadEmail,
			spotifyauth.ScopeUserReadPlaybackState,
			spotifyauth.ScopeUserModifyPlaybackState,
			spotifyauth.ScopeUserReadCurrentlyPlaying,
		),
	)
	State = utils.RandomString(16, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"))
)

func ReqSpotifyAuth(c *gin.Context) {
	url := AuthConfig.AuthURL(State)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)
	c.Redirect(302, url)
}

func CompleteSpotifyAuth(c *gin.Context) {
	if code := c.Query("code"); code != "" {
		tok, err := AuthConfig.Token(context.Background(), State, c.Request)
		if err != nil {
			c.String(400, "当前页面已失效，请重新登录！")
			c.Abort()
			return
		}
		pipe := model.Rdb().TxPipeline()
		pipe.Set(context.Background(), "access_token", tok.AccessToken, 3600*time.Second)
		pipe.Set(context.Background(), "refresh_token", tok.RefreshToken, 0)
		_, err = pipe.Exec(context.Background())
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		userProfile, _ := GetSpotifyUserProfile()
		model.Rdb().Set(context.Background(), "market", userProfile.Country, 0)
		c.HTML(200, "login_success.tmpl", gin.H{})
		return
	}
	c.JSON(401, gin.H{"error": "用户未授权"})
}

type SpotifyPlayerQueue struct {
	CurrentlyPlaying struct {
		Album struct {
			AlbumType string `json:"album_type"`
			Artists   []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			AvailableMarkets []string `json:"available_markets"`
			ExternalUrls     struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name                 string `json:"name"`
			ReleaseDate          string `json:"release_date"`
			ReleaseDatePrecision string `json:"release_date_precision"`
			TotalTracks          int    `json:"total_tracks"`
			Type                 string `json:"type"`
			URI                  string `json:"uri"`
		} `json:"album"`
		Artists []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"artists"`
		AvailableMarkets []string `json:"available_markets"`
		DiscNumber       int      `json:"disc_number"`
		DurationMs       int      `json:"duration_ms"`
		Explicit         bool     `json:"explicit"`
		ExternalIds      struct {
			Isrc string `json:"isrc"`
		} `json:"external_ids"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href        string `json:"href"`
		ID          string `json:"id"`
		IsLocal     bool   `json:"is_local"`
		Name        string `json:"name"`
		Popularity  int    `json:"popularity"`
		PreviewURL  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		URI         string `json:"uri"`
	} `json:"currently_playing"`
	Queue []struct {
		Album struct {
			AlbumType string `json:"album_type"`
			Artists   []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			AvailableMarkets []string `json:"available_markets"`
			ExternalUrls     struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name                 string `json:"name"`
			ReleaseDate          string `json:"release_date"`
			ReleaseDatePrecision string `json:"release_date_precision"`
			TotalTracks          int    `json:"total_tracks"`
			Type                 string `json:"type"`
			URI                  string `json:"uri"`
		} `json:"album"`
		Artists []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"artists"`
		AvailableMarkets []string `json:"available_markets"`
		DiscNumber       int      `json:"disc_number"`
		DurationMs       int      `json:"duration_ms"`
		Explicit         bool     `json:"explicit"`
		ExternalIds      struct {
			Isrc string `json:"isrc"`
		} `json:"external_ids"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href        string `json:"href"`
		ID          string `json:"id"`
		IsLocal     bool   `json:"is_local"`
		Name        string `json:"name"`
		Popularity  int    `json:"popularity"`
		PreviewURL  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		URI         string `json:"uri"`
	} `json:"queue"`
}

func GetSpotifyQueueAPI(c *gin.Context) {
	qu, err := GetSpotifyQueue()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, qu)
}

func GetSpotifyQueue() (*SpotifyPlayerQueue, error) {
	response, err := api.SpotifyInstance().Get("/me/player/queue")
	if err != nil {
		return nil, err
	}
	var qu SpotifyPlayerQueue
	json.Unmarshal(response.Data, &qu)
	return &qu, nil
}

type SpotifySearchResult struct {
	Tracks struct {
		Href  string `json:"href"`
		Items []struct {
			Album struct {
				AlbumType string `json:"album_type"`
				Artists   []struct {
					ExternalUrls struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href string `json:"href"`
					ID   string `json:"id"`
					Name string `json:"name"`
					Type string `json:"type"`
					URI  string `json:"uri"`
				} `json:"artists"`
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href   string `json:"href"`
				ID     string `json:"id"`
				Images []struct {
					Height int    `json:"height"`
					URL    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"images"`
				Name                 string `json:"name"`
				ReleaseDate          string `json:"release_date"`
				ReleaseDatePrecision string `json:"release_date_precision"`
				TotalTracks          int    `json:"total_tracks"`
				Type                 string `json:"type"`
				URI                  string `json:"uri"`
			} `json:"album"`
			Artists []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			DiscNumber  int  `json:"disc_number"`
			DurationMs  int  `json:"duration_ms"`
			Explicit    bool `json:"explicit"`
			ExternalIds struct {
				Isrc string `json:"isrc"`
			} `json:"external_ids"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href        string `json:"href"`
			ID          string `json:"id"`
			IsLocal     bool   `json:"is_local"`
			IsPlayable  bool   `json:"is_playable"`
			Name        string `json:"name"`
			Popularity  int    `json:"popularity"`
			PreviewURL  string `json:"preview_url"`
			TrackNumber int    `json:"track_number"`
			Type        string `json:"type"`
			URI         string `json:"uri"`
		} `json:"items"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
	} `json:"tracks"`
}

func GetSpotifySearchResult(q string) (*SpotifySearchResult, error) {
	market, _ := model.Rdb().Get(context.Background(), "market").Result()
	values := make(url.Values)
	values.Set("q", q)
	values.Set("type", "track")
	values.Set("market", market)
	values.Set("limit", "10")
	resp, err := api.SpotifyInstance().Get("/search", values)
	if err != nil {
		return nil, err
	}
	var r SpotifySearchResult
	json.Unmarshal(resp.Data, &r)
	log.Printf("%v", string(resp.Data))
	return &r, nil
}

type TrackReturnVal struct {
	URI       string `json:"uri"`
	TrackName string `json:"trackName"`
	Artists   string `json:"artists"`
}

type SpotifyUserProfile struct {
	Country         string `json:"country"`
	DisplayName     string `json:"display_name"`
	Email           string `json:"email"`
	ExplicitContent struct {
		FilterEnabled bool `json:"filter_enabled"`
		FilterLocked  bool `json:"filter_locked"`
	} `json:"explicit_content"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href  interface{} `json:"href"`
		Total int         `json:"total"`
	} `json:"followers"`
	Href    string        `json:"href"`
	ID      string        `json:"id"`
	Images  []interface{} `json:"images"`
	Product string        `json:"product"`
	Type    string        `json:"type"`
	URI     string        `json:"uri"`
}

func GetSpotifyUserProfile() (*SpotifyUserProfile, error) {
	resp, err := api.SpotifyInstance().Get("/me")
	var u SpotifyUserProfile
	json.Unmarshal(resp.Data, &u)
	return &u, err
}

type SpotifyDevicesList struct {
	Devices []struct {
		ID               string `json:"id"`
		IsActive         bool   `json:"is_active"`
		IsPrivateSession bool   `json:"is_private_session"`
		IsRestricted     bool   `json:"is_restricted"`
		Name             string `json:"name"`
		Type             string `json:"type"`
		VolumePercent    int    `json:"volume_percent"`
	} `json:"devices"`
}

func GetSpotifyDevicesList() (*SpotifyDevicesList, error) {
	resp, err := api.SpotifyInstance().Get("/me/player/devices")
	var dl SpotifyDevicesList
	json.Unmarshal(resp.Data, &dl)
	return &dl, err
}
