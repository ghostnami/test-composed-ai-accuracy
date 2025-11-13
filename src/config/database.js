// src/config/database.js
// VULNERABILITY 10: Hardcoded Database Credentials (Line 3)
const dbConfig = {
  host: 'localhost',
  port: 5432,
  user: 'admin',
  password: 'SuperSecretPassword123!',
  database: 'production_db'
};

module.exports = dbConfig;

