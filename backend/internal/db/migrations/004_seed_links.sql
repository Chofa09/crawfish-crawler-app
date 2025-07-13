INSERT INTO links (url_id, href, is_internal, status_code, is_broken)
VALUES
  -- For URL ID 1 (example.com)
  (1, 'https://example.com/about', true, 200, false),
  (1, 'https://external.com', false, 404, true),

  -- For URL ID 2 (news.site.com)
  (2, 'https://news.site.com/latest', true, 200, false),
  (2, 'https://ads.network.com', false, 500, true),
  (2, 'https://news.site.com/contact', true, 200, false),

  -- For URL ID 3 (login.portal.com)
  (3, 'https://login.portal.com/reset-password', true, 200, false),
  (3, 'https://cdn.portal.com/assets', false, 200, false);
