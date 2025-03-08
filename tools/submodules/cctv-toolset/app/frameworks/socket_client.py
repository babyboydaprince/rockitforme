import socket
from interfaces.network_handler import NetworkHandler

class SocketClient(NetworkHandler):
    def send_rtsp_request(self, ip: str, port: int) -> str:
        try:
            s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
            s.settimeout(5)
            s.connect((ip, port))
            request = f"DESCRIBE rtsp://{ip}:{port} RTSP/1.0\r\nCSeq: 1\r\n\r\n"
            s.sendall(request.encode())
            data = s.recv(4096)
            s.close()
            return data.decode('utf-8', errors='ignore')
        except Exception:
            return ""

    def send_http_request(self, ip: str, port: int, username: str, password: str) -> bool:
        return False