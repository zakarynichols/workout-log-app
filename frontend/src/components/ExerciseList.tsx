import { useEffect, useState } from "react";

interface Exercise {
  id: number;
  name: string;
  notes: string;
}

export default function ExerciseList({ sessionId }: { sessionId: number }) {
  const [exercises, setExercises] = useState<Exercise[]>([]);

  useEffect(() => {
    fetch(`http://localhost:8081/api/sessions/${sessionId}/exercises`)
      .then((res) => res.json())
      .then(setExercises)
      .catch((err) => console.error("Error fetching exercises:", err));
  }, [sessionId]);

  if (exercises.length === 0) return <p>No exercises yet.</p>;

  return (
    <div className="space-y-3">
      {exercises.map((ex) => (
        <div key={ex.id} className="bg-gray-50 rounded-xl p-3 shadow">
          <h3 className="font-semibold">{ex.name}</h3>
          {ex.notes && <p className="text-gray-600 text-sm">{ex.notes}</p>}
        </div>
      ))}
    </div>
  );
}
