-- 20241104000000_alter_tasks_add_user_id.sql
ALTER TABLE tasks ADD COLUMN user_id INTEGER REFERENCES users(id) ON DELETE CASCADE;
