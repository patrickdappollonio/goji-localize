// Package goji-localize is a Goji middleware to detect the current language
// of a user who visits a Goji-powered website. It works by appending `goji.Use()`
// which will affect all requests. It **does not perform localization (translation)**,
// but only two-letter language detection only.
//
// The package uses cookies in the user browser to keep track of the language in
// use. Overwriting a cookie will change the language, as well as passing a querystring
// value.
//
// Language detection process
//
// The code has two function: `SetLanguageCookie`, which is the middleware function you'll
// pass to `goji.Use()` and `GetLanguageCode` that'll return the two-letter language code from
// the request itself, or the default value if no language code was found.
//
// Usage
//
// You can either create your own setup...
//		loc := localize.Default()
//
//		// the line above is equivalent to this...
//		loc := &localize.Localize{
//			CookieName:      "__i18n", // the name of the cookie where language will be stored
//			CookieExpires:   time.Now().AddDate(1, 0, 0), // the cookie's expiration date
//			HttpOnly:        true, // if the cookie will be read via Javascript you can set this to false
//			DefaultLanguage: "en", // the GET parameter - querystring that, if present, will change the language
// 			AvailableLanguages: []string{"en"}, // a slice of available languages to use
//			GetParamName:    "lang", // the default language, this can be any string
//		}
//
// or you can use a custom configuration, like:
//		loc := &localize.Localize{
//			CookieName:      "my_i18n_cookie",
//			CookieExpires:   time.Now().Add(50 * 24 * time.Hour),
//			HttpOnly:        true,
//			GetParamName:    "lang",
//			DefaultLanguage: "en",
// 			AvailableLanguages: []string{"en", "fr", "es"},
//		}
//
// Then, later in your Goji initialization, you can write:
// 		goji.Use(loc.SetLanguageCookie)
//		goji.Serve()
//
// Once you have the `SetLanguageCookie` function in place, you can ask any time the
// language currently in use by issuing a `loc.GetLanguageCode()`. That will return either
// the language in use or the default language.
package localize
