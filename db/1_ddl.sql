CREATE TABLE user (
    id TEXT NOT NULL,
    name TEXT NOT NULL,
    password TEXT NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE purchase (
    product_id TEXT NOT NULL,
    product_name TEXT NOT NULL,
    user_id TEXT,
    purchase_date NOT NULL,
    PRIMARY KEY(product_id),
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
        REFERENCES user(id)
)