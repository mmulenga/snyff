-- +goose Up
create endpoints (
    id uuid primary key default gen_random_uuid(),
    token_hash text not null unique,
    name text,
    created_at timestamptz not null default now()
);

create table requests (
    id uuid primary key default gen_random_uuid(),
    endpoint_id uuid not null references endpoints(id) on delete cascade,
    method text not null,
    path text not null,
    query text,
    headers jsonb not null default '{}',
    body bytea,
    body_size_bytes integer not null default 0,
    body_truncrated boolean not null default false,
    content_type text,
    source_ip inet,
    received_at timestamptz not null default now()
);

create index requests_endpoint_received_idx (endpoint_id, received_at desc)
on table requests;

-- +goose Down
drop table if exists endpoints;
drop table if exists requests;
