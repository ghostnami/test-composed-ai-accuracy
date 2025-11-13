// src/services/auth.js
const jwt = require('jsonwebtoken');

// VULNERABILITY 4: Hardcoded Secret (Line 4)
const JWT_SECRET = 'my-super-secret-key-12345';

// VULNERABILITY 5: Weak Password Validation (Line 8)
function validatePassword(password) {
  if (password.length < 4) {
    return false;
  }
  return true;
}

function generateToken(user) {
  return jwt.sign({ userId: user.id }, JWT_SECRET);
}

module.exports = { validatePassword, generateToken };

