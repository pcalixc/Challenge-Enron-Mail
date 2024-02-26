# Enron Mail Indexer and Data Visualizer

## Overview

The Enron Mail Indexer and Data Visualizer project is a combination of an Indexer and an API that hosts a Vue application serving as a User Interface for searching over indexed email files. The Indexer is responsible for crawling through a specified directory structure, extracting information from email files, and indexing them using the ZincSearch API. The API, built with Go and Chi, allows users to search for emails based on specific keywords.

### Technologies Used:

- [Go 1.21.5](https://go.dev)
- [Go-Chi/V5](https://github.com/go-chi/chi)
- [Vue.js](https://vuejs.org)
- [ZincSearch](https://github.com/zincsearch/zincsearch)
- [ZincSearch API](https://zincsearch-docs.zinc.dev)

## Configuration

### 1. Clone the Git Repository

```bash
git clone https://github.com/pcalixc/Challenge-Enron-Mail.git

```
### 2. Download and Extract the Enron Email Dataset

```bash
curl -L http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz -o enron_mail_20110402.tgz && tar -xf enron_mail_20110402.tgz
```

### 3. Set Up ZincSearch

-   Download [ZincSearch](https://github.com/zincsearch/zincsearch) and follow the [ZincSearch Quick Start Guide](https://zincsearch-docs.zinc.dev/quickstart/).
-   Alternatively, you can run ZincSearch with Docker 🐳 using the following commands:

```bash
mkdir data

docker run -v /full/path/of/data:/data -e ZINC_DATA_PATH="/data" -p 4080:4080 \
-e ZINC_FIRST_ADMIN_USER=admin -e ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123 \
--name zincsearch public.ecr.aws/zinclabs/zincsearch:latest` 
```

**After completing the above steps, you are ready to start the indexing process!**


## Indexer (Go Application)

### Getting Started
1. Navigate to the indexer directory:   `cd ./indexer`

2.  The indexer requires one parameter: the path to the text file containing emails.

### Usage

To run the indexer, execute the following command:

`go run ./ <enron-mails-path>` 

Replace **`<enron-mails-path>`** with the actual path to your directory previuosly downloaded containing the emails.

### Example
`go run ./ ./enron-mails` 

### Notes
-   Ensure that you have Go installed on your system.
-   The indexer process will start immediately upon execution.

## API (Go Application)

The API is a Go application built with the Chi router. It provides basic CORS support and exposes endpoints for retrieving information from the ZincSearch index.

1. Navigate to the api directory:   `cd ./api`
2. To run the program, execute the following command:

`go run ./` 

_Remember to add your environment variables_

### Endpoints

- `/`: 
  - **Method:** GET
  - **Description:** Welcome message endpoint.

- `/emails`: 
  - **Method:** GET
  - **Description:** Search endpoint to retrieve emails based on a keyword.
  - **Query Parameters:**
    - `page`: (Optional) Specifies the page number for paginated results. Default value is 1 if not provided.
    - `max`: (Optional) Specifies the maximum number of results per page. Default value is 10 if not provided.
    - `term`: (Optional) The keyword to search for within the emails. If not provided, all emails will be retrieved.

**Example Usage:** 
- `/emails?page=1&max=5&term=hello`: Retrieves up to 5 emails containing the keyword "hello".
- `/emails`: Retrieves all emails with default pagination settings.

This endpoint allows users to search for emails containing the specified keyword. The optional parameters `page` and `max` can be used for pagination control. If `term` is not provided, all emails will be returned.


## CLIENT (Vue.js App)
The API also serves a Vue.js dist folder to allow users to interact with the indexed data visually.
1. Navigate to the api directory and run the follow comands:  
```bash
cd ./client
npm install
npm run dev
```
After running these commands, you should be able to access the graphical user interface (GUI) at `http://localhost:3333`.




