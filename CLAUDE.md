# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Core Commands

### Development & Build
```bash
# Build the Go application
task -g build

# Deploy configuration files and restart services
task -g bench

# View application logs
task -g log

# Connect to MySQL
task -g mysql
```

### Performance Analysis
```bash
# Analyze nginx access logs with alp
task -g alp

# View slow query analysis
task -g slow-query

# Send analysis results to Discord
task -g analyze

# Profile with pprof/fgprof
task -g prof
task -g prof-check
```

### Multi-Server Management
```bash
# Execute command on all servers
task -g all -- <command>

# Execute task on all servers
task -g all-<task> -- <args>

# Switch git branch
task -g switch -- <branch>
```

## Architecture Overview

### Application Structure
- **Language**: Go (with Ruby, PHP, Python, Node.js alternatives available)
- **Web Framework**: Chi router
- **Database**: MySQL
- **Session Store**: Memcached (via gorilla-sessions-memcache)
- **Reverse Proxy**: Nginx
- **Process Manager**: systemd

### Key Components
1. **Web Application** (`/home/isucon/private_isu/webapp/golang/`)
   - Handles image uploads, posts, comments
   - Uses template rendering for HTML responses
   - Implements session-based authentication

2. **Benchmarker** (`/home/isucon/private_isu/benchmarker/`)
   - Measures application performance
   - Validates response correctness
   - Calculates scores based on successful requests

3. **Infrastructure Configuration**
   - Server configs stored in `s1/`, `s2/`, etc. directories
   - Environment variables in `/home/isucon/env.sh` and `env` files
   - Nginx configured for LTSV logging format
   - MySQL configured with slow query logging (long_query_time=0)

### Database Schema
- `users`: User accounts with authentication
- `posts`: Image posts with metadata
- `comments`: Comments on posts

### Performance Monitoring Setup
- **Nginx**: Logs in LTSV format to `/var/log/nginx/access.log`
- **MySQL**: Slow query log at `/var/log/mysql/mysql-slow.log`
- **alp**: Configured with URI pattern matching in `tool-config/alp/config.yaml`
- **Profiling**: pprof and fgprof endpoints available at ports 6060, 8090, 9090

### Service Management
- Go app: `isu-go.service` (port 8080)
- Ruby app: `isu-ruby.service`
- PHP app: `php8.3-fpm`
- Python app: `isu-python.service`

### Critical Paths
- `/initialize`: Database reset endpoint (must complete within 10 seconds)
- Image uploads: Limited to 10MB
- Static files served from `/home/isucon/private_isu/webapp/public/`

## Development Workflow
1. Make changes to application code
2. Run `task -g build` to compile
3. Run `task -g deploy-conf` to update configurations
4. Run `task -g restart` to restart services
5. Monitor with `task -g log` and analyze with `task -g analyze`