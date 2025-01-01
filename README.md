*documentation for running each docker-compose*

## Creating development, testing and production environments

### Prerequisites

- Set environment variables in `.env` file. The `.env.example` file is provided as a template.

### Development

1. Clone the repository
2. Run the following command to start the development environment:
```bash
docker-compose -f docker-compose.dev.yml up
```

### Testing

1. Clone the repository
2. Run the following command to start the production environment:
```bash
docker-compose -f docker-compose.test.yml up
```

### Production

1. Clone the repository
2. Run the following command to start the production environment:
```bash
docker-compose -f docker-compose.prod.yml up
```

