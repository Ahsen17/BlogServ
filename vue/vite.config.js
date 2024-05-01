import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    open: true, // 启动服务自动访问主页
    port: 8088,
  },
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'), // 声明项目根目录
    }
  },
  plugins: [vue()],
})
