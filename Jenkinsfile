pipeline {
    agent any
    environment {
        GIT_URL = 'https://github.com/ShlokManandhar5/golang_kubefile.git'
        GIT_CREDENTIALS_ID = 'GIT_CREDENTIALS_ID'
        DOCKER_HUB_USERNAME = 'shlokmndr'
        DOCKER_HUB_CREDENTIALS = 'docker-hub-creds'
        IMAGE_NAME = 'golang-app'
        IMAGE_VERSION = "0.0.${BUILD_NUMBER}"
    }
    stages {
        stage('Git checkout') {
            steps {
                git branch: 'main',
                    credentialsId: env.GIT_CREDENTIALS_ID,
                    url: 'https://github.com/ShlokManandhar5/Devops_project_4.git'
            }
        }
        stage('Making Image') {
            steps {
                sh '''
                    docker build -t ${IMAGE_NAME}:${IMAGE_VERSION} .
                '''
            }
        }
        stage('Pushing to Docker Hub') {
            steps {
                withCredentials([
                    usernamePassword(
                        credentialsId: "${DOCKER_HUB_CREDENTIALS}",
                        usernameVariable: 'USERNAME',
                        passwordVariable: 'PASSWORD'
                    )
                ]) {
                    sh '''
                        echo "$PASSWORD" | docker login -u "$USERNAME" --password-stdin
                        docker tag ${IMAGE_NAME}:${IMAGE_VERSION} \
                            ${USERNAME}/${IMAGE_NAME}:${IMAGE_VERSION}
                        docker push ${USERNAME}/${IMAGE_NAME}:${IMAGE_VERSION}
                        docker logout
                    '''
                }
            }
        }
        stage('Cleanup local images') {
            steps {
                sh '''
                    docker rmi ${IMAGE_NAME}:${IMAGE_VERSION} || true
                    docker rmi ${DOCKER_HUB_USERNAME}/${IMAGE_NAME}:${IMAGE_VERSION} || true
                '''
            }
        }
        stage('Kubefile Git checkout') {
            steps {
                git branch: 'main',
                    credentialsId: env.GIT_CREDENTIALS_ID,
                    url: env.GIT_URL
            }
        }
        stage('Update Image Tag') {
            steps {
                sh '''
                    sed -i "s|image: shlokmndr/golang-app:.*|image: shlokmndr/golang-app:${IMAGE_VERSION}|" deploy-golang.yaml
                    echo "New image:"
                    grep "image:" deploy-golang.yaml
                '''
            }
        }
        stage('Push deploy.yaml') {
            steps {
                withCredentials([
                    usernamePassword(
                        credentialsId: "${GIT_CREDENTIALS_ID}",
                        usernameVariable: 'GIT_USERNAME',
                        passwordVariable: 'GIT_PASSWORD'
                    )
                ]) {
                    sh '''
                        git config user.name "shlok"
                        git config user.email "shlokmanandhar5@example.com"
                        git add deploy-golang.yaml
                        git commit -m "Update image tag to ${IMAGE_VERSION}" || true
        
                        CLEAN_USER=$(printf '%s' "$GIT_USERNAME" | tr -d '[:space:]')
                        CLEAN_PASS=$(printf '%s' "$GIT_PASSWORD" | tr -d '[:space:]')
        
                        git push "https://${CLEAN_USER}:${CLEAN_PASS}@github.com/ShlokManandhar5/golang_kubefile.git" HEAD:main
                    '''
                }
            }
        }
    }
}
