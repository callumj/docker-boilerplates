from flask import Flask
from os import getenv

app = Flask(__name__)

# setup our X-App-Version "middleware"
@app.after_request
def apply_caching(response):
    response.headers["X-App-Version"] = getenv("GIT_SHA", "None")
    return response

@app.route("/hello_world")
def hello_world():
    return "Hello World"