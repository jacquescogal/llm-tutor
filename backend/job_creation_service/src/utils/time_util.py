import time
from datetime import datetime
from src.utils.configuration import Configuration

class TimeUtil:
    @staticmethod
    def get_current_unix_time_from_epoch():
        current_unix_time = int(time.time())
        epoch_timestamp = Configuration.get_instance().epoch_timestamp
        return current_unix_time - epoch_timestamp