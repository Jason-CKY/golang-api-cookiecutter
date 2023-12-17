# Go htmx server

## Dependencies

- docker (docker-desktop if you are using windows)
- docker-compose (comes with docker-desktop, but can install [here](https://docs.docker.com/compose/install/standalone/) if you are not on windows)
- [>=Node LTS v18](https://nodejs.org/en/download)
- [>=Go v1.21](https://go.dev/doc/install)
- [Air](https://github.com/cosmtrek/air)
- [templ](https://github.com/a-h/templ)

## Features

- [swagger-ui](https://swagger.io/tools/swagger-ui/) for JSON api routes, declared with [swaggo/swag declarative comments format](https://github.com/swaggo/swag#declarative-comments-format)
- [air](https://github.com/cosmtrek/air) for code reloading in dev environment
- [echo](https://echo.labstack.com/) web server that serves html on htmx endpoints
- [templ](https://templ.guide/) templates
- [HTMX](https://htmx.org/) for interactivity, minimal js needed
- Lazy loading with HTMX
- [tailwind](https://tailwindcss.com/) for CSS Styling
- [DaisyUI](daisyui.com/) with [theme-changing library](https://github.com/saadeghi/theme-change) for CSS styling and themes
- [SortableJS](https://github.com/SortableJS/Sortable) for drag and drop of tasks (sorting and updates)
- [Directus](https://directus.io/) for headless CMS and API routes for CRUD operations

## Quickstart (development mode)
{% if cookiecutter.html_templating and cookiecutter.use_oauth %}
[Create github oauth application](https://docs.github.com/en/apps/oauth-apps/building-oauth-apps/creating-an-oauth-app).
{% endif %}
Run `cp .env.example .env`, and fill in the relevant information

```sh
# Run install-deps once to install all dev dependencies, including air and templ
make install-deps
# Start any dev dependencies with docker-compose
make build-dev
# make sure directus is up on http://localhost:8055 before running migrations for directus
make initialize-db
# start golang server with code reloading using air
air
```

There will be swagger documentation being served from `http://localhost:8080/swagger/index.html`.

{% if cookiecutter.html_templating %}
You can view the web app on `http://localhost:8080`.
{% endif %}


## Format on save

Add the following config into your vscode `settings.json` to enable format on save of a file in vscode:

```json
"editor.formatOnSave": true,
```

## VS-code extensions for good developer experience

- [Prettier - Code formatter](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode)
- [Tailwind CSS IntelliSense](https://marketplace.visualstudio.com/items?itemName=bradlc.vscode-tailwindcss)

### Syntax highlighting of golang template files on vscode

- Download [templ-vscode](https://marketplace.visualstudio.com/items?itemName=a-h.templ) vscode extension for go-templ syntax highlighting
- Add the following into your vscode `settings.json` to allow for tailwind syntax highlighting in your go `templ` files:

```json
"tailwindCSS.includeLanguages": {
"templ": "html"
}
```
