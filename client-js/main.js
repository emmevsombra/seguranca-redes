'use strict';

var tls = require('tls');
var fs = require('fs');


const PORT = 443;
const HOST = '129.159.50.106'

var options = {
    key: fs.readFileSync('./ssl/CLIENT_JS_privkey.key'),
    cert: fs.readFileSync('./ssl/CLIENT_JS_cert.pem'),
    ca: fs.readFileSync('./ssl/server.pem'),
    requestCert: true,
    rejectUnauthorized: false
};

var timeBefore = Date.now();
console.log("time before::: ", timeBefore);
var client = tls.connect(PORT, HOST, options, function() {
    if (client.authorized) {
        console.log("Connection authorized by a Certificate Authority.");
    } else {
        console.log("Connection not authorized: " + client.authorizationError)
    }
    var timeAfter= Date.now();
    console.log("time after::: ", timeAfter);
    console.log("transmission time::: ", timeAfter - timeBefore);

    //envio da mensagem
    client.write("teste JS");
    return

});

client.on("data", function(data) {
    console.log('Received: %s [it is %d bytes long]',
        data.toString().replace(/(\n)/gm,""),
        data.length);

    client.end();
});

client.on('close', function() {
    console.log("Connection closed");
});

client.on('error', function(error) {
    console.error(error);
    client.destroy();
});