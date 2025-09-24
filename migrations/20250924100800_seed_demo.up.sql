-- ================================
-- Seed demo data
-- ================================

-- Ensure base dictionary exercises exist
INSERT INTO workout.dictionary_exercises (name, category)
VALUES
    ('Bench Press', 'Chest'),
    ('Dumbbell Bench Press', 'Chest'),
    ('Squat', 'Legs'),
    ('Deadlift', 'Back'),
    ('Overhead Press', 'Shoulders'),
    ('Barbell Row', 'Back'),
    ('Pull-Up', 'Back'),
    ('Rice Bucket Hold', 'Forearms')
ON CONFLICT (name) DO NOTHING;

-- Ensure demo user exists
WITH demo_user AS (
    INSERT INTO workout.users (username)
    VALUES ('demo_user')
    ON CONFLICT (username) DO UPDATE SET username = EXCLUDED.username
    RETURNING id
),

-- ================================
-- Push session: Dumbbell Bench
-- ================================
push_session AS (
    INSERT INTO workout.sessions (user_id, session_date, session_type, notes)
    VALUES (
        (SELECT id FROM demo_user),
        CURRENT_DATE,
        'push',
        'Demo push workout session'
    )
    RETURNING id
),
db_bench_exercise AS (
    INSERT INTO workout.exercises (session_id, dictionary_exercise_id, variation, notes)
    VALUES (
        (SELECT id FROM push_session),
        (SELECT id FROM workout.dictionary_exercises WHERE name = 'Dumbbell Bench Press'),
        'flat bench',
        'Controlled tempo'
    )
    RETURNING id
),

-- ================================
-- Pull session: Pull-Up + Rice Bucket
-- ================================
pull_session AS (
    INSERT INTO workout.sessions (user_id, session_date, session_type, notes)
    VALUES (
        (SELECT id FROM demo_user),
        CURRENT_DATE,
        'pull',
        'Demo pull workout session'
    )
    RETURNING id
),
pullup_exercise AS (
    INSERT INTO workout.exercises (session_id, dictionary_exercise_id, variation, notes)
    VALUES (
        (SELECT id FROM pull_session),
        (SELECT id FROM workout.dictionary_exercises WHERE name = 'Pull-Up'),
        'bodyweight strict',
        'Focus on form'
    )
    RETURNING id
),
ricebucket_exercise AS (
    INSERT INTO workout.exercises (session_id, dictionary_exercise_id, notes)
    VALUES (
        (SELECT id FROM pull_session),
        (SELECT id FROM workout.dictionary_exercises WHERE name = 'Rice Bucket Hold'),
        'Grip training'
    )
    RETURNING id
)

-- ================================
-- Final inserts for sets
-- ================================
INSERT INTO workout.sets (exercise_id, set_number, weight, reps, duration, notes)
VALUES
    -- Dumbbell Bench Press
    ((SELECT id FROM db_bench_exercise), 1, 70, '6', NULL, 'solid'),
    ((SELECT id FROM db_bench_exercise), 2, 70, '6', NULL, 'still strong'),
    ((SELECT id FROM db_bench_exercise), 3, 70, '5', NULL, 'close to failure'),

    -- Pull-Ups
    ((SELECT id FROM pullup_exercise), 1, NULL, '10', NULL, 'solid'),
    ((SELECT id FROM pullup_exercise), 2, NULL, '8', NULL, 'harder'),
    ((SELECT id FROM pullup_exercise), 3, NULL, '6', NULL, 'to failure'),

    -- Rice Bucket Hold
    ((SELECT id FROM ricebucket_exercise), 1, NULL, NULL, '2 minutes', 'burned like crazy');
