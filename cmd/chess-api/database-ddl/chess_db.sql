create table players (
    last_name  varchar(100) not null,
    first_name varchar(100) not null,
    primary key ( last_name, first_name )
);

create table locations (
    town varchar(100) primary key
);

create table openings (
    eco_code     varchar(10)   primary key,
    name         varchar(100)  not null,
    variation    varchar(200),
    opening_line varchar(500)
);

create table results (
    result varchar(10) primary key
);

create table games (
    id               serial        primary key,
    location         varchar(100)  not null,
    round            int,
    white_last_name  varchar(100)  not null,
    white_first_name varchar(100)  not null,
    black_last_name  varchar(100)  not null,
    black_first_name varchar(100)  not null,
    white_rank       int,
    black_rank       int,
    date             date,
    opening          varchar(10)   not null,
    result           varchar(10)   not null,
    moves            varchar(2000) not null,
    constraint games_locations_fk
        foreign key ( location )
        references locations ( town )
        on update cascade
        on delete cascade,
    constraint white_games_players_fk
        foreign key ( white_last_name, white_first_name )
        references players ( last_name, first_name )
        on update cascade
        on delete cascade,
    constraint black_games_players_fk
        foreign key ( black_last_name, black_first_name )
        references players   ( last_name, first_name )
        on update cascade
        on delete cascade,
    constraint games_openings_fk
        foreign key ( opening )
        references openings  ( eco_code )
        on update cascade
        on delete cascade,
    constraint games_results
        foreign key ( result )
        references results ( result )
        on update cascade
        on delete cascade
);
