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
    stage('Build') {         
        sh "'${mvnHome}/bin/mvn' -Dmaven.test.failure.ignore clean package"
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