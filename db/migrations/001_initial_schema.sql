-- Migration: 001_initial_schema
-- Create tables for the initial schema
-- Enable foreign key constraints
PRAGMA foreign_keys = ON;

-- Create Driver table
CREATE TABLE IF NOT EXISTS driver (
    id TEXT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    identification VARCHAR(255) NOT NULL,
    license_number VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (datetime('now')),
    updated_at TIMESTAMP NOT NULL DEFAULT (datetime('now'))
);

-- Create Vehicle table
CREATE TABLE IF NOT EXISTS vehicle (
    id TEXT PRIMARY KEY,
    plate_number VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (datetime('now')),
    updated_at TIMESTAMP NOT NULL DEFAULT (datetime('now'))
);

-- Create Route table
CREATE TABLE IF NOT EXISTS route (
    id TEXT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    status VARCHAR(255) NOT NULL CHECK (status IN ('pending', 'started', 'completed')),
    created_at TIMESTAMP NOT NULL DEFAULT (datetime('now')),
    updated_at TIMESTAMP NOT NULL DEFAULT (datetime('now')),
    vehicle_id TEXT NOT NULL,
    driver_id TEXT NOT NULL,
    FOREIGN KEY (vehicle_id) REFERENCES vehicle(id),
    FOREIGN KEY (driver_id) REFERENCES driver(id)
);

-- Create RoutePoint table
CREATE TABLE IF NOT EXISTS route_point (
    id TEXT PRIMARY KEY,
    purchase_order_id VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL CHECK (status IN ('pending', 'in_route', 'completed')),
    latitude REAL NOT NULL,
    longitude REAL NOT NULL,
    address VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT (datetime('now')),
    updated_at TIMESTAMP NOT NULL DEFAULT (datetime('now')),
    route_id TEXT NOT NULL,
    FOREIGN KEY (route_id) REFERENCES route(id)
);

-- Create indexes for better performance
CREATE INDEX idx_route_status ON route(status);
CREATE INDEX idx_route_point_route_id ON route_point(route_id);
CREATE INDEX idx_route_point_status ON route_point(status); 