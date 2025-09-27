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

  return (
    <div className="p-6 space-y-6">
      <div className="bg-white rounded-2xl shadow p-4">
        <h1 className="text-xl font-bold">
          {session.session_type} —{" "}
          {new Date(session.session_date).toDateString()}
        </h1>
        <p className="text-gray-600">{session.notes}</p>
      </div>

      <ExerciseList exercises={exercises} />

      <div className="bg-white rounded-2xl shadow p-4">
        <h2 className="text-lg font-semibold mb-2">Add Exercise</h2>
        <ExerciseForm sessionId={session.id} onAdded={handleExerciseAdded} />
      </div>
    </div>
  );
}
