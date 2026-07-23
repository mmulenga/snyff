-- +goose Up
create table endpoints (
    id uuid primary key default uuidv7(),
    token_hash text not null unique,
    name text,
    created_at timestamptz not null default now()
);

create table requests (
    id uuid primary key default uuidv7(),
    method text not null,
    path text not null,
    query text,
    headers jsonb not null default '{}',
    body bytea,
    body_size_bytes integer not null default 0,
    body_truncated boolean not null default false,
    content_type text,
    source_ip inet,
    received_at timestamptz not null default now()
);

INSERT INTO requests
  (method, path, query, headers, body, body_size_bytes, body_truncated, content_type, source_ip, received_at)
SELECT
  method,
  path,
  query,
  headers::jsonb,
  convert_to(body_text, 'UTF8'),
  -- claimed_size wins when set (row 5); otherwise fall back to actual stored bytes; 0 when no body
  coalesce(claimed_size, octet_length(convert_to(body_text, 'UTF8')), 0),
  truncated,
  content_type,
  source_ip::inet,
  received_at::timestamptz
FROM (VALUES
  -- 1. stripe: clean json POST, ipv4
  ('POST', '/webhooks/stripe', NULL,
   '{"Content-Type":["application/json"],"User-Agent":["Stripe/1.0"],"Stripe-Signature":["t=1718442723,v1=5257a869e7ecebeda32affa62cdca3fa"]}',
   '{"id":"evt_1PabcXYZ","type":"charge.succeeded","data":{"amount":2000,"currency":"usd"}}',
   NULL, false, 'application/json', '203.0.113.45', '2026-06-15 09:12:03.412+00'),

  -- 2. github: catch-all root path
  ('POST', '/', NULL,
   '{"Content-Type":["application/json"],"User-Agent":["GitHub-Hookshot/abc123"],"X-GitHub-Event":["push"],"X-GitHub-Delivery":["72d3162e-cc78-11e3-81ab-4c9367dc0958"]}',
   '{"ref":"refs/heads/main","commits":[{"id":"a1b2c3","message":"fix ingest"}]}',
   NULL, false, 'application/json', '198.51.100.22', '2026-06-15 09:12:41.087+00'),

  -- 3. meta verification GET: query present, NO body, NO content-type
  ('GET', '/webhooks/meta', 'hub.mode=subscribe&hub.challenge=1158201444&hub.verify_token=my_token',
   '{"User-Agent":["facebookexternalua"],"Accept":["*/*"]}',
   NULL, 0, false, NULL, '192.0.2.7', '2026-06-15 09:13:10.900+00'),

  -- 4. form-encoded + duplicate Set-Cookie (array-valued jsonb the flat-map design would drop)
  ('POST', '/webhooks/legacy', 'debug=true&retry=1',
   '{"Content-Type":["application/x-www-form-urlencoded"],"Set-Cookie":["session=abc; Path=/; HttpOnly","tracking=xyz; Path=/"],"User-Agent":["curl/8.7.1"]}',
   'event=order.created&order_id=8821&total=49.99',
   NULL, false, 'application/x-www-form-urlencoded', '203.0.113.201', '2026-06-15 09:14:55.331+00'),

  -- 5. TRUNCATED: client sent 5 MiB, body holds only the capped prefix, ipv6 source
  ('POST', '/webhooks/custom', NULL,
   '{"Content-Type":["application/json"],"User-Agent":["MyApp/2.3"]}',
   '{"status":"updated","items":[{"sku":"A1","qty":3},{"sku":"B2","qty',
   5242880, true, 'application/json', '2001:db8::1a2b', '2026-06-15 09:15:02.008+00')
) AS seed(method, path, query, headers, body_text, claimed_size, truncated, content_type, source_ip, received_at);


--create index requests_endpoint_received_idx on requests(endpoint_id, received_at desc);


-- +goose Down
drop table if exists endpoints;
drop table if exists requests;
