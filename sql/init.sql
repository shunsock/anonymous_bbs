CREATE DATABASE bbs;

\c bbs

CREATE TABLE IF NOT EXISTS bbs_threads (
    thread_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL DEFAULT '風吹けば名無しさん',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS bbs_comments (
    id SERIAL PRIMARY KEY,
    comment TEXT NOT NULL,
    commenter_ip_address INET NOT NULL,
    thread_id INT NOT NULL REFERENCES bbs_threads(thread_id) ON DELETE CASCADE,
    username VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS reported_comments (
    id SERIAL PRIMARY KEY,
    comment_id INT NOT NULL REFERENCES bbs_comments(id) ON DELETE CASCADE,
    reporter_ip_address INET NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

