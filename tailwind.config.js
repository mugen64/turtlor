// tailwind.config.js

/** @type {import('tailwindcss').Config} */

const sans = [
  "Maven Pro",
  "system-ui",
  "-apple-system",
  "Lucida Sans",
  "Lucida Sans Regular",
  "Lucida Grande",
  "Lucida Sans Unicode",
  "BlinkMacSystemFont",
  "Roboto",
  "Geneva",
  "Verdana",
  "sans-serif",
];

module.exports = {
  content: ["./internal/app/ui/**/*.{html,js,svelte,ts,templ}"],
  darkMode: "class",
  theme: {
    fontFamily: {
      headline: sans,
      body: sans,
      sans,
      serif: ["Crimson Pro", ...sans],
      mono: [
        "DM Mono",
        "Space Mono",
        "JetBrains Mono",
        "Courier New",
        "Courier",
        "monospace",
      ],
    },
    extend: {
      colors: {
        primary: {
          50: "var(--color-primary-50)",
          100: "var(--color-primary-100)",
          200: "var(--color-primary-200)",
          300: "var(--color-primary-300)",
          400: "var(--color-primary-400)",
          500: "var(--color-primary-500)",
          600: "var(--color-primary-600)",
          700: "var(--color-primary-700)",
          800: "var(--color-primary-800)",
          900: "var(--color-primary-900)",
          950: "var(--color-primary-950)",
        },
        theme: {
          surface: {
            1: "var(--color-surface-50)",
            2: "var(--color-surface-100)",
            3: "var(--color-surface-200)",
            4: "var(--color-surface-300)",
            5: "var(--color-surface-400)",
            6: "var(--color-surface-500)",
            7: "var(--color-surface-600)",
          },
          foreground: {
            1: "var(--foreground-1)",
            2: "var(--foreground-2)",
            3: "var(--foreground-3)",
            4: "var(--foreground-4)",
          },
          border: {
            1: "var(--color-surface-100)",
            2: "var(--color-surface-200)",
            3: "var(--color-surface-300)",
          },
        },
      },

      screens: {
        xxs: "320px",
        xs: "480px",
      },
    },
  },
  plugins: [],
};
