# News Portal

This project is a news portal application built with Go (Golang).

## Getting Started

Follow these instructions to set up and run the project on your local machine.

### Prerequisites

- Go (version 1.16 or later)
- MySQL (or any other preferred database)

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/Danangoffic/news-portal-golang.git
    cd news-portal
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

### Database Setup

1. Create a new database named `news_portal_go`:
    ```sql
    CREATE DATABASE news_portal_go;
    ```

2. Update the database configuration in the `config` file (e.g., `config.yaml` or `config.json`).

### Running the Application

1. Run the application:
    ```sh
    go run main.go
    ```

2. Open your browser and navigate to `http://localhost:8080` to see the application in action.

### Contributing

Contributions are welcome! Please fork the repository and create a pull request.

### License

This project is licensed under the MIT License.

### Contact

For any inquiries, please contact [your email].
