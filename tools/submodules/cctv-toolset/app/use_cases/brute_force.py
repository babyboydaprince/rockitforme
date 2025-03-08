from interfaces.file_handler import FileHandler
from interfaces.network_handler import NetworkHandler

class BruteForce:
    def __init__(self, file_handler: FileHandler, network_handler: NetworkHandler):
        self.file_handler = file_handler
        self.network_handler = network_handler

    def brute_force(self, ip: str, port: int, username: str, user_wordlist: str, pass_wordlist: str, delay: float):
        if username:
            self._single_user_brute_force(ip, port, username, pass_wordlist, delay)
        elif user_wordlist:
            self._user_wordlist_brute_force(ip, port, user_wordlist, pass_wordlist, delay)

    def _single_user_brute_force(self, ip: str, port: int, username: str, pass_wordlist: str, delay: float):
        passwords = self.file_handler.read_file(pass_wordlist)
        for password in passwords:
            if self.network_handler.send_http_request(ip, port, username, password):
                print(f"Success! Username: {username}, Password: {password}")
                return

    def _user_wordlist_brute_force(self, ip: str, port: int, user_wordlist: str, pass_wordlist: str, delay: float):
        users = self.file_handler.read_file(user_wordlist)
        passwords = self.file_handler.read_file(pass_wordlist)
        for user in users:
            for password in passwords:
                if self.network_handler.send_http_request(ip, port, user, password):
                    print(f"Success! Username: {user}, Password: {password}")
                    return