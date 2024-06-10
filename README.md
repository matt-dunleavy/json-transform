![](C:\Users\matth\OneDrive\Documents\GitHub\json-transform\assets\images\json-transform.png)

# json-transform

json-transform is a microservice and standalone CLI tool for performing AI batch operations on JSON files using OpenAI (Chat GPT) and Google Gemini. This package is part of the Napoleon suite of microservices.



## Installation

Ensure you have [Go](https://golang.org/doc/install) installed on your system.

#### Linux

1. **Download the binary:**

   ```sh
   wget https://github.com/matt-dunleavy/json-transform/releases/download/v1.0.0/json-transform-linux-amd64 -O json-transform
   ```

2. **Make it executable:**

   ```sh
   chmod +x json-transform
   ```

3. **Move it to a directory in your PATH:**

   ```sh
   sudo mv json-transform /usr/local/bin/
   ```

#### Windows

1. **Download the binary:**
   - Go to the [Releases page](https://github.com/matt-dunleavy/json-transform/releases) and download `json-transform-windows-amd64.exe`.

2. **Add the executable to your PATH:**
   - Move the `json-transform-windows-amd64.exe` to a directory included in your PATH, e.g., `C:\Windows\System32`.

#### MacOS

1. **Download the binary:**

   ```sh
   curl -L https://github.com/matt-dunleavy/json-transform/releases/download/v1.0.0/json-transform-darwin-amd64 -o json-transform
   ```

2. **Make it executable:**

   ```sh
   chmod +x json-transform
   ```

3. **Move it to a directory in your PATH:**

   ```sh
   sudo mv json-transform /usr/local/bin/
   ```



## Compiling from Source

If you prefer to build the binary from source, follow these steps:

1. **Clone the repository:**

   ```sh
   git clone https://github.com/matt-dunleavy/json-transform.git
   cd json-transform
   ```

2. **Install dependencies:**

   ```sh
   go mod tidy
   ```

3. **Build the binary:**

   ```sh
   go build -o json-transform
   ```

4. **Move it to a directory in your PATH:**

   ```sh
   sudo mv json-transform /usr/local/bin/  # For Linux/MacOS
   move json-transform C:\Windows\System32  # For Windows
   ```

## Usage

```
json-transform [command]
```



### Standard Console Commands

#### Batch Operations

Perform batch operations on JSON files with the `batch` command by providing the directory containing the source json files, the desired output directory, and the desired batch operations.

```
json-transform batch -i <input_directory> -o <output_directory> -p <operation> -k <api_key> -f <prompt_file> -s <api_service> -m <model>
```

##### **Flags**

- `-i, --input` (required): Input directory containing JSON files.
- `-o, --output` (required): Output directory for processed JSON files.
- `-p, --operation` (required): Batch operation to perform (merge, split, filter, api-process).
- `-k, --api-key` (required): API key for external API.
- `-f, --prompt` (required): File containing the prompt for API processing.
- `-s, --api-service`: API service to use (chatgpt or gemini). Defaults to `chatgpt`.
- `-m, --model`: AI model to use.

##### **Example**

```
json-transform batch -i ./input -o ./output
```



#### Transform

Transform JSON files using specified rules.

```
json-transform transform -i <input_file> -o <output_file> -r <rules_file>
```

**Flags:**

- `-i, --input` (required): Input JSON file.
- `-o, --output` (required): Output JSON file.
- `-r, --rules` (required): Transformation rules file.

```
json-transform transform -i input.json -o output.json -r rules.json
```



#### Validate

Validate JSON files against a schema.

```
json-transform validate -i <input_file> -s <schema_file>
```

**Flags:**

- `-i, --input` (required): Input JSON file.
- `-s, --schema` (required): JSON schema file.

**Example:**

```
json-transform validate -i input.json -s schema.json
```



#### Help

To get help for any command, you can use the `--help` flag.

```
json-transform --help
```

or

```
json-transform [command] --help
```



## Batch Operations Powered by AI

The `json-transform` tool leverages AI APIs (e.g. ChatGPT, Google Gemini, etc.) to perform advanced operations on JSON files. 

### Usage

#### Instructions

1. **Setup API Keys:**

   - Ensure you have a valid API key for the chosen AI service (ChatGPT or Gemini).
   - Keep your API key secure and do not expose it in your code or public repositories.

2. **Prepare the Prompt File:**

   - Create a text file containing the prompt that will be used by the AI service. This prompt should include instructions on what you want the AI to do with the input JSON content.
   - Use placeholders like `{{input}}` to denote where the content from the JSON file should be inserted in the prompt.

   **Run the Batch Command:**

   - Use the `batch` command with the appropriate flags to process JSON files using the AI service. Specify the input directory, output directory, operation (`api-process`), API key, prompt file, API service (`chatgpt` or `gemini`), and the model if applicable.

```
json-transform batch -i <input_directory> -o <output_directory> -p api-process -k <api_key> -f <prompt_file> -s <api_service> -m <model>
```

**Example:**

```
json-transform batch -i ./input -o ./output -p api-process -k your_api_key -f prompt.txt -s chatgpt -m text-davinci-003
```

**Review the Output:**

- The processed JSON files will be saved in the specified output directory.
- Responses from the AI service will be logged to the console and written to the output files.

### Best Practices for AI-Powered Operations

1. **Prompt Design:**
   - Craft clear and concise prompts to guide the AI in performing the desired task.
   - Test and iterate on your prompts to improve the accuracy and relevance of the AI's responses.
2. **Input Data Quality:**
   - Ensure the input JSON files are well-structured and contain the necessary data for the AI to process.
   - Clean and preprocess your data to avoid errors during processing.
3. **Model Selection:**
   - Choose the appropriate AI model based on the complexity and nature of the task. For example, use `text-davinci-003` for more complex and nuanced tasks with ChatGPT.
4. **Error Handling:**
   - Implement robust error handling to manage issues such as network errors, API rate limits, and invalid input data.
   - Log errors for troubleshooting and to understand any issues with the AI processing.
5. **Security:**
   - Keep your API key secure and avoid hardcoding it in scripts or sharing it publicly.
   - Follow best practices for managing sensitive information, such as using environment variables to store API keys.
6. **Usage Limits:**
   - Be aware of the API usage limits and quotas for your AI service. Monitor your usage to avoid exceeding these limits.
   - Consider implementing rate limiting in your application to stay within the allowed usage limits.
7. **Scalability:**
   - For large-scale processing, consider batching requests and using asynchronous processing to handle high volumes of JSON files efficiently.
   - Optimize the performance of your application to handle the load without affecting the AI service's response time.

By following these instructions and best practices, you can effectively leverage the AI capabilities of the `json-transform` tool to perform advanced operations on your JSON files.

#### Text Summaries

Utilize the AI API to summarize long text documents stored in JSON files. This is particularly useful for processing large amounts of textual data such as reports, articles, or meeting transcripts.

```
json-transform batch -i ./documents -o ./summaries -p api-process -k your_api_key -f summarize_prompt.txt -s chatgpt -m text-davinci-003
```

**Prompt:** (e.g. `summarize_prompt.txt`):

```
Summarize the following text:

{{input}}
```

#### Entity Extraction from Logs

Use the AI API to extract entities (like names, dates, locations, etc.) from JSON files containing log data. This can help in parsing and structuring unstructured log data for further analysis.

```
json-transform batch -i ./logs -o ./structured_logs -p api-process -k your_api_key -f entity_extraction_prompt.txt -s chatgpt -m text-davinci-003
```

**Prompt** (e.g. `entity_extraction_prompt.txt`):

```
Extract all the entities (names, dates, locations, etc.) from the following log entry:

{{input}}
```

#### Code Generation from Specifications

Integrate the AI API to generate code snippets or scripts from JSON files containing specifications. This can be useful for automating code generation tasks based on defined requirements.

```
json-transform batch -i ./specifications -o ./generated_code -p api-process -k your_api_key -f code_generation_prompt.txt -s chatgpt -m text-davinci-003
```

**Prompt** (e.g. `code_generation_prompt.txt`):

```
Generate a Python script based on the following specification:

{{input}}
```

These examples illustrate how the AI API integration can be utilized to enhance the functionality of the `json-transform` tool, providing powerful automation capabilities for a variety of tasks.

#### Language Translation

Leverage the AI API to translate JSON files containing text in one language to another language. This is useful for global businesses needing to process multi-language documents.

```
json-transform batch -i ./spanish_texts -o ./english_texts -p api-process -k your_api_key -f translate_prompt.txt -s chatgpt -m text-davinci-003
```

**Prompt** (e.g. `translate_prompt.txt`):

```
Translate the following text from Spanish to English:

{{input}}
```

#### Sentiment Analysis on Customer Feedback

Use the AI API to perform sentiment analysis on JSON files containing customer feedback or reviews. This can help businesses understand customer sentiments and improve their products or services.

```
json-transform batch -i ./feedback -o ./processed -p api-process -k your_api_key -f sentiment_prompt.txt -s chatgpt -m text-davinci-003
```

**Prompt:** (e.g. `sentiment_prompt.txt`)

```
Analyze the following customer feedback and determine the sentiment (positive, negative, neutral):

{{input}}
```

#### 

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on [GitHub](https://github.com/matt-dunleavy/json-transform).



## License

This package is distributed under the [MIT License](LICENSE.md).
