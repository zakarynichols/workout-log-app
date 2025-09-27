import { useEffect, useState } from "react";
import AddSessionForm from "./AddSessionForm";
import { Link } from "react-router-dom";

type Session = {
  id: number;
  session_date: string;
  session_type: string;
  notes: string;
};

export default function Sessions() {
  const [sessions, setSessions] = useState<Session[]>([]);
  const [loading, setLoading] = useState(true);
  const [editingId, setEditingId] = useState<number | null>(null);
  const [formData, setFormData] = useState<Partial<Session>>({});
  const [saving, setSaving] = useState(false);

  const loadSessions = () => {
    setLoading(true);
    fetch("/api/sessions")
      .then((res) => res.json())
      .then((data) => setSessions(data))
      .finally(() => setLoading(false));
  };

  const deleteSession = async (id: number) => {
    if (!confirm("Delete this session?")) return;

    try {
      await fetch(`/api/sessions/${id}`, { method: "DELETE" });
      loadSessions();
    } catch (err) {
      console.error("Failed to delete session:", err);
    }
  };

  const startEdit = (s: Session) => {
    setEditingId(s.id);
    setFormData({
      session_date: s.session_date.split("T")[0], // format for <input type="date" />
      session_type: s.session_type,
      notes: s.notes,
    });
  };

  const cancelEdit = () => {
    setEditingId(null);
    setFormData({});
  };

  const saveEdit = async () => {
    if (!editingId) return;
    setSaving(true);
    try {
      const res = await fetch(`/api/sessions/${editingId}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(formData),
      });
      if (!res.ok) throw new Error(`HTTP ${res.status}`);
      loadSessions();
      cancelEdit();
    } catch (err) {
      console.error("Failed to update session:", err);
    } finally {
      setSaving(false);
    }
  };

  useEffect(() => {
    loadSessions();
  }, []);

  return (
    <div className="p-4">
      <AddSessionForm onCreated={loadSessions} />

      {loading ? (
        <p>Loading sessions...</p>
      ) : (
        <ul className="space-y-2">
          {sessions.map((s) => (
            <li key={s.id} className="p-3 border rounded">
              {editingId === s.id ? (
                <div className="space-y-2">
                  <input
                    type="date"
                    value={formData.session_date || ""}
                    onChange={(e) => setFormData({ ...formData, session_date: e.target.value })}
                    className="border p-1 rounded w-full"
                  />
                  <select
                    value={formData.session_type || ""}
                    onChange={(e) => setFormData({ ...formData, session_type: e.target.value })}
                    className="border p-1 rounded w-full"
                  >
                    <option value="">Select type</option>
                    <option value="push">Push</option>
                    <option value="pull">Pull</option>
                    <option value="legs">Legs</option>
                    <option value="other">Other</option>
                  </select>
                  <textarea
                    placeholder="Notes"
                    value={formData.notes || ""}
                    onChange={(e) => setFormData({ ...formData, notes: e.target.value })}
                    className="border p-1 rounded w-full"
                  />
                  <div className="flex gap-2">
                    <button
                      onClick={saveEdit}
                      disabled={saving}
                      className="bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700 disabled:opacity-50"
                    >
                      {saving ? "Saving..." : "Save"}
                    </button>
                    <button
                      onClick={cancelEdit}
                      className="bg-gray-400 text-white px-3 py-1 rounded hover:bg-gray-500"
                    >
                      Cancel
                    </button>
                  </div>
                </div>
              ) : (
                <div className="flex justify-between items-start">
                  <div>
                    <div className="font-semibold">
                      {new Date(s.session_date).toLocaleDateString()} – {s.session_type}
                    </div>
                    <div className="text-gray-600">{s.notes}</div>
                  </div>
                  <div className="flex gap-2">
                    <Link
                      to={`/sessions/${s.id}`}
                      className="bg-blue-500 text-white px-3 py-1 rounded hover:bg-blue-600"
                    >
                      View
                    </Link>
                    <button
                      onClick={() => startEdit(s)}
                      className="bg-yellow-500 text-white px-3 py-1 rounded hover:bg-yellow-600"
                    >
                      Edit
                    </button>
                    <button
                      onClick={() => deleteSession(s.id)}
                      className="bg-red-600 text-white px-3 py-1 rounded hover:bg-red-700"
                    >
                      Delete
                    </button>
                  </div>
                </div>
              )}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
