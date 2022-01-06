# TODO APP API

## Usage
### Prerequisites
- Code Editor / IDE
- [Go Installation](https://golang.org/dl/)

## Installation

1. Clone the repository
    ```bash
    git clone https://github.com/harisfi/todo-app-api.git
    ```

2. Configure .env files, => copy .env.example and rename it to .env
    ```bash
    cp .env.example .env
    ```

3. Set your database configuration in .env files

4. Download and verify dependencies
    ```bash
    go get -d -v
    ```

5. Run server
    ```bash
    go run main.go
    ```

### Building the project for deployment

1. Run build command
    ```bash
    go build
    ```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)