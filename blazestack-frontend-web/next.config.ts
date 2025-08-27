import type { NextConfig } from 'next';

const nextConfig: NextConfig = {
  /* config options here */
  output: 'export',
  images: {
    // Allow remote images for the incidents thumbnails
    remotePatterns: [
      { protocol: 'https', hostname: 'via.placeholder.com' },
      { protocol: 'https', hostname: '**' },
    ],
    unoptimized: true, // needed for static export so next/image works without the Image Optimization API
  },
};

export default nextConfig;
