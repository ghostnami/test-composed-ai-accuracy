// src/services/apiClient.js
const https = require('https');

// VULNERABILITY 11: SSRF
function fetchUserData(url) {
  return new Promise((resolve, reject) => {
    https.get(url, (res) => {
      let data = '';
      res.on('data', (chunk) => {
        data += chunk;
      });
      res.on('end', () => resolve(data));
    }).on('error', reject);
  });
}

module.exports = { fetchUserData };

