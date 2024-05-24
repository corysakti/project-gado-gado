// RITA_APP_SIT, RITA_SIT_13 -- alternative node if not running
node('RITA_SIT_13') {
    stage('Git Repo') {
        git branch: params.GITBRANCH, credentialsId: 'jenkinsforgitlabce', url: 'http://10.64.17.56/vas1/iot.git'
    }
    stage('Build and Test') {
        sh "mvn clean package -DskipTests"
        echo "${currentBuild.projectName}"
        echo "BUILD_TAG = ${BUILD_TAG}"
        echo "JOB_NAME = ${JOB_NAME}"
        echo "JOB_BASE_NAME = ${JOB_BASE_NAME}"
    }

    stage('Deploy') {
        sh "pwd"
        def PROJECT_PATH = sh(script: 'pwd', returnStdout: true).trim()
        def TARGET_WAR = sh(script: 'ls target/iot*.war', returnStdout: true).trim()

        // Error handling: Check if WAR file exists
        if (TARGET_WAR.empty) {
            error "No WAR file found. Deployment aborted."
        }

        def JBOSS_USER = sh(script: 'echo iot', returnStdout: true).trim()
        def JBOSS_PASSWORD = sh(script: 'echo iot123', returnStdout: true).trim()
        def CONTROLLER_IP = '10.64.25.121'

        sh "TARGET_WAR=${TARGET_WAR}"
        sh "PROJECT_PATH=${PROJECT_PATH}"

        // Deployment
        try {
            sh "sh /app/EAP-7.0.0/bin/jboss-cli.sh -c --controller=${CONTROLLER_IP} --user=${JBOSS_USER} --password=${JBOSS_PASSWORD} --command=\" undeploy iot-*.war --keep-content \" "
            sh "sh /app/EAP-7.0.0/bin/jboss-cli.sh -c --controller=${CONTROLLER_IP} --user=${JBOSS_USER} --password=${JBOSS_PASSWORD} --command=\" deploy ${PROJECT_PATH}/${TARGET_WAR} \" "
        } catch (Exception e) {
            // Error handling: Deployment failed
            error "Deployment failed: ${e.message}"
        }

        // Verification
        try {
            sh "sh /app/EAP-7.0.0/bin/jboss-cli.sh -c --controller=${CONTROLLER_IP} --user=${JBOSS_USER} --password=${JBOSS_PASSWORD} --command=\" deployment-info --name=*war \"  "
        } catch (Exception e) {
            // Error handling: Verification failed
            error "Verification failed: ${e.message}"
        }
    }
}
