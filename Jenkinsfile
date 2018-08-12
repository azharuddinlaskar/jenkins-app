pipeline {
	agent none
	
    stages {
    	stage('Test') {
    		agent {
				docker {
					image 'golang:latest'
					args '-e "GOCACHE=/tmp/gocache"'
					}
				}
            steps {
                echo "Testing..."
                sh "go test ./"
            }
        }
        stage('SonarQube Analysis') {
        	steps {
        		script {
        			def scannerHome = tool 'SonarQube Scanner 2.8';
					withSonarQubeEnv('SONARQUBE-1') {
						sh "${scannerHome}/bin/sonar-scanner"
					}
        		}
        	}
        }
        stage('Build') {
    		agent {
				docker {
					image 'golang:latest'
					args '-e "GOCACHE=/tmp/gocache"'
					}
				}
            steps {
                echo "Building..."
                checkout scm
                sh "go build -o matrix"
            }
        }
    }
}
