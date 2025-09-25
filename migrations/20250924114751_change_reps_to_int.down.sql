BEGIN;

-- 1. Add back a text column
ALTER TABLE workout.sets ADD COLUMN reps_text TEXT;

-- 2. Copy int reps back into text
UPDATE workout.sets
SET reps_text = reps::TEXT;

-- 3. Drop the int column
ALTER TABLE workout.sets DROP COLUMN reps;

-- 4. Rename back
ALTER TABLE workout.sets RENAME COLUMN reps_text TO reps;

COMMIT;
