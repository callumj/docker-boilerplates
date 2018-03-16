# Python web app `Dockerfile`

This Python example is just a simple Flask web application powered by Gunicorn. We pass down `GIT_SHA` as a Docker build arg as a environment variable and expose it as `X-App-Version` header using an `after_request` hook in Flask.

The `Dockerfile` first copies just the `requirements.txt` file and install this via pip, if we skipped this and just did a `COPY . .` we'd be missing Docker's caching layers and waiting for a `pip install` everytime we change something inside our directory.