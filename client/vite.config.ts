import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  base: "/",
  preview: {
    port: 80,
    strictPort: true,
  },
  server: {
    // for development
    proxy: {
      '/api': {
        target: 'http://server-dev:8080/api/v1',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
    port: 3000,
    strictPort: true,
    host: true,
    origin: "http://0.0.0.0:3000",
    hmr: {
      port: 3000,
    }
  },
})
