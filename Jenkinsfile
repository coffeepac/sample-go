podTemplate(label: 'sample-go', containers: [
    containerTemplate(name: 'jnlp', image: 'jenkinsci/jnlp-slave:2.62-alpine', args: '${computer.jnlpmac} ${computer.name}'),
    containerTemplate(name: 'golang', image: 'golang:1.7.5', ttyEnabled: true, command: 'cat'),
    containerTemplate(name: 'docker', image: 'docker', command: 'cat', ttyEnabled: true)
  ], volumes: [
    hostPathVolume(hostPath: '/var/run/docker.sock', mountPath: '/var/run/docker.sock'),
    secretVolume(mountPath: '/home/jenkins/.docker/', secretName: 'coffeepac-quay-robot-dockercfg')
  ]) {
    node('sample-go') {
        container('golang'){

            stage('checkout') {
                git url: 'https://github.com/coffeepac/sample-go'
            }    

            stage('build') {
                sh 'go build -v'
            }

            stage('the hell mk II') {
                sh 'ls $HOME; ls ~'
            }

            stage('test') {
                sh 'go test -v'
            }
        }

        container('docker') {
            stage('docker build') {
                sh 'docker build -t quay.io/coffeepac/sample-go:jenkins .'
            }

            stage('docker push') {
                sh 'docker push quay.io/coffeepac/sample-go:jenkins'
            }
        }
    }
  }  