package localize

import (
	"net/http"
	"strings"
	"time"

	"github.com/zenazn/goji/web"
)

type Localize struct {
	CookieName      string
	CookieExpires   time.Time
	HttpOnly        bool
	DefaultLanguage string
	GetParamName    string
}

// Default creates a default localize struct with prefilled default values.
func Default() *Localize {
	return &Localize{
		CookieName:      "__i18n",
		CookieExpires:   time.Now().AddDate(1, 0, 0),
		HttpOnly:        true,
		DefaultLanguage: "en",
		GetParamName:    "lang",
	}
}

// SetCookieLanguage sets a cookie on every request based on a certain verifications
// made by the `GetLanguageFromRequest` function.
func (l *Localize) SetCookieLanguage(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		l.createCookie(w, l.GetLanguageFromRequest(r))
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// GetLanguageFromRequest finds a language by checking three elements:
// a GET querystring value; a cookie found in the request; or finally a
// default language predefined.
func (l *Localize) GetLanguageFromRequest(r *http.Request) string {
	// Check if there's a "lang" parameter somewhere
	if r.FormValue(l.GetParamName) != "" {
		return r.FormValue(l.GetParamName)
	}

	// If we couldn't find a language, check the cookie
	if cookie, _ := r.Cookie(l.CookieName); cookie != nil {
		return cookie.Value
	}

	// If we still can't find a language, check the headers
	if acceptLanguage := r.Header.Get("Accept-Language"); acceptLanguage != "" {
		if languages := strings.Split(acceptLanguage, ","); len(languages) >= 1 {
			return languages[0]
		}
	}

	return l.DefaultLanguage
}

func (l *Localize) createCookie(w http.ResponseWriter, lang string) {
	lang = strings.ToLower(lang)

	langs := strings.Split(lang, "-")
	if len(langs) > 0 {
		lang = langs[0]
	}

	cookieLang := &http.Cookie{
		Name:     l.CookieName,
		Value:    lang,
		Expires:  l.CookieExpires,
		HttpOnly: l.HttpOnly,
		Path:     "/",
	}

	http.SetCookie(w, cookieLang)
}
