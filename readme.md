# Qualabs Challenge

## Instructions to Run the Code
1. Ensure Go is installed on your system. If not, follow the installation instructions for your operating system:
   - For Linux: `sudo apt update && sudo apt install golang-go`
   - For Windows: Download and install Go from the official website: https://golang.org/dl/
   - For Mac: `brew install go`

2. Clone the repository to your local machine.
3. Navigate to the `/qualabs` directory.
4. Run the following command to build the application:
   ```
   Linux -> CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./qualabs
   Windows -> CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./qualabs.exe
   Mac -> CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ./qualabs
   ```
5. Once the build is successful, you can run the application with the following command:
   ```
   ./linux -source <path_to_data_directory> -operation <operation_type>
   ./mac -source <path_to_data_directory> -operation <operation_type>
   ./windows.exe -source <path_to_data_directory> -operation <operation_type> 
   ```
   Replace `<path_to_data_directory>` (it uses ./data by default) with the path to your data directory and `<operation_type>` with the operation you want to run (A or B).
6. The application will load the data, perform the specified operation, and display the result.
