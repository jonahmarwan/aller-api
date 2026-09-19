from multiprocessing.connection import Client

SOCK_PATH = r'\\.\pipe\my_pipe'
with Client(SOCK_PATH, family='AF_PIPE') as conn:
    conn.send("Hello from the client!")