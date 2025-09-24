-- Drop in reverse order
DROP TABLE IF EXISTS workout.sets CASCADE;
DROP TABLE IF EXISTS workout.exercises CASCADE;
DROP TABLE IF EXISTS workout.custom_exercises CASCADE;
DROP TABLE IF EXISTS workout.dictionary_exercises CASCADE;
DROP TABLE IF EXISTS workout.sessions CASCADE;
DROP TABLE IF EXISTS workout.users CASCADE;

DROP SCHEMA IF EXISTS workout CASCADE;