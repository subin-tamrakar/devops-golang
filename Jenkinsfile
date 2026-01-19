pipeline {
    agent any
	tools { go 'go-1.23.10' }

    stages {
        stage('Verify Go') {
            steps {
		echo "verifying go"
            }
        }

        stage('Test') {
            steps {
		echo "testing go"
            }
        }

        stage('Build') {
            steps {
		echo "building"
            }
        }

        stage('Archive') {
            steps {
			echo "archiving"
                }
            }
        }
    }


