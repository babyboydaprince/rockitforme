import os
import subprocess
import sys
from colorama import Fore

class DependencyChecker:
    INSTALLED_FLAG = ".installed"

    @staticmethod
    def is_package_installed(package_name):
        try:
            import importlib.util
            spec = importlib.util.find_spec(package_name)
            return spec is not None
        except (ImportError, AttributeError):
            return False

    @staticmethod
    def install_missing_packages(packages):
        if os.path.exists(DependencyChecker.INSTALLED_FLAG):
            return

        checkmark = '\u2713'
        missing_packages = [pkg for pkg in packages if not DependencyChecker.is_package_installed(pkg)]

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

                print(f'\nInstallation finished. {Fore.GREEN}[{checkmark}]{Fore.RESET}')
                time.sleep(2)

                for pkg in missing_packages:
                    if not DependencyChecker.is_package_installed(pkg):
                        print(f"\n{Fore.RED}[!] Error: Package {pkg} failed to install. Please check your internet connection and try again manually with: 'pip install {pkg}'{Fore.RESET}")
                        sys.exit(1)

                with open(DependencyChecker.INSTALLED_FLAG, "w") as f:
                    f.write("Dependencies installed.")

            except subprocess.CalledProcessError as e:
                print(f'\nAn exception occurred while installing packages: {e}')
                print(f"Installation of one or more packages failed. Check the error message above and try again manually.")
                sys.exit(1)
            except Exception as ex:
                print('\nAn exception occurred: \n', ex)
                sys.exit(1)
        else:
            print(f'\n             Requirements already installed. {Fore.GREEN}[{checkmark}]{Fore.RESET}')
            time.sleep(4)

            with open(DependencyChecker.INSTALLED_FLAG, "w") as f:
                f.write("Dependencies installed.")