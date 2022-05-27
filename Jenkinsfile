pipeline {
    agent any
    tools {
        go 'Go 1.18'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage("build") {
            steps {
                go build 'building'
            }
        }
        stage("test") {
            steps {
                go test
            }
        }
    }
}
