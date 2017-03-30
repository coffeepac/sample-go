/*Jenkinsfile (Declarative Pipeline)
pipeline {
    agent any
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
}*/
node {
    // Install the desired Go version
    def root = tool name: 'Go 1.8', type: 'go'

    // Export environment variables pointing to the directory where Go was installed
    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        sh 'go version'
    }
    
    stage('build') {         
        sh 'go build'
    }
    stage('test') {
        sh 'go test'        
    }
    stage('docker build') {
        sh 'docker build -t quay.io/coffeepac/sample-go:jenkins .'
    }
    stage('docker push') {
        sh 'docker push -t quay.io/coffeepac/sample-go:jenkins'
    }

}