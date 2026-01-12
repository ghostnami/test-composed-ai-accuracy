// VULNERABLE: Hardcoded credentials
const config = {
  apiKey: 'sk_live_1234567890abcdef',
  databasePassword: 'SuperSecret123!',
  jwtSecret: 'my-secret-key-dont-tell-anyone',
  awsAccessKey: 'AKIAIOSFODNN7EXAMPLE',
  awsSecretKey: 'wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY'
};

module.exports = config;
