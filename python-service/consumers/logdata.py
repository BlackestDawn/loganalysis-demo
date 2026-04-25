import json
import pika
import datetime
from services.log_analysis import LogAnalyser
from models.logdata import log_data


def callback():
    analyzer: LogAnalyser = LogAnalyser()

    def process(ch, method, properties, body):
        decoded: dict = json.loads(body)
        data: log_data = log_data(
            timestamp=datetime.datetime.strptime(decoded["timestamp"], "%Y-%m-%dT%H:%M:%SZ"),
            level=decoded["level"],
            message=decoded["message"],
            service=decoded["service"]
        )
        analyzer.add_logdata(data)
    return process


def start_consumer(conf):
    print("connecting to: ", conf.queue.conn.host + ":" + str(conf.queue.conn.port))
    connection = pika.BlockingConnection(pika.ConnectionParameters(
        host=conf.queue.conn.host,
        port=conf.queue.conn.port,
        credentials=pika.PlainCredentials(conf.queue.conn.user, conf.queue.conn.password)
    ))

    channel = connection.channel()
    channel.exchange_declare(exchange=conf.queue.exchange, exchange_type='fanout', durable=True)
    channel.queue_declare(queue=conf.queue.queue)
    channel.queue_bind(exchange=conf.queue.exchange, queue=conf.queue.queue)

    channel.basic_consume(queue=conf.queue.queue, on_message_callback=callback(), auto_ack=True)

    print(' [*] Waiting for logs. To exit press CTRL+C')
    channel.start_consuming()
