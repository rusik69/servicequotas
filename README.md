# Service Quotas

## Overview
Service Quotas is a powerful tool that helps you increate AWS service quotas

## Building the Tool

To build the tool, you can use the provided Makefile. Follow these steps:

1. Clone the repository to your local machine:
    ```
    $ git clone https://github.com/rusik69/servicequotas.git
    ```

2. Navigate to the project directory:
    ```
    $ cd servicequotas
    ```

3. Install the required dependencies:
    ```
    $ make get
    ```

4. Build the tool using the Makefile:
    ```
    $ make build
    ```

## Configuring Quota Increases

To configure quota increases, you need to modify the `config.json` file. Follow these steps:

1. Open the `config.json` file in a text editor.

2. Locate the section for the specific AWS service you want to configure quota increases for.

3. Update the values in the `quota` field to specify the desired quota increase values.

4. Save the `config.json` file.

## Running the Tool

To run the tool and request quota increases based on the configuration in `config.json`, follow these steps:

1. Make sure you have configured your AWS credentials in you environment.

2. Run the tool using the following command:
    ```
    $ ./bin/quotas --region region --config ./test/quotas.json
    ```

3. The tool will read the `config.json` file and automatically request quota increases for the specified AWS services.

4. Monitor the tool's output for the status of the quota increase requests.

That's it! You have successfully built the tool using the Makefile and configured quota increases using the `config.json` file.
