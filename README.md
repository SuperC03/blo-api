# Banana Lounge API

## Development
- Using [Swaggo](https://github.com/swaggo/swag) for documentation generation
- Using [Goose](https://github.com/pressly/goose) for database migrations

## Environment Variables

### Application Variables
#### Server Params
- `HOSTNAME`
- `PORT`
- `PUBLIC_URL`
- `IS_PRODUCTION`: 1 if in production environment, 0 otherwise

#### PostgreSQL Database
- `POSTGRES_USER`
- `POSTGRES_PASSWORD`
- `POSTGRES_NAME`
- `POSTGRES_HOST`
- `POSTGRES_PORT`

#### Gotify Notifications
- `GOTIFY_ENABLED`
- `GOTIFY_URL`
- `GOTIFY_TOKEN`

#### Google Cloud Credentials
- `GOOGLE_KEY_PATH`

#### Swagger Configuration
- `SWAGGER_HOST`
- `SWAGGER_BASE_PATH`

### Google Sheets Variables
#### Banana Log Values
- `BANANALOG_DOC_ID`
- `BANANALOG_OVERVIEW_DATA_RANGE`: Range for relavent data
