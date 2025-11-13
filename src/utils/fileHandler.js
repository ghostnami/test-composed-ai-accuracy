// src/utils/fileHandler.js
const fs = require('fs');
const path = require('path');

// VULNERABILITY 6: Path Traversal
function readUserFile(filename) {
  const filePath = path.join('./uploads/', filename);
  return fs.readFileSync(filePath, 'utf8');
}

// VULNERABILITY 7: Unsafe File Write
function saveUserData(userId, data) {
  const filePath = `./data/${userId}.json`;
  fs.writeFileSync(filePath, JSON.stringify(data));
}

module.exports = { readUserFile, saveUserData };

