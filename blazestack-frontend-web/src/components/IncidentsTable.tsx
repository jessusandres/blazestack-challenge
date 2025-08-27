'use client';

import Image from 'next/image';
import { useEffect, useState } from 'react';

export type Incident = {
  id: string | number;
  title: string;
  incidentType: string;
  date?: string | number | Date;
  image?: string | null;
};

function formatDate(value?: string | number | Date) {
  if (!value) return '-';
  try {
    const date = value instanceof Date ? value : new Date(value);
    if (isNaN(date.getTime())) return '-';
    return new Intl.DateTimeFormat(undefined, {
      year: 'numeric',
      month: 'short',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
    }).format(date);
  } catch {
    return '-';
  }
}

export default function IncidentsTable() {
  const [incidents, setIndicents] = useState<Incident[]>([]);

  const fetchIncidents = async () => {
    const response = await fetch('http://localhost:8080/api/v1/incidents', {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`,
      },
    });
    const data = await response.json();
    console.log({ data });
    setIndicents(data.data?.incidents || []);
  };

  useEffect(() => {
    fetchIncidents();
  }, []);

  return (
    <div className="overflow-x-auto">
      <table className="min-w-full border border-zinc-200 dark:border-zinc-800 rounded-lg overflow-hidden">
        <thead className="bg-zinc-50 dark:bg-zinc-900/40">
          <tr>
            <th className="text-left text-xs font-semibold text-zinc-600 dark:text-zinc-300 uppercase tracking-wide px-4 py-3 border-b border-zinc-200 dark:border-zinc-800">
              Thumbnail
            </th>
            <th className="text-left text-xs font-semibold text-zinc-600 dark:text-zinc-300 uppercase tracking-wide px-4 py-3 border-b border-zinc-200 dark:border-zinc-800">
              Title
            </th>
            <th className="text-left text-xs font-semibold text-zinc-600 dark:text-zinc-300 uppercase tracking-wide px-4 py-3 border-b border-zinc-200 dark:border-zinc-800">
              Incident Type
            </th>
            <th className="text-left text-xs font-semibold text-zinc-600 dark:text-zinc-300 uppercase tracking-wide px-4 py-3 border-b border-zinc-200 dark:border-zinc-800">
              Date
            </th>
          </tr>
        </thead>
        <tbody>
          {incidents.length === 0 ? (
            <tr>
              <td
                colSpan={4}
                className="px-4 py-6 text-center text-sm text-zinc-500"
              >
                No incidents to display.
              </td>
            </tr>
          ) : (
            incidents.map((inc) => (
              <tr
                key={inc.id}
                className="hover:bg-zinc-50/60 dark:hover:bg-zinc-900/30"
              >
                <td className="px-4 py-3 align-middle">
                  <div className="h-12 w-12 relative rounded-md overflow-hidden bg-zinc-100 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-800">
                    {inc.image ? (
                      <Image
                        src={`data:image/png;base64,${inc.image}`}
                        alt={inc.title ? `${inc.title} thumbnail` : 'Thumbnail'}
                        fill
                        sizes="48px"
                        className="object-cover"
                      />
                    ) : (
                      <div className="h-full w-full flex items-center justify-center text-xs text-zinc-400">
                        N/A
                      </div>
                    )}
                  </div>
                </td>
                <td className="px-4 py-3">
                  <span className="font-medium">{inc.title || '-'}</span>
                </td>
                <td className="px-4 py-3">
                  <span className="inline-flex items-center rounded-md border border-zinc-200 dark:border-zinc-800 px-2 py-0.5 text-xs text-zinc-700 dark:text-zinc-300">
                    {inc.incidentType || '-'}
                  </span>
                </td>
                <td className="px-4 py-3">
                  <span className="text-sm text-zinc-700 dark:text-zinc-300">
                    {formatDate(inc.date)}
                  </span>
                </td>
              </tr>
            ))
          )}
        </tbody>
      </table>
    </div>
  );
}
