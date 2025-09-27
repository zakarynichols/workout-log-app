import { useState } from "react";

type Props = {
  onCreated: () => void; // callback to refresh list after add
};

export default function AddSessionForm({ onCreated }: Props) {
  const [sessionDate, setSessionDate] = useState("");
  const [sessionType, setSessionType] = useState("");
  const [notes, setNotes] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError(null);

    try {
      const res = await fetch("/api/sessions", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          session_date: sessionDate,
          session_type: sessionType,
          notes,
        }),
      });

      if (!res.ok) throw new Error(`HTTP ${res.status}`);

      setSessionDate("");
      setSessionType("");
      setNotes("");
      onCreated(); // refresh parent list
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="p-4 border rounded mb-4 space-y-3">
      <h2 className="text-lg font-semibold">Add New Session</h2>

      <div>
        <label className="block text-sm">Date</label>
        <input
          type="date"
          value={sessionDate}
          onChange={(e) => setSessionDate(e.target.value)}
          className="border p-2 rounded w-full"
          required
        />
      </div>

      <div>
        <label className="block text-sm">Type</label>
        <input
          type="text"
          value={sessionType}
          onChange={(e) => setSessionType(e.target.value)}
          placeholder="push / pull / legs"
          className="border p-2 rounded w-full"
          required
        />
      </div>

      <div>
        <label className="block text-sm">Notes</label>
        <textarea
          value={notes}
          onChange={(e) => setNotes(e.target.value)}
          className="border p-2 rounded w-full"
        />
      </div>

      {error && <p className="text-red-500 text-sm">Error: {error}</p>}

      <button
        type="submit"
        disabled={loading}
        className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
      >
        {loading ? "Saving..." : "Add Session"}
      </button>
    </form>
  );
}
