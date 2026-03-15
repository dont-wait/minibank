const collectionName = "customers";
const customer_account_id_idx_name_1 = "customer_account_id_idx_name_1";
const customer_email_idx_name_1 = "customer_email_idx_name_1";
const customer_phone_idx_name_1 = "customer_phone_idx_name_1";

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
            await db.collection(collectionName)
                .createIndex({ account_id: 1 },
                    { name: customer_account_id_idx_name_1 });

            await db.collection(collectionName)
                .createIndex({ email: 1 },
                    { name: customer_email_idx_name_1 });

            await db.collection(collectionName)
                .createIndex({ phone: 1 },
                    { name: customer_phone_idx_name_1 });
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
            await db.collection(collectionName)
                .dropIndex(customer_account_id_idx_name_1);

            await db.collection(collectionName)
                .dropIndex(customer_email_idx_name_1);

            await db.collection(collectionName)
                .dropIndex(customer_phone_idx_name_1);

        } catch (error) {
            log(`error dropping index on ${collectionName}: ${error.message}`);
        }
    }
};
