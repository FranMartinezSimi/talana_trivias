Create or replace INDEXCREATE INDEX question_text_idx ON questions USING GIN (to_tsvector('english', text));
