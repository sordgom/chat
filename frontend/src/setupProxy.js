const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function(app) {
  app.use(
    '/voice/*',
    createProxyMiddleware({
      target: 'http://localhost:8083',
      secure: false,
      logLevel: 'debug',
      changeOrigin: true,
    })
  );
};
