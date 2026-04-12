from consumers.logdata import start_consumer
from config.config import config


def main():
    conf = config()
    start_consumer(conf)


if __name__ == '__main__':
    main()
