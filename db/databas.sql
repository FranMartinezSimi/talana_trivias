CREATE INDEX question_text_idx ON questions USING GIN (to_tsvector('english', question));
