node('master') {
    docker.withServer('unix:///var/run/docker.sock') {
        stage('Code Checkout') {
            
        }
        stage('Code Checkout') {
            docker
                .image('golang:latest')
                .inside('--volumes-from jenkins') {
                    sh """
                        git 'https://github.com/azharuddinlaskar/jenkins-app.git'
                    """
                }
        }
        stage('Unit Test') {
            docker
                .image('golang:latest')
                .inside('--volumes-from jenkins') {
                    sh """
                        go test ./
                    """
                }
        }
        stage('SonarQube Analysis') {
            docker
                .image('sonarqube-sonarscanner')
                .inside('--volumes-from jenkins') {
                    sh """
                        sonar-scanner
                    """
                }
        }
        stage('Build') {
            docker
                .image('golang:latest')
                .inside('--volumes-from jenkins') {
                    echo "Deployinh"
                }
        }
    }
}
