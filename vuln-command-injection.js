const { exec } = require('child_process');

// VULNERABLE: Command injection
function backupDatabase(filename) {
  exec(`tar -czf ${filename} /var/db/`);
}

module.exports = { backupDatabase };
