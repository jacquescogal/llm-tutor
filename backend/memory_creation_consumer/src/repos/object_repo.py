from src.db.object_store import ObjectStore

class ObjectRepository:
    def __init__(self, object_store: ObjectStore):
        self.object_store = object_store

    def get_object(self, object_key):
        return self.object_store.get_object(object_key)

    def delete_object(self, object_key):
        self.object_store.delete_object(object_key)
