package handler

import (
	"fmt"
	"net/http"

	"github.com/morning-night-dream/platform-app/internal/driver/firebase"
	"github.com/morning-night-dream/platform-app/internal/driver/public"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
	"github.com/morning-night-dream/platform-app/pkg/log"
	"github.com/morning-night-dream/platform-app/pkg/openapi"
)

var _ openapi.ServerInterface = (*Handler)(nil)

const HeaderApiKey = "Api-Key"

type Auth struct {
	signIn  port.APIAuthSignIn
	signOut port.APIAuthSignOut
	signUp  port.APIAuthSignUp
	verify  port.APIAuthVerify
	refresh port.APIAuthRefresh
	code    port.APIAuthGenerateCode
}

func NewAuth(
	signIn port.APIAuthSignIn,
	signOut port.APIAuthSignOut,
	signUp port.APIAuthSignUp,
	verify port.APIAuthVerify,
	refresh port.APIAuthRefresh,
	code port.APIAuthGenerateCode,
) *Auth {
	return &Auth{
		signIn:  signIn,
		signOut: signOut,
		signUp:  signUp,
		verify:  verify,
		refresh: refresh,
		code:    code,
	}
}

type Handler struct {
	version  string
	key      string
	auth     *Auth
	client   *Client
	firebase *firebase.Client
	public   *public.Public
}

func New(
	version string,
	key string,
	auth *Auth,
	client *Client,
	firebase *firebase.Client,
	public *public.Public,
) *Handler {
	return &Handler{
		version:  version,
		key:      key,
		auth:     auth,
		client:   client,
		firebase: firebase,
		public:   public,
	}
}

func (hdl *Handler) IsUnauthorizedAPIKey(w http.ResponseWriter, r *http.Request) bool {
	key := r.Header.Get(HeaderApiKey)

	if key == hdl.key {
		return false
	}

	w.WriteHeader(http.StatusUnauthorized)

	log.GetLogCtx(r.Context()).Warn(fmt.Sprintf("invalid api key: %s", key))

	return true
}
