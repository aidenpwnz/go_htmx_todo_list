/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: "jit",
  darkMode: "media",
  content: [
    "./internal/views/**/*.templ",
    "./dist/output.css",
    "./node_modules/flowbite/**/*.js",
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ["Inter var"],
      },
    },
  },
  plugins: [require("flowbite/plugin")],
};
