<!-- Generator: Widdershins v4.0.1 -->

<h1 id="morning-night-dream-appgateway">Morning Night Dream - AppGateway v0.0.1</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

This is the AppGateway API documentation.

Base URLs:

* <a href="http://localhost:8082/api">http://localhost:8082/api</a>

<a href="https://example.com">Terms of service</a>
Email: <a href="mailto:morning.night.dream@example.com">Support</a> 
 License: MIT

# Authentication

* API Key (cookieAuth)
    - Parameter Name: **auth**, in: cookie. 

<h1 id="morning-night-dream-appgateway-auth">auth</h1>

認証

## v1AuthSignUp

<a id="opIdv1AuthSignUp"></a>

`POST /v1/auth/signup`

*サインアップ*

サインアップ

> Body parameter

```json
{
  "email": "morning.night.dream@example.com",
  "password": "password"
}
```

<h3 id="v1authsignup-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|サインアップリクエストボディ|
|» email|body|string(email)|true|メールアドレス|
|» password|body|string(password)|true|パスワード|

<h3 id="v1authsignup-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|None|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|None|

<aside class="success">
This operation does not require authentication
</aside>

## v1AuthSignIn

<a id="opIdv1AuthSignIn"></a>

`POST /v1/auth/signin`

*サインイン*

サインイン

> Body parameter

```json
{
  "email": "morning.night.dream@example.com",
  "password": "password",
  "publicKey": "string"
}
```

<h3 id="v1authsignin-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|サインインリクエストボディ|
|» email|body|string(email)|true|メールアドレス|
|» password|body|string(password)|true|パスワード|
|» publicKey|body|string(base64)|true|公開鍵|

<h3 id="v1authsignin-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|None|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|None|

<aside class="success">
This operation does not require authentication
</aside>

## v1AuthVerify

<a id="opIdv1AuthVerify"></a>

`GET /v1/auth/verify`

*検証*

検証

> Example responses

> 401 Response

```json
{
  "code": "f5d62b05-370e-48be-a755-8675ca146431"
}
```

<h3 id="v1authverify-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|None|
|302|[Found](https://tools.ietf.org/html/rfc7231#section-6.4.3)|Redirect|None|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|[UnauthorizedResponse](#schemaunauthorizedresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|None|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
cookieAuth
</aside>

## v1AuthRefresh

<a id="opIdv1AuthRefresh"></a>

`GET /v1/auth/refresh`

*リフレッシュ*

リフレッシュ

<h3 id="v1authrefresh-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|code|query|string|true|署名付きコード|

<h3 id="v1authrefresh-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|None|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|None|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
cookieAuth
</aside>

## v1AuthSignOut

<a id="opIdv1AuthSignOut"></a>

`POST /v1/auth/signout`

*サインアウト*

サインアウト

<h3 id="v1authsignout-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|None|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|None|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
cookieAuth
</aside>

## v1AuthResign

<a id="opIdv1AuthResign"></a>

`DELETE /v1/auth`

*リサイン(退会)*

リサイン(退会)

> Body parameter

```json
{
  "password": "password"
}
```

<h3 id="v1authresign-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|リサインリクエストボディ|
|» password|body|string(password)|true|パスワード|

<h3 id="v1authresign-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|None|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|None|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
cookieAuth
</aside>

## v1Sign

<a id="opIdv1Sign"></a>

`GET /v1/sign`

*署名検証*

署名検証

<h3 id="v1sign-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|code|query|string|true|署名付きコード|
|signature|query|string|true|署名|

<h3 id="v1sign-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|None|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|None|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="morning-night-dream-appgateway-article">article</h1>

記事

## v1ListArticles

<a id="opIdv1ListArticles"></a>

`GET /v1/article`

*List articles*

List articles

<h3 id="v1listarticles-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|pageToken|query|string|false|トークン|
|maxPageSize|query|integer|true|ページサイズ|

> Example responses

> 200 Response

```json
{
  "articles": [
    {
      "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
      "url": "https://example.com",
      "title": "sample title",
      "description": "sample description",
      "thumbnail": "https://example.com",
      "tags": [
        "tag"
      ]
    }
  ],
  "nextPageToken": "string"
}
```

<h3 id="v1listarticles-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[ListArticleResponse](#schemalistarticleresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|サーバーエラー|None|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="morning-night-dream-appgateway-health">health</h1>

ヘルスチェック

## v1Health

<a id="opIdv1Health"></a>

`GET /v1/health`

*ヘルスチェック*

ヘルスチェック

<h3 id="v1health-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|None|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|None|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_Article">Article</h2>
<!-- backwards compatibility -->
<a id="schemaarticle"></a>
<a id="schema_Article"></a>
<a id="tocSarticle"></a>
<a id="tocsarticle"></a>

```json
{
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
  "url": "https://example.com",
  "title": "sample title",
  "description": "sample description",
  "thumbnail": "https://example.com",
  "tags": [
    "tag"
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|string(uuid)|false|none|id|
|url|string(uri)|false|none|記事のURL|
|title|string|false|none|タイトル|
|description|string|false|none|description|
|thumbnail|string(uri)|false|none|サムネイルのURL|
|tags|[string]|false|none|タグ|

<h2 id="tocS_ListArticleResponse">ListArticleResponse</h2>
<!-- backwards compatibility -->
<a id="schemalistarticleresponse"></a>
<a id="schema_ListArticleResponse"></a>
<a id="tocSlistarticleresponse"></a>
<a id="tocslistarticleresponse"></a>

```json
{
  "articles": [
    {
      "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
      "url": "https://example.com",
      "title": "sample title",
      "description": "sample description",
      "thumbnail": "https://example.com",
      "tags": [
        "tag"
      ]
    }
  ],
  "nextPageToken": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|articles|[[Article](#schemaarticle)]|false|none|none|
|nextPageToken|string|false|none|次回リクエスト時に指定するページトークン|

<h2 id="tocS_UnauthorizedResponse">UnauthorizedResponse</h2>
<!-- backwards compatibility -->
<a id="schemaunauthorizedresponse"></a>
<a id="schema_UnauthorizedResponse"></a>
<a id="tocSunauthorizedresponse"></a>
<a id="tocsunauthorizedresponse"></a>

```json
{
  "code": "f5d62b05-370e-48be-a755-8675ca146431"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|code|string(uuid)|true|none|コード|

