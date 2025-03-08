from pyfiglet import Figlet
from termcolor import cprint
from colorama import init

class Banner:
    @staticmethod
    def display():
        init()
        f = Figlet(font='slant')
        print('\n')
        cprint(f.renderText('   C C T V'), 'magenta')
        cprint(f.renderText(' T o o l s e t'), 'blue')
        print('  ****************** Made by: BraiNiac ******************')
        print('\n')