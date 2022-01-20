import socket
import ssl
import time

HOST = "127.0.0.1"
PORT = 443

SERVER_HOST = "129.159.50.106"
SERVER_PORT = 443

client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
client.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)

client = ssl.wrap_socket(client, keyfile="CLIENT_PY_privkey.key", certfile="CLIENT_PY_cert.pem")

if __name__ == "__main__":

    client.connect((SERVER_HOST, SERVER_PORT))
    client.send("Hello World!".encode("utf-8"))