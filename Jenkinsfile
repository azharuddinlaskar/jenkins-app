pipeline {
    agent { docker { image 'golang:latest' } }
    stages {
        stage('build') {
            steps {
                sh 'go run matrix.go'
            }
        }
    }
}
