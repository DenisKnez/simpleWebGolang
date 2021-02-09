DO $$
    BEGIN
        IF EXISTS (SELECT * FROM information_schema.tables WHERE
            table_schema = 'public' AND table_name = 'books') THEN

            DROP TABLE books;

        END IF;
    END
$$