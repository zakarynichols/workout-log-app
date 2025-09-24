-- ================================
-- Seed teardown: remove demo data
-- ================================

-- Delete demo user and cascade their sessions/exercises/sets
DELETE FROM workout.users
WHERE username = 'demo_user';
