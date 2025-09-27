interface Exercise {
  id: number;
  name: string;
  notes?: string;
  variation?: string;
}

export default function ExerciseList({ exercises }: { exercises: Exercise[] }) {
  if (exercises.length === 0) return <p>No exercises yet.</p>;

  return (
    <div className="space-y-3">
      {exercises.map((ex) => (
        <div key={ex.id} className="bg-gray-50 rounded-xl p-3 shadow">
          <h3 className="font-semibold">
            {ex.name}
            {ex.variation && (
              <span className="text-gray-500"> ({ex.variation})</span>
            )}
          </h3>
          {ex.notes && <p className="text-gray-600 text-sm">{ex.notes}</p>}
        </div>
      ))}
    </div>
  );
}
