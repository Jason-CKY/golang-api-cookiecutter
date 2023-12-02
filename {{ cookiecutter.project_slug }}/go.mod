module github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}

go {{ cookiecutter.go_version }}

require (
	github.com/a-h/templ v0.2.476
	github.com/google/uuid v1.4.0
	github.com/labstack/echo/v4 v4.11.2
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)
