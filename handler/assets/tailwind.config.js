/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["index.html"], 
  theme: {
    screens:{
      sm:'480PX',
      md:'768px',
      lg:'970px',
      xl: '1440px'
    },
    colors:{
      'blue': '#1fb6ff',
      'purple': '#7e5bef',
      'pink': '#ff49db',
      'orange': '#ff7849',
      'green': '#13ce66',
      'yellow': '#ffc82c',
      'gray-dark': '#273444',
      'gray': '#8492a6',
      'gray-light': '#d3dce6',
    },
    fontFamily: {
      sans: ['Helvetica', 'sans-serif'],
      serif: ['Merriweather', 'serif'],
      body: ['Play'],
    },
    extend: {
      spacing:{
        '128': '32rem',
        '144': '36rem',
      },
      borderRadius: {
        '4x1':'2rem',
      }
    },
  },
  plugins: [],
}