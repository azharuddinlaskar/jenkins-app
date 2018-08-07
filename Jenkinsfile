pipeline {
    agent { docker { image 'golang:latest' } }
    stages {
        stage('Build') {
            steps {
                sh "pwd"
                sh "ls -l"
                sh "go version"
            }
        }
        stage('Test') {
            steps {
                sh "pwd"
            }
        }
        stage('Deploy') {
            steps {
                sh "pwd"
            }
        }
    }
}
