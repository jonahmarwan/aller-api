import os
from multiprocessing.connection import Listener

SOCK_PATH = r'\\.\pipe\my_pipe'

with Listener(SOCK_PATH, family='AF_PIPE') as listener:
    print(f"Listening on {SOCK_PATH}...")
    while True:
        with listener.accept() as conn:
            print("Connection accepted from", listener.last_accepted)
            while True:
                try:
                    msg = conn.recv()
                    print("Received:", msg)
                except EOFError:
                    print("Connection closed.")
                    break