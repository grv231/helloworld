node {
    def myGradleContainer = docker.image('gradle:jdk8-alpine')
    myGradleContainer.pull()

    stage('prep') {
        //ws("/var/lib/jenkins/goworkspace/src/github.com/helloworld"){
        git url: 'https://github.com/grv231/helloworld.git'
      //}
    }

    stage('build') {
      myGradleContainer.inside("-v ${env.HOME}/.gradle:/home/gradle/.gradle"){
        sh '/opt/gradle/bin/gradle build'
        sh 'go test -coverprofile=coverage.out'
      }
    }

    stage('sonar-scanner') {
      def sonarqubeScannerHome = tool name: 'sonar', type: 'hudson.plugins.sonar.SonarRunnerInstallation'
      withCredentials([string(credentialsId: 'sonar', variable: 'sonarLogin')]) {
        sh "${sonarqubeScannerHome}/bin/sonar-scanner -e -Dsonar.host.url=http://localhost:9000 -Dsonar.go.coverage.reportPaths=./coverage.out -Dsonar.login=${sonarLogin} -Dsonar.projectName=hello-test -Dsonar.projectVersion=${env.BUILD_NUMBER} -Dsonar.projectKey=HW -Dsonar.sources=. -Dsonar.tests=."
      }
    }
}
