node('master') {
    docker.withServer('unix:///var/run/docker.sock') {
        stage('Code Checkout') {
            
        }
        stage('Code Checkout') {
            docker
                .image('golang:latest')
                .inside('--volumes-from jenkins-master') {
                    sh """
                        git 'https://github.com/azharuddinlaskar/jenkins-app.git'
                    """
                }
        }
        stage('Unit Test') {
            docker
                .image('golang:latest')
                .inside('--volumes-from jenkins-master') {
                    sh """
                        go test ./
                    """
                }
        }
        stage('SonarQube Analysis') {
            docker
                .image('sonarqube-sonarscanner')
                .inside('--volumes-from jenkins-master') {
                    sh """
                        sonar-scanner
                    """
                }
        }
        stage('Build') {
            docker
                .image('golang:latest')
                .inside('--volumes-from jenkins-master') {
                    echo "Deployinh"
                }
        }
    }
}
