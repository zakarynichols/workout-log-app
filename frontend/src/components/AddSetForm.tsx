import { useState } from "react";

interface WorkoutSet {
  id: number;
  exercise_id: number;
  weight: number;
  reps: number;
  notes?: string;
}

export default function AddSetForm({
  exerciseId,
  onAdded,
}: {
  exerciseId: number;
  onAdded: (set: WorkoutSet) => void;
}) {
  const [weight, setWeight] = useState("");
  const [reps, setReps] = useState("");
  const [notes, setNotes] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const payload = {
      weight: Number(weight),
      reps: Number(reps),
      notes: notes || undefined,
    };

    const res = await fetch(`/api/exercises/${exerciseId}/sets`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload),
    });

    if (!res.ok) {
      console.error("Failed to add set");
      return;
    }

    const created = await res.json();
    onAdded(created);

    setWeight("");
    setReps("");
    setNotes("");
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="flex gap-2 items-center text-sm mt-2"
    >
      <input
        type="number"
        placeholder="Weight"
        value={weight}
        onChange={(e) => setWeight(e.target.value)}
        className="w-20 border p-1 rounded"
        required
      />
      <input
        type="number"
        placeholder="Reps"
        value={reps}
        onChange={(e) => setReps(e.target.value)}
        className="w-16 border p-1 rounded"
        required
      />
      <input
        type="text"
        placeholder="Notes"
        value={notes}
        onChange={(e) => setNotes(e.target.value)}
        className="flex-1 border p-1 rounded"
      />
      <button
        type="submit"
        className="bg-blue-600 text-white px-3 py-1 rounded shadow"
      >
        Add
      </button>
    </form>
  );
}
