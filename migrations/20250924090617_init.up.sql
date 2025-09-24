CREATE SCHEMA IF NOT EXISTS workout;

CREATE TABLE workout.users (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    username TEXT UNIQUE NOT NULL
);

CREATE TABLE workout.sessions (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT NOT NULL REFERENCES workout.users(id) ON DELETE CASCADE,
    session_date DATE NOT NULL,
    session_type TEXT, -- push, pull, legs, rest, custom
    notes TEXT
);

CREATE TABLE workout.dictionary_exercises (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    category TEXT,
    UNIQUE(name)
);

-- Per-user defined
CREATE TABLE workout.custom_exercises (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT NOT NULL REFERENCES workout.users(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    UNIQUE(user_id, name)
);

-- Per session
CREATE TABLE workout.exercises (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    session_id INT NOT NULL REFERENCES workout.sessions(id) ON DELETE CASCADE,
    dictionary_exercise_id INT REFERENCES workout.dictionary_exercises(id),
    custom_exercise_id INT REFERENCES workout.custom_exercises(id),
    variation TEXT,
    notes TEXT,
    CHECK (
        (dictionary_exercise_id IS NOT NULL AND custom_exercise_id IS NULL)
        OR (dictionary_exercise_id IS NULL AND custom_exercise_id IS NOT NULL)
    )
);

CREATE TABLE workout.sets (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    exercise_id INT NOT NULL REFERENCES workout.exercises(id) ON DELETE CASCADE,
    set_number INT,
    weight NUMERIC,
    weight_unit TEXT DEFAULT 'lbs',
    reps TEXT,
    duration INTERVAL,
    notes TEXT
);
