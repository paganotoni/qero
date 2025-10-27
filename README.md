# Qero

This is a simple web application that generates QR codes from text or URLs. It is built with Go, HTMX, and Tailwind CSS.

## Getting Started

### Prerequisites

- Go 1.24 or later

### Installation

1. Clone the repository:

```sh
git clone https://github.com/paganotoni/qero.git
```

2. Navigate to the project directory:

```sh
cd qero
```

3. Download the dependencies:

```sh
go mod download
```

### Running the application

To run the application in development mode, use the following command:

```sh
go tool dev
```

This will start the application on http://localhost:3000. The application will automatically restart when changes are detected in the Go files.

## Building the application

To build the application for production, run the following command:

```sh
go build -o bin/app ./cmd/app
```

This will create a binary file named `app` in the `bin` folder.