import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  async rewrites() {
    return [
      {
        source: '/api/backend/:path*',
        destination: 'http://localhost:8080/zenith-financial/:path*',
      },
    ];
  },
};

export default nextConfig;
