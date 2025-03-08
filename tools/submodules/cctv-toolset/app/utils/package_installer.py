import subprocess
import sys
from colorama import Fore

class PackageInstaller:
    @staticmethod
    def install(packages: list):
        checkmark = '\u2713'
        try:
            missing_packages = [pkg for pkg in packages if not PackageInstaller.is_installed(pkg)]
            if missing_packages:
                print(f"Installing missing packages: {', '.join(missing_packages)}")
                subprocess.check_call([sys.executable, '-m', 'pip', 'install', *missing_packages])
                print(f'\nInstallation finished. {Fore.GREEN}[{checkmark}]{Fore.RESET}')
            else:
                print(f'\nRequirements already installed. {Fore.GREEN}[{checkmark}]{Fore.RESET}')
        except Exception as ex:
            print('\nAn exception occurred: \n', ex)

    @staticmethod
    def is_installed(package_name: str):
        try:
            import importlib
            importlib.import_module(package_name)
            return True
        except ImportError:
            return False