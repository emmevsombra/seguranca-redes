const express = require('express');
const server = express();
const fs = require('fs');
const path = require('path');
const http = require('http');
const https = require('https');

const port = 80
const portSSL = 443

const privateKey = fs.readFileSync(path.resolve('sslcert/server.key'), 'utf8');
const certificate = fs.readFileSync(path.resolve('sslcert/server.cert'), 'utf8');
const options = { key: privateKey, cert: certificate };


server.use('/', (req, res) => {
    res.send('Hello World!')
})

const httpServer = http.createServer(server);
const httpsServer = https.createServer(options, server);
  
httpServer.listen(port, () => {
  console.info(`BACKEND is running on port ${port}.`);
});
  
httpsServer.listen(portSSL, () => {
  console.info(`BACKEND is running on port ${portSSL}.`);
});
  
