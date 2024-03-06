-- +goose Up
-- +goose StatementBegin

CREATE TABLE contact_person
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    vk_id      VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP    NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE contact_person;

-- +goose StatementEnd
