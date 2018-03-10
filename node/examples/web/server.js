var express = require('express');
var app = express();
var router = express.Router();

router.use(function (req, res, next) {
  res.append("X-App-Version", process.env.GIT_SHA);
  next();
});

router.get('/hello_world', function (req, res) {
  res.send('Hello World');
});

app.use('/', router);
app.listen(3000);