INSERT INTO apps (id, name, secret)
VALUES (1, 'runner', 'runner-secret')
ON CONFLICT DO NOTHING;