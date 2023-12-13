package core

var (
	LogLevel           = "info"
	DirectusHost       = "http://localhost:8055"
	WebPort            = 8080
	{% if cookiecutter.use_oauth %}	
	AccessTokenCookie  = "golang-simple-api-access-token"
	GithubHost         = "https://github.com"
	GithubClientID     = ""
	GithubClientSecret = ""
	GithubRedirectUri  = ""
	GithubScope        = []string{"user"}
	{% endif %}
)
