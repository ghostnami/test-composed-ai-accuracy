// Clean code with no vulnerabilities
const express = require('express');
const { execFile } = require('child_process');

// Safe parameterized query
async function getUserById(pool, userId) {
  const query = 'SELECT * FROM users WHERE id = $1';
  const result = await pool.query(query, [userId]);
  return result.rows[0];
}

// Safe HTML escaping
function escapeHtml(text) {
  const map = {
    '&': '&amp;',
    '<': '&lt;',
    '>': '&gt;',
    '"': '&quot;',
    "'": '&#039;'
  };
  return text.replace(/[&<>"']/g, m => map[m]);
}

module.exports = { getUserById, escapeHtml };
