var PouchDB = require('pouchdb');
var express = require('express');
var app = express();

app.use('/', require('express-pouchdb')(PouchDB));

app.listen(3000);