const collectionName = "transactions";
const transaction_id_idx_name_1 = "transaction_id_idx_name_1";
const transaction_from_account_id_idx_name_1 = "transaction_from_account_id_idx_name_1";
const transaction_to_account_id_idx_name_1 = "transaction_to_account_id_idx_name_1";
const transaction_date_issued_idx_name_sub1 = "transaction_date_issued_idx_name_sub1";

function log(message) {
    console.log(`[${new Date().toISOString()}] ${message}`);
}
module.exports = {
    /**
     * @param db {import('mongodb').Db}
     * @param client {import('mongodb').MongoClient}
     * @returns {Promise<void>}
     */
    async up(db, client) {

        try {
            await db.collectionName(collectionName)
                .createIndex({ transaction_id: 1 },
                    { name: transaction_id_idx_name_1 });
            await db.collectionName(collectionName)
                .createIndex({ from_account_id: 1 },
                    { name: transaction_from_account_id_idx_name_1 });
            await db.collectionName(collectionName)
                .createIndex({ to_account_id: 1 },
                    { name: transaction_to_account_id_idx_name_1 });
            await db.collectionName(collectionName)
                .createIndex({ date_issued: -1 },
                    { name: transaction_date_issued_idx_name_sub1 });
            log(`Indexes created on ${collectionName} collection successfully.`);

        } catch (error) {
            log(`error creating index on ${collectionName}: ${error.message}`);
        }
    },

    /**
     * @param db {import('mongodb').Db}
     * @param client {import('mongodb').MongoClient}
     * @returns {Promise<void>}
     */
    async down(db, client) {

        try {
            await db.collectionName(collectionName)
                .dropIndex({ name: transaction_id_idx_name_1 });

            await db.collectionName(collectionName)
                .dropIndex({ name: transaction_from_account_id_idx_name_1 });

            await db.collectionName(collectionName)
                .createIndex({ name: transaction_to_account_id_idx_name_1 });

            await db.collectionName(collectionName)
                .dropIndex({ name: transaction_date_issued_idx_name_sub1 });

            log(`Indexes created on ${collectionName} collection successfully.`);

        } catch (error) {
            log(`error deleting index on ${collectionName}: ${error.message}`);
        }
    }
};
