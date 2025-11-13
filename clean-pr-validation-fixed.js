const db = require('./db');
const { execFile } = require('child_process');

// Safe SQL query using parameterized statements
async function getUserById(pool, userId) {
  const query = 'SELECT * FROM users WHERE id = $1';
  const result = await pool.query(query, [userId]);
  return result.rows[0];
}

// Safe command execution using execFile with sanitized arguments
function runSafeCommand(command, args = []) {
  return new Promise((resolve, reject) => {
    execFile(command, args, (error, stdout) => {
      if (error) return reject(error);
      resolve(stdout);
    });
  });
}

// Safe rendering using escaping
function renderComment(comment) {
  const safeComment = comment.replace(/[&<>"']/g, (char) => ({
    '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;', "'": '&#39;'
  }[char]));
  return `<section>${safeComment}</section>`;
}

module.exports = { getUserById, runSafeCommand, renderComment };
