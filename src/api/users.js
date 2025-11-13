// src/api/users.js
const express = require('express');
const { exec } = require('child_process');
const db = require('../services/db');

const router = express.Router();

// VULNERABILITY 1: SQL Injection
router.get('/user/:id', async (req, res) => {
  const userId = req.params.id;
  const query = `SELECT * FROM users WHERE id = ${userId}`;
  const result = await db.query(query);
  res.json(result.rows[0]);
});

// VULNERABILITY 2: Command Injection
router.post('/backup', (req, res) => {
  const dbName = req.body.database;
  exec(`pg_dump ${dbName} > backup.sql`, (error) => {
    if (error) {
      return res.status(500).send('Backup failed');
    }
    res.send('Backup complete');
  });
});

// VULNERABILITY 3: XSS
router.get('/search', (req, res) => {
  const query = req.query.q;
  res.send(`<h1>Results for: ${query}</h1>`);
});

module.exports = router;

