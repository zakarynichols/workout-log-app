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

export default function SessionDetails() {
  const { id } = useParams();
  const [session, setSession] = useState<Session | null>(null);

  useEffect(() => {
    fetch(`/api/sessions/${id}`)
      .then((res) => res.json())
      .then(setSession)
      .catch((err) => console.error("Error fetching session:", err));
  }, [id]);

  if (!session) return <div>Loading...</div>;

  return (
    <div className="p-6 space-y-6">
      <div className="bg-white rounded-2xl shadow p-4">
        <h1 className="text-xl font-bold">
          {session.session_type} — {new Date(session.session_date).toDateString()}
        </h1>
        <p className="text-gray-600">{session.notes}</p>
      </div>

      <ExerciseList sessionId={session.id} />

      <div className="bg-white rounded-2xl shadow p-4">
        <h2 className="text-lg font-semibold mb-2">Add Exercise</h2>
        <ExerciseForm sessionId={session.id} />
      </div>
    </div>
  );
}
