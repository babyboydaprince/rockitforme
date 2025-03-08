from abc import ABC, abstractmethod

class NetworkHandler(ABC):
    @abstractmethod
    def send_rtsp_request(self, ip: str, port: int) -> str:
        pass

    @abstractmethod
    def send_http_request(self, ip: str, port: int, username: str, password: str) -> bool:
        pass