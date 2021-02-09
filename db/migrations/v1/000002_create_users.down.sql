DO $$
    BEGIN
        IF EXISTS (SELECT * FROM information_schema.tables WHERE
            table_schema = 'public' AND table_name = 'users') THEN

            DROP TABLE users;
        END IF;
    END
$$