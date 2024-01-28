/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./public/**/*.{html,js}",
    "./node_modules/flowbite/**/*.js"
  ],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        'mainColorLight': "#E1EFFE",
        'contentScratchLight': "#E5E7EB",
      },
      // fontFamily: {
      //   main: ['Roboto', 'sans-serif'],
      // },
    },
  },
  plugins: [
    require('flowbite/plugin')
  ],
}

