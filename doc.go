// Package goji-localize is a Goji middleware you can append by using `goji.Use()`
// to all requests. It **does not perform localization** only language detection.
// The package uses cookies in the user browser to keep track of the language in
// use. Overwriting a cookie will change the language, as well as passing a querystring
// value.
//
// Usage
//
// You can either create your own setup...
//		loc := localize.Default()
//
//		// this will create the following struct
//		loc := &localize.Localize{
//			CookieName:      "__i18n", // the name of the cookie where language will be stored
//			CookieExpires:   time.Now().AddDate(1, 0, 0), // the cookie's expiration date
//			HttpOnly:        true, // if the cookie will be read via Javascript you can set this to false
//			DefaultLanguage: "en", // the GET parameter - querystring that, if present, will change the language
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
//		}
//
// Then, later in your Goji initialization, you can write:
// 		goji.Use(loc)
//		goji.Serve()
//
package localize
