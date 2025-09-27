export type Session = {
  id: number;
  user_id: number;
  session_date: string;
  session_type: string;
  notes: string;
}

export async function getSessions(): Promise<Session[]> {
  const res = await fetch("/api/sessions", {
    headers: { "Content-Type": "application/json" },
  });

  if (!res.ok) {
    throw new Error(`Failed to fetch sessions: ${res.status}`);
  }

  return res.json();
}
