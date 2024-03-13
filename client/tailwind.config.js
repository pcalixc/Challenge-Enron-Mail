/** @type {import('tailwindcss').Config} */
export default  {
  content: ["./app/**/*.html",
  "./src/**/*.vue", ],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        "secundary_dark": "#13161C", 
        "primary_dark": "#010101",
        "ligth_blue": "#1a1a40",
        "deep_blue": "#082032",
        "vivid_blue": "#1a1a40",
        "transparent_blue": "#28287552",
        "royal_purple": "#9292FF",
        "royal_blue": "#3A3AFF"
      },
  },
  plugins: [],
}
}

