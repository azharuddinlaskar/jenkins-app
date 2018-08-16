node('master') {
    docker.withServer('unix:///var/run/docker.sock') {
        stage('Code Checkout') {
            docker
                .image('golang:latest')
                .inside('--volumes-from jenkins-ci') {
                    checkout scm
                }
        }
        stage('Unit Test') {
            docker
                .image('golang:latest')
                .inside('--volumes-from jenkins-ci') {
                    sh """
                        go test ./
                    """
                }
        }
        stage('SonarQube Analysis') {
            docker
                .image('azharuddinlaskar/sonarscanner')
                .inside('--volumes-from jenkins-ci --link sonarqube --net ci-pipeline-net') {
                    sh """
                        sonar-scanner
                    """
                }
        }
        stage('Build') {
            docker
                .image('golang:latest')
                .inside('--volumes-from jenkins-ci') {
                    echo "Deployinh"
                }
        }
    }
}
