pipeline {
    agent any

    environment {
        GO_VERSION = "1.23.1"
        CGO_ENABLED = "0"   // Disable C bindings for reliable cross-compilation
    }

    stages {
        stage('Verify Go') {
            steps {
                sh 'go version'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Build') {
            steps {
                sh '''
                  GOOS=linux   GOARCH=amd64 go build -o app-linux-amd64
                  GOOS=windows GOARCH=amd64 go build -o app-windows-amd64.exe
                '''
            }
        }

        stage('Archive') {
            steps {
                script {
                    def version = ""
                    if (env.TAG_NAME) {
                        version = env.TAG_NAME
                    } else if (env.BRANCH_NAME) {
                        version = "${env.BRANCH_NAME}-${env.BUILD_NUMBER}"
                    } else {
                        version = "dev-${env.BUILD_NUMBER}"
                    }

                    sh '''
                      mv app-linux-amd64 app-${version}-linux-amd64
                      mv app-windows-amd64.exe app-${version}-windows-amd64.exe
                    '''

                    archiveArtifacts artifacts: "app-${version}-*", fingerprint: true
                }
            }
        }
    }
}

