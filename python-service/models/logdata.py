from dataclasses import dataclass
import datetime


@dataclass
class log_data:
    timestamp: datetime.datetime
    level: str
    message: str
    service: str
