'use client';
import { useContext } from 'react';

/* Project */
import { redirect, RedirectType } from 'next/navigation';
import IncidentsTable from '@/components/IncidentsTable';
import { AuthContext } from '@/context/auth.context';

export default function App() {
  const currentUser = useContext(AuthContext);

  if (!currentUser?.user?.token) {
    redirect('/login', RedirectType.replace);
  }

  return (
    <div className="p-6">
      <h1 className="text-2xl font-semibold mb-4">Incidents App</h1>
      <IncidentsTable />
    </div>
  );
}
