node {
    def myGradleContainer = docker.image('gradle:jdk8-alpine')
    myGradleContainer.pull()

    stage('prep') {
        ws("/var/lib/jenkins/goworkspace/src/github.com/helloworld"){
          git url: 'https://github.com/grv231/helloworld.git'
          sh 'go test -coverprofile=coverage.out'
      }
    }

    stage('build') {
      myGradleContainer.inside("-v ${env.HOME}/.gradle:/home/gradle/.gradle"){
        sh '/opt/gradle/bin/gradle build'
      }
    }

    stage('sonar-scanner') {
      def sonarqubeScannerHome = tool name: 'sonar', type: 'hudson.plugins.sonar.SonarRunnerInstallation'
      withCredentials([string(credentialsId: 'sonar', variable: 'sonarLogin')]) {
        sh "${sonarqubeScannerHome}/bin/sonar-scanner -e -Dsonar.host.url=http://sonarqube:9000 -Dsonar.go.coverage.reportPaths=/var/lib/jenkins/goworkspace/src/github.com/helloworld/coverage.out -Dsonar.login=${sonarLogin} -Dsonar.projectName=hello-test -Dsonar.projectVersion=${env.BUILD_NUMBER} -Dsonar.projectKey=HW -Dsonar.sources=. -Dsonar.tests=."
      }
    }
}
