var PouchDB = require('pouchdb');
var express = require('express');
var app = express();

app.use(
    '/', 
    require('express-pouchdb') (
        PouchDB, {
            auth: {
                username: 'john',
                password: 'smith'
            }
        }
    )
);

app.listen(5984);