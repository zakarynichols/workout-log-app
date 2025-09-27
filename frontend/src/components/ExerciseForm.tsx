import { useState } from "react";

export default function ExerciseForm({ sessionId }: { sessionId: number }) {
  const [mode, setMode] = useState<"dictionary" | "custom">("dictionary");
  const [dictionaryId, setDictionaryId] = useState("");
  const [customName, setCustomName] = useState("");
  const [variation, setVariation] = useState("");
  const [notes, setNotes] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const payload: any = {
      notes: notes || undefined,
      variation: variation || undefined,
    };

    if (mode === "dictionary") {
      payload.dictionary_exercise_id = dictionaryId ? Number(dictionaryId) : null;
    } else {
      payload.custom_exercise_id = customName ? 0 : null; 
      payload.name = customName; // you’ll need a backend route to insert custom_exercises first
    }

    await fetch(`/api/sessions/${sessionId}/exercises`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload),
    });

    // reset
    setDictionaryId("");
    setCustomName("");
    setVariation("");
    setNotes("");
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-3 border p-4 rounded-md">
      <div className="flex gap-4">
        <label>
          <input
            type="radio"
            value="dictionary"
            checked={mode === "dictionary"}
            onChange={() => setMode("dictionary")}
          />
          Dictionary
        </label>
        <label>
          <input
            type="radio"
            value="custom"
            checked={mode === "custom"}
            onChange={() => setMode("custom")}
          />
          Custom
        </label>
      </div>

      {mode === "dictionary" ? (
        <select
          value={dictionaryId}
          onChange={(e) => setDictionaryId(e.target.value)}
          className="w-full border p-2 rounded"
          required
        >
          <option value="">Select exercise</option>
          {/* TODO: populate with dictionary_exercises from API */}
          <option value="1">Bench Press</option>
          <option value="2">Deadlift</option>
        </select>
      ) : (
        <input
          type="text"
          placeholder="Custom exercise name"
          value={customName}
          onChange={(e) => setCustomName(e.target.value)}
          className="w-full border p-2 rounded"
          required
        />
      )}

      <input
        type="text"
        placeholder="Variation (optional)"
        value={variation}
        onChange={(e) => setVariation(e.target.value)}
        className="w-full border p-2 rounded"
      />

      <textarea
        placeholder="Notes (optional)"
        value={notes}
        onChange={(e) => setNotes(e.target.value)}
        className="w-full border p-2 rounded"
      />

      <button
        type="submit"
        className="bg-blue-600 text-white px-4 py-2 rounded-xl shadow"
      >
        Add Exercise
      </button>
    </form>
  );
}
