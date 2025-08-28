'use client';

import { useCallback, useContext, useEffect, useState } from 'react';
import { redirect, RedirectType } from 'next/navigation';

/* External */
import axios from 'axios';

/* Project */
import IncidentsTable, { Incident } from '@/components/IncidentsTable';
import { AuthContext } from '@/context/auth.context';
import CreateIncidentModal from '@/components/CreateIncidentModal';

export default function App() {
  const { user, setUser } = useContext(AuthContext);

  if (!user?.token) {
    redirect('/login', RedirectType.replace);
  }

  const [incidentTypes, setIncidentTypes] = useState<string[]>([]);
  const [isOpen, setIsOpen] = useState(false);

  const fetchIncidents = useCallback(async () => {
    const { data } = await axios.get('http://localhost:8080/api/v1/incidents', {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`,
      },
    });

    setIncidents(data.data?.incidents || []);
  }, []);

  const [incidents, setIncidents] = useState<Incident[]>([]);

  const fetchIncidentTypes = useCallback(async () => {
    const { data } = await axios.get('http://localhost:8080/api/v1/', {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`,
      },
    });

    setIncidentTypes(data.data?.incidentTypes || []);
  }, []);

  const logout = () => {
    localStorage.removeItem('token');
    setUser({
      token: '',
      name: '',
      email: '',
    });
  };

  useEffect(() => {
    fetchIncidents();
    fetchIncidentTypes();
  }, []);

  return (
    <div className="p-6">
      <div className="overflow-x-auto">
        <div className="flex items-center justify-between mb-3">
          <h1 className="text-2xl font-semibold mb-4">Incidents</h1>
          <button
            onClick={() => setIsOpen(true)}
            className="inline-flex items-center gap-2 rounded-md bg-blue-900 hover:bg-blue-950 text-white text-sm px-3 py-2 cursor-pointer"
            aria-haspopup="dialog"
            aria-expanded={isOpen}
          >
            + Create Incident
          </button>
        </div>
      </div>

      <IncidentsTable incidents={incidents} />

      <CreateIncidentModal
        incidentTypes={incidentTypes}
        isOpen={isOpen}
        onClose={() => setIsOpen(false)}
        onCreated={fetchIncidents}
      />

      <div className="mt-6 flex justify-end">
        <button className="bg-blue-900 cursor-pointer p-2" onClick={logout}>
          Logout
        </button>
      </div>
    </div>
  );
}
