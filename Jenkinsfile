pipeline {
    agent { docker { image 'golang:latest' } }
    stages {
        
        stage('Test') {
            steps {
                echo "Testing..."
                sh "go build -o matrix"
            }
        }
        stage('Build') {
            steps {
                echo "Building..."
                sh "go build -o matrix"
            }
        }
        stage('Deploy') {
            steps {
                echo "Deploying"
            }
        }
    }
}
