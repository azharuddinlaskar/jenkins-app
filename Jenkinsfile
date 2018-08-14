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
                .image('sonarscanner')
                .inside('--volumes-from jenkins-ci --link sonarqube --net development_pipeline_go_default') {
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
