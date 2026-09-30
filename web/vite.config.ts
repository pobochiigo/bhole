import { defineConfig } from 'vite';

export default defineConfig({
  // A relative base keeps asset URLs valid both at a domain root and under a
  // project path such as https://<owner>.github.io/bhole/ on GitHub Pages.
  base: './',
});
