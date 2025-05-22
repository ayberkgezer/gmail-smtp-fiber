# gmail-smtp-fiber

A simple, production-ready REST API for sending emails via Gmail SMTP, built with [Go Fiber](https://gofiber.io/).
This project demonstrates clean architecture, middleware usage, request validation, and graceful shutdown in Go.

---

## Features

- **Send Emails** via Gmail SMTP
- **API Key Authentication** middleware
- **Request Validation** with detailed error responses
- **Request Logging** with request IDs for traceability
- **Graceful Shutdown** on SIGINT/SIGTERM
- **Environment-based Configuration** via `.env`
- **Health Check Endpoint**
- **Clean, Modular Architecture**

---

## Getting Started

### Prerequisites

- [Go 1.20+](https://go.dev/dl/)
- A Gmail account with [App Passwords enabled](https://support.google.com/accounts/answer/185833?hl=en)
- (Optional) [Make](https://www.gnu.org/software/make/) for build/run commands

---

### Clone the Repository

```bash
git clone https://github.com/ayberkgezer/gmail-smtp-fiber.git
cd gmail-smtp-fiber
```

---

### Configuration

1. **Copy the example environment file:**

   ```bash
   cp .env.example .env
   ```

2. **Edit `.env` with your settings:**

   ```
   ENV=development
   PORT=3000
   API_KEY=your_api_key_here

   # Gmail SMTP Settings
   SMTP_HOST=smtp.gmail.com
   SMTP_PORT=587
   SMTP_USERNAME=your_gmail@gmail.com
   SMTP_PASSWORD=your_gmail_app_password
   EMAIL_FROM=recipient@example.com
   ```

   - `SMTP_USERNAME`: Your Gmail address (must have App Password enabled)
   - `SMTP_PASSWORD`: Your Gmail App Password (not your normal password)
   - `EMAIL_FROM`: The recipient email address (where emails will be sent)

---

### Build & Run

#### With Make

```bash
make build
make run
```

#### Or with Go

```bash
make dev
```

---

## API Documentation

### Authentication

All endpoints require an `X-API-KEY` header matching your `.env` `API_KEY`.

---

### Health Check

**GET** `/api/v1/health`

- **Response:**
  ```json
  {
    "statusCode": 200,
    "message": "Server Running",
    "success": true
  }
  ```

---

### Send Email

**POST** `/api/v1/email/send`

#### Headers

- `Content-Type: application/json`
- `X-API-KEY: your_api_key_here`

#### Request Body

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "emailMessage": "Hello, this is a test message!"
}
```

- `name`: Sender's name (min 3, max 100 chars)
- `email`: Sender's email (must be valid)
- `emailMessage`: Message body (min 10, max 1000 chars)

#### Success Response

```json
{
  "statusCode": 200,
  "message": "Email successfully sent",
  "success": true
}
```

#### Validation Error Example

```json
{
  "statusCode": 400,
  "success": false,
  "errorMessage": "email is invalid (tag=email, param=)"
}
```

#### Unauthorized Example

```json
{
  "statusCode": 401,
  "success": false,
  "errorMessage": "Invalid API Key"
}
```

---

## Project Structure

```
gmail-smtp-fiber/
├── cmd/app/                # Application entrypoint
├── internal/
│   ├── app/
│   │   ├── controller/     # HTTP controllers
│   │   ├── handler/        # Business logic handlers
│   │   ├── model/          # Request/response models
│   │   ├── router/         # Route definitions
│   │   └── services/       # Email sending logic
│   ├── base/               # Base response types
│   ├── common/
│   │   ├── httpfilter/     # Error handling
│   │   ├── middleware/     # Middleware (API key, logger, request ID)
│   │   └── validation/     # Request validation
│   ├── config/             # Environment config loader
│   ├── domain/             # Domain models
│   └── server/             # Server startup & graceful shutdown
├── .env.example
├── Makefile
├── go.mod
├── go.sum
└── README.md
```

---

## Security Notes

- **Never commit your real `.env` file or credentials.**
  The `.env` file is in `.gitignore` by default.
- Use [Gmail App Passwords](https://support.google.com/accounts/answer/185833?hl=en) for SMTP authentication.
- Change your `API_KEY` to a strong, random value.

---

## Troubleshooting

- **"Error loading .env file"**
  Ensure you have a `.env` file in the project root.

- **"535-5.7.8 Username and Password not accepted"**
  - Make sure you are using an App Password, not your Gmail login password.
  - Check that "Less secure app access" is enabled or use App Passwords.

- **Cannot send email to arbitrary addresses**
  This starter sends all emails to the `EMAIL_FROM` address for demo purposes.
  You can modify the code to allow dynamic recipients if needed.

---

## License

MIT License © 2025 Ayberk Gezer

---

## Author

- [Ayberk Gezer](https://github.com/ayberkgezer)

---

## Contributing

Pull requests and issues are welcome!

---

## Acknowledgements

- [Fiber Web Framework](https://gofiber.io/)
- [Go Playground Validator](https://github.com/go-playground/validator)
- [gocolorlog](https://github.com/ayberkgezer/gocolorlog)

---

Happy Coding! 🚀
