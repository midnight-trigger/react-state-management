CREATE TABLE IF NOT EXISTS tags (
  id SERIAL PRIMARY KEY,
  name varchar(20) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tasks (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title varchar(100) NOT NULL DEFAULT '',
  tag_id INT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_tags FOREIGN KEY(tag_id) REFERENCES tags(id)
);


-- SEEDING
-- INSERT INTO tasks (title, tag_id) values ('Task1', 1), ('Task2', 2), ('Task3', 3);
-- INSERT INTO tags (name) values ('Name of Tag1'), ('Name of Tag2'), ('Name of Tag3');