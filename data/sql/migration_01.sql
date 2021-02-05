DO $$
    BEGIN
        IF NOT EXISTS (SELECT * FROM information_schema.tables WHERE
            table_schema = 'public' AND table_name = 'users') THEN

            CREATE TABLE users (
                id uuid primary key,
                "name" varchar(60) not null CONSTRAINT name_short_lenght CHECK (length("name") >= 3), 
                lastname varchar(60) not null CONSTRAINT lastname_short_lenght CHECK (length(lastname) >= 3),
                age smallint,
                email varchar(255) not null unique CONSTRAINT email_short_lenght CHECK (length(email) > 5),
                "password" varchar(255) not null CONSTRAINT password_short_lenght CHECK (length("password") >= 7),
                created_at timestamp not null,
                updated_at timestamp not null,
                deleted_at timestamp,
                is_deleted boolean default false not null
            );

        END IF;

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