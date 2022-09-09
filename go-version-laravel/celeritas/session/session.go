package session

import (
	"github.com/alexedwards/scs/v2"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Session struct {
	CookieLifeTime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	CookieSecure   string
	SessionType    string
}

func (s *Session) InitSession() *scs.SessionManager {
	var persist, secure bool

	minutes, err := strconv.Atoi(s.CookieLifeTime)

	if err != nil {
		minutes = 60
	}

	if strings.ToLower(s.CookiePersist) == "true" {
		persist = true
	}

	if strings.ToLower(s.CookieSecure) == "true" {
		secure = true
	}

	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Domain = s.CookieDomain
	session.Cookie.Secure = secure
	session.Cookie.Name = s.CookieName
	session.Cookie.SameSite = http.SameSiteLaxMode

	switch strings.ToLower(s.SessionType) {
	case "redis":
	case "mysql", "mariadb":
	case "postgres", "postgresql":
	default:

	}

	return session
}
