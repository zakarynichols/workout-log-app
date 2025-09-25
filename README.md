User Workout Flow – Frontend Call Sequence
==========================================

This document describes how a user interacts with the workout logging app, from login to progression tracking.
It outlines the sequence of API calls the frontend will make, and the expected request/response shapes.

------------------------------------------------------------
1. User Authentication (future – optional for MVP)
------------------------------------------------------------
POST /users         -> Create new user
GET /users/:id      -> Fetch user info

------------------------------------------------------------
2. Create a Workout Session
------------------------------------------------------------
POST /sessions
Request:
{
  "user_id": 1,
  "session_date": "2025-09-24",
  "session_type": "push",
  "notes": "Felt strong today"
}

Response:
{
  "id": 10,
  "user_id": 1,
  "session_date": "2025-09-24",
  "session_type": "push",
  "notes": "Felt strong today"
}

------------------------------------------------------------
3. Add Exercises to Session
------------------------------------------------------------
Dictionary Exercise:
POST /sessions/:id/exercises
Request:
{
  "dictionary_exercise_id": 5,
  "variation": "incline",
  "notes": "slow negative"
}

Response:
{
  "id": 22,
  "session_id": 10,
  "dictionary_exercise_id": 5,
  "variation": "incline",
  "notes": "slow negative"
}

Custom Exercise:
POST /sessions/:id/exercises
Request:
{
  "custom_exercise_name": "Rice Bucket Hold",
  "notes": "grip burn"
}

Response:
{
  "id": 23,
  "session_id": 10,
  "custom_exercise_id": 7,
  "notes": "grip burn"
}

------------------------------------------------------------
4. Add Sets to Exercise
------------------------------------------------------------
POST /exercises/:id/sets
Request:
{
  "set_number": 1,
  "weight": 70,
  "weight_unit": "lbs",
  "reps": 8,
  "notes": "solid"
}

Response:
{
  "id": 100,
  "exercise_id": 22,
  "set_number": 1,
  "weight": 70,
  "weight_unit": "lbs",
  "reps": 8,
  "notes": "solid"
}

------------------------------------------------------------
5. View Past Sessions
------------------------------------------------------------
GET /sessions        -> List all sessions with exercises & sets
GET /sessions/:id    -> Single session detail

------------------------------------------------------------
6. Edit data (patch)
------------------------------------------------------------
### Edit Session
PATCH /sessions/10
{
  "notes": "Chest + triceps"
}

### Edit Exercise
PATCH /exercises/20
{
  "variation": "incline bench"
}

### Edit Set
PATCH /sets/50
{
  "weight": 75,
  "reps": 7
}

---

## 6. Delete Data (DELETE)
### Delete Set
DELETE /sets/50

### Delete Exercise (removes its sets too)
DELETE /exercises/20

### Delete Session (removes all exercises & sets)
DELETE /sessions/10

------------------------------------------------------------
6. Progression Tracking
------------------------------------------------------------
GET /progression?exercise=Bench%20Press

Response:
{
  "exercise": "Bench Press",
  "sessions": [
    {
      "session_id": 1,
      "date": "2025-09-10",
      "top_set": { "weight": 65, "reps": 8 },
      "total_volume": 1560
    },
    {
      "session_id": 10,
      "date": "2025-09-24",
      "top_set": { "weight": 70, "reps": 8 },
      "total_volume": 1680
    }
  ],
  "percent_change": {
    "top_set_weight": 7.7,
    "total_volume": 7.7
  }
}

------------------------------------------------------------
Notes
------------------------------------------------------------
- session_date: stored in DB as DATE, returned as ISO 8601 string (YYYY-MM-DD).
- duration: stored in DB as INTERVAL, returned as integer duration_ms.
- reps: always INT (no text). Non-numeric reps are preserved in notes.

