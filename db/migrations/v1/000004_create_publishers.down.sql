DO $$
    BEGIN
        IF EXISTS (SELECT * FROM information_schema.tables WHERE
            table_schema = 'public' AND table_name = 'publishers') THEN

            DROP TABLE publishers;

        END IF;

        
        ALTER TABLE books DROP COLUMN publisher_id uuid;

        ALTER TABLE books DROP CONSTRAINT books_publisher_ids_fkey;

    COMMIT;
END
$$