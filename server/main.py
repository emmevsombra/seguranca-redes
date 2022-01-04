from OpenSSL import crypto
from OpenSSL.SSL import WantReadError as socket_timeout
from connection.tls import create_socket
from threading import Thread
from json import dumps, loads
from json.decoder import JSONDecodeError
from base64 import b64decode as b64

conns = {}


def send_command(technique, agent_con):
    agent_con.send(technique.to_json)


def auth_ssl():
    hostname = '0.0.0.0'
    port = 443
    ssl_sock = create_socket('server.key', 'server.pem')
    ssl_sock.bind((hostname, port))
    ssl_sock.listen(5)
    print(f'[+] Escutando em {hostname}:{port}...')
    while True:
        print('\n[+] Aguardando conexão...)\n')
        conn, addr = ssl_sock.accept()
        print(f'[+] Conexão iniciada por {addr}')
        print('[+] Checando certificado...')
        conn.do_handshake()
        cert = conn.get_peer_certificate()
        print(f'cert => {cert}')
        if cert is not None:
            print('[+] Certificado recebido!')
            org_name = cert.get_subject().O
        else:
            print('[-] Certificado não recebido!')
            org_name = 'None'
        print(f'[~] ORG NAME: {org_name}')
        cn = cert.get_subject().CN
        
        print('[0]', conns)
        conns[cn] = conn
        print('[1]', conns)


if __name__ == '__main__':
    ssl_proc = Thread(target=auth_ssl)
    ssl_proc.start()
    ssl_proc.join()
