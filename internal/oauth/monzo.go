package oauth

// Shamelessly stolen from https://github.com/romain-h/gone-fishing/blob/master/internal/oauth/monzo.go
// MIT License
// Copyright (c) 2020 Romain Hardy

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package oauth

import (
	"encoding/json"

	"github.com/romain-h/gone-fishing/internal/cache"
	"github.com/romain-h/gone-fishing/internal/config"
	"golang.org/x/oauth2"
)

var EndpointMonzo = oauth2.Endpoint{
	AuthURL:  "https://auth.monzo.com",
	TokenURL: "https://api.monzo.com/oauth2/token",
}

const endpointMonzoMe = "https://api.monzo.com/ping/whoami"

func NewMonzo(cfg config.Config, cache cache.CacheManager) AuthProvider {
	cacheKey := "monzo_tk"
	oauthCfg := &oauth2.Config{
		ClientID:     cfg.Monzo.AuthProvider.ClientID,
		ClientSecret: cfg.Monzo.AuthProvider.ClientSecret,
		RedirectURL:  cfg.AppURL + "package oauth

import (
	"encoding/json"

	"github.com/romain-h/gone-fishing/internal/cache"
	"github.com/romain-h/gone-fishing/internal/config"
	"golang.org/x/oauth2"
)

var EndpointMonzo = oauth2.Endpoint{
	AuthURL:  "https://auth.monzo.com/",
	TokenURL: "https://api.monzo.com/oauth2/token",
}

const endpointMonzoMe = "https://api.monzo.com/ping/whoami"

func NewMonzo(cfg config.Config, cache cache.CacheManager) AuthProvider {
	cacheKey := "monzo_tk"
	oauthCfg := &oauth2.Config{
		ClientID:     cfg.Monzo.AuthProvider.ClientID,
		ClientSecret: cfg.Monzo.AuthProvider.ClientSecret,
		RedirectURL:  cfg.AppURL + "/api/oauth2/monzo/redirect",
		Endpoint:     EndpointMonzo,
	}
	config := Config{
		Config:   oauthCfg,
		cache:    cache,
		cacheKey: cacheKey,
	}
	var tk oauth2.Token

	strToken, _ := cache.Get(cacheKey)
	json.Unmarshal([]byte(strToken), &tk)

	client := config.Client(oauth2.NoContext, &tk)

	return &provider{
		config:     config,
		client:     client,
		endpointMe: endpointMonzoMe,
	}
}",
		Endpoint:     EndpointMonzo,
	}
	config := Config{
		Config:   oauthCfg,
		cache:    cache,
		cacheKey: cacheKey,
	}
	var tk oauth2.Token

	strToken, _ := cache.Get(cacheKey)
	json.Unmarshal([]byte(strToken), &tk)

	client := config.Client(oauth2.NoContext, &tk)

	return &provider{
		config:     config,
		client:     client,
		endpointMe: endpointMonzoMe,
	}
}
