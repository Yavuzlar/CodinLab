/** @type {import('next').NextConfig} */
const i18n = {
  locales: ["en", "tr"],
  defaultLocale: "en",
};
const nextConfig = {
  reactStrictMode: false,
  swcMinify: true,
  webpack: (config) => {
    config.module.rules.push({
      test: /\.svg$/,
      use: ["@svgr/webpack"],
    });

    return config;
  },
  i18n,
};

module.exports = {
  nextConfig,
};
