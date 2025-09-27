import AddSetForm from "./AddSetForm";

interface Exercise {
  id: number;
  name: string;
  notes?: string;
  variation?: string;
  sets?: WorkoutSet[];
}

interface WorkoutSet {
  id: number;
  exercise_id: number;
  weight: number;
  reps: number;
  notes?: string;
}

export default function ExerciseList({
  exercises,
  onSetAdded,
}: {
  exercises: Exercise[];
  onSetAdded: (exerciseId: number, set: WorkoutSet) => void;
}) {
  if (exercises.length === 0) return <p>No exercises yet.</p>;

  return (
    <div className="space-y-4">
      {exercises.map((ex) => (
        <div
          key={ex.id}
          className="bg-gray-50 rounded-xl p-4 shadow space-y-3"
        >
          <div>
            <h3 className="font-semibold">
              {ex.name}
              {ex.variation && (
                <span className="text-gray-500"> ({ex.variation})</span>
              )}
            </h3>
            {ex.notes && <p className="text-gray-600 text-sm">{ex.notes}</p>}
          </div>

          {/* Show sets if they exist */}
          {ex.sets && ex.sets.length > 0 && (
            <div className="space-y-1">
              {ex.sets.map((set) => (
                <div
                  key={set.id}
                  className="text-sm flex justify-between bg-white rounded p-2 border"
                >
                  <span>
                    {set.weight} lbs × {set.reps} reps
                  </span>
                  {set.notes && (
                    <span className="text-gray-500">{set.notes}</span>
                  )}
                </div>
              ))}
            </div>
          )}

          {/* Inline Add Set form */}
          <AddSetForm exerciseId={ex.id} onAdded={(set) => onSetAdded(ex.id, set)} />
        </div>
      ))}
    </div>
  );
}
