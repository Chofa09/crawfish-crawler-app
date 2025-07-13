INSERT INTO urls (address, page_title, html_version, headings, has_login_form, crawl_status)
VALUES
  (
    'https://example.com',
    'Example Domain',
    'HTML5',
    JSON_OBJECT('h1', 1, 'h2', 2, 'h3', 0, 'h4', 0, 'h5', 0, 'h6', 0),
    false,
    'done'
  ),
  (
    'https://news.site.com',
    'Daily News',
    'HTML4.01',
    JSON_OBJECT('h1', 2, 'h2', 5, 'h3', 1, 'h4', 0, 'h5', 0, 'h6', 0),
    true,
    'done'
  ),
  (
    'https://login.portal.com',
    'Secure Login Portal',
    'HTML5',
    JSON_OBJECT('h1', 1, 'h2', 1, 'h3', 0, 'h4', 0, 'h5', 0, 'h6', 0),
    true,
    'queued'
  );
