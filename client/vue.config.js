module.exports = {
  devServer: {
    proxy: {
      "/login": {
        target: "http://localhost:8080",
        changeOrigin: true
      },
      "/signup": {
        target: "http://localhost:8080",
        changeOrigin: true
      },
      "/api": {
        target: "http://localhost:8080",
        changeOrigin: true
      },
      "/api/no-auth": {
        target: "http://localhost:8080",
        changeOrigin: true
      }
    }
  }
};
