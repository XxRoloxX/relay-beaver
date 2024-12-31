# Relay Beaver

Relay Beaver is a web-managed reverse proxy designed to facilitate
HTTP/HTTPS traffic debugging, analysis,
and redirection through a web-based configuration panel.

## Architecture

The project is structured into several key components:

- Proxy: A high-performance HTTP/HTTPS server that forwards traffic to target servers.
- Routing Configuration: A web panel that allows adding, removing, and modifying proxy rules (e.g., changing HTTP headers, load-balancing).
- Traffic Logging/Debugging: A module for real-time traffic inspection and request replay.
- Analytics: A module that aggregates traffic information for analysis through the web panel.

## Features

- Forward HTTP/HTTPS traffic to specified hosts.
- Add and remove proxy rules.
- Modify HTTP headers based on defined rules.
- Perform application-layer load balancing across multiple servers.
- Log traffic passing through the proxy.
- Aggregate traffic information for analysis in the web panel.
- Replay HTTP requests to specified hosts.
- Authenticate using SSO (GitHub, Google, etc.).

## Backend Overview

The backend is implemented in Go and includes several modules for handling different aspects of the application:

- auth: Authentication middleware.
- client_event: Handles client-specific events.
- common: Common utilities and handlers.
- database: Database initialization and interactions.
- loadbalancer: Load balancing CRUD logic.
- logger: Logging middleware.
- proxy_event: Handles proxy events.
- proxy_rule: Manages proxy rules.
- stats: Aggregates and provides analytics.

## Proxy Overview

The proxy is implemented as a separate binary using Go.
It is responsible for handling HTTP/HTTPS traffic redirection.

## Frontend Overview

The frontend is implemented using React/TypeScript and Vite for development and build processes.
It includes configurations for ESLint to ensure code quality and consistency.
