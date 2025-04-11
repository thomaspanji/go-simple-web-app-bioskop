-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE bioskop (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    lokasi VARCHAR(255) NOT NULL,
    rating DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Fungsi trigger untuk update updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$ language 'plpgsql';

-- Buat trigger-nya
CREATE TRIGGER set_updated_at
BEFORE UPDATE ON bioskop
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

-- +migrate StatementEnd
-- +migrate Down
DROP TRIGGER IF EXISTS set_updated_at ON bioskop;
DROP FUNCTION IF EXISTS update_updated_at_column();
DROP TABLE IF EXISTS bioskop;
