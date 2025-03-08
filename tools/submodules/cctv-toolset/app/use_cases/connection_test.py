from interfaces.network_handler import NetworkHandler

class ConnectionTest:
    def __init__(self, network_handler: NetworkHandler):
        self.network_handler = network_handler

    def test_connection(self, ip: str, port: int):
        return self.network_handler.send_rtsp_request(ip, port)