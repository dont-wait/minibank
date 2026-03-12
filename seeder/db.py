import logging
import os

import pymongo
from dotenv import load_dotenv
from pymongo import MongoClient

load_dotenv("../.env")
logging.basicConfig(level=logging.INFO)
MONGO_URL = os.getenv("MONGO_URL")
DB_NAME = os.getenv("MONGO_DB_NAME")


def connect_db():
    try:
        client = pymongo.MongoClient(MONGO_URL)
    except Exception as e:
        logging.error(f"Error connecting to MongoDB: {e}")
        return None
    return client


def clean_db():
    try:
        client = pymongo.MongoClient(MONGO_URL)
    except Exception as e:
        logging.error(f"Error connecting to MongoDB: {e}")
        return None

    client.db[DB_NAME]["customers"].drop()
    client.db[DB_NAME]["accounts"].drop()
    client.db[DB_NAME]["loans"].drop()
    client.db[DB_NAME]["transactions"].drop()
    client.db[DB_NAME]["branches"].drop()
    logging.info("Database cleaned.")
