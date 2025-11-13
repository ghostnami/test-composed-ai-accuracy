// src/app.js
const express = require('express');
const userRoutes = require('./api/users');
const authService = require('./services/auth');

const app = express();
app.use(express.json());
app.use('/api/users', userRoutes);

app.listen(3000, () => {
  console.log('Server running on port 3000');
});

module.exports = app;

