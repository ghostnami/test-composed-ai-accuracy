// src/services/db.js
// Mock database module for testing
class Database {
  async query(sql) {
    return { rows: [] };
  }

  raw(sql) {
    return this.query(sql);
  }
}

module.exports = new Database();

