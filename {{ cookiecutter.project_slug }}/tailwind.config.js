/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./pkg/**/*.{html,js,templ}"],
  theme: {
    extend: {},
  },
  darkMode: "class",
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
  daisyui: {
    themes: ["garden", "dracula"],
  },
};
