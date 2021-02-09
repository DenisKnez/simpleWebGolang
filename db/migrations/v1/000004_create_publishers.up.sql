DO $$
    BEGIN
        IF NOT EXISTS (SELECT * FROM information_schema.tables WHERE
            table_schema = 'public' AND table_name = 'publishers') THEN

            CREATE TABLE publishers (
                id uuid primary key,
                "name" varchar(25) not null CONSTRAINT name_short_lenght CHECK (length("name") >= 3), 
                date_founded date,
                created_at timestamp not null,
                updated_at timestamp not null,
                deleted_at timestamp,
                is_deleted boolean default false not null
            );
        END IF;

        ALTER TABLE books ADD COLUMN publisher_id uuid;

        ALTER TABLE books ADD CONSTRAINT books_publisher_ids_fkey FOREIGN KEY (publisher_id) REFERENCES publishers (id);
    COMMIT;
END
$$