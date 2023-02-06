INSERT INTO bookmarks 
  (id, url, title, description) 
VALUES 
  ('e6fbaf41-0688-41d4-aaa0-4e0c2cfba6a8', 'https://tabo-syu.com', 'Blog', 'My Blog'),
  ('c77c048a-5c89-4da5-9fb5-b7e54dd00df4', 'https://twitter.com/tabo_web', 'Twitter', 'My Twitter account.'),
  ('abfbce21-b123-4345-ba5c-0f957e7a28e6', 'https://github.com/tabo-syu', 'GitHub', 'My GitHub account.');

INSERT INTO tags
  (id, name, color)
VALUES
  ('92449b60-90ac-4161-9be1-200b07854a46', '技術', '000000'),
  ('c84cbdf8-58c3-4a39-88be-921264daa8d2', '趣味', 'ffffff'),
  ('baa234d5-8c65-4a32-912e-224752be823c', 'その他', '123456');

INSERT INTO comments
  (bookmark_id, comment)
VALUES
  ('e6fbaf41-0688-41d4-aaa0-4e0c2cfba6a8', 'コメント1'),
  ('e6fbaf41-0688-41d4-aaa0-4e0c2cfba6a8', 'コメント2'),
  ('e6fbaf41-0688-41d4-aaa0-4e0c2cfba6a8', 'コメント3'),
  ('c77c048a-5c89-4da5-9fb5-b7e54dd00df4', 'コメント4'),
  ('abfbce21-b123-4345-ba5c-0f957e7a28e6', 'コメント5'),
  ('abfbce21-b123-4345-ba5c-0f957e7a28e6', 'コメント6');

INSERT INTO bookmark_has_tags
  (bookmark_id, tag_id)
VALUES
  ('e6fbaf41-0688-41d4-aaa0-4e0c2cfba6a8', '92449b60-90ac-4161-9be1-200b07854a46'),
  ('e6fbaf41-0688-41d4-aaa0-4e0c2cfba6a8', 'c84cbdf8-58c3-4a39-88be-921264daa8d2'),
  ('e6fbaf41-0688-41d4-aaa0-4e0c2cfba6a8', 'baa234d5-8c65-4a32-912e-224752be823c'),
  ('c77c048a-5c89-4da5-9fb5-b7e54dd00df4', '92449b60-90ac-4161-9be1-200b07854a46'),
  ('abfbce21-b123-4345-ba5c-0f957e7a28e6', 'c84cbdf8-58c3-4a39-88be-921264daa8d2'),
  ('abfbce21-b123-4345-ba5c-0f957e7a28e6', 'baa234d5-8c65-4a32-912e-224752be823c');
