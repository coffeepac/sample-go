Jenkinsfile (Declarative Pipeline)
pipeline {
    stages {
        stage('build') {
            steps {
                sh 'go build'
            }
        }
        stage('test') {
            steps {
                sh 'go test'
            }
        }
        stage('docker build') {
            steps {
                sh 'docker build -t quay.io/coffeepac/sample-go:jenkins .'
            }
        }
        stage('docker push') {
            steps {
                sh 'docker push -t quay.io/coffeepac/sample-go:jenkins'
            }
        }

    }
}