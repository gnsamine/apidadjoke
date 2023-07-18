# Random Dad Joke API

This Go program utilizes the Random Dad Joke API to fetch and display a random dad joke. It is built using the Fiber framework for handling HTTP requests.

## Getting Started

To get started with this program, follow the steps below.

### Prerequisites

Make sure you have Go installed on your system. You can download and install Go from the official website: [https://golang.org](https://golang.org)

### Installation

1. Clone the repository to your local machine or download the source code as a ZIP file.

```bash
git clone https://github.com/your-username/random-dad-joke-api.git
```

2. Navigate to the project directory.

```bash
cd random-dad-joke-api
```

3. Install the required dependencies using the `go get` command.

```bash
go get github.com/gofiber/fiber/v2
```

4. Create a file named `key.json` in the same directory as the Go file. Put your Random Dad Joke API key inside this file in the following format:

```json
{
  "api_key": "your-api-key"
}
```

### Usage

To use this program, follow the steps below.

1. Build the Go program.

```bash
go build
```

2. Run the executable.

```bash
./random-dad-joke-api
```

3. Open your web browser and navigate to `http://localhost:8000` to see a random dad joke displayed.

### Configuration

You can modify the program's behavior by changing the following configurations:

- Port: By default, the program listens on port `8000`. If you want to use a different port, modify the `app.Listen()` function call in the `main()` function.

### API Key

Make sure to obtain an API key from the Random Dad Joke API before running the program. The API key should be placed in the `key.json` file as mentioned in the installation steps.

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

### Acknowledgments

- This program uses the Random Dad Joke API to fetch random dad jokes. Visit their website for more information: [https://random-dad-joke-api.herokuapp.com](https://random-dad-joke-api.herokuapp.com)

## Contact

For any questions or inquiries, please contact amine0gunes@gmail.com.
