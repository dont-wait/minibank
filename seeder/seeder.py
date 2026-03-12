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


def gen_accounts(locales, amount):
    fs = Fieldset(locales, amount)
    df = pd.DataFrame.from_dict(
        {
            "account_id": fs("uuid"),
            "balance": fs("integer_number", start=0, end=10_000_000),
            "pin_code": fs("integer_number", start=100000, end=999999),
            # 1: active, -1: closed, 0: frozen
            "status": fs("random.weighted_choice", choices={1: 0.7, -1: 0.2, 0: 0.1}),
            # 1: checking, 0: savings
            "type": fs("random.weighted_choice", choices={1: 0.6, 0: 0.4}),
            "created_at": fs("datetime", start=2020, end=2026),
            "updated_at": fs("datetime", start=2020, end=2026),
        }
    )
    return df.to_dict(orient="records")


def gen_customers(locales, amount):
    fs = Fieldset(locales, amount)
    customer_ids = fs("uuid")
    df = pd.DataFrame.from_dict(
        {
            "customer_id": customer_ids,
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
            "accounts": [gen_accounts(locales, 2) for _ in range(amount)],
        }
    )
    logging.info("Generated %d customers.", amount)
    return df.to_dict(orient="records")


def gen_branchs(locales, amount):
    fs = Fieldset(locales, amount)
    df = pd.DataFrame.from_dict(
        {
            "branch_id": fs("uuid"),
            "branch_name": fs("company"),
            "branch_location": fs("city"),
            "created_at": fs("datetime", start=2020, end=2026),
            "updated_at": fs("datetime", start=2020, end=2026),
        }
    )
    logging.info("Generated %d branches.", amount)
    return df.to_dict(orient="records")


def gen_loans(locales, amount, db):
    fs = Fieldset(locales, amount)
    date_issued = fs("datetime", start=2020, end=2026)

    customer_ids = list(
        db[DB_NAME]["customers"].aggregate([{"$sample": {"size": amount}}])
    )
    if not customer_ids:
        logging.error("No customers found in DB to generate loans.")
        return []

    df = pd.DataFrame.from_dict(
        {
            "loan_id": fs("uuid"),
            "customer_id": fs(
                "choice", items=[str(c["customer_id"]) for c in customer_ids]
            ),
            "loan_amount": fs("float_number", start=1000, end=50000, precision=2),
            # 1: personal, 2: mortgage, 3: auto
            "loan_type": fs("random.weighted_choice", choices={1: 0.5, 2: 0.3, 3: 0.2}), 
            "date_issued": date_issued,
            "created_at": date_issued,
            "updated_at": fs("datetime", start=2020, end=2026),
        }
    )
    logging.info("Generated %d loans.", amount)
    return df.to_dict(orient="records")


if __name__ == "__main__":
    db = connect_db()
    if not db:
        logging.error("Failed to connect to the database. Aborting.")
        exit(1)

    logging.info("Connected to database. Starting data insertion...")

    # 1. Branches
    branchs = gen_branchs(Locale.EN, 5)
    db[DB_NAME]["branchs"].insert_many(branchs)
    logging.info("Inserted %d branches.", len(branchs))

    # 2. Customers (gộp accounts bên trong)
    customers = gen_customers(Locale.EN, 10)
    db[DB_NAME]["customers"].insert_many(customers)
    logging.info("Inserted %d customers.", len(customers))

    # 3. Loans (cần customers đã có trong DB)
    loans = gen_loans(Locale.EN, 20, db)
    if loans:
        db[DB_NAME]["loans"].insert_many(loans)
        logging.info("Inserted %d loans.", len(loans))

    logging.info("Data insertion completed.")
