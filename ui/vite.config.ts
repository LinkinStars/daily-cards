import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  base: '/card',
  plugins: [vue()],
  build: {
    outDir: "../static/"
  },
  resolve: {
    alias: {
      '@/': `${path.resolve(__dirname, 'src')}/`
    }
  }
})
