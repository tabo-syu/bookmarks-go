CREATE TABLE bookmarks (
  id         UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  url        TEXT        NOT NULL,
  created_at TIMESTAMPTZ NOT NULL    DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE tags (
  id         UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
  name       VARCHAR(100) NOT NULL,
  created_at TIMESTAMPTZ  NOT NULL    DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ  NOT NULL
);

CREATE TABLE comments (
  id         UUID                 DEFAULT gen_random_uuid(),
  comment    TEXT        NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE bookmark_has_tags (
  bookmark_id UUID,
  tag_id      UUID,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMPTZ NOT NULL
);

CREATE TABLE bookmark_has_comments (
  bookmark_id UUID,
  tag_id      UUID,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMPTZ NOT NULL
);
