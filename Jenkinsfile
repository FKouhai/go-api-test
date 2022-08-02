pipeline {
  agent any
  tools {
    go '1.18'
  }
  environment {
    GO111MODULE="on"
    def CURRDATE = sh(script: "echo `date +%m_%d_%y`", returnStdout: true).trim()
    GOPATH="${WORKSPACE}"
    BINDEST="${JENKINS_HOME}/binaries"
    AGENTBIN="api_${BUILD_ID}_${CURRDATE}"
  }
  stages {
    stage("build"){
      steps {
        echo 'Building the api binary'
        sh 'ls $GOPATH'
        sh 'rm $GOPATH/go.mod'
        sh 'go build -o ${AGENTBIN} main'
      }
    }
  }
  post {
    success {
      echo 'Building has finished successfully'
      sh "cp ${AGENTBIN} ${env.BINDEST}"
    }
    failure {
      echo 'Building has failed'
    }
  }
}