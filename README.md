# Go Calculator Web API

A simple Go-based calculator application providing a Web API for basic arithmetic operations. This project is designed for teaching Declarative Pipelines in Jenkins.

## 1. Local Setup & Running

### Prerequisites
- Go 1.16 or higher installed on your machine.

### Running Tests
To run the valid unit tests:
```bash
go test ./...
```
Expected output:
```
ok      devops-go       0.xxx s
```

### Running the Application
To start the web server:
```bash
go run main.go
```
The server will start on port `8080`.

### Testing Endpoints
You can use `curl` or a web browser to test the API:

- **Health Check**:
  ```bash
  curl http://localhost:8080/
  ```
- **Addition**:
  ```bash
  curl "http://localhost:8080/add?a=10&b=20"
  ```
- **Subtraction**:
  ```bash
  curl "http://localhost:8080/sub?a=10&b=5"
  ```
- **Multiplication**:
  ```bash
  curl "http://localhost:8080/mul?a=5&b=5"
  ```
- **Division**:
  ```bash
  curl "http://localhost:8080/div?a=20&b=4"
  ```

---

## 2. Environment Setup

### Installing Go on Ubuntu
Run the following commands to install Go on an Ubuntu system (e.g., your Jenkins agent):

```bash
# 1. Update package list
sudo apt-get update

# 2. Download the Go binary (check https://go.dev/dl/ for the latest version)
wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz

# 3. Remove any previous Go installation and extract the archive into /usr/local
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz

# 4. Add Go binary to your PATH (add this to ~/.profile or ~/.bashrc for persistence)
export PATH=$PATH:/usr/local/go/bin

# 5. Verify installation
go version
```

### Setting up Go on Jenkins
You can configure Go directly on the Jenkins server or on specific agents.

**Option A: Using Global Tool Configuration (Recommended)**
1.  Go to **Manage Jenkins** -> **Global Tool Configuration**.
2.  Scroll down to **Go**.
3.  Click **Add Go**.
4.  Name it `go-1.21` (or your preferred version).
5.  Check "Install automatically".
6.  Save.
7.  *Note*: In your `Jenkinsfile`, you would then use `tools { go 'go-1.21' }`.

**Option B: Manual Installation on Agent**
1.  Follow the Ubuntu installation steps above on the machine running the Jenkins agent.
2.  Ensure `go` is in the `PATH` of the user running the Jenkins agent service.

---

## 3. Student Assignments

Complete the following assignments to practice working with Jenkins and Declarative Pipelines.

### Assignment 1: Enhance the Build Stage
**Task**: Modify the `Jenkinsfile` to build the application for multiple operating systems (Cross-Compilation).
- **Goal**: Create binaries for both Linux and Windows.
- **Hint**: Use Go environment variables `GOOS` and `GOARCH`.
- **Expected Change**:
  ```groovy
  sh 'GOOS=linux go build -o calculator-linux'
  sh 'GOOS=windows go build -o calculator-windows.exe'
  ```

### Assignment 2: Expand the Test Stage
**Task**: Add a new test case to `main_test.go` and unsure the pipeline runs it.
- **Goal**: Add a test for a new scenario (e.g., testing that `/add` with invalid characters returns a 400 Bad Request).
- **Verify**: Commit the change and watch the "Test" stage in Jenkins pass.

### Assignment 3: Improve the Deploy Stage
**Task**: Currently, the deploy stage just echoes a message. Make it more "real" by simulating a release.
- **Goal**: Archive not just one, but ALL binaries created in Assignment 1.
- **Action**: Update the `archiveArtifacts` step to include `calculator-linux` and `calculator-windows.exe`.

### Assignment 4: Pipeline Resilience
**Task**: Add notification steps.
- **Goal**: Print a special message if the pipeline fails.
- **Hint**: Look at the `post` section in the `Jenkinsfile`.
- **Action**: Add a `failure` block that prints "The build has failed! Please check the logs." using the `echo` step.
