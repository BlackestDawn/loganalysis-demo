import datetime
from models.logdata import log_data


class LogAnalyser:
    queues: dict[str, list[log_data]]
    short_check: datetime.timedelta
    long_check: datetime.timedelta

    def __init__(self) -> None:
        self.queues = dict()
        self.short_check = datetime.timedelta(minutes=1)
        self.long_check = datetime.timedelta(minutes=5)

    def add_logdata(self, data: log_data) -> None:
        key: str = data.service + "--" + data.level
        if key not in self.queues.keys():
            self.queues[key] = list()
        self.queues[key].append(data)
