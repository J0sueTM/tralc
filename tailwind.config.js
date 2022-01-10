const production = !process.env.ROLLUP_WATCH

module.exports = {
  purge: {
    content: [
      './src/svelte/App.svelte'
    ],
    enabled: production
  },
  darkMode: 'class',
  theme: {
    extend: {},
  },
  variants: {
    extend: {},
  },
  plugins: [],
  future: {
    purgeLayersByDefault: true,
    removeDeprecatedGapUtilities: true
  }
}
