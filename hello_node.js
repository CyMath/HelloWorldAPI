var express = require('express');
var app = express();
var port= process.env.PORT || 3000;

app.listen(port, () =>
{
    console.log('Server started on: ' + port);
});

const REPLY = 
    {
        "message": "Hello message"
    };

app.get('/', (req, res) => index(req, res));
app.get('/JSON', (req, res) => JSONMessage(req, res));
app.get('/:path', (req, res) => pathMessage(req, res));

function index(request, response)
{
    response.end("Hello, World!");
}

function JSONMessage(request, response)
{
    response.setHeader('content-type', 'application/json');
    response.end(JSON.stringify(REPLY));
}

function pathMessage(request, response)
{
    response.end("Hello, " + request.params.path);
}
