DO $$
    BEGIN
        IF NOT EXISTS (SELECT * FROM information_schema.tables WHERE
            table_schema = 'public' AND table_name = 'books') THEN

            CREATE TABLE books (
                id uuid primary key,
                title varchar(25) not null CONSTRAINT title_short_lenght CHECK (length(title) >= 3), 
                author varchar(80) not null CONSTRAINT author_short_lenght CHECK (length(author) >= 3),
                release_date date,
                created_at timestamp not null,
                updated_at timestamp not null,
                deleted_at timestamp,
                is_deleted boolean default false not null
            );

        END IF;
    END
$$