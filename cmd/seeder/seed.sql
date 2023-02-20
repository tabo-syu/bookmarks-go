INSERT INTO bookmarks 
  (id, url, title, description) 
VALUES 
  ('e6fbaf41-0688-41d4-aaa0-4e0c2cfba6a8', 'https://tabo-syu.com', 'Blog', 'My Blog'),
  ('c77c048a-5c89-4da5-9fb5-b7e54dd00df4', 'https://twitter.com/tabo_web', 'Twitter', 'My Twitter account.'),
  ('abfbce21-b123-4345-ba5c-0f957e7a28e6', 'https://github.com/tabo-syu', 'GitHub', 'My GitHub account.');

INSERT INTO tags
  (id, name, color)
VALUES
  ('92449b60-90ac-4161-9be1-200b07854a46', '技術', '#000000'),
  ('c84cbdf8-58c3-4a39-88be-921264daa8d2', '趣味', '#ffffff'),
  ('baa234d5-8c65-4a32-912e-224752be823c', 'その他', '#123456');

INSERT INTO comments
  (id, bookmark_id, body)
VALUES
  ('ac70bae9-fb15-2831-2c38-a6574396a178', 'e6fbaf41-0688-41d4-aaa0-4e0c2cfba6a8', 'コメント1'),
  ('a282ee39-654b-09ed-eee5-4ba48ef4a5ff', 'e6fbaf41-0688-41d4-aaa0-4e0c2cfba6a8', 'コメント2'),
  ('865ba8dc-1b38-4f36-446b-7e2563f44bfc', 'e6fbaf41-0688-41d4-aaa0-4e0c2cfba6a8', 'コメント3'),
  ('3d623f3a-8b88-ff36-983d-63e350759443', 'c77c048a-5c89-4da5-9fb5-b7e54dd00df4', 'コメント4'),
  ('fac4c2ce-0b2a-b018-7ba3-0cb2696040d1', 'abfbce21-b123-4345-ba5c-0f957e7a28e6', 'コメント5'),
  ('4066c545-1c11-9cc4-21d1-cd0eafb1296d', 'abfbce21-b123-4345-ba5c-0f957e7a28e6', 'コメント6');

INSERT INTO bookmark_has_tags
  (bookmark_id, tag_id)
VALUES
  ('e6fbaf41-0688-41d4-aaa0-4e0c2cfba6a8', '92449b60-90ac-4161-9be1-200b07854a46'),
  ('e6fbaf41-0688-41d4-aaa0-4e0c2cfba6a8', 'c84cbdf8-58c3-4a39-88be-921264daa8d2'),
  ('e6fbaf41-0688-41d4-aaa0-4e0c2cfba6a8', 'baa234d5-8c65-4a32-912e-224752be823c'),
  ('c77c048a-5c89-4da5-9fb5-b7e54dd00df4', '92449b60-90ac-4161-9be1-200b07854a46'),
  ('abfbce21-b123-4345-ba5c-0f957e7a28e6', 'c84cbdf8-58c3-4a39-88be-921264daa8d2'),
  ('abfbce21-b123-4345-ba5c-0f957e7a28e6', 'baa234d5-8c65-4a32-912e-224752be823c');
