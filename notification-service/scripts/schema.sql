CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY,
    type VARCHAR(50) NOT NULL,
    payload JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
    );

CREATE TABLE IF NOT EXISTS notifications (
    id UUID PRIMARY KEY,
    event_id UUID NOT NULL REFERENCES events(id),
    type VARCHAR(50) NOT NULL,
    recipient VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
    );

CREATE INDEX IF NOT EXISTS idx_notifications_status ON notifications(status);
CREATE INDEX IF NOT EXISTS idx_notifications_event_id ON notifications(event_id);