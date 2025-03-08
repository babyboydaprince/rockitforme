import requests
from interfaces.network_handler import NetworkHandler

class RequestsClient(NetworkHandler):
    def send_rtsp_request(self, ip: str, port: int) -> str:
        url = f"http://{ip}:{port}"
        try:
            response = requests.get(url, timeout=5)
            return response.text if response.status_code == 200 else ""
        except requests.RequestException:
            return ""

    def send_http_request(self, ip: str, port: int, username: str, password: str) -> bool:
        url = f"http://{ip}:{port}"
        try:
            response = requests.get(url, auth=(username, password), timeout=5)
            return response.status_code == 200
        except requests.RequestException:
            return False