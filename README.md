# Mango CRM - Orchard Management System

A Go-based microservice for managing orchards and their related information. This system provides CRUD operations for
orchard entities, enabling efficient tracking and management of agricultural properties.

## Features

- Create new orchards with detailed information
- Retrieve orchard details by ID
- List all registered orchards
- Update existing orchard information
- Delete orchard records
- Structured logging system

## Technical Details

### Architecture

The project follows a clean architecture pattern with the following structure:

- `internal/orchard/`
    - `domain/` - Core business logic and entities
    - `application/` - Use cases and business rules
    - `infrastructure/` - External implementations

### Dependencies

- Go SDK 1.23.0
- zerolog - For structured logging

## Usage Examples
