import logging
import os

import pymongo
from dotenv import load_dotenv
from pymongo import MongoClient

load_dotenv("../.env")
logging.basicConfig(level=logging.INFO)
MONGO_URL = os.getenv("MONGO_URL")


def connect_db():
    try:
        client = pymongo.MongoClient(MONGO_URL)
    except Exception as e:
        logging.error(f"Error connecting to MongoDB: {e}")
        return None
    return client
