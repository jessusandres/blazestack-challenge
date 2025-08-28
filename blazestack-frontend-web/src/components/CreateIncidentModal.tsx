'use client';

import { useEffect, useState, useCallback } from 'react';

/* External */
import axios from 'axios';

export type CreateIncidentModalProps = {
  incidentTypes: string[];
  isOpen: boolean;
  onClose: () => void;
  onCreated?: () => Promise<void> | void;
};

export default function CreateIncidentModal({
  isOpen,
  onClose,
  onCreated,
  incidentTypes,
}: CreateIncidentModalProps) {
  const [title, setTitle] = useState('');
  const [incidentType, setIncidentType] = useState('');
  const [description, setDescription] = useState('');
  const [location, setLocation] = useState('');
  const [imageFile, setImageFile] = useState<File | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const resetForm = useCallback(() => {
    setTitle('');
    setIncidentType('');
    setDescription('');
    setLocation('');
    setImageFile(null);
    setError(null);
  }, []);

  useEffect(() => {
    if (!isOpen) return;

    const onKey = (ev: KeyboardEvent) => {
      if (ev.key === 'Escape') onClose();
    };

    window.addEventListener('keydown', onKey);

    return () => window.removeEventListener('keydown', onKey);
  }, [isOpen, onClose]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    if (!title.trim() || !incidentType.trim()) {
      setError('Title and Incident Type are required.');

      return;
    }

    setLoading(true);
    try {
      const formData = new FormData();
      formData.append('title', title.trim());
      formData.append('incidentType', incidentType.trim());
      if (description.trim())
        formData.append('description', description.trim());
      if (location.trim()) formData.append('location', location.trim());
      if (imageFile) formData.append('image', imageFile);

      const res = await axios.post(
        'http://localhost:8080/api/v1/incidents',
        formData,
        {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`,
          },
          validateStatus: () => true,
        },
      );

      if (res.status < 200 || res.status >= 300) {
        const errMsg =
          (res.data && (res.data.message as string)) ||
          'Failed to create incident';
        throw new Error(errMsg);
      }

      if (onCreated) await onCreated();
      resetForm();
      onClose();
    } catch (e: unknown) {
      if (typeof e === 'object' && e !== null && 'message' in e) {
        const msg = (e as { message?: unknown }).message;
        setError(typeof msg === 'string' ? msg : 'Unexpected error');
      } else {
        setError('Unexpected error');
      }
    } finally {
      setLoading(false);
    }
  };

  if (!isOpen) return null;

  return (
    <div
      className="fixed inset-0 z-50 flex items-center justify-center"
      role="dialog"
      aria-modal="true"
    >
      <div className="absolute inset-0 bg-black/50" onClick={onClose} />
      <div className="relative z-10 w-full max-w-md rounded-lg border border-zinc-200 dark:border-zinc-800 bg-white dark:bg-zinc-900 p-4 shadow-xl">
        <div className="flex items-center justify-between mb-3">
          <h3 className="text-base font-semibold">Create Incident</h3>
          <button
            className="rounded-md p-1 text-zinc-500 hover:text-zinc-900 dark:hover:text-white"
            onClick={onClose}
            aria-label="Close"
          >
            ✕
          </button>
        </div>
        {error && <div className="mb-3 text-sm text-red-600">{error}</div>}
        <form onSubmit={handleSubmit} className="space-y-3">
          <div>
            <label className="block text-sm mb-1" htmlFor="title">
              Title
            </label>
            <input
              id="title"
              type="text"
              className="w-full rounded-md border border-zinc-300 dark:border-zinc-700 bg-white dark:bg-zinc-950 px-3 py-2 text-sm"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              required
            />
          </div>
          <div>
            <label className="block text-sm mb-1" htmlFor="incidentType">
              Incident Type
            </label>
            <select
              id="incidentType"
              className="w-full rounded-md border border-zinc-300 dark:border-zinc-700 bg-white dark:bg-zinc-950 px-3 py-2 text-sm"
              value={incidentType}
              onChange={(e) => setIncidentType(e.target.value)}
              required
            >
              <option value="" disabled>
                Select type…
              </option>
              {incidentTypes.map((type) => (
                <option key={type} value={type}>
                  {type}
                </option>
              ))}
            </select>
          </div>
          <div>
            <label className="block text-sm mb-1" htmlFor="description">
              Description
            </label>
            <textarea
              id="description"
              className="w-full min-h-[80px] rounded-md border border-zinc-300 dark:border-zinc-700 bg-white dark:bg-zinc-950 px-3 py-2 text-sm"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              placeholder="Optional description"
            />
          </div>
          <div>
            <label className="block text-sm mb-1" htmlFor="location">
              Location
            </label>
            <input
              id="location"
              type="text"
              className="w-full rounded-md border border-zinc-300 dark:border-zinc-700 bg-white dark:bg-zinc-950 px-3 py-2 text-sm"
              value={location}
              onChange={(e) => setLocation(e.target.value)}
              placeholder="Optional location"
            />
          </div>
          <div>
            <label className="block text-sm mb-1" htmlFor="image">
              Image (optional)
            </label>
            <input
              id="image"
              type="file"
              accept="image/*"
              className="block w-full text-sm"
              onChange={(e) => setImageFile(e.target.files?.[0] || null)}
            />
          </div>
          <div className="flex items-center justify-end gap-2 pt-2">
            <button
              type="button"
              className="rounded-md border border-zinc-300 dark:border-zinc-700 px-3 py-2 text-sm"
              onClick={() => {
                onClose();
                resetForm();
              }}
              disabled={loading}
            >
              Cancel
            </button>
            <button
              type="submit"
              className="rounded-md bg-blue-600 hover:bg-blue-700 text-white px-3 py-2 text-sm disabled:opacity-60"
              disabled={loading}
            >
              {loading ? 'Creating...' : 'Create'}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
