BEGIN;

-- 1. Add a temporary int column
ALTER TABLE workout.sets ADD COLUMN reps_int INT;

-- 2. Copy over numeric reps
UPDATE workout.sets
SET reps_int = reps::INT
WHERE reps ~ '^[0-9]+$';

-- 3. Preserve non-numeric reps in notes
UPDATE workout.sets
SET notes = COALESCE(notes || ' | reps: ' || reps, reps)
WHERE reps !~ '^[0-9]+$' AND reps IS NOT NULL;

-- 4. Drop old column
ALTER TABLE workout.sets DROP COLUMN reps;

-- 5. Rename temp column
ALTER TABLE workout.sets RENAME COLUMN reps_int TO reps;

COMMIT;
