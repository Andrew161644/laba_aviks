CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE if not exists Bank_Account(
    id uuid DEFAULT uuid_generate_v4 (),
    currencyId integer,
    userId integer,
    value integer,
    FOREIGN KEY (currencyId) REFERENCES Currency(id),
    FOREIGN KEY (userId) REFERENCES Users(id)
)