// src/utils/validator.js
// SAFE: Proper input validation
function validateEmail(email) {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return emailRegex.test(email);
}

// SAFE: Parameterized query
async function getUserById(pool, userId) {
  const query = 'SELECT * FROM users WHERE id = $1';
  const result = await pool.query(query, [userId]);
  return result.rows[0];
}

module.exports = { validateEmail, getUserById };

