/* eslint-env node */
require("@rushstack/eslint-patch/modern-module-resolution");

module.exports = {
  root: true,
  extends: [
    "plugin:vue/vue3-essential",
    "eslint:recommended",
    "@vue/eslint-config-typescript",
    "@vue/eslint-config-prettier",
    "plugin:@intlify/vue-i18n/recommended",
  ],
  parserOptions: {
    ecmaVersion: "latest",
  },
  rules: {
    "vue/multi-word-component-names": 0,
    "prettier/prettier": ["warn", { endOfLine: "auto" }],
    "@intlify/vue-i18n/no-raw-text": [
      "error",
      {
        ignorePattern: "^[!@#$%^&*()_\\-+=№;:?{}]+$",
        ignoreText: [],
      },
    ],
  },
  ignorePatterns: ["dist/*", "node_modules/*"],
  settings: {
    "vue-i18n": {
      localeDir: "./path/to/locales/*.{json,json5,yaml,yml}",
      messageSyntaxVersion: "^9.0.0",
    },
  },
};
