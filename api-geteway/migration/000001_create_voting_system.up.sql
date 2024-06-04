create table if not exists election(
    id uuid primary key not null,
    name varchar(40),
    date TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table if not exists party(
    id uuid primary key not null,
    name varchar(30) not null,
    slogan text DEFAULT null,
    opened_date TIMESTAMP not null,
    description text DEFAULT null,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create type genders as enum("male","female");

create table if not exists public(
    id uuid primary key not null,
    first_name varchar(30) not null,
    last_name varchar(30) not null,
    birthday TIMESTAMP not null,
    gender genders not null,
    nation varchar(30) not null,
    party_id uuid references party(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
)

create table if not exists candidate(
    id uuid primary key not null,
    election_id uuid references election(id) not null,
    public_id uuid references public(id) not null,
    party_id uuid references party(id) not null,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
)

create table if not exists public_vote(
    id uuid primary key not null,
    election_id uuid references election(id) not null,
    public_id uuid references public(id) not null,
    created_at TIMESTAMP DEFAULT NOW()
)

create table if not exists votes(
    id uuid primary key not null,
    candidate_id uuid references candidate(id)
    created_at TIMESTAMP DEFAULT NOW()
)