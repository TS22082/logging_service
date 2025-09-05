import { defineConfig } from "vite";
import { resolve } from "path";

export default defineConfig({
  root: "static",
  build: {
    outDir: "dist/scripts/",
    emptyOutDir: true,
    rollupOptions: {
      input: {
        stream_test_page: resolve(
          __dirname,
          "static/src/scripts/stream_test_page.ts"
        ),
        docs_page: resolve(__dirname, "static/src/scripts/docs_page.ts"),
        login_page: resolve(__dirname, "static/src/scripts/login_page.ts"),
        home_page: resolve(__dirname, "static/src/scripts/home_page.ts"),
      },
      output: {
        entryFileNames: "[name].js",
        chunkFileNames: "[name]-[hash].js",
      },
    },
  },
});
