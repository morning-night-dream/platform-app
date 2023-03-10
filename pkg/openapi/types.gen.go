// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package openapi

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

const (
	CookieAuthScopes = "cookieAuth.Scopes"
)

// Article defines model for Article.
type Article struct {
	// Description description
	Description *string `json:"description,omitempty"`

	// Id id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Tags タグ
	Tags *[]string `json:"tags,omitempty"`

	// Thumbnail サムネイルのURL
	Thumbnail *string `json:"thumbnail,omitempty"`

	// Title タイトル
	Title *string `json:"title,omitempty"`

	// Url 記事のURL
	Url *string `json:"url,omitempty"`
}

// V1ListArticleResponse defines model for V1ListArticleResponse.
type V1ListArticleResponse struct {
	Articles *[]Article `json:"articles,omitempty"`

	// NextPageToken 次回リクエスト時に指定するページトークン
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

// V1UnauthorizedResponse defines model for V1UnauthorizedResponse.
type V1UnauthorizedResponse struct {
	// Code コード
	Code openapi_types.UUID `json:"code"`
}

// V1ListArticlesParams defines parameters for V1ListArticles.
type V1ListArticlesParams struct {
	// PageToken トークン
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	// MaxPageSize ページサイズ
	MaxPageSize int `form:"maxPageSize" json:"maxPageSize"`
}

// V1AuthResignJSONBody defines parameters for V1AuthResign.
type V1AuthResignJSONBody struct {
	// Password パスワード
	Password string `json:"password"`
}

// V1AuthRefreshParams defines parameters for V1AuthRefresh.
type V1AuthRefreshParams struct {
	// Code 署名付きコード
	Code string `form:"code" json:"code"`

	// Signature 署名
	Signature string `form:"signature" json:"signature"`
	ExpiresIn *int   `form:"expiresIn,omitempty" json:"expiresIn,omitempty"`
}

// V1AuthSignInJSONBody defines parameters for V1AuthSignIn.
type V1AuthSignInJSONBody struct {
	// Email メールアドレス
	Email openapi_types.Email `json:"email"`

	// ExpiresIn トークン有効期限(秒)
	ExpiresIn *int `json:"expiresIn,omitempty"`

	// Password パスワード
	Password string `json:"password"`

	// PublicKey 公開鍵
	PublicKey string `json:"publicKey"`
}

// V1AuthSignUpJSONBody defines parameters for V1AuthSignUp.
type V1AuthSignUpJSONBody struct {
	// Email メールアドレス
	Email openapi_types.Email `json:"email"`

	// Password パスワード
	Password string `json:"password"`
}

// V1SignParams defines parameters for V1Sign.
type V1SignParams struct {
	// Code 署名付きコード
	Code string `form:"code" json:"code"`

	// Signature 署名
	Signature string `form:"signature" json:"signature"`
}

// V1AuthResignJSONRequestBody defines body for V1AuthResign for application/json ContentType.
type V1AuthResignJSONRequestBody V1AuthResignJSONBody

// V1AuthSignInJSONRequestBody defines body for V1AuthSignIn for application/json ContentType.
type V1AuthSignInJSONRequestBody V1AuthSignInJSONBody

// V1AuthSignUpJSONRequestBody defines body for V1AuthSignUp for application/json ContentType.
type V1AuthSignUpJSONRequestBody V1AuthSignUpJSONBody
