import datetime
from collections import deque
from models.logdata import log_data


class LogAnalyser:
    queues: dict[str, deque[log_data]]
    max_window: datetime.timedelta
    latest: datetime.datetime

    def __init__(self) -> None:
        self.queues = dict()
        self.max_window = datetime.timedelta(minutes=5)
        self.latest = datetime.datetime.now() - datetime.timedelta(weeks=10)

    def add_logdata(self, data: log_data) -> None:
        if data.level.lower() not in ["error", "err"]:
            return

        if data.timestamp > self.latest:
            self.latest = data.timestamp

        cutoff: datetime.datetime = self.latest - self.max_window
        if data.timestamp < cutoff:
            return

        key: str = data.service.lower() + "--" + data.level.lower()
        if key not in self.queues.keys():
            self.queues[key] = deque()
        self.queues[key].append(data)

        self.queues[key] = deque(
            log for log in self.queues[key]
            if log.timestamp >= cutoff
        )

        self.check_counts(key)
        print("DEBUG:: added:", data)

    def check_timespan_count(self, key: str, delta: datetime.timedelta = datetime.timedelta(minutes=1)) -> int:
        cutoff: datetime.datetime = self.latest - delta
        count: int = 0
        for log in reversed(self.queues[key]):
            if log.timestamp < cutoff:
                break
            count += 1
        return count

    def check_counts(self, key: str) -> None:
        if len(self.queues[key]) > 7 or self.check_timespan_count(key) >= 3:
            self.raise_alarm(key)

    def raise_alarm(self, key: str) -> None:
        print("Alarm raised for:", key.replace("--", " "))
