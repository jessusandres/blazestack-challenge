'use client';

import { useContext, useEffect, useState } from 'react';
import Link from 'next/link';
import { redirect, RedirectType } from 'next/navigation';

/* External */
import axios from 'axios';

/* Project */
import { AuthContext } from '@/context/auth.context';

export default function LoginPage() {
  const { user, setUser } = useContext(AuthContext);
  const [loading, setLoading] = useState(true);

  const [email, setEmail] = useState('demo@example.com');
  const [password, setPassword] = useState('admin');
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const saveUser = (user: { token: string; name: string; email: string }) => {
    localStorage.setItem('token', user.token);

    setUser({
      token: user.token,
      name: user.name,
      email: user.email,
    });
  };

  const getProfile = async (token: string) => {
    try {
      const { data } = await axios.get(
        'http://localhost:8080/api/v1/auth/profile',
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        },
      );

      if (data?.data?.token) {
        saveUser({
          token: data.data.token,
          name: data.data.name,
          email: data.data.email,
        });
      }
    } catch (err) {
      setLoading(false);
    }
  };

  useEffect(() => {
    const token = localStorage.getItem('token');
    const userToken = user?.token;

    if (!userToken && !token) {
      setLoading(false);
      return;
    }

    if (userToken) {
      setLoading(false);
      redirect('/', RedirectType.replace);
    }

    if (token) {
      getProfile(token);
    }
  }, [getProfile, user]);

  async function onSubmit(e: React.FormEvent) {
    e.preventDefault();
    setError(null);
    setIsSubmitting(true);

    try {
      const { data } = await axios.post(
        'http://localhost:8080/api/v1/auth/login',
        {
          email,
          password,
        },
      );

      if (data?.data?.token) {
        saveUser({
          token: data.data.token,
          name: data.data.name,
          email: data.data.email,
        });
      } else {
        throw new Error('Login failed');
      }
    } catch (err) {
      setError('Login failed. Please try again.');
    } finally {
      setIsSubmitting(false);
    }
  }

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="min-h-screen flex items-center justify-center p-6">
      <div className="w-full max-w-sm border rounded-xl p-6 shadow-sm bg-white dark:bg-zinc-900 border-zinc-200 dark:border-zinc-800">
        <h1 className="text-2xl font-semibold mb-1">
          Welcome to Blaze Incidents App
        </h1>
        <p className="text-sm text-zinc-500 mb-6">Sign in</p>
        <form onSubmit={onSubmit} className="space-y-4">
          <div>
            <label htmlFor="email" className="block text-sm mb-1">
              Email
            </label>
            <input
              id="email"
              type="email"
              required
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="w-full rounded-md border border-zinc-300 dark:border-zinc-700 bg-transparent px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="you@example.com"
            />
          </div>
          <div>
            <label htmlFor="password" className="block text-sm mb-1">
              Password
            </label>
            <input
              id="password"
              type="password"
              required
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full rounded-md border border-zinc-300 dark:border-zinc-700 bg-transparent px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="••••••••"
            />
          </div>
          {error && (
            <p className="text-sm text-red-600" role="alert">
              {error}
            </p>
          )}
          <button
            type="submit"
            disabled={isSubmitting}
            className="w-full rounded-md bg-blue-600 text-white py-2.5 font-medium hover:bg-blue-700 disabled:opacity-50"
          >
            {isSubmitting ? 'Signing in…' : 'Sign in'}
          </button>
        </form>
        <p className="text-xs text-zinc-500 mt-4">
          Don’t have an account?
          <Link href="#" className="underline">
            Sign up
          </Link>
        </p>
      </div>
    </div>
  );
}
