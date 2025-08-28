'use client';

import { useState } from 'react';
import { Geist, Geist_Mono } from 'next/font/google';

/* Project */
import { AuthContext, IAuthUser } from '@/context/auth.context';
import './globals.css';

const geistSans = Geist({
  variable: '--font-geist-sans',
  subsets: ['latin'],
});

const geistMono = Geist_Mono({
  variable: '--font-geist-mono',
  subsets: ['latin'],
});

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const [currentUser, setCurrentUser] = useState<IAuthUser>({
    name: '',
    email: '',
    token: '',
  });

  const setUser = (user: IAuthUser) => {
    setCurrentUser(user);
  };

  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        <AuthContext.Provider
          value={{
            user: currentUser,
            setUser: setUser,
          }}
        >
          {children}
        </AuthContext.Provider>
      </body>
    </html>
  );
}
