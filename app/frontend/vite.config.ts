import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import { resolve } from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      '@assets': resolve(__dirname, './src/assets'),
      '@components': resolve(__dirname, './src/components'),
      '@stores': resolve(__dirname, './src/stores'),
      '@utils': resolve(__dirname, './src/utils'),
      '@wailsjs/runtime': resolve(__dirname, './wailsjs/runtime/runtime'),
      '@wailsjs/go': resolve(__dirname, './wailsjs/go')
    }
  },
  plugins: [svelte()]
})
