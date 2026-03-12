import datetime
import json
import logging
import os

import pandas as pd
from db import connect_db
from dotenv import load_dotenv
from mimesis import Fieldset
from mimesis.locales import Locale

logging.basicConfig(level=logging.INFO)
load_dotenv("../.env")
DB_NAME = os.getenv("MONGO_DB_NAME")


def gen_customers(locales, amount):
    fs = Fieldset(locales, amount)
    df = pd.DataFrame.from_dict(
        {
            "customer_id": fs("uuid"),
            "first_name": fs("first_name"),
            "last_name": fs("last_name"),
            "email": fs("email"),
            "city": fs("city"),
            "phone": fs("telephone", mask="+84 (9##) ###-####"),
            "birthdate": [
                dt.strftime("%Y-%m-%d")
                for dt in fs("birthdate", min_year=1950, max_year=2026)
            ],
            "created_at": fs("datetime", start=2020, end=2026),
            "updated_at": fs("datetime", start=2020, end=2026),
        }
    )
    return df.to_dict(orient="records")


def gen_accounts(locales, amount):
    fs = Fieldset(locales, amount)
    df = pd.DataFrame.from_dict(
        {
            "account_id": fs("uuid"),
            "customer_id": fs("uuid"),
            "account_type": fs("choice", items=["savings", "checking"]),
            "balance": fs("float_number", start=0, end=10000, precision=2),
            "created_at": fs("datetime", start=2020, end=2026),
            "updated_at": fs("datetime", start=2020, end=2026),
        }
    )
    return df.to_dict(orient="records")


if __name__ == "__main__":
    customers = gen_customers(Locale.EN, 10)
    db = connect_db()
    # print(json.dumps(customers, indent=2, default=str))
    if db:
        logging.info("Successfully connected to the database. Inserting data...")
        customers_collection = db[DB_NAME]["customers"]
        customers_collection.insert_many(customers)
        logging.info("Data insertion completed with %d records.", len(customers))
    else:
        logging.error("Failed to connect to the database. Data insertion aborted.")
