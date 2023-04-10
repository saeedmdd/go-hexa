    CREATE TABLE tasks (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255),
        description TEXT,
        color VARCHAR(7) DEFAULT \'#FFFFFF\',
        starts_at TIMESTAMP NULL,
        done_at TIMESTAMP NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        deleted_at TIMESTAMP NULL
    )