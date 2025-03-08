import os
import json
import base64
import sqlite3
import shutil
import time
import itertools
import threading
import sys
from colorama import Fore
from datetime import datetime, timedelta
import win32.win32crypt as win32crypt  # pypiwin32
from Crypto.Cipher import AES  # pycryptodome

banner = """\n                                     .-. .-')     ('-.                .-')    
                                     \  ( OO )  _(  OO)              ( OO ).  
   .-----.  .-'),-----.  .-'),-----. ,--. ,--. (,------. ,--.   ,--.(_)---\_) 
  '  .--./ ( OO'  .-.  '( OO'  .-.  '|  .'   /  |  .---'  \  `.'  / /    _ |  
  |  |('-. /   |  | |  |/   |  | |  ||      /,  |  |    .-')     /  \  :` `.  
 /_) |OO  )\_) |  |\|  |\_) |  |\|  ||     ' _)(|  '--.(OO  \   /    '..`''.) 
 ||  |`-'|   \ |  | |  |  \ |  | |  ||  .   \   |  .--' |   /  /\_  .-._)   \ 
(_'  '--'\    `'  '-'  '   `'  '-'  '|  |\   \  |  `---.`-./  /.__) \       / 
   `-----'      `-----'      `-----' `--' '--'  `------'  `--'       `-----'  
   """


def get_chrome_datetime(chromedate):
    """Grab datetime.datetime object from chrome format datetime"""
    if chromedate != 86400000000 and chromedate:
        try:
            return datetime(1601, 1, 1) + timedelta(microseconds=chromedate)
        except Exception as e:
            print(f'[ERROR] {e}, chromedate: {chromedate}')
            return chromedate
    else:
        return ""


def get_encryption_key():
    local_state_path = os.path.join(os.environ['USERPROFILE'],
                                    'AppData', 'Local', 'Google', 'Chrome',
                                    'User Data', 'Local State')
    with open(local_state_path, 'r', encoding='utf-8') as f:
        local_state = f.read()
        local_state = json.loads(local_state)

    # decode encryption key from Base64
    key = base64.b64decode(local_state['os_crypt']['encrypted_key'])
    # remove 'DPAPI' str
    key = key[5:]
    # return decrypted key that was originally encrypted
    # using a session key derived from current user's logon credentials
    return win32crypt.CryptUnprotectData(key, None, None, None, 0)[1]


def decrypt_data(data, key):
    try:
        # get the initialization vector
        iv = data[3:15]
        data = data[15:]
        # generate cipher
        cipher = AES.new(key, AES.MODE_GCM, iv)
        # decrypt password
        return cipher.decrypt(data)[:-16].decode()
    except:
        try:
            return str(win32crypt.CryptUnprotectData(data,
                                                     None, None, None, 0)[1])
        except:
            # not supported
            return ""


# loading animation


def main():

    print(Fore.RED + banner + Fore.RESET)
    print(Fore.GREEN + '\n        Do you want a cookie?\n' + Fore.RESET)
    time.sleep(2)

    done = False

    def animate():
        for c in itertools.cycle(['|', '/', '-', '\\']):
            if done:
                break
            sys.stdout.write(
                Fore.YELLOW + '\r        Grabbing you some... ' + Fore.RESET + c)
            sys.stdout.flush()
            time.sleep(0.1)
    sys.stdout.write('\rDone!     ')

    t = threading.Thread(target=animate)
    t.start()
    time.sleep(4)
    done = True

    # local sqlite Chrome cookie database path
    db_path = os.path.join(os.environ["USERPROFILE"], "AppData", "Local",
                           "Google", "Chrome", "User Data", "Default",
                           "Network", "Cookies")
    # Copy file to current directory
    # as the DB will be locked if chrome is currently open
    filename = 'Cookies.db'
    if not os.path.isfile(filename):
        # copy file when it does not exist in the current directory
        shutil.copyfile(db_path, filename)

    db = sqlite3.connect(filename)
    # ignore decoding errors
    db.text_factory = lambda b: b.decode(errors='ignore')
    cursor = db.cursor()
    # get the cookies from 'cookies' table
    cursor.execute("""
    SELECT host_key, name, value, creation_utc, last_access_utc, expires_utc,
    encrypted_value FROM cookies""")
    # it is also possible to search by domain, e.g thepythoncode.com
    # cursor.execute("""
    # SELECT host_key, name, value, creation_utc, last_access_utc, expires_utc,
    # encrypted_value FROM cookies
    # WHERE host_key like '%thepythoncode.com%'""")

    # get AES key
    key = get_encryption_key()
    for host_key, name, value, creation_utc, \
            last_access_utc, expires_utc, encrypted_value in cursor.fetchall():
        if not value:
            decrypted_value = decrypt_data(encrypted_value, key)
        else:
            # already decrypted
            decrypted_value = value
        print(f"""\n
        Host: {host_key}
        Cookie name: {name}
        Cookie value (decrypted): {decrypted_value}
        Creation datetime (UTC): {get_chrome_datetime(creation_utc)}
        Last access datetime (UTC): {get_chrome_datetime(last_access_utc)}
        Expires datetime (UTC): {get_chrome_datetime(expires_utc)}
        ===============================================================
        """)
        # update the cookies table with the decrypted value
        # and make session cookie persistent
        cursor.execute("""
        UPDATE cookies SET value = ?, has_expires = 1,
        expires_utc = 99999999999999999, is_persistent = 1, is_secure = 0
        WHERE host_key = ?
        AND name = ?""", (decrypted_value, host_key, name))
    # commit changes
    db.commit()
    # close connection
    db.close()


if __name__ == "__main__":
    main()
