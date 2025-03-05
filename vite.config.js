import { defineConfig } from 'vite';

export default defineConfig({
  base: '/',
  build: {
    outDir: 'dist'
  },
  publicDir: 'public' // This is default, but good to be explicit
});