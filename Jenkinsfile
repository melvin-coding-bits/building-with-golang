pipeline {
  environment {
    registry = "melvinodsa/kube-user-service"
    registryCredential = 'dockerhub'
    dockerImage = ''
  }
  agent {
    kubernetes {
      inheritFrom 'mypod'
    }
  }
  stages {
    stage('Building image') {
      steps{
        script {
          dockerImage = docker.build(registry, "-f user-service/Dockerfile user-service")
        }
      }
    }
    stage('Deploy Image') {
      steps{
        script {
          docker.withRegistry( '', registryCredential ) {
            dockerImage.push()
           }
        }
      }
    }
    stage('Remove Unused docker image') {
      steps{
        sh "docker rmi $registry"
      }
    }

    stage('Pushing to k8s') {
      steps {
        withKubeConfig([credentialsId: "clusterk8sconfig"]) {
          sh "kubectl version"
          sh 'kubectl rollout restart deployment/user-service-deployment -n dezerv'
        }
      }
    }
  }    
}
