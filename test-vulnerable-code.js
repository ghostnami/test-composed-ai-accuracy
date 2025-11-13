// Vulnerable code for testing progress comment update
const express = require('express');
const { exec } = require('child_process');
const app = express();

// SQL injection vulnerability
app.get('/user', async (req, res) => {
  const userId = req.query.id;
  const query = `SELECT * FROM users WHERE id = ${userId}`; // VULNERABLE
  const result = await pool.query(query);
  res.json(result.rows[0]);
});

// Command injection vulnerability
app.post('/run', (req, res) => {
  const command = req.body.command;
  exec(command, (error, stdout) => { // VULNERABLE
    res.send(stdout);
  });
});

// XSS vulnerability
app.get('/search', (req, res) => {
  const query = req.query.q;
  res.send(`<h1>Results for: ${query}</h1>`); // VULNERABLE - no escaping
});

module.exports = app;
