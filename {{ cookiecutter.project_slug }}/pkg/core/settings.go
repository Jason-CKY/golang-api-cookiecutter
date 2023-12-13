package core

var (
	DirectusHost       = "http://localhost:8055"
	WebPort            = 8080
	StateTokenCookie   = "{{ cookiecutter.project_slug }}-state-token"
	AccessTokenCookie  = "{{ cookiecutter.project_slug }}-access-token"
	RefreshTokenCookie = "{{ cookiecutter.project_slug }}-refresh-token"
)
