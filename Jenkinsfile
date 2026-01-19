pipeline {
    agent any
	tools { go 'go1.23.10' }

    stages {
        stage('Verify Go') {
            steps {
		echo "verifying go"
            }
        }

        stage('Test') {
            steps {
		sh 'go test ./...'
		echo "testing go"
            }
        }

        stage('Build') {
            steps {
		sh '''
			mkdir -p test
			go build -o test/calculator
		'''
		echo "building"
            }
        }

        stage('Archive') {
            steps {
		archiveArtifacts artifacts:"test/calculator*", fingerprint:true
			echo "archiving"
                }
            }
        }
    }


