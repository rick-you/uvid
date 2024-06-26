import { defineConfig } from "vitepress";

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "UVID",
  description: "Observable Platform for Frontend Websites",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: "Home", link: "/" },
      { text: "Get Started", link: "/documentation" },
    ],

    sidebar: [],

    socialLinks: [{ icon: "github", link: "https://github.com/rick-you/uvid" }],
    logo: "logo.svg",
  },
  head: [
    [
      "script",
      { type: "module", src: '/uvid.js' },
    ],
  ],
  ignoreDeadLinks: "localhostLinks",
});
