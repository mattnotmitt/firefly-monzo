package oauth

// Shamelessly stolen from https://github.com/romain-h/gone-fishing/blob/master/internal/oauth/config.go
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

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mattnotmitt/firefly-monzo/internal/cache"
	"golang.org/x/oauth2"
)

type Config struct {
	*oauth2.Config
	cache    cache.CacheManager
	cacheKey string
}

func (c *Config) StoreToken(token *oauth2.Token) error {
	t, _ := json.Marshal(token)
	err := c.cache.Set(c.cacheKey, string(t))
	return err
}

func (c *Config) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := c.Config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	if err := c.StoreToken(token); err != nil {
		return nil, err
	}
	return token, nil
}

func (c *Config) Client(ctx context.Context, t *oauth2.Token) *http.Client {
	return oauth2.NewClient(ctx, c.TokenSource(ctx, t))
}

func (c *Config) TokenSource(ctx context.Context, t *oauth2.Token) oauth2.TokenSource {
	rts := &CacheTokenSource{
		source: c.Config.TokenSource(ctx, t),
		config: c,
	}
	return oauth2.ReuseTokenSource(t, rts)
}

type CacheTokenSource struct {
	source oauth2.TokenSource
	config *Config
}

func (t *CacheTokenSource) Token() (*oauth2.Token, error) {
	token, err := t.source.Token()
	if err != nil {
		return nil, err
	}
	if err := t.config.StoreToken(token); err != nil {
		return nil, err
	}
	return token, nil
}
