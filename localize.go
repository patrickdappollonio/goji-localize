package localize

import (
	"net/http"
	"strings"
	"time"

	"github.com/zenazn/goji/web"
)

type Localize struct {
	CookieName         string
	CookieExpires      time.Time
	HttpOnly           bool
	DefaultLanguage    string
	AvailableLanguages []string
	GetParamName       string
}

// Default creates a default localize struct with prefilled default values.
func Default() *Localize {
	return &Localize{
		CookieName:         "__i18n",
		CookieExpires:      time.Now().AddDate(1, 0, 0),
		HttpOnly:           true,
		DefaultLanguage:    "en",
		AvailableLanguages: []string{"en"},
		GetParamName:       "lang",
	}
}

// SetLanguageCookie sets a cookie on every request based on a certain verifications
// made by the `GetLanguageCode` function.
func (l *Localize) SetLanguageCookie(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		l.createCookie(w, l.GetLanguageCode(r))
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// GetLanguageCode finds a two-letter language by checking four elements:
// a GET querystring value; a cookie found in the request; the Accept-Language
// header, or finally a default language predefined.
func (l *Localize) GetLanguageCode(r *http.Request) string {
	// Check if there's a "lang" parameter somewhere
	if r.FormValue(l.GetParamName) != "" {
		return l.langOrDefault(r.FormValue(l.GetParamName))
	}

	// If we couldn't find a language, check the cookie
	if cookie, _ := r.Cookie(l.CookieName); cookie != nil {
		return l.langOrDefault(cookie.Value)
	}

	// If we still can't find a language, check the headers
	if acceptLanguage := r.Header.Get("Accept-Language"); acceptLanguage != "" {
		if languages := strings.Split(acceptLanguage, ","); len(languages) >= 1 {
			return l.langOrDefault(languages[0])
		}
	}

	return l.DefaultLanguage
}

func (l *Localize) langOrDefault(lang string) string {
	for _, v := range l.AvailableLanguages {
		if lang == v {
			return lang
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
