pipeline {
    agent any

    stages {
        stage('Build and Test') {
            steps {
                script {
                    // Checkout code from your Git repository
                    checkout scm

                    // Set up Go
                    tool name: 'Go', type: 'GoInstallation'

                    // Build and test your Go application
                    sh 'go build'
                    sh 'go test ./...'
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    // Copy files to the deployment server using SCP
                    sh 'scp -i $DEPLOY_SSH_KEY -r ./* $DEPLOY_USERNAME@$DEPLOY_HOST:/path/to/deployment/directory'

                    // SSH into the server and restart the application
                    sshagent(['DEPLOY_SSH_KEY']) {
                        sh "ssh $DEPLOY_USERNAME@$DEPLOY_HOST 'cd /path/to/deployment/directory && systemctl restart your-app.service'"
                    }
                }
            }
        }
    }
}
