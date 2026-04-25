import json
import pika
from services.log_analysis import LogAnalyser
from models.logdata import log_data


def callback():
    analyzer: LogAnalyser = LogAnalyser()

    def process(ch, method, properties, body):
        data: log_data = log_data(**json.loads(body))
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
