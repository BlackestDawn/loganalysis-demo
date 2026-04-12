from dataclasses import dataclass
from dotenv import load_dotenv, find_dotenv
import os


@dataclass
class pika_conn:
    host: str
    port: int
    user: str
    password: str


@dataclass
class queue_config:
    conn: pika_conn
    exchange: str
    key: str
    queue: str


class config():
    queue: queue_config

    def __init__(self):
        load_dotenv(find_dotenv())
        tmp_conn_string = os.getenv('QUEUE_SERVER')
        if tmp_conn_string is not None:
            conn_string = tmp_conn_string.split("/")[2]
        else:
            conn_string = "guest:guest@localhost:5672"

        pika = pika_conn(
            host=conn_string.split(':')[1].split("@")[1],
            port=int(conn_string.split(':')[2]),
            user=conn_string.split(':')[0],
            password=conn_string.split(':')[1].split("@")[0]
        )

        self.queue = queue_config(
            conn=pika,
            exchange=os.getenv('QUEUE_EXCHANGE') or "logs",
            key=os.getenv('QUEUE_KEY') or "logs",
            queue=os.getenv('QUEUE_NAME') or "logs"
        )
