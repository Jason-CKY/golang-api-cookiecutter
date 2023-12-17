import os
import shutil

REMOVE_PATHS = [
    '{% if not cookiecutter.html_templating %} pkg/components {% endif %}',
    '{% if not cookiecutter.html_templating %} pkg/handlers/htmx.go {% endif %}',
    '{% if not cookiecutter.html_templating %} pkg/handlers/page.go {% endif %}',
    '{% if not cookiecutter.html_templating %} pkg/schemas/api.go {% endif %}',
    '{% if not cookiecutter.html_templating or not cookiecutter.use_oauth %} pkg/schemas/oauth.go {% endif %}',
    '{% if not cookiecutter.html_templating or cookiecutter.use_oauth %} pkg/core/cookie.go {% endif %}',
    '{% if not cookiecutter.html_templating or cookiecutter.use_oauth %} pkg/core/oauth.go {% endif %}',
]

for path in REMOVE_PATHS:
    path = path.strip()
    if path and os.path.exists(path):
        if os.path.isdir(path):
            shutil.rmtree(path)
        else:
            os.unlink(path)