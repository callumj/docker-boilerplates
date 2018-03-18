FROM python:3.6.4

# setup our app directory
WORKDIR /app

# copy in just the requirements.txt file
# this will keep a cache around if we don't change our requirements (keeping the build size down)
COPY requirements.txt ./
# now install it
RUN pip install --no-cache-dir -r requirements.txt

# now copy all of our code
COPY . .

# pass down the GIT_SHA that this was built with, to surface in our X-App-Version header
ARG GIT_SHA
ENV GIT_SHA ${GIT_SHA:-unknown}

# use gunicorn with 2 workers, binding on port 3000 to serve
CMD ["gunicorn", "-w", "2", "-b", "0.0.0.0:3000", "main:app"]