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
        'primaryColor': "rgb(191 219 254)",
        'secondaryColor': "rgb(167 139 250)",
        'secondaryColor2': "rgb(91 33 182)",
        'tertiaryColor': "rgb(190 242 100)",
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

