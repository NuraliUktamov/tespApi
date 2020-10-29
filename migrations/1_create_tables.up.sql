CREATE TABLE IF NOT EXISTS person (user_id uuid NOT NULL, first_name TEXT NOT NUll, last_name TEXT NOT NULL, UNIQUE(user_id));

CREATE TABLE IF NOT EXISTS person_addresses (user_id uuid NOT NULL, address TEXT);

ALTER TABLE person_addresses ADD CONSTRAINT person_addresses_id_fkey FOREIGN KEY (user_id) REFERENCES person (user_id);
