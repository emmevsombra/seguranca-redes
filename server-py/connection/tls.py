from socket import socket
from OpenSSL import SSL


def callback(conn, cert, errno, depth, result):
    return True


def create_socket(privkey, certfile):
    context = SSL.Context(SSL.TLSv1_2_METHOD)
    context.set_options(SSL.OP_NO_SSLv2)

    # Abaixo utilizaremos o VERIFY_PEER passando nossa função de callback para validar e podermos utilizar
    # o certificado autoassinado do agente
    context.set_verify(SSL.VERIFY_PEER, callback)
    
    # Conserto do Bug de conexão: usar as duas funções abaixo para carregar a chave e o certificado do servidor
    context.use_privatekey_file(privkey)
    context.use_certificate_file(certfile)

    return SSL.Connection(context, socket())