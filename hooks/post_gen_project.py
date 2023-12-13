import os
import sys

REMOVE_PATHS = [
    '{% if not cookiecutter.use_oauth %} pkg/core/cookie.go {% endif %}',
    '{% if not cookiecutter.use_oauth %} pkg/core/oauth.go {% endif %}',
]

for path in REMOVE_PATHS:
    path = path.strip()
    if path and os.path.exists(path):
        if os.path.isdir(path):
            os.rmdir(path)
        else:
            os.unlink(path)