pipeline {
    agent any  // This specifies that the pipeline can run on any available agent (Jenkins slave).

    stages {
        stage('Checkout') {
            steps {
                git(url: 'https://github.com/aamuzakii/arukari-go-api', branch: 'master')
            }
        }
    }

    post {
        success {
            // Actions to perform if the pipeline succeeds.
            echo 'Pipeline succeeded! Deploying to production...'
            // You can add additional steps here.
        }

        failure {
            // Actions to perform if the pipeline fails.
            echo 'Pipeline failed! Notify the team...'
            // You can add additional steps here, like sending notifications.
        }
    }
}
