import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import ExerciseList from "../components/ExerciseList";
import ExerciseForm from "../components/ExerciseForm";

interface Session {
  id: number;
  session_date: string;
  session_type: string;
  notes: string;
}

interface Exercise {
  id: number;
  name: string;
  notes?: string;
  variation?: string;
  sets: WorkoutSet[]
}

interface WorkoutSet {
  id: number;
  exercise_id: number;
  weight: number;
  reps: number;
  notes?: string;
}

export default function SessionDetails() {
  const { id } = useParams();
  const [session, setSession] = useState<Session | null>(null);
  const [exercises, setExercises] = useState<Exercise[]>([]);

  // Fetch session
  useEffect(() => {
    fetch(`/api/sessions/${id}`)
      .then((res) => res.json())
      .then(setSession)
      .catch((err) => console.error("Error fetching session:", err));
  }, [id]);

  // Fetch exercises
  useEffect(() => {
    if (!id) return;
    fetch(`/api/sessions/${id}/exercises`)
      .then((res) => res.json())
      .then(setExercises)
      .catch((err) => console.error("Error fetching exercises:", err));
  }, [id]);

  if (!session) return <div>Loading...</div>;

  const handleExerciseAdded = (exercise: Exercise) => {
    setExercises((prev) => [...prev, exercise]); // instant update
  };

const handleSetAdded = (exerciseId: number, set: WorkoutSet) => {
  setExercises((prev) =>
    prev.map((ex) =>
      ex.id === exerciseId
        ? { ...ex, sets: [...(ex.sets || []), set] }
        : ex
    )
  );
};

return (
  <div className="p-6 space-y-6">
    {/* session card */}
    <ExerciseList exercises={exercises} onSetAdded={handleSetAdded} />
    <ExerciseForm sessionId={session.id} onAdded={handleExerciseAdded} />
  </div>
);
}
