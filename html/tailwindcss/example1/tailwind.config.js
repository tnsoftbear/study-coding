/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./public/**/*.{html,js}"],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        mainColor: "yellow",
      },
      padding: {
        big: '50px',
      },
      screens: {
        xm: '480px',
        md: '768px',
        lg: '1024px',
        xl: '1280px',
      }
    },
  },
  plugins: [],
}

