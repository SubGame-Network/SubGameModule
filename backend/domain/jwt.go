package domain

const CookieName = "subgame_module_token"
const CookieNameFront = "subgame_module_front_token"

type JwtUtil interface {
	GetAccessTokenKey(account string, isAdmin bool) string
	GenToken(params map[string]string) (string, string, error)
	Parse(tokenString string) (map[string]string, error)
	RenewToken(claim map[string]string) (string, error)
	CheckTokenJti(claim map[string]string, infoJsonStr string) bool
}
